<template>
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
</template>

<script>
import { w3cwebsocket } from 'websocket'
const W3CWebSocket = w3cwebsocket
export default {
  data() {
    return {
      socket: new W3CWebSocket('ws://localhost/api/ws'),
      // クライアントから送る値
      message: '',
      // サーバから受け取る値
      messages: [],
      // Token
      IAM: {
        token: null,
        name: null,
        is_join: false
      }
    }
  },
  mounted() {
    // メッセージの更新があれば受け取る
    this.socket.onmessage = (event) => {
      // json
      let response = JSON.parse(event.data)
      if (response.event === "token") {
        console.log(response.token)
        this.IAM.token = response.token
      } else if (response.event === "message-post") {
        console.log(response.message)
        this.messages.push(response.message)
      }
    }
  },
  methods: {
    // 送信ボタン関数
    submit: function () {
      // 存在するか１文字以上存在するかを確認 all true
      if (this.message && this.message != "") {
        let obj = { event: 'post', token: this.IAM.token, message: this.message}
        this.socket.send(JSON.stringify(obj))
        this.message = ""
      } else {
        // なければ理由を送信
        console.log("Empty messages cannot be sent.")
      }
    },
  }
}
</script>
