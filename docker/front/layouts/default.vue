<template>
  <v-app style="background-color: #e6e6e6">
    <header>
      <div class="header-box">
        <h1>
          <nuxt-link to="/" class="n-link">AniChat</nuxt-link>
        </h1>
        <nav class="pc-nav">
          <ul>
            <li>
              <nuxt-link v-if="this.$route.path === '/'" to="#merit" class="n-link">特徴</nuxt-link>
              <nuxt-link v-else to="/" class="n-link">特徴</nuxt-link>
            </li>
            <li><nuxt-link to="/terms" class="n-link">利用規約</nuxt-link></li>
            <li><nuxt-link to="/support" class="n-link">サポート</nuxt-link></li>
            <li>
              <nuxt-link v-if="this.$store.getters['isLoggedIn']" to="/room" class="n-link btn-gradient-radius">チャットを開始</nuxt-link>
              <nuxt-link v-else to="/auth/signin" class="n-link btn-gradient-radius">ログイン</nuxt-link>
            </li>
          </ul>
        </nav>
        <div id="hamburger" @click="toggle">
          <span></span>
        </div>
        <nav class="sp-nav" v-bind:class="{toggle: Toggle}">
          <ul>
            <li>
              <nuxt-link v-if="this.$route.path === '/'" to="#merit" class="n-link" @click="toggle">特徴</nuxt-link>
              <nuxt-link v-else to="/" class="n-link" @click.native="toggle">特徴</nuxt-link>
            </li>
            <li><nuxt-link to="/terms" class="n-link" @click.native="toggle">利用規約</nuxt-link></li>
            <li><nuxt-link to="/support" class="n-link" @click.native="toggle">サポート</nuxt-link></li>
            <li><nuxt-link to="/room" class="n-link" @click="toggle">チャットを開始</nuxt-link></li>
            <li class="close" @click="toggle"><span>閉じる</span></li>
          </ul>
        </nav>
      </div>
      <div class="main-visual" v-if="this.$route.path === '/'">
        <div class="catch-box">
          <div class="catch-sub-box catch-text">
            <h3>いっしょに、</h3>
            <h1>リアタイの感想を共有しよう！</h1>
            <p>AniChatは、地上波放送などのアニメの視聴感想をリアルタイムで発信できるチャットサービスです</p>
            <nuxt-link to="#" class="btn-gradient-radius n-link">チャットを開始</nuxt-link>
          </div>
          <div class="catch-sub-box">
            <img src="../static/img/header-chat.svg">
          </div>
        </div>
      </div>
    </header>
    <main>
      <Nuxt />
    </main>
    <footer>
      <div class="foot-wrap">
        <div class="menu">
          <h3>AniChatについて</h3>
          <ul class="foot">
            <li><a href="#">開発者詳細</a></li>
            <li><a href="#">Twitter</a></li>
            <li><a href="#">利用規約</a></li>
          </ul>
        </div>
        <div class="menu">
          <h3>サポート</h3>
          <ul class="foot">
            <li><a href="#">ヘルプセンター</a></li>
            <li><a href="#">コミュニティ</a></li>
          </ul>
        </div>
      </div>
    </footer>
  </v-app>
</template>

<script>
export default {
  name: "index",
  data() {
    return {
      Toggle: false
    }
  },
  methods: {
    toggle: function () {
      this.Toggle = !this.Toggle
    }
  }
}
</script>

<style>
body {
  position: relative;
}
header {
  background: url("~@/static/img/bg-img-1.svg") top center / cover no-repeat;
}
header .header-box {
  padding: 30px 4% 10px;
  position: relative;
  width: 100%;
  top: 0;
  display: flex;
  align-items: center;
}
header h1 {
  margin: 0; padding: 0;
  font-size: 25px;
}
header .n-link {
  text-decoration: none;
  color: #FFFFFF;
  font-weight: bold;
}
header nav {
  margin: 0 10% 0 auto;
}
header ul {
  list-style: none;
  margin: 0;
  display: flex;
  align-items: center;
}
header li {
  margin: 0 2rem 0 2rem;
  font-size: 16px;
}

.sp-nav {
  display: none;
}

.btn-gradient-radius {
  display: inline-block;
  padding: 7px 20px;
  border-radius: 25px;
  text-decoration: none;
  color: #FFF;
  background-image: linear-gradient(#6795fd 0%, #67ceff 100%);
  transition: .4s;
}

.btn-gradient-radius:hover {
  background-image: linear-gradient(#6795fd 0%, #67ceff 70%);
}

@media screen and (max-width: 840px) {
  .pc-nav {
    display: none;
  }
  .sp-nav {
    z-index: 1;
    position: fixed;
    top: 0;
    left: 0;
    width: 100%;
    height: 100vh;
    display: block;
    width: 100%;
    background: rgba(0, 0, 0, .8);
    transition: all .2s ease-in-out;
    opacity: 0;
    transform: translateY(-100%);
  }
  .sp-nav ul {
    padding: 0;
    display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: center;
    height: 100%;
  }
  .sp-nav li {
    font-size: 15px;
    color: #fff;
  }
  .sp-nav li a, .sp-nav li span {
    display: block;
    padding: 20px 0;
  }
  .sp-nav a {
    position: relative;
    display: inline-block;
    text-decoration: none;
  }
  .sp-nav a::after {
    position: absolute;
    bottom: 2px;
    left: 0;
    content: '';
    width: 100%;
    height: 2px;
    background: #fff;
    opacity: 0;
    visibility: hidden;
    transition: .3s;
  }
  a:hover::after {
    bottom: -4px;
    opacity: 1;
    visibility: visible;
  }
  .sp-nav .close {
    position: relative;
    padding-left: 20px;
  }
  .sp-nav .close::before {
    content: '';
    position: absolute;
    top: 50%;
    left: 0;
    display: block;
    width: 16px;
    height: 1px;
    background: #fff;
    transform: rotate( 45deg );
  }
  .sp-nav .close::after {
    content: '';
    position: absolute;
    top: 50%;
    left: 0;
    display: block;
    width: 16px;
    height: 1px;
    background: #fff;
    transform: rotate( -45deg );
  }
  #hamburger {
    position: relative;
    display: block;
    width: 30px;
    height: 25px;
    margin: 0 0 0 auto;
  }
  #hamburger span {
    position: absolute;
    top: 50%;
    left: 0;
    display: block;
    width: 100%;
    height: 2px;
    background-color: #fff;
    transform: translateY(-50%);
  }
  #hamburger::before {
    content: '';
    display: block;
    position: absolute;
    top: 0;
    left: 0;
    width: 100%;
    height: 2px;
    background-color: #fff;
  }
  #hamburger::after {
    content: '';
    display: block;
    position: absolute;
    bottom: 0;
    left: 0;
    width: 70%;
    height: 2px;
    background-color: #fff;
  }
  .toggle {
    transform: translateY(0);
    opacity: 1;
  }
}

/* catch copy */

.main-visual {
  display: flex;
  justify-content: center;
  align-items: center;
  height: 80vh;
  padding:0 4%;
  color: #fff;
}

.main-visual .catch-box {
  width: 100vw;
  display: flex;
  justify-content: center;
  flex-wrap: wrap;
  padding: 56px 24px;
}

.main-visual .catch-sub-box {
  width: 550px;
  height: auto;
}

.main-visual .catch-text {
  padding: 0 10px;
  margin: auto auto;
}

.main-visual .catch-text h1 {
  font-size: 36px;
}
footer {
  width: 100%;
  background-color: #242729;
}
.foot-wrap {
  display: flex;
  flex-direction: row;
  justify-content: center;
}

.foot-wrap .menu {
  margin: 0 20px;
  padding: 20px 0;
}

.foot-wrap ul {
  list-style: none;
  padding: 0;
  line-height: 40px;
}

.foot-wrap h3 {
  color: #5B7DDE;
}

.foot-wrap a {
  color: #FFFFFF;
  text-decoration: none;
}
</style>
