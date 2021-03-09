<template>
  <div>
    <nuxt-link v-if="this.$store.getters['isLoggedIn']" to="/room">Go to AniChat.jp</nuxt-link>
    <nuxt-link v-else to="/auth/signin">Login</nuxt-link>
  </div>
</template>

<script>
export default {
  layout: 'default',
}
</script>
