<template>
  <div class="site">
    <div class="site-header has-text-centered">
      <h2 class="is-size-2">{{ site.name }} ({{ site.code | uppercase }})</h2>
      <h4 class="is-size-4" v-if="site.code === 'tfn'">Weather data from SONG site.</h4>
      <p class="subtitle">
        <span class="heading">
          Elevation: {{ site.elevation }}m
          Location: {{ site.lat | ns}} {{ site.lng | ew }}
        </span>
        <span title="sunset">sunset: {{ sunset }}</span>
        <small>UTC</small>
        &nbsp;&nbsp;
        <span title="sunrise">sunrise: {{ sunrise }}</span>
        <small>UTC</small>
      </p>
    </div>
    <div>
      <p class="level-heading heading">Current Values</p>
      <nav class="level">
        <div class="level-item has-text-centered">
          <div>
            <p class="heading">Air Temp &deg;C</i></p>
            <p class="title">{{ datums['Weather Air Temperature Value'].data | latestVal }}</p>
          </div>
        </div>
        <div class="level-item has-text-centered">
          <div>
            <p class="heading">Pressure mbar</p>
            <p class="title">{{ datums['Weather Barometric Pressure Value'].data | latestVal }}</p>
          </div>
        </div>
        <div class="level-item has-text-centered">
          <div>
            <p class="heading">Humidity %</p>
            <p class="title">{{ datums['Weather Humidity Value'].data | latestVal }}</p>
          </div>
        </div>
        <div class="level-item has-text-centered">
          <div>
            <p class="heading">Wind meters per sec</p>
            <p class="title">{{ datums['Weather Wind Speed Value'].data | latestVal }}</p>
          </div>
        </div>
        <div class="level-item has-text-centered">
          <div>
            <p class="heading">Wind &deg;E of N</p>
            <p class="title">{{ datums['Weather Wind Direction Value'].data | latestVal }}&deg;</p>
          </div>
        </div>
        <div class="level-item has-text-centered">
          <div>
            <p class="heading">Brightness mag/arcsec<sup>^</sup>2</p>
            <p class="title">{{ datums['Weather Sky Brightness Value'].data | latestVal }}</p>
          </div>
        </div>
        <div class="level-item has-text-centered">
          <div>
            <p class="heading">Sky Temp &deg;C</p>
            <p class="title">{{ datums['Boltwood Sky Minus Ambient Temperature'].data | latestVal }}</p>
          </div>
        </div>
      </nav>
    </div>

    <section class="section section-xsmall">
      <p class="heading">Last {{ this.$store.state.range }}</p>
      <h4 class="is-size-4 helptoggle">OK to Open
        <a class="helptoggle is-pulled-right"><sup><small>?</small></sup></a><span class="help is-pulled-right">All weather conditions are within acceptable range to allow observing.</span>
      </h4>
      <figure class="image">
        <Timeline datumid="oktoopen" datumname="Weather Ok To Open" :cdata="datums['Weather Ok To Open'].data"></Timeline>
      </figure>
    </section>

    <section class="section section-xsmall ">
      <h4 class="is-size-4">Air Temperature
        <a class="helptoggle is-pulled-right"><sup><small>?</small></sup></a><span class="help is-pulled-right"> Ambient temperature measured by HMP45C-L temperature probe at the site's weather station.</span>
      </h4>
      <figure class="image">
          <TimeChart datumid="airtemp" datumname="Air Temperature" unit="C"
                     :cdata="datums['Weather Air Temperature Value'].data"
                     :limit="limit('Weather Air Temperature Value')">
           </TimeChart>
      </figure>
    </section>

    <section class="section section-xsmall ">
      <h4 class="is-size-4">Sky - Ambient Temp
        <a class="helptoggle is-pulled-right"><sup><small>?</small></sup></a><span class="help is-pulled-right">Sky Temperature is inferred from 8-14µm irradiance measure by a Boltwood II cloud sensor at the site's weather station.</span>
      </h4>
      <figure class="image">
          <TimeChart datumid="skytemp" datumname="Boltwood Sky Minus Ambient Temperature" unit="C"
                    :cdata="datums['Boltwood Sky Minus Ambient Temperature'].data"
                    :limit="limit('Boltwood Sky Minus Ambient Temperature')">
          </TimeChart>
      </figure>
    </section>

    <section class="section section-xsmall ">
      <h4 class="is-size-4">Humidity
        <a class="helptoggle is-pulled-right"><sup><small>?</small></sup></a><span class="help is-pulled-right">Relative humidity measured by HMP45C-L humidity probe at the site's weather station.</span>
      </h4>
      <figure class="image">
          <TimeChart datumid="humidity" datumname="Weather Humidity Value" unit="%"
                     :cdata="datums['Weather Humidity Value'].data"
                     :limit="limit('Weather Humidity Value')">
          </TimeChart>
      </figure>
    </section>

    <section class="section section-xsmall ">
      <h4 class="is-size-4">Pressure
        <a class="helptoggle is-pulled-right"><sup><small>?</small></sup></a><span class="help is-pulled-right">Barometric pressure measured by Vaisala PTB110 barometer at the site's weather station.</span>
      </h4>
      <figure class="image">
          <TimeChart datumid="pressure" datumname="Weather Barometric Pressure Value" unit="mbar"
                     :cdata="datums['Weather Barometric Pressure Value'].data"
                     :limit="limit('Weather Barometric Pressure Value')">
          </TimeChart>
      </figure>
    </section>

    <section class="section section-xsmall ">
      <h4 class="is-size-4">Wind Speed
        <a class="helptoggle is-pulled-right"><sup><small>?</small></sup></a><span class="help is-pulled-right">Wind speed measured by Windsonic1-L wind sensor at the site's weather station.</span>
      </h4>
      <figure class="image">
          <TimeChart datumid="windspeed" datumname="Weather Wind Speed Value" unit="m/s"
                     :cdata="datums['Weather Wind Speed Value'].data"
                     :limit="limit('Weather Wind Speed Value')">
          </TimeChart>
      </figure>
    </section>

    <section class="section section-xsmall ">
      <h4 class="is-size-4">Wind Direction
        <a class="helptoggle is-pulled-right"><sup><small>?</small></sup></a><span class="help is-pulled-right">Wind direction, in degrees East of North, measured by Windsonic1-L wind sensor at the site's weather station.</span>
      </h4>
      <figure class="image">
          <TimeChart datumid="winddirection" datumname="Weather Wind Direction Value" unit="deg"
                     :cdata="datums['Weather Wind Direction Value'].data"
                     :limit="limit('Weather Wind Direction Value')">
          </TimeChart>
      </figure>
    </section>

    <section class="section section-xsmall ">
      <h4 class="is-size-4">Sky Brightness
        <a class="helptoggle is-pulled-right"><sup><small>?</small></sup></a><span class="help is-pulled-right">Sky brightness, in magnitudes per square arcsecond, measured by Unihedron SQM-LE at the site's weather station.</span>

      </h4>
      <figure class="image">
          <TimeChart datumid="brightness" datumname="Weather Sky Brightness Value" unit="mag/arcsec^2"
                     :cdata="datums['Weather Sky Brightness Value'].data"
                     :limit="limit('Weather Sky Brightness Value')">
          </TimeChart>
      </figure>
    </section>

     <section class="section section-xsmall ">
      <h4 class="is-size-4">Transparency
        <a class="helptoggle is-pulled-right"><sup><small>?</small></sup></a><span class="help is-pulled-right">τ =[Offset-(Tsky - Tamb)]*100/Scale , where the scale and offset are empirically determined for each Boltwood sensor.</span>
      </h4>
      <figure class="image">
          <TimeChart datumid="transparency" datumname="Boltwood Transparency Average" unit="%"
                     :cdata="datums['Boltwood Transparency Average'].data"
                     :limit="limit('Boltwood Transparency Average')">
          </TimeChart>
      </figure>
    </section>
  </div>
</template>
<script>
import suncalc from 'suncalc';
import moment from 'moment';
import {sites} from '../config';
import TimeChart from './TimeChart';
import Timeline from './Timeline';
export default {
  name: 'Site',
  props: ['sitecode'],
  components: {TimeChart, Timeline},
  data(){
    return {
      site: {},
      datums: {
        'Weather Air Temperature Value': {
          data: [],
          limit: {
            default: -20,
            coj: -4,
            ogg: 0
          }
        },
        'Weather Barometric Pressure Value': {
          data: [],
          limit: {
            default: null
          }
        },
        'Weather Humidity Value': {
          data: [],
          limit: {
            default: 92.5,
            coj: 85,
            tfn: 75,
            ogg: 85
          }
        },
        'Weather Wind Speed Value': {
          data: [],
          limit: {
            default: 15,
            ogg: 17
          }
        },
        'Weather Wind Direction Value': {
          data: [],
          limit: {
            default: null
          }
        },
        'Weather Sky Brightness Value': {
          data: [],
          limit: {
            default: 23
          }
        },
        'Boltwood Sky Minus Ambient Temperature': {
          data: [],
          limit: {
            default: null
          }
        },
        'Boltwood Transparency Average': {
          data: [],
          limit: {
            default: 25
          }
        },
        'Weather Ok To Open': {
          data: [],
          limit: {
            default: null
          }
        }
      }
    };
  },
  watch: {
    '$route'(){
      this.initialize();
    },
    start(){
      this.fetchDatums();
    },
    end(){
      this.fetchDatums();
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
      this.fetchDatums();
    },
    fetchDatums(){
      Object.keys(this.datums).forEach((key) => {
        this.fetchDatum(key, (resp) => {
          this.datums[key]['data'] = resp;
        });
      });
    },
    fetchDatum(datumName, cb){
      let request = new XMLHttpRequest();
      let url = 'https://weather-api.lco.global/query?site=' + this.site.code + '&datumname=' + datumName;
      if(datumName === 'Weather Ok To Open'){
        url += '&agg=False';
      }
      url += '&start=' + this.start.format() + '&end=' + this.end.format();
      request.open('GET', url, true);
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
    },
    limit(datumName){
      if(this.datums[datumName].limit.hasOwnProperty(this.site.code)){
        return this.datums[datumName].limit[this.site.code];
      }else{
        return this.datums[datumName].limit.default;
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
    start(){
      // Make sure we get a little data leading up to the time range to avoid gaps
      let start = this.$store.getters.start.clone();
      start.subtract(moment.duration(3, 'hours'));
      return start;
    },
    end(){
      return this.$store.getters.end;
    }
  },
  filters: {
    latestVal(values){
      if (!values || values.length < 1) return 0;
      let val = values[values.length - 1].Value;
      return val.toFixed(1);
    },
    cardinal(val){
      const num = Math.floor((val / 22.5) + 0.5);
      const compass = ['N', 'NNE', 'NE', 'ENE', 'E', 'ESE', 'SE', 'SSE', 'S', 'SSW', 'SW', 'WSW', 'W', 'WNW', 'NW', 'NNW'];
      return compass[num % 16];
    },
    ns(val){
      if(val > 0){
        return val + 'N'
      } else {
        return Math.abs(val) + 'S'
      }
    },
    ew(val){
      if(val > 0){
        return val + 'E'
      }
      else{
        return Math.abs(val) + 'W'
      }
    },
    uppercase(str){
      return str.toUpperCase();
    }
  }
};
</script>
<style lang="scss">
  .help {
    display: none;
    margin-right: 20px;
  }
  .helptoggle:hover + .help {
    display: inline;
  }
  .level-heading {
    padding: 0.5rem 1.5rem;
  }
  .subtitle {
    small {
      font-size: 0.8rem;
    }
  }
</style>
