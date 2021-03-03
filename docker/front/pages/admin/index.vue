<template>
  <div>
    <h1>Admin Page</h1>
  </div>
</template>

<script>
export default {
  layout: 'management',
  name: "index",
  middleware: 'handle-admin-route',
}
</script>

<style scoped>

</style>
