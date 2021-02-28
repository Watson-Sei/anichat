<template>
  <div class="container">
    <div class="hello">
      <h1 v-if="user">Welcome to {{ user.email }}🎉</h1>
      <nuxt-link to="/admin">Admin Link</nuxt-link><br>
      <nuxt-link to="/room">Room List</nuxt-link><br>
      <button @click="Logout">Sign out</button>
    </div>
  </div>
</template>

<script>
export default {
  middleware: 'authenticated',
  data() {
    return {
      user: this.$store.state.authUser
    }
  },
  methods: {
    async Logout() {
      await this.$fire.auth.signOut()
      localStorage.delete('access_token')
    },
  }
}
</script>
