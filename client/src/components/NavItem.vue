<template>
<div class="navitem">
  <div class="columns level mini-info">
    <div class="column place is-two-thirds">
      {{site.name}} {{ site.country }}
    </div>
    <div class="column time is-one-fifth">{{ offsetSign }}{{ site.tz }}</div>
    <div class="column mini-weather"><i :class="[isNight ? 'wi-night-clear' : 'wi-day-sunny', 'wi']"></i></div>
  </div>
  <div class="columns">
    <div class="column minimap">
      <img width="400" :src="mapimg" alt="Google Map of site">
    </div>
  </div>
</div>
</template>
<script>
  import suncalc from 'suncalc';
  import moment from 'moment';
  export default {
    name: 'NavItem',
    props: ['site'],
    data: function() {

      return {site_code: this.site.code, site_status: '??'}

    },

    computed: {
      mapimg(){
        return require(`../assets/maps/${this.site.code}.png`); // eslint-disable-line no-undef
      },
      isNight(){
        const position = suncalc.getPosition(moment.utc().valueOf(), this.site.lat, this.site.lng);
        return position.altitude < 0;
      },
      offsetSign(){
        return this.site.tz >= 0 ? '+' : '';
      }
    }
  };
</script>
<style lang="scss">
  .minimap {
    display: none;
  }
  .mini-info {
    margin-right: 0rem;
  }
</style>
