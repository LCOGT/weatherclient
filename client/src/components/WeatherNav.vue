<template>
<aside class="menu">
  <ul class="menu-list">
    <li>
      <div class="navheader">
        <div class="columns">
          <div class="column">
            <img src="https://cdn.lco.global/mainstyle/img/LCO-logo-web.jpg" alt="Las Cumbres Observatory Weater" class="image logo">
          </div>
          <div class="column align-bottom">
            <h2 class="is-size-3-desktop is-size-4-mobile">Weather</h2>
          </div>
        </div>
        <div class="columns level mini-info">
          <div class="column place is-two-thirds heading">Location</div>
          <div class="column time is-one-fifth heading">UTC</div>
          <div class="column mini-weather heading">Night</div>
        </div>
      </div>
    </li>
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
.navheader {
  padding: 0.5em 0.75em;
}
.align-bottom {
  display: flex;
  align-items: center;
}

.logo {
  max-width: 9rem;
}
</style>
