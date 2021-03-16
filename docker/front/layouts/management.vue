<template>
  <v-app > <!-- dark -->
    <v-navigation-drawer app v-model="drawer">
      <v-list-item>
        <v-list-item-title class="title">
          Application <a href="/room"><v-icon>mdi-vector-link</v-icon></a>
        </v-list-item-title>
      </v-list-item>
      <v-divider />
      <v-list nav>
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
          <v-btn block class="black white--text" @click="Logout">
            Logout
          </v-btn>
        </div>
      </template>
    </v-navigation-drawer>

    <v-app-bar app>
      <v-app-bar-nav-icon @click.stop="drawer = !drawer" />
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
  data() {
    return {
      drawer: false,
      menus: [
        { title: 'Dashboard', icon: 'mdi-web', url: '/admin/'},
        { title: 'Room', icon: 'mdi-google-cloud', url: '/admin/room/'},
        { title: 'User', icon: 'mdi-account-box', url: '/admin/user/'}
      ]
    }
  },
  methods: {
    async Logout() {
      await this.$fire.auth.signOut()
      localStorage.removeItem('access_token')
      localStorage.removeItem('refresh_token')
    },
  }
}
</script>
