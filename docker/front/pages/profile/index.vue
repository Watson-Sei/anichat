<template>
  <v-container>
    <v-row style="padding: 54px 24px">
      <v-col class="v-r">
        <v-avatar
          size="128"
        >
          <img :src="user.photoURL">
        </v-avatar>
      </v-col>
      <v-col style="margin: auto" class="v-l">
        <h2>{{ user.displayName }}</h2>
        <p>{{ user.email }}</p>
        <v-btn
          color="green"
          dark
          @click="editUser"
        >
          編集
        </v-btn>
        <v-dialog
          v-model="dialog"
          persistent
          max-width="600px"
        >
          <v-card>
            <v-card-title>
              <span class="headline">Profile Edit</span>
            </v-card-title>
            <v-card-text>
              <v-container>
                <v-row>
                  <v-col
                    cols="12"
                    sm="6"
                    md="4"
                  >
                    <v-text-field
                      v-model="editedItem.displayName"
                      label="Display Name*"
                      required
                    ></v-text-field>
                  </v-col>
                  <v-col
                    cols="12"
                    sm="6"
                    md="4"
                  >
                    <v-text-field
                      v-model="editedItem.email"
                      label="Email Address*"
                      required
                    ></v-text-field>
                  </v-col>
                </v-row>
              </v-container>
            </v-card-text>
            <v-card-actions>
              <v-spacer></v-spacer>
              <v-btn
                color="blue darken-1"
                text
                @click="close"
              >
                Close
              </v-btn>
              <v-btn
                color="blue darken-1"
                text
                @click="save"
              >
                Save
              </v-btn>
            </v-card-actions>
          </v-card>
        </v-dialog>
      </v-col>
    </v-row>
  </v-container>
</template>

<script>
import firebase from 'firebase'
import { mapMutations } from 'vuex'

export default {
  layout: "protected",
  name: "index",
  data() {
    return {
      dialog: false,
      user: {
        displayName: "",
        email: "",
        photoURL: "",
      },
      editedItem: {
        displayName: "",
        email: "",
      },
      defaultItem: {
        displayName: "",
        email: ""
      }
    }
  },
  mounted() {
    this.user = firebase.auth().currentUser;
  },
  watch: {
    dialog (val) {
      val || this.close()
    },
  },
  methods: {
    ...mapMutations({
      update: 'UPDATE_USER'
    }),
    editUser () {
      this.editedItem = Object.assign({}, this.user)
      this.dialog = true
    },
    close () {
      this.dialog = false
      this.$nextTick(() => {
        this.editedItem = Object.assign({}, this.defaultItem)
      })
    },
    save() {
      let user = firebase.auth().currentUser;
      user.updateProfile({
        displayName: this.editedItem.displayName,
        email: this.editedItem.email
      }).then(function () {
        console.log("update successful")
      }).catch(function (error) {
        console.log(error)
      })
      this.update(firebase.auth().currentUser)
      this.close()
      location.reload()
    },
    ...mapMutations({
      update: 'UPDATE_USER'
    })
  }
}
</script>

<style scoped>
.v-r {
  text-align: right;
}
.v-l {
  text-align: left;
}
@media screen and (max-width: 415px) {
  .v-r {
    text-align: center;
  }
  .v-l {
    text-align: center;
  }
}
</style>
