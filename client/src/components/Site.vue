<template>
  <div class="site">
    <div class="site-header has-text-centered">
      <h2 class="is-size-2">{{ site.name }} ({{ site.code }})</h2>
      <p class="subtitle"><i class="wi wi-sunrise"></i>{{ sunrise }}&nbsp;&nbsp; <i class="wi wi-sunset"></i>{{ sunset }}</p>
      <nav class="level">
        <div class="level-item has-text-centered">
          <div>
            <p class="heading">Air Temp</p>
            <p class="title">{{ airTemp | latestVal }} <sup><i class="wi wi-celsius"></i></sup></p>
          </div>
        </div>
        <div class="level-item has-text-centered">
          <div>
            <p class="heading">Pressure</p>
            <p class="title">{{ pressure | latestVal }} <sup><small>mbar</small></sup></p>
          </div>
        </div>
        <div class="level-item has-text-centered">
          <div>
            <p class="heading">Humidity</p>
            <p class="title">{{ humidity | latestVal }}<sup><small>%</small></sup></p>
          </div>
        </div>
        <div class="level-item has-text-centered">
          <div>
            <p class="heading">Wind</p>
            <p class="title">{{ windSpeed | latestVal }}<sup><small>m/s {{ windDirection | latestVal | cardinal }}</small></sup></p>
          </div>
        </div>
        <div class="level-item has-text-centered">
          <div>
            <p class="heading">Brigtness</p>
            <p class="title">{{ brightness | latestVal }} <sup><small>mag/arcs<sup>^</sup>2</small></sup> </p>
          </div>
        </div>
        <div class="level-item has-text-centered">
          <div>
            <p class="heading">Sky Temp</p>
            <p class="title">{{ skyTemp | latestVal }} <sup><i class="wi wi-celsius"></i></sup></p>
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
      site: {},
      airTemp: [],
      pressure: [],
      humidity: [],
      windSpeed: [],
      windDirection: [],
      brightness: [],
      skyTemp: []
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
      this.fetchDatum('Weather Air Temperature Value', (resp) => {
        this.airTemp = resp;
      });
      this.fetchDatum('Weather Barometric Pressure Value', (resp) => {
        this.pressure = resp;
      });
      this.fetchDatum('Weather Humidity Value', (resp) => {
        this.humidity = resp;
      });
      this.fetchDatum('Weather Wind Speed Value', (resp) => {
        this.windSpeed = resp;
      });
      this.fetchDatum('Weather Wind Direction Value', (resp) => {
        this.windDirection = resp;
      });
      this.fetchDatum('Weather Sky Brightness Value', (resp) => {
        this.brightness = resp;
      });
      this.fetchDatum('Boltwood Sky Minus Ambient Temperature', (resp) => {
        this.skyTemp = resp;
      });
    },
    fetchDatum(datumName, cb){
      let request = new XMLHttpRequest();
      request.open('GET', 'http://127.0.0.1:8080/query?site=' + this.site.code + '&datumname=' + datumName, true);
      request.onload = () => {
        if (request.status >=200 && request.status < 400) {
          cb(JSON.parse(request.responseText));
        } else {
          console.log('error:' + request.responseText);
        }
      };

      request.onerror = function() {
        console.log('There was a connection error');
      };

      request.send();
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
  },
  filters: {
    latestVal(values){
      if (values.length < 1) return 0;
      let val = values[values.length - 1].Value;
      return val.toFixed(2);
    },
    cardinal(val){
      const num = Math.floor((val / 22.5) + 0.5);
      const compass = ['N', 'NNE', 'NE', 'ENE', 'E', 'ESE', 'SE', 'SSE', 'S', 'SSW', 'SW', 'WSW', 'W', 'WNW', 'NW', 'NNW'];
      return compass[num % 16];
    }
  }
};
</script>
