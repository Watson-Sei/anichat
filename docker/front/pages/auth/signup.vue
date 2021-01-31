<template>
  <div class="container">
    <div class="signup">
      <h2>Sign up</h2>
      <form @submit.prevent="signUp">
        <label for="usernameTxt">Username:</label>
        <input id="usernameTxt" type="text" v-model="email">
        <label for="passwordTxt">Password:</label>
        <input id="passwordTxt" type="password" v-model="password">
        <button type="submit">Sign Up</button>
      </form>
      <p>Do you have an account?
        <nuxt-link to="/auth/signin">sign in now!!</nuxt-link>
      </p>
    </div>
  </div>
</template>

<script>
import { mapActions } from 'vuex'
import firebase from '~/plugins/firebase'
export default {
  name: "signup",
  data() {
    return {
      email: '',
      password: '',
    }
  },
  methods: {
    ...mapActions('modules/user', [ 'login' ]),
    async signUp() {
      try {
        const firebaseUser = await firebase.auth().createUserWithEmailAndPassword(this.email, this.password)
        await this.writeUserData(firebaseUser.user.uid, firebaseUser.user.email)
        await this.login(firebaseUser.user.uid)
        await this.$router.push('/')
      } catch (error) {
        console.log(error.message)
        console.log('クリエイト時のエラー')
      }
    },
    writeUserData (userId, email) {
      return firebase.database().ref('users/' + userId).set({
        email: email
      })
    }
  }
}
</script>

<style scoped>
</style>
