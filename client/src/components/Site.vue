<template>
  <div class="site">
    <div class="site-header has-text-centered">
      <h2 class="is-size-2">{{ site.name }} ({{ site.code | uppercase }})</h2>
      <p class="subtitle">
        <span class="heading">Elevation: {{ site.elevation }}m Location: {{ site.lat }}N {{ site.lng }}E</span>
        <span title="sunset"><i class="wi wi-sunset"></i>{{ sunset }}</span>
        <small>UTC</small>
        &nbsp;&nbsp;
        <span title="sunrise"><i class="wi wi-sunrise"></i>{{ sunrise }}</span>
        <small>UTC</small>
      </p>
    </div>
    <div>
      <p class="level-heading heading">Current Values</p>
      <nav class="level">
        <div class="level-item has-text-centered">
          <div>
            <p class="heading">Air Temp &deg;C</i></p>
            <p class="title">{{ datums['Weather Air Temperature Value'] | latestVal }}</p>
          </div>
        </div>
        <div class="level-item has-text-centered">
          <div>
            <p class="heading">Pressure mbar</p>
            <p class="title">{{ datums['Weather Barometric Pressure Value'] | latestVal }}</p>
          </div>
        </div>
        <div class="level-item has-text-centered">
          <div>
            <p class="heading">Humidity %</p>
            <p class="title">{{ datums['Weather Humidity Value'] | latestVal }}</p>
          </div>
        </div>
        <div class="level-item has-text-centered">
          <div>
            <p class="heading">Wind m/s</p>
            <p class="title">{{ datums['Weather Wind Speed Value'] | latestVal }} <small>{{ datums['Weather Wind Direction Value'] | latestVal | cardinal }}</small></p>
          </div>
        </div>
        <div class="level-item has-text-centered">
          <div>
            <p class="heading">Brightness mag/arcsec<sup>^</sup>2</p>
            <p class="title">{{ datums['Weather Sky Brightness Value'] | latestVal }}</p>
          </div>
        </div>
        <div class="level-item has-text-centered">
          <div>
            <p class="heading">Sky Temp &deg;C</p>
            <p class="title">{{ datums['Boltwood Sky Minus Ambient Temperature'] | latestVal }}</p>
          </div>
        </div>
      </nav>
    </div>

    <section class="section section-small">
      <p class="heading">Last {{ this.$store.state.range }}</p>
      <h4 class="is-size-4 helptoggle">OK to Open
        <a class="helptoggle is-pulled-right"><sup><small>?</small></sup></a><span class="help is-pulled-right">If the weather subsystem at site considers the weather conditions sufficient for observing.</span>
      </h4>
      <figure class="image">
        <Timeline datumid="oktoopen" datumname="Weather Ok To Open" :cdata="datums['Weather Ok To Open']"></Timeline>
      </figure>
    </section>

    <section class="section section-small ">
      <h4 class="is-size-4">Air Temperature
        <a class="helptoggle is-pulled-right"><sup><small>?</small></sup></a><span class="help is-pulled-right">Temperature of the air at the weather station.</span>
      </h4>
      <figure class="image">
          <TimeChart datumid="airtemp" datumname="Air Temperature" :cdata="datums['Weather Air Temperature Value']" unit="C"></TimeChart>
      </figure>
    </section>

    <section class="section section-small ">
      <h4 class="is-size-4">Pressure
        <a class="helptoggle is-pulled-right"><sup><small>?</small></sup></a><span class="help is-pulled-right">Barometric (atmospheric) pressure, is a measurement of the force per unit area exerted against a surface by the weight the atmosphere.</span>
      </h4>
      <figure class="image">
          <TimeChart datumid="pressure" datumname="Weather Barometric Pressure Value" :cdata="datums['Weather Barometric Pressure Value']" unit="mbar"></TimeChart>
      </figure>
    </section>

    <section class="section section-small ">
      <h4 class="is-size-4">Humidity
        <a class="helptoggle is-pulled-right"><sup><small>?</small></sup></a><span class="help is-pulled-right">Relative humidity is a percentage representation of the amount of moisture present in the air.</span>
      </h4>
      <figure class="image">
          <TimeChart datumid="humidity" datumname="Weather Humidity Value" :cdata="datums['Weather Humidity Value']" unit="%"></TimeChart>
      </figure>
    </section>

    <section class="section section-small ">
      <h4 class="is-size-4">Wind Speed
        <a class="helptoggle is-pulled-right"><sup><small>?</small></sup></a><span class="help is-pulled-right">The speed of wind in meteres per second.</span>
      </h4>
      <figure class="image">
          <TimeChart datumid="windspeed" datumname="Weather Wind Speed Value" :cdata="datums['Weather Wind Speed Value']" unit="m/s"></TimeChart>
      </figure>
    </section>

    <section class="section section-small ">
      <h4 class="is-size-4">Wind Direction
        <a class="helptoggle is-pulled-right"><sup><small>?</small></sup></a><span class="help is-pulled-right">The direction of the wind, in degrees from North.</span>
      </h4>
      <figure class="image">
          <TimeChart datumid="winddirection" datumname="Weather Wind Direction Value" :cdata="datums['Weather Wind Direction Value']" unit="deg"></TimeChart>
      </figure>
    </section>

    <section class="section section-small ">
      <h4 class="is-size-4">Sky Brightness
        <a class="helptoggle is-pulled-right"><sup><small>?</small></sup></a><span class="help is-pulled-right">As measured by the sky quality meter. Values increase as the sky gets darker.</span>

      </h4>
      <figure class="image">
          <TimeChart datumid="brightness" datumname="Weather Sky Brightness Value" :cdata="datums['Weather Sky Brightness Value']" unit="mag/arcsec^2"></TimeChart>
      </figure>
    </section>

    <section class="section section-small ">
      <h4 class="is-size-4">Sky Temp
        <a class="helptoggle is-pulled-right"><sup><small>?</small></sup></a><span class="help is-pulled-right">Higher temperatures generally indicate the presence of clouds.</span>
      </h4>
      <figure class="image">
          <TimeChart datumid="skytemp" datumname="Boltwood Sky Minus Ambient Temperature" :cdata="datums['Boltwood Sky Minus Ambient Temperature']" unit="C"></TimeChart>
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
        'Weather Air Temperature Value': [],
        'Weather Barometric Pressure Value': [],
        'Weather Humidity Value': [],
        'Weather Wind Speed Value': [],
        'Weather Wind Direction Value': [],
        'Weather Sky Brightness Value': [],
        'Boltwood Sky Minus Ambient Temperature': [],
        'Weather Ok To Open': [],
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
          this.datums[key] = resp;
        });
      });
    },
    fetchDatum(datumName, cb){
      let request = new XMLHttpRequest();
      let url = 'https://weather-api.lco.global/query?site=' + this.site.code + '&datumname=' + datumName;
      if(datumName === 'Weather Ok To Open'){
        url += '&agg=False';
      }
      url += '&start=' + this.start + '&end=' + this.end;
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
      return this.$store.getters.startStr;
    },
    end(){
      return this.$store.getters.endStr;
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
