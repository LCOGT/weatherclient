<template>
  <div class="site">
    <div class="site-header has-text-centered">
      <h2 class="is-size-2">{{ site.name }} ({{ site.code }})</h2>
      <p class="subtitle"><i class="wi wi-sunrise"></i>{{ sunrise }}&nbsp;&nbsp; <i class="wi wi-sunset"></i>{{ sunset }}</p>
      <nav class="level">
        <div class="level-item has-text-centered">
          <div>
            <p class="heading">Air Temp &deg;C</i></p>
            <p class="title">{{ airTemp | latestVal }}</p>
          </div>
        </div>
        <div class="level-item has-text-centered">
          <div>
            <p class="heading">Pressure mbar</p>
            <p class="title">{{ pressure | latestVal }}</p>
          </div>
        </div>
        <div class="level-item has-text-centered">
          <div>
            <p class="heading">Humidity %</p>
            <p class="title">{{ humidity | latestVal }}</p>
          </div>
        </div>
        <div class="level-item has-text-centered">
          <div>
            <p class="heading">Wind m/s</p>
            <p class="title">{{ windSpeed | latestVal }} <small>{{ windDirection | latestVal | cardinal }}</small></p>
          </div>
        </div>
        <div class="level-item has-text-centered">
          <div>
            <p class="heading">Brigtness mag/arcsec<sup>^</sup>2</p>
            <p class="title">{{ brightness | latestVal }}</p>
          </div>
        </div>
        <div class="level-item has-text-centered">
          <div>
            <p class="heading">Sky Temp &deg;C</p>
            <p class="title">{{ skyTemp | latestVal }}</p>
          </div>
        </div>
      </nav>
    </div>
    <section class="section section-small has-text-centered">
      <h4 class="is-size-4">Air Temperature</h4>
      <figure class="image">
          <TimeChart datumid="airtemp" datumname="Air Temperature" :cdata="airTemp"></TimeChart>
      </figure>
    </section>

    <section class="section section-small has-text-centered">
      <h4 class="is-size-4">Pressure</h4>
      <figure class="image">
          <TimeChart datumid="pressure" datumname="Weather Barometric Pressure Value" :cdata="pressure"></TimeChart>
      </figure>
    </section>

    <section class="section section-small has-text-centered">
      <h4 class="is-size-4">Humidity</h4>
      <figure class="image">
          <TimeChart datumid="humidity" datumname="Weather Humidity Value" :cdata="humidity"></TimeChart>
      </figure>
    </section>

    <section class="section section-small has-text-centered">
      <h4 class="is-size-4">Wind Speed</h4>
      <figure class="image">
          <TimeChart datumid="windspeed" datumname="Weather Wind Speed Value" :cdata="windSpeed"></TimeChart>
      </figure>
    </section>

    <section class="section section-small has-text-centered">
      <h4 class="is-size-4">Wind Direction</h4>
      <figure class="image">
          <TimeChart datumid="winddirection" datumname="Weather Wind Direction Value" :cdata="windDirection"></TimeChart>
      </figure>
    </section>

    <section class="section section-small has-text-centered">
      <h4 class="is-size-4">Brightness</h4>
      <figure class="image">
          <TimeChart datumid="brightness" datumname="Weather Sky Brightness Value" :cdata="brightness"></TimeChart>
      </figure>
    </section>

    <section class="section section-small has-text-centered">
      <h4 class="is-size-4">Sky Temp</h4>
      <figure class="image">
          <TimeChart datumid="skytemp" datumname="Boltwood Sky Minus Ambient Temperature" :cdata="skyTemp"></TimeChart>
      </figure>
    </section>
  </div>
</template>
<script>
import suncalc from 'suncalc';
import moment from 'moment';
import {sites} from '../config';
import TimeChart from './TimeChart';
export default {
  name: 'Site',
  props: ['sitecode'],
  components: {TimeChart},
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
