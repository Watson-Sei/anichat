<template>
  <v-container>
    <v-row style="height: 200px;" justify="center" align-content="center">
      <v-col class="ma-9" cols="6">
        <v-text-field
          class="pa-4"
          v-model="search"
          label="タイトル検索"
          type="text"
        >
          <template v-slot:append-outer>
            <v-btn color="primary" @click="letsearch">検索</v-btn>
          </template>
        </v-text-field>
      </v-col>
    </v-row>
    <v-row justify="center">
      <v-card max-width="500px" v-for="(item, index) in dataset.room" :key="index" class="ma-2">
        <v-row>
          <v-col cols="4">
            <v-img :src="item.img" />
          </v-col>
          <v-col>
            <v-card-title>
              {{ item.title }}
            </v-card-title>
            <v-card-subtitle>
              {{ item.time }}
            </v-card-subtitle>
            <v-card-actions>
              <v-btn class="primary" v-if="item.public" @click="returnChat(`${item.chatId}`)">Go to Chat</v-btn>
              <v-btn disabled v-else>Go to Chat</v-btn>
            </v-card-actions>
          </v-col>
        </v-row>
      </v-card>
    </v-row>
  </v-container>
</template>

<script>
export default {
  layout: 'protected',
  middleware: 'authenticated',
  name: "index",
  data() {
    return {
      search: '',
      dataset: {
        room: [
          { chatId: 1, title: "ドクターストーン#1", time: "22:30 ~ 23:00", public: true, img: "https://eiga.k-img.com/images/anime/program/104831/photo/777ee2e888036ca0/320.jpg?1456654190"},
          { chatId: 2, title: "進撃の巨人 Season 3 #11", time: "23:00 ~ 23:30", public: false, img: "https://eiga.k-img.com/images/anime/program/106280/photo/d581dab61c13028f/320.jpg?1524821461"}
        ]
      }
    }
  },
  methods: {
    letsearch () {
      console.log(this.search)
    },
    returnChat (val) {
      const path = `/room/${val}/`
      this.$router.push(path)
    }
  }
}
</script>

<style scoped>
</style>
