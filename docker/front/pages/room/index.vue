<template>
  <v-container>
    <v-row class="search-box" style="margin: auto" justify="center" align-content="center">
      <v-col>
        <v-text-field
          class="pa-2"
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
      <v-card v-for="(item, index) in dataset.searchrooms" :key="index" class="ma-2 card-w">
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
              <v-btn class="primary" v-if="item.public" @click="returnChat(`${item.id}`)">Go to Chat</v-btn>
              <v-btn disabled v-else>Go to Chat</v-btn>
            </v-card-actions>
          </v-col>
        </v-row>
      </v-card>
    </v-row>
  </v-container>
</template>

<script>
import axios from "axios";

export default {
  layout: 'protected',
  middleware: 'authenticated',
  name: "index",
  data() {
    return {
      search: '',
      dataset: {
        defaultrooms: [],
        searchrooms: [],
      }
    }
  },
  methods: {
    letsearch () {
      console.log(this.search)
      if (this.search === "") {
        console.log("検索内容は空です")
        this.dataset.searchrooms = this.dataset.defaultrooms
      } else {
        if (this.dataset.defaultrooms.length !== null) {
          console.log("配列データはあります。")
          let reset = []
         this.dataset.defaultrooms.forEach((elem, index) => {
           if (elem.title.includes(this.search)) {
             reset.push(elem)
           }
         })
          if (!reset) {
            this.dataset.searchrooms = this.dataset.defaultrooms
          } else {
            this.dataset.searchrooms = reset
          }
        }
      }
    },
    returnChat (val) {
      const path = `/room/${val}/`
      this.$router.push(path)
    }
  },
  mounted() {
    axios.get(`${process.env.BASE_URL}/api/rooms`)
    .then((response) => {
      this.dataset.defaultrooms = response.data.data
      this.dataset.searchrooms = response.data.data
    })
  }
}
</script>

<style scoped>
.card-w {
  max-width: 500px;
}
.search-box {
  max-width: 800px;
}
@media screen and (max-width: 540px) {
  .card-w {
    max-width: 350px;
  }
}
</style>
