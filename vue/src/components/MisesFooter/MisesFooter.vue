<template>
  <div class="footer">
    <div class="footer-content">
      <div class="content flex">
        <router-link to="/">
          <img src='/images/index/logo_bottom@2x.png' alt="" class="logo"/>
          <p class="logo-title">Decentralized social media</p>
        </router-link>
        <div class="center-content flex">
          <div>
            <p class="official">Official Website:</p>
            <a href="https://www.mises.site" target="_blank" class="site">https://www.mises.site</a>
          </div>
          <div class="flex download-app">
            <a href="https://d48fg6wpq9lc4.cloudfront.net/playstore/222071211.apk" target="_blank" class="flex btn download">
              <img src="/images/index/download@2x.png" alt="" >
              <span>Download</span>
            </a>
            <a href="https://play.google.com/store/apps/details?id=site.mises.browser" target="_blank"  class="flex btn store">
              <img src="/images/index/GooglePlay@2x.png" alt="" >
              <span>Android</span>
            </a>
            <a href="https://testflight.apple.com/join/Tk1BxD1i" target="_blank"  class="flex btn store">
              <img src="/images/index/iOS@2x.png" alt="" >
              <span>iOS</span>
            </a>
          </div>
        </div>
        <div class="icons">
          <a target="_blank" v-for="(val,index) in linkList" :class="val.className" class="iconStyle" :key="index" :href="val.url">
          </a>
        </div>
      </div>
    </div>
    <div class="footer-copyright">© 2021 Mises technology</div>
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue'
export default defineComponent({
  name: 'MisesFooter',
  setup() {
    const linkList = [{
      url: 'https://twitter.com/Mises001',
      className: 'twitter'
    }, {
      url: 'https://github.com/mises-id',
      className: 'github'
    }, {
      url: 'https://t.me/Misesofficial',
      className: 'telegram'
    }, {
      url: 'https://discord.gg/bwg7hTTYKz',
      className: 'discord'
    }, {
      url: 'https://medium.com/@Mises001',
      className: 'medium'
    }]
    return {
      linkList
    }
  }
})
</script>

<style scoped lang="scss">
.iconStyle{
  width: 34px;
  height: 34px;
  display: inline-block;
  margin-right: 10px;
  margin:5px;
}

$icons: twitter,facebook,instagram,youtube,github,linkedin,reddit,telegram,discord,medium;
@each $icon in $icons {
  .#{$icon}{
    background-image: url('/images/index/#{$icon}@2x.png');
    background-size: 34px 34px;
    &:hover{
      background-image: url('/images/index/#{$icon}_hover@2x.png');
    }
  }
}
.footer-copyright{
  height: 50px;
  background-color: #1F1F25;
  line-height: 50px;
  font-size: 14px;
  font-family: 'hlt-45';
  color: rgba($color: #fff, $alpha: 0.5);
  text-align: center;
}
.footer-content{
  background-color: #111115;
  padding: 20px 0;
  .logo{
    width: 122px;
  }
  .logo-title{
    font-size: 16px;
    font-family: 'hlt-45';
    color: #FFFFFF;
    margin-top: 12px;
  }
}
.icons{
  display: flex;
}
.official{
  font-size: 14px;
  font-family: 'hlt-45';
  color: rgba($color: #fff, $alpha: 0.5);
  margin: 0;
}
.site{
  font-size: 16px;
  font-family: 'hlt-45';
  color: #FFFFFF;
  margin-top: 12px;
}
@media (max-width: 768px) {
  .content{
    flex-direction: column;
    text-align: center;
  }
  .center-content{
    flex:1;
    display: flex;
    flex-direction: column;
    justify-content: center;
  }
  .download-app{
    flex-direction: column;
    width: 90vw;
    .btn{
      width: 100%;
      display: flex;
      height: 44px;
      border: 1px solid #FFFFFF;
      border-radius: 50px;
      justify-content: center;
      margin: 10px 0;
      img{
        display: block;
        width: 20px;
        margin-right: 10px;
      }
      &.download{
        color: white;
      }
      &.store{
        background: #FFFFFF;
        color:#16161D;
      }
    }
  }
}
@media (min-width: 769px) {
  .center-content{
    flex:1;
    display: flex;
    justify-content: center;
  }
  .download-app{
    margin-left: 40px;
    .btn{
      display: flex;
      margin-left: 10px;
      img{
        display: block;
        width: 20px;
        margin-right: 10px;
      }
      &.download{
        width: 130px;
        height: 34px;
        border: 1px solid #FFFFFF;
        border-radius: 17px;
        color: white;
        justify-content: center;
      }
      &.store{
        width: 130px;
        height: 34px;
        background: #FFFFFF;
        border: 1px solid #FFFFFF;
        border-radius: 17px;
        color:#16161D;
        justify-content: center;
      }
    }
  }
}
</style>