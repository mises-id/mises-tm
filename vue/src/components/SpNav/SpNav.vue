<template>
  <div>
    <Suspense><SpSystemBar /></Suspense>
    <div class="navbar-wrapper">
      <div class="navbar-section content">
        <slot name="logo">
          <router-link :to="'/'" class="hide-on-small logo-link" :alt="'Home'" :title="'Home'">
            <img src="/images/index/logo@2x.png" alt="logo" srcset="" class="logo" />
          </router-link>
        </slot>
        <div class="right-nav-link">
          <div v-for="(link, lid) in links" :key="`link-${lid}`" class="nav-link">
            <router-link v-if="!link.children" :to="link.url" class="sp-nav-link" :alt="link.name" :title="link.name">
              <div :class="link.url === activeRouteName ? 'link-active' : ''">
                {{ link.name }}
              </div>
            </router-link>

            <div v-if="link.children" class="link-list-container">
              <div class="flex">
                <span class="sp-nav-link" :class="link.children.some((val) => val.url === activeRouteName) ? 'link-active' : ''">
                  {{ link.name }}
                </span>
                <img src="/images/index/down_black@2x.png" alt="" class="downblock" />
              </div>
              <div class="link-list">
                <router-link v-for="(item, id) in link.children" :key="`link-item-${id}`" :to="item.url" :alt="item.name" :title="item.name">
                  <div :class="item.url === activeRouteName ? 'link-active' : ''">
                    {{ item.name }}
                  </div>
                </router-link>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent, PropType } from 'vue'
import { SpSystemBar } from '@starport/vue'

export interface NavbarLink {
  name: string
  url: string
  children?: NavbarLink[]
}

export default defineComponent({
  name: 'SpNavbar',

  components: {
    SpSystemBar
  },

  props: {
    links: {
      type: Object as PropType<NavbarLink[]>,
      required: true
    },
    activeRoute: {
      type: String,
      required: false
    }
  },
  computed:{
    activeRouteName(){
      return this.activeRoute==='/' ? '/portfolio' : this.activeRoute
    }
  },
})
</script>

<style scoped lang="scss">
$height: 80px;
.navbar-wrapper {
  background: #ffffff;
}
.logo {
  width: 130px;
}
.logo-link {
  display: flex;
  align-items: center;
}
.navbar-section {
  display: flex;
  justify-content: space-between;
  height: $height;
  align-content: center;
  .right-nav-link {
    display: flex;
    align-items: center;
  }
}

.sp-nav-link {
  font-size: 16px;
  color: #16161d;
  font-weight: 400;
  text-decoration: none;
  cursor: pointer;
  transition: font-weight 0.2s ease, color 0.2s ease;
  display: block;
}

.sp-nav-link:hover {
  color: #5d61ff;
}

.sp-nav-link.selected {
  font-weight: 600;
  color: #000000;
}

.description-grey {
  font-size: 13px;
  line-height: 153.8%;
  color: rgba(0, 0, 0, 0.667);
}

.external-link {
  font-weight: 600;
  font-size: 16px;
  cursor: pointer;
}

.external-link:hover {
  opacity: 0.8;
}

.link-active {
  color: #5d61ff;
}
.link-list {
  min-width: 100px;
  border-radius: 5px;
  background: white;
  border: 1px solid #eeeeee;
  box-shadow: 0px 5px 10px 0px rgba(51, 51, 51, 0.05);
  border-radius: 10px;
  position: absolute;
  z-index: 55;
  padding: 15px;
  font-size: 16px;
  top: 60px;
  display: none;
  a {
    white-space: nowrap;
    color: #16161d;
    text-decoration: none;
    line-height: 40px;
    &:hover {
      color: #5d61ff;
    }
  }
}
.link-list-container {
  position: relative;
  height: $height;
  display: flex;
  align-items: center;
}
.nav-link {
  &:not(:last-child) {
    margin-right: 50px;
  }
  cursor: pointer;
  &:hover {
    .sp-nav-link {
      color: #5d61ff;
    }
    .link-list {
      display: block;
    }
    .downblock {
      transform: rotate(180deg);
    }
  }
}
.flex {
  display: flex;
  align-items: center;
}
.downblock {
  width: 11px;
  height: 6px;
  margin-left: 5px;
  transition: all 0.3s;
}
@media (max-width: 600px) {
  .hide-on-small {
    display: none;
  }
  $height: 66px;
  .navbar-section{
    height: $height;
    justify-content: center;
  }
  .link-list-container{
    height: $height;
  }
}
</style>
