<template>
  <div>
    <textarea v-model="message"></textarea>
    <br/>
    <button v-on:click="send">送信</button>
    <p>message: {{ messages }}</p>
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
      messages: []
    }
  },
  mounted: function () {
    this.socket.onmessage = function (event) {
      console.log("websocket.onmessage()")

      if (event && event.data) {
        console.log(event.data)
        this.messages.push(event.data.string)
      }
    }
  },
  methods: {
    send: function () {
      this.socket.send(this.message)
    }
  }
}
</script>
