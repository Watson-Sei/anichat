import firebase from "firebase";

export const state = () => ({
  authUser: null
})

export const getters = {
  isLoggedIn: state => !!state.authUser,
  isAdminIn: state => !!state.authUser.isAdmin
}

export const actions = {
  async nuxtServerInit({ dispatch }, ctx) {
    if (this.$fire.auth === null) {
      throw 'nuxtServerInit Example not working - this.$fire.auth cannot be accessed.'
    }

    if (ctx.$fire.auth === null) {
      throw 'nuxtServerInit Example not working - ctx.$fire.auth cannot be accessed.'
    }

    if (ctx.app.$fire.auth === null) {
      throw 'nuxtServerInit Example not working - ctx.$fire.auth cannot be accessed.'
    }

    if (ctx.res && ctx.res.locals && ctx.res.locals.user) {
      const { allClaims: claims, ...authUser } = ctx.res.locals.user

      console.info(
        'Auth User verified on server-side. User: ',
        authUser,
        'Claims:',
        claims
      )

      await dispatch('onAuthStateChanged', {
        authUser,
        claims,
      })
    }
  },

  async onAuthStateChanged({ commit }, { authUser, claims }) {
    if (!authUser) {
      await this.$router.push("/auth/signin")
      commit('RESET_STORE')
      return
    }

    if (authUser && claims) {
      try {
        await this.$router.push("/room")
      } catch (e) {
        console.error(e)
      }
    }

    commit('SET_USER', { authUser, claims })
  },

  async refreshToken() {

  }
}

export const mutations = {
  RESET_STORE(state) {
    state.authUser = null
  },
  SET_USER(state, { authUser, claims }) {
    state.authUser = {
      uid: authUser.uid,
      email: authUser.email,
      emailVerified: authUser.emailVerified,
      displayName: authUser.displayName,
      photoURL: claims.picture,
      isAdmin: claims.admin
    }
  }
}
