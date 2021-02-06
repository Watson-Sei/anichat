<template>
  <div class="container">
    <div class="hello">
      <h1 v-if="user">Welcome to {{ user.name }}🎉</h1>
      <button @click="hello">Hello</button>
      <button @click="Logout">Sign out</button>
    </div>
  </div>
</template>
<script>
import Cookies from 'js-cookie'
export default {
  middleware: 'authenticated',
  computed: {
    user(state) {
      return this.$store.getters['modules/user/user'];
    }
  },
  methods: {
    async Logout() {
      await this.$store.dispatch('modules/user/logout')
      await this.$router.push("/auth/signin")
    },
    hello() {
      const response = this.$axios.$get('http://localhost/api/hello',{
        headers: {'Authorization': `Bearer ${Cookies.get('access_token')}`}
      })
      .then(function (response) {
        console.log(response)
      })
    }
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
