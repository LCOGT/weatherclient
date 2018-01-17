<template>
<aside class="menu">
  <p class="menu-label">
    All Sites
  </p>
  <ul class="menu-list">
    <li v-for="site in sites">
      <router-link :to="`/${site.code}`" active-class="is-active">
        <NavItem :site="site"/>
      </router-link>
    </li>
  </ul>
</aside>
</template>
<script>
  import {sites} from '../config.js';
  import NavItem from './NavItem';
  export default {
    name: 'WeatherNav',
    components: {NavItem},
    data: function(){
      return {
        sites: sites
      };
    }
  };
</script>
<style lang="scss">
  .is-active {
    .minimap {
      display: block;
  }
}
</style>
