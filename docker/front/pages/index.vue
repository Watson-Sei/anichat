<template>
  <div>
    <div class="merit" id="merit">
      <div class="box">
        <div class="box-text">
          <h2>あなたのアニメライフを共有</h2>
          <p>AniChatは、あなたのご視聴になるアニメの感想をリアルタイムで共有する空間を提供します。</p>
        </div>
        <div class="box-img">
          <img src="../static/img/online-share.svg" alt="">
        </div>
      </div>
    </div>
    <div class="merit" style="background-color: #F5F5F5">
      <div class="box" style="flex-direction: row;">
        <div class="box-text">
          <h2>放送時間だけのチャット</h2>
          <p>チャットルームは、アニメの放送時間だけしかご利用いただけませんその時間だけをお楽しみください。</p>
        </div>
        <div class="box-img">
          <img src="../static/img/open-chat.svg" alt="">
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  layout: 'default',
}
</script>

<style>

.merit {
  width: 100%;
  padding: 54px 24px;
}

.merit .box {
  width: 100%;
  height: auto;
  display: flex;
  flex-direction: row-reverse;
  flex-wrap: wrap;
  justify-content: center;
  align-items: center;
}

.merit .box .box-text {
  width: 500px;
  height: auto;
  text-align: center;
  padding: 0 20px;
}

.merit h2 {
  font-size: 25px;
  padding: 1rem 3rem;
  -webkit-transform: skew(-15deg);
  transform: skew(-15deg);
  color: #fff;
  background-image: -webkit-gradient(linear, left top, right top, from(#209cff), to(#68e0cf));
  background-image: -webkit-linear-gradient(left, #209cff 0%, #68e0cf 80%);
  background-image: linear-gradient(to right, #209cff 0%, #68e0cf 80%);
}

.merit p {
  margin-top: 20px;
  font-weight: bold;
}

.merit .box .box-img {
  width: 400px;
  height: auto;
}
</style>
