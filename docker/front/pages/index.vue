<template>
  <div class="container">
    <div class="hello">
      <h1>Welcome to {{ user.name }}🎉</h1>
      <button @click="signOut">Sign out</button>
    </div>
  </div>
</template>
<script>
import Cookies from 'js-cookie'
import firebase from '~/plugins/firebase'
export default {
  middleware: 'authenticated',
  computed: {
    user(state) {
      return this.$store.getters['modules/user/user'];
    }
  },
  methods: {
    signOut() {
      firebase.auth().signOut().then(() => {
        // localStorage.removeItem('jwt')
        Cookies.remove('access_token')
        this.$router.push('/auth/signin')
      })
    },
  }
}
</script>

<style scoped>
h1, h2 {
  font-weight: normal;
}
ul {
  list-style-type: none;
  padding: 0;
}
li {
  display: inline-block;
  margin: 0 10px;
}
a {
  color: #42b983;
}
button {
  margin: 10px 0;
  padding: 10px;
}
</style>
