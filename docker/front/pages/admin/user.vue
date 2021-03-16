<template>
  <User :users="users" @get-users="GetUsers"/>
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
    this.GetUsers()
  },
  methods: {
    GetUsers(key, value) {
      this.$axios.get("http://localhost/api/admin/users", {
        headers: {
          'Authorization': `Bearer ${localStorage.getItem('access_token')}`
        }
      }).then((response) => {
        const obj = JSON.parse(response.data)
        this.users = obj.users
        console.log("Get Users実行されました")
      }).catch((error) => {
        console.log(error)
      })
    },
  }
}
</script>

<style scoped>

</style>
