<template>
  <div>
    <h1>Room一覧</h1>
  </div>
</template>

<script>
import firebaseApp from '~/plugins/firebase'
export default {
  // layout: 'protected',
  middleware: 'authenticated',
  name: "index",
}
</script>

<style scoped>

</style>
