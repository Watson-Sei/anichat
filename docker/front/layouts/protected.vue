<template>
  <v-app>
    <v-navigation-drawer app v-model="drawer">
      <v-list-item>
        <v-list-item-title class="title">
          Application
        </v-list-item-title>
      </v-list-item>
      <v-divider />
      <v-list nav>
        <v-list-item class="px-2">
          <v-list-item-avatar>
            <v-img :src="this.$store.state.authUser.photoURL"></v-img>
          </v-list-item-avatar>

          <v-list-item-title>Sei Watson</v-list-item-title>

          <v-list-item-icon>
            <nuxt-link to="/profile">
              <v-icon>mdi-cog-outline</v-icon>
            </nuxt-link>
          </v-list-item-icon>
        </v-list-item>
        <v-list-item v-for="menu in menus" :key="menu.title" :to="menu.url">
          <v-list-item-icon>
            <v-icon>{{ menu.icon }}</v-icon>
          </v-list-item-icon>
          <v-list-item-content>
            <v-list-item-title>{{ menu.title }}</v-list-item-title>
          </v-list-item-content>
        </v-list-item>
      </v-list>

      <template v-slot:append>
        <div class="pa-2">
          <v-btn block class="black white--text">
            Logout
          </v-btn>
        </div>
      </template>
    </v-navigation-drawer>

    <v-app-bar app>
      <v-app-bar-nav-icon @click.stop="drawer = !drawer"></v-app-bar-nav-icon>
    </v-app-bar>

    <v-main>
      <Nuxt />
    </v-main>

    <v-footer app>

    </v-footer>
  </v-app>
</template>

<script>
export default {
  name: "protected",
  data() {
    return {
      drawer: false,
      menus: [
        { title: 'Room List', icon: 'mdi-format-list-bulleted-triangle', url: '/room/'},
      ]
    }
  },
}
</script>

<style scoped>

</style>
