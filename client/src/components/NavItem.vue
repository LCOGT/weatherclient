<template>
<div class="navitem">
  <div class="columns level mini-info">
    <div class="column status is-one-fifth"> {{ site.status }} {{getSiteStatus }} </div>
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
  import {EventBus} from "../event-bus";
  export default {
    name: 'NavItem',
    props: ['site', 'sitestatus'],
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
      },
      getSiteStatus()
      {
        /*
        Can't figure out how to retrieve the value from inside the async(?) callback, I can't figure out how to use
        Promises here
         */
        var status_letter;
        var that = this;
        EventBus.$on(this.site.code, function(payLoad) {
          console.log("payload is: " + payLoad); // this prints out the correct value, meaning it was received
          status_letter = payLoad;
          that.site.status = payLoad;
        });
        return this.site.status; /* this returns undefined because the earlier callback hasn't finished running yet? */
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
