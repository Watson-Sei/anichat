<template>
  <div>
    <div class="search-box">
      <div class="search-title">
        <h1>ヘルプセンター</h1>
        <form action="/" name="search1" method="post">
          <dl class="search1">
            <dt><input type="text" name="search" value="" placeholder="Search"></dt>
            <dd><button><span></span></button></dd>
          </dl>
        </form>
      </div>
    </div>
    <div class="card-box">
      <a href="#">
        <div class="card">
          <div class="card-img">
            <img src="../static/img/beginner.svg">
          </div>
          <div class="card-title">
            <h2>使い方</h2>
          </div>
        </div>
      </a>
    </div>
  </div>
</template>

<script>
export default {
  layout: 'default',
  name: "support"
}
</script>

<style scoped>
.search-box {
  width: 100vw;
  height: 30vh;
  background: url("~@/static/img/help-center.png") no-repeat;
  background-size: cover;
  display: flex;
  flex-direction: column;
}

.search-title {
  text-align: center;
  color: #FFFFFF;
  margin: auto;
  width: 500px;
}

form {
  margin: 0 auto;
}

dl.search1{
  position:relative;
  background-color:#fff;
  border:1px solid #aaa;
  -webkit-border-radius:6px;
  -moz-border-radius:6px;
  -o-border-radius:6px;
  -ms-border-radius:6px;
  border-radius:6px;
}
dl.search1 dt{
  margin-right:40px;
  padding:8px 0 8px 8px;
}
dl.search1 dt input{
  width:100%;
  height:24px;
  line-height:24px;
  background:none;
  border:none;
}
dl.search1 dd{
  position:absolute;
  top:0;
  right:0;
}
dl.search1 dd button{
  display:block;
  padding:10px;
  background:none;
  border:none;
}
dl.search1 dd button span{
  display:block;
  width:20px;
  height:20px;
  background:url('~@/static/img/search.svg') no-repeat scroll 0 0;
}

@media screen and (max-width: 830px) {
  .search-title {
    width: 300px;
  }
}

.card-box {
  width: 100%;
  display: flex;
  flex-direction: row;
  justify-content: center;
  padding: 54px 24px;
  height: 50vh;
}

.card-box a {
  text-decoration: none;
  color: black;
}

.card-box .card {
  display: flex;
  flex-direction: row;
  width: 300px;
  height: 100px;
  box-shadow: 2px 2px 4px gray;
}

.card-box .card .card-img {
  width: 100px;
  height: 100px;
  background-color: #C8E2F0;
  box-shadow: 2px 2px 4px gray;
}

.card-box .card .card-img img {
  width: 100px;
  height: 100px;
}

.card-box .card .card-title {
  width: 200px;
  height: 100px;
  text-align: center;
  display: flex;
  justify-content: center;
  align-items: center;
}
</style>
