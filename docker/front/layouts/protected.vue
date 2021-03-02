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
      </v-list>
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
    }
  },
}
</script>

<style scoped>

</style>
