<template>
  <div>
    <h1>{{ $route.params.roomId }}ルームへようこそ</h1>
    <div>
      <textarea v-model="message"></textarea>
      <br/>
      <button v-on:click="submit">送信</button>
      <ul>
        <li v-for="(message, index) in messages" :key="index">
          {{ message.name }} => {{ message.message }}
        </li>
      </ul>
    </div>
  </div>
</template>

<script>
import { w3cwebsocket } from 'websocket'
const W3CWebSocket = w3cwebsocket
export default {
  // layout: 'protected',
  middleware: 'authenticated',
  name: "roomId",
  computed: {
    user(state) {
      return this.$store.getters['modules/user/user']
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
      Member: {
        0: "マスター"
      },
    }
  },
  mounted() {
    // 接続したかを確認する
    this.socket.onopen = (event) => {
      console.log("安全に接続しました")
    }
    // メッセージの更新があれば受け取る
    this.socket.onmessage = (event) => {
      // 受け取った文字列をjsonにparseします
      const response = JSON.parse(event.data)
      // イベント処理
      // トークンを受け取った際の処理
      if (response.event === "token") {
        console.log("tokenを受け取りました")
        this.IAM.token = response.token
      }
      // ルームの誰かが送信した場合に処理されます
      if (response.event === "member-post") {
        console.log("メッセージを受信しました")
        this.messages.push({name: response.name, message: response.message})
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
        const obj = {event: 'post', token: this.IAM.token, name: this.user.name, message: this.message}
        this.socket.send(JSON.stringify(obj))
        this.message = ""
      } else {
        return false;
      }
    }
  }
}
</script>

<style scoped>

</style>
