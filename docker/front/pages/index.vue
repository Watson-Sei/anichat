<template>
  <div>
    <textarea v-model="message"></textarea>
    <br/>
    <button v-on:click="send">送信</button>
    <ul>
      <li v-for="(message, index) in messages" :key="index">
        {{ message }}
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
      messages: []
    }
  },
  mounted() {
    this.socket.onmessage = (event) => {
      console.log(typeof event.data)
      this.messages.push(event.data)
      console.log(this.messages)
    }
  },
  methods: {
    send: function () {
      this.socket.send(this.message)
    },
    add: function (msg) {
      this.messages.push(msg)
    }
  }
}
</script>
