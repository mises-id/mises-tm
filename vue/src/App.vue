<template>
  <div class="home">
    <div class="fixed"><SpNavbar :links="navbarLinks" :active-route="router.currentRoute.value.path" /></div>
    <div class="page">
      <router-view />
    </div>
    <MisesFooter />
  </div>
</template>

<script lang="ts">
import './scss/app.scss'
import { computed, onBeforeMount } from 'vue'
import { useRouter } from 'vue-router'
import { useStore } from 'vuex'
import SpNavbar from './components/SpNav'
import MisesFooter from './components/MisesFooter'
export default {
  components: { SpNavbar,MisesFooter },
  setup() {
    // store
    let $s = useStore()
    // router
    let router = useRouter()
    // state
    let navbarLinks = [
      { name: 'Home', url: '/portfolio' },
      {
        name: 'BlockChain',
        children: [
          {
            name: 'View Holders',
            url: '/holders'
          },
          {
            name: 'View Blocks',
            url: '/blocks'
          },
          {
            name: 'View Txns',
            url: '/tx'
          }
        ]
      },
      {
        name: 'Validators',
        url: '/validators'
      }
    ]
    // computed
    let address = computed(() => $s.getters['common/wallet/address'])
    // lh
    onBeforeMount(async () => {
      await $s.dispatch('common/env/init')
      //router.push('portfolio')
    })
    return {
      navbarLinks,
      // router
      router,
      // computed
      address
    }
  }
}
</script>

<style scoped lang="scss">
body {
  margin: 0;
}
.home{
  display: flex;
  flex-direction: column;
  min-height: 100vh;
  .page{
    flex: 1;
  }
}
.fixed{
  position: fixed;
  top: 0;
  left: 0;
  right: 0;
  z-index: 55;
}
.page{
  padding-top: 80px;
}
@media (max-width: 768px) {
  .page{
    padding-top: 66px;
  }
}
</style>