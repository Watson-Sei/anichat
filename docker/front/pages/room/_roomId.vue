<template>
  <v-app>
    <v-container class="fill-height">
      <v-row class="fill-height pb-14" align="end">
        <v-col>
          <div v-for="(item, index) in messages" :key="index" :class="['d-flex flex-row align-center my-2', item.class == 'msg-me' ? 'justify-end': null]">
            <span v-if="item.class == 'msg-me'" class="blue--text mr-3">{{ item.message }}</span>
            <v-avatar :color="item.class == 'msg-me' ? 'indigo':'red'" size="36">
              <span class="white--text">{{ item.name[0] }}</span>
            </v-avatar>
            <span v-if="item.class != 'msg-me'" class="blue--text ml-3">{{ item.message }}</span>
          </div>
        </v-col>
      </v-row>
    </v-container>
    <v-footer inset fixed app>
      <v-container class="ma-0 pa-0">
        <v-row no-gutters>
          <v-col>
            <div class="d-flex flex-row align-center">
              <v-text-field v-model="message" placeholder="Type Something" @keypress.enter="send"></v-text-field>
              <v-btn icon class="ml-4" @click="send"><v-icon>mdi-send</v-icon></v-btn>
            </div>
          </v-col>
        </v-row>
      </v-container>
    </v-footer>
  </v-app>
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
  watch: {
    messages: function () {
      this.$vuetify.goTo(99999)
    }
  },
  data() {
    return {
      socket: new W3CWebSocket(`wss://${process.env.WebSocket_URL}/ws/${this.$route.params.roomId}`),
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
    send: function () {
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
  },
  beforeRouteLeave(to, from, next) {
    const answer = window.confirm("本当に退室しますか？")
    if (answer) {
      this.socket.send(JSON.stringify({
        event: "quit",
        token: this.IAM.token
      }))

      // ボタンを無効にする
      this.isQuit = true
      next();
    } else {
      next(false)
    }
  },
}
</script>

<style scoped>
.msg-master {
  margin-right: auto;
}
.msg-member {
  margin-right: auto;
}
.msg-me {
  margin-left: auto;
}
</style>
