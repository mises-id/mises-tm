<template>
  <div>
    <SpTheme>
      <div class="fixed"><SpNavbar :links="navbarLinks" :active-route="router.currentRoute.value.path" /></div>
      <div class="page">
        <router-view />
      </div>
    </SpTheme>
  </div>
</template>

<script lang="ts">
import './scss/app.scss'
import { SpNavbar, SpTheme } from '@starport/vue'
import { computed, onBeforeMount } from 'vue'
import { useRouter } from 'vue-router'
import { useStore } from 'vuex'
export default {
  components: { SpTheme, SpNavbar },
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
        url: '/latestblocks',
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
            url: '/txns'
          }
        ]
      },
      {
        name: 'Validators',
        url: '/validators',
        children: [
          {
            name: 'Validators Leaderboard',
            url: '/leaderboard'
          },
          {
            name: 'Validators Set Info',
            url: '/setInfo'
          }
        ]
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
</style>