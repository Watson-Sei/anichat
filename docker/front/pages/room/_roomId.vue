<template>
  <div>
    <h1>{{ $route.params.roomId }}ルームへようこそ</h1>
    <div>
      <textarea v-model="message"></textarea>
      <button id="post-button" v-on:click="submit">送信</button>
      <button v-bind:disabled="isQuit" id="quit-button" v-on:click="quit">退室</button>
      <!-- 発言ログ -->
      <ul>
        <li v-for="(data, index) in messages" :key="index">
          <span v-bind:class="data.class"><span class="name">{{ data.name }}</span>> {{ data.message }}</span>
        </li>
      </ul>
      <!-- メンバー一覧 -->
    </div>
  </div>
</template>

<script>
import { w3cwebsocket } from 'websocket'
const W3CWebSocket = w3cwebsocket
export default {
  layout: 'protected',
  middleware: 'authenticated',
  name: "roomId",
  computed: {
    user(state) {
      return this.$store.state.authUser
    }
  },
  data() {
    return {
      socket: new W3CWebSocket(`ws://localhost/api/ws/${this.$route.params.roomId}`),
      // 送信するメッセージ
      message: '',
      // 受信したメッセージ
      messages: [],
      // Token etc...
      IAM: {
        token: null,
        name: null,
        is_join: false
      },
      // メンバーリスト
      Member: {
        0: "マスター"
      },
      isQuit: false
    }
  },
  mounted() {
    // 接続したかを確認する
    this.socket.onopen = (event) => {
      console.log("安全に接続しました")
    }
    // 切断ハンドラ
    this.socket.onclose = (event) => {
      console.log("切断されました")
      this.$router.push({path: "/room"})
    }
    // メッセージの更新があれば受け取る
    this.socket.onmessage = (event) => {
      // 受け取った文字列をjsonにparseします
      const data = JSON.parse(event.data)
      console.log(data)
      // イベント処理
      // トークンを受け取った際の処理
      if (data.event === "token") {
        console.log("tokenを受け取りました")
        this.IAM.token = data.token
        this.IAM.name = this.user.displayName

        // event:join送信
        console.log("join発火")
        this.socket.send(JSON.stringify({
          event: "join",
          token: this.IAM.token,
          name: this.IAM.name
        }))
      }
      // 入室結果が返ってきた
      if ( data.event === "join-result") {
        // 正常に入室できた
        if( data.status ){
          // 入室フラグを立てる
          this.IAM.is_join = true;

          // すでにログイン中のメンバー一覧を反映
          console.log("すでにログインしているユーザー:", data.list)
          if( data.list !== null ){
            for(let i=0; i<data.list.length; i++){
              const cur = data.list[i];
              if( ! (cur.token in this.Member) ){
                this.addMemberList(cur.token, cur.name)
              }
            }
          }
        } else {
          alert("入室できませんでした")
        }
      }
      // 誰かが入室した
      if ( data.event === "member-join") {
        if( this.IAM.is_join ){
          this.addMessageFromMaster(`${data.name}さんが入室しました`);
          this.addMemberList(data.token, data.name)
        }
      }
      // ルームの誰かが送信した場合に処理されます
      if (data.event === "member-post") {
        console.log("メッセージを受信しました")
        console.log(data)
        if( this.IAM.is_join ){
          const is_me = (data.token === this.IAM.token);
          this.addMessage(data, is_me);
        }
      }
      // 退室処理の結果が返ってきた
      if (data.event === "quit-result") {
        if( data.status ){
          this.gotoSTEP1();
        } else {
          alert("退室できませんでした");
        }

        // ボタンを有効に戻す
        this.isQuit = false
      }
      // 誰かが退室したら
      if (data.event === "member-quit") {
        if( this.IAM.is_join ){
          const name = this.Member[data.token]
          this.addMessageFromMaster(`${name}さんが退室しました`);
          this.removeMemberList(data.token);
        }
      }
    }
  },
  methods: {
    // ルームにユーザー情報を渡す
    // 送信ボタン関数
    submit: function () {
      // 接続は維持されているか、チャット内容は空で実行されて無いかなどチェックをして送信をします。
      if (!this.socket) {
        return false;
      }
      if (this.message && this.message != "") {
        const obj = {event: 'post', token: this.IAM.token, message: this.message}
        this.socket.send(JSON.stringify(obj))
        this.message = ""
      } else {
        return false;
      }
    },
    addMessage(msg, is_me=false){
      const name = this.Member[msg.token]

      if( msg.token === 0 ){
        this.messages.push({class: "msg-master", name: name, message: msg.message})
      } else if( is_me ) {
        this.messages.push({class: "msg-me", name: name, message: msg.message})
      } else {
        this.messages.push({class: "msg-member", name: name, message: msg.message})
      }
    },
    addMessageFromMaster(msg){
      this.addMessage({token: 0, message: msg})
    },
    addMemberList(token, name){
      // 内部変数に保存
      this.Member[token] = name;
      console.log("内部変数が保存されました", this.Member)
    },
    quit() {
      if( confirm("本当に退室しますか？") ){
        this.socket.send(JSON.stringify({
          event: "quit",
          token: this.IAM.token
        }))

        // ボタンを無効にする
        this.isQuit = true
      }
    },
    removeMemberList(token) {
      // 内部変数から削除
      delete this.Member[token];
    },
    gotoSTEP1() {
      // 自分の情報を初期化
      this.IAM.token = null;
      this.IAM.name = null;
      this.IAM.is_join = null;

      //　メンバー一覧を初期化
      for( let key in this.Member ){
        if( key !== "0"){
          delete this.Member[key];
        }
      }

      this.$router.push("/room")
    }
  }
}
</script>

<style scoped>

</style>
