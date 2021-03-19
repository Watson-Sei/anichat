<template>
  <div>
    <Room :rooms="rooms" @get-rooms="GetRooms" @delete-rooms="DeleteRoom"/>
  </div>
</template>

<script>
import Room from "@/components/Room";
import axios from "axios";

export default {
  layout: 'management',
  name: "room",
  components: {Room},
  data() {
    return {
      rooms: []
    }
  },
  mounted() {
    this.GetRooms()
  },
  methods: {
    GetRooms() {
      axios.get(`${process.env.API_URL}/rooms`)
      .then((response) => {
        this.rooms = response.data.data
      }).catch((error) => {
        console.log(error)
      })
    },
    DeleteRoom(id) {
      axios.delete(`${process.env.API_URL}/admin/rooms/${id}`, {
        headers: {
          'Authorization': `Bearer ${localStorage.getItem('access_token')}`
        }
      }).then(() => {
        console.log("delete room")
      })
    }
  }
}
</script>

<style scoped>

</style>
