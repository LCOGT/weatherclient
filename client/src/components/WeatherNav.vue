<template>
<aside class="menu">
  <ul class="menu-list">
    <li>
      <div class="navheader">
        <div class="columns">
          <div class="column">
            <img src="https://cdn.lco.global/mainstyle/img/LCO-logo-web.jpg" alt="Las Cumbres Observatory Weather" class="image logo">
          </div>
          <div class="column align-bottom">
            <h2 class="is-size-3-desktop is-size-4-mobile">Weather</h2>
          </div>
        </div>
        <div class="columns level">
          <div class="column">
            Time Range:
            <select v-model="timeRange">
              <option value="1:hours">1 Hour</option>
              <option value="24:hours">24 Hours</option>
              <option value="2:days">2 Days</option>
              <option value="7:days">7 Days</option>
              <option value="30:days">30 Days</option>
            </select>
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
  import moment from 'moment';
  import {sites} from '../config.js';
  import NavItem from './NavItem';
  export default {
    name: 'WeatherNav',
    components: {NavItem},
    data: function(){
      return {
        sites: sites,
        timeRange: '24:hours'
      };
    },
    computed:{
      num(){
        return Number(this.timeRange.split(':')[0])
      },
      unit(){
        return this.timeRange.split(':')[1]
      },
      start(){
        return moment.utc().subtract(this.num, this.unit);
      }
    },
    watch:{
      timeRange(){
        this.$store.commit('setStart', this.start);
        this.$store.commit('setRange', '' + this.num + ' ' + this.unit);
      }
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
