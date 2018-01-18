<template>
  <div class="site">
    <div class="site-header has-text-centered">
      <h2 class="is-size-2">{{ site.name }} ({{ site.code }})</h2>
      <p class="subtitle"><i class="wi wi-sunrise"></i>{{ sunrise }}&nbsp;&nbsp; <i class="wi wi-sunset"></i>{{ sunset }}</p>
      <nav class="level">
        <div class="level-item has-text-centered">
          <div>
            <p class="heading">Air Temp</p>
            <p class="title">24 <i class="wi wi-celsius"></i></p>
          </div>
        </div>
        <div class="level-item has-text-centered">
          <div>
            <p class="heading">Pressure</p>
            <p class="title">456 mmHg</p>
          </div>
        </div>
        <div class="level-item has-text-centered">
          <div>
            <p class="heading">Humidity</p>
            <p class="title">47%</p>
          </div>
        </div>
        <div class="level-item has-text-centered">
          <div>
            <p class="heading">Wind</p>
            <p class="title">1.3m/s N</p>
          </div>
        </div>
        <div class="level-item has-text-centered">
          <div>
            <p class="heading">Brigtness</p>
            <p class="title">0 W/m2 </p>
          </div>
        </div>
        <div class="level-item has-text-centered">
          <div>
            <p class="heading">Cloud</p>
            <p class="title">Wet</p>
          </div>
        </div>
      </nav>
    </div>
  </div>
</template>
<script>
import suncalc from 'suncalc';
import moment from 'moment';
import {sites} from '../config';
export default {
  name: 'Site',
  props: ['sitecode'],
  data(){
    return {
      site: {}
    };
  },
  watch: {
    '$route' (){
      this.initialize();
    }
  },
  methods:{
    initialize(){
      for (var i = 0; i < sites.length; i++) {
        if(sites[i].code === this.sitecode){
          this.site = sites[i];
          break;
        }
      }
    }
  },
  created(){
    this.initialize();
  },
  computed:{
    suntTimes(){
      return suncalc.getTimes(moment.utc().valueOf(), this.site.lat, this.site.lng);
    },
    sunrise(){
      return moment.utc((this.suntTimes.sunrise)).format('HH:mm');
    },
    sunset(){
     return moment.utc((this.suntTimes.sunset)).format('HH:mm');
    },
  }
};
</script>
