<template>
  <v-card class="ma-4">
    <v-card-title>
      Room List
      <v-spacer></v-spacer>
    </v-card-title>

    <v-data-table
      :headers="headers"
      :items="rooms"
      :search="search"
      class="elevation-1"
    >
      <template v-slot:top>
        <v-text-field
          v-model="search"
          label="Search User"
          class="mx-4"
        ></v-text-field>
        <v-spacer></v-spacer>

        <!-- 更新処理 -->
        <v-dialog
          v-model="dialog"
          max-width="500px"
        >
          <template v-slot:activator="{ on, attrs }">
            <v-btn
              color="primary"
              dark
              class="mb-2"
              v-bind="attrs"
              v-on="on"
            >
              New Item
            </v-btn>
          </template>
          <v-card>
            <v-card-title>
              <span class="handline">{{ formTitle }}</span>
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
                      v-model="editedItem.title"
                      label="Room title"
                    ></v-text-field>
                  </v-col>
                  <v-col
                    cols="12"
                    sm="6"
                    md="4"
                  >
                    <v-text-field
                      v-model="editedItem.time"
                      label="Room Time"
                    ></v-text-field>
                  </v-col>
                  <v-col
                      cols="12"
                      sm="6"
                      md="4"
                  >
                    <v-text-field
                      v-model="editedItem.img"
                      label="Room Image Link"
                    ></v-text-field>
                  </v-col>
                  <v-col
                    cols="12"
                    sm="6"
                    md="4"
                  >
                    <v-checkbox
                      v-model="editedItem.public"
                      label="Room Public"
                    ></v-checkbox>
                  </v-col>
                </v-row>
              </v-container>
            </v-card-text>

            <v-card-subtitle>
              <v-spacer></v-spacer>
              <v-btn
                color="blue darken-1"
                text
                @click="close"
              >
                Cancel
              </v-btn>
              <v-btn
                color="blue darken-1"
                text
                @click="save"
              >
                Save
              </v-btn>
            </v-card-subtitle>
          </v-card>
        </v-dialog>

        <v-dialog v-model="dialogDelete" max-width="500px">
          <v-card>
            <v-card-title class="headline">
              Are you sure you want to delete this item?
            </v-card-title>
            <v-card-actions>
              <v-spacer></v-spacer>
              <v-btn color="blue darken-1" text @click="closeDelete">Cancel</v-btn>
              <v-btn color="blue darken-1" text @click="deleteItemConfirm">OK</v-btn>
              <v-spacer></v-spacer>
            </v-card-actions>
          </v-card>
        </v-dialog>

      </template>
      <template v-slot:item.time="{ item }">
        {{ item.time }}
      </template>
      <template v-slot:item.public="{ item }">
        <v-simple-checkbox
          v-model="item.public"
          disabled
        ></v-simple-checkbox>
      </template>

      <template v-slot:item.actions="{ item }">
        <v-icon
          class="mr-2"
          @click="editItem(item)"
        >
          mdi-pencil
        </v-icon>
        <v-icon
          @click="deleteItem(item)"
        >
          mdi-delete
        </v-icon>
      </template>
    </v-data-table>
  </v-card>
</template>

<script>
import axios from 'axios'

export default {
  name: "Room",
  props: ['rooms'],
  data() {
    return {
      dialog: false,
      dialogDelete: false,
      headers: [
        {text: 'ID', value: 'id'},
        {text: 'Title', value: 'title'},
        {text: 'Time', value: 'time'},
        {text: 'Public', value: 'public'},
        {text: 'Image', value: 'img'},
        {text: '操作', value: 'actions', align: 'end', sortable: false}
      ],
      editedIndex: -1,
      editedItem: {
        id: 0,
        title: '',
        time: '',
        img: '',
        public: false
      },
      defaultItem: {
        id: 0,
        title: '',
        time: '',
        img: '',
        public: false
      },
      search: '',
    }
  },
  computed: {
    formTitle () {
      if(this.editedIndex === -1) {
        return 'Edit Item'
      }
    },
  },
  watch: {
    dialog (val) {
      val || this.close()
    },
    dialogDelete (val) {
      val || this.closeDelete()
    }
  },
  methods: {
    editItem (item) {
      this.editedIndex = this.rooms.indexOf(item)
      this.editedItem = Object.assign({}, item)
      this.dialog = true
    },
    deleteItem (item) {
      this.editedIndex = this.rooms.indexOf(item)
      this.editedItem = Object.assign({}, item)
      this.dialogDelete = true
    },
    deleteItemConfirm () {
      this.rooms.splice(this.editedIndex, 1)
      this.$nextTick(() => {
        this.$emit("delete-rooms", this.editedItem.id)
      })
      this.closeDelete()
    },
    close () {
      this.dialog = false
      this.$nextTick(() => {
        this.editedItem = Object.assign({}, this.defaultItem)
        this.editedIndex = -1
      })
    },
    closeDelete () {
      this.dialogDelete = false
      this.$nextTick(() => {
        this.editedItem = Object.assign({}, this.defaultItem)
        this.editedIndex = -1
      })
    },
    save () {
      if (this.editedIndex > -1) {
        const data = {
          id: this.editedItem.id,
          title: this.editedItem.title,
          time: this.editedItem.time,
          public: this.editedItem.public
        }
        axios.put(`http://localhost/api/admin/rooms/${this.editedItem.id}`, data, {
          headers: {
            'Authorization': `Bearer ${localStorage.getItem('access_token')}`
          }
        }).then(() => {
          console.log("room info update")
          this.$nextTick(() => {
            this.$emit("get-rooms")
          })
        })
      } else {
        const data = {
          title: this.editedItem.title,
          time: this.editedItem.time,
          img: this.editedItem.img,
          public: this.editedItem.public
        }
        axios.post("http://localhost/api/admin/rooms", data, {
          headers: {
            'Authorization': `Bearer ${localStorage.getItem('access_token')}`
          }
        }).then((response) => {
          console.log(response)
        })
        this.rooms.push(this.editedItem)
      }
      this.close()
    }
  }
}
</script>

<style scoped>

</style>
