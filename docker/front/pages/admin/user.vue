<template>
  <User :users="users"/>
</template>

<script>
import User from '@/components/User'

export default {
  layout: 'management',
  name: "user",
  components: {User},
  data() {
    return {
      users: []
    }
  },
  mounted() {
    this.$axios.get("http://localhost/api/admin/users", {
      headers: {
        'Authorization': `Bearer ${localStorage.getItem('access_token')}`
      }
    }).then((response) => {
      const obj = JSON.parse(response.data)
      this.users = obj.users
    }).catch((error) => {
      console.log(error)
    })
  }
}
</script>

<style scoped>

</style>
