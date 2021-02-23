<template>
  <div>
    <h1>Room一覧</h1>
    <nuxt-link to="/room/ドクターストーン">Go to Chat</nuxt-link>
  </div>
</template>

<script>
export default {
  // layout: 'protected',
  middleware: 'authenticated',
  name: "index",
}
</script>

<style scoped>

</style>
