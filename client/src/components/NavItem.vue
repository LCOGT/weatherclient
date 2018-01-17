<template>
<div class="navitem">
  <div class="columns level mini-info">
    <div class="column place is-two-thirds">{{site.name}} {{ site.country }}</div>
    <div class="column time">{{ site.tz }}</div>
    <div class="column mini-weather"><i class="wi wi-night-clear"></i></div>
  </div>
  <div class="columns">
    <div class="column minimap">
      <img width="400" :src="mapimg" alt="Google Map of site">
    </div>
  </div>
</div>
</template>
<script>
  export default {
    name: 'NavItem',
    props: ['site'],
    computed: {
      mapimg(){
        return `https://maps.googleapis.com/maps/api/staticmap?center=${this.site.lat},${this.site.lng}
                &zoom=2&scale=false&size=400x200&maptype=hybrid&format=png&visual_refresh=true
                &markers=size:mid%7Ccolor:0x009ec3%7Clabel:%7C${this.site.lat},${this.site.lng}`;
      }
    }
  };
</script>
<style lang="scss">
  .minimap {
    display: none;
  }
</style>
