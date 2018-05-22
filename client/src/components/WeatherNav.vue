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
              <option value="24:hours">24 Hours</option>
              <option value="2:days">2 Days</option>
              <option value="7:days">7 Days</option>
              <option value="28:days">28 days</option>
            </select>
          </div>
        </div>

        <div class="columns level-left mini-info">
          <div class="column status is-one-fifth heading">Open?</div>
          <div class="column place is-two-thirds heading">Location</div>
          <div class="column time is-one-fifth heading">UTC</div>
          <div class="column mini-weather heading">Night</div>
        </div>
      </div>
    </li>
    <li v-for="site in sites">
      <router-link :to="`/${site.code}`" active-class="is-active">
        <NavItem :site_status="status(site.code)" :site="site"/>
      </router-link>
    </li>
  </ul>

  <div class="columns is-centered">
    <div class="column is-narrow">
      <button class="button is-success"> {{numOpened}}  open </button>
    </div>
    <div class="column is-narrow">
      <button class="button is-warning"> {{numClosed}} closed </button>
    </div>
  </div>

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
        timeRange: '24:hours',
        site_statuses: {}
      };
    },
    methods:
      {
        fetch_site_status(site_code, start, end, callback)
        { // will return a status letter?
          let request = new XMLHttpRequest();
          //console.log("fetch_site_status called");
          //console.log("site_code: " + site_code);
          let url = 'https://weather-api.lco.global/query?site=' + site_code + '&datumname=' + 'Weather Ok To Open' + '&agg=False' + '&start=' + start.format() + '&end=' + end.format();
          //console.log("making request to URL:" + url);
          request.open('GET', url, false);
          request.onload = () => {
           // console.log("on load event starting");
            if (request.status >= 200 && request.status < 400)
            {
              //console.log("retrieved response from fetch_site_status");
              callback(JSON.parse(request.responseText));

            }
            else
            {
              //console.log("error, couldn't receive response");
            }
          };

          request.onerror = function()
          {
            //console.log("A connection error occurred.");
          };

          request.send();
        },

        status(site_code)
        {
          //console.log("in status() of methods");
          //console.log("site_code: " + site_code);
          //console.log(this.start);
          //console.log(this.$store.getters.end);

          var my_status;
          this.fetch_site_status(site_code, this.start, this.$store.getters.end, (resp) =>
          {
             //console.log("callback beginning");
             //console.log("response: " );
            // console.log(resp);
            if (resp.length < 1)
            {
              return '?';
            }
            else {
              var last_val = resp[resp.length - 1].ValueString;
              //console.log("last value was: " + last_val);
              // TODO: Refactor this into a global filter? We use this logic a lot in different places
              var last_letter = (last_val === 'true' || last_val === "Unknown") ? 'Y' : 'N';
             // console.log("last_letter = " + last_letter);
              this.site_statuses[site_code] = last_letter;
              my_status = last_letter;
              return last_letter; // this line probably isnt needed since returning from a callback here is useless
            }
          });
          //console.log("returning from status() of methods");
         // console.log("my status = " + my_status);
          return my_status;
        }
      }
  ,
    computed: {
      num() {
        return Number(this.timeRange.split(':')[0])
      },
      unit() {
        return this.timeRange.split(':')[1]
      },
      start() {
        return moment.utc().subtract(this.num, this.unit);
      },
      // TODO: Refactor these two methods
      numOpened()
      {
        var opened = 0;
        for (var prop in this.site_statuses)
        {
          if (this.site_statuses[prop] === 'Y')
          {
            opened++;
          }
        }
        return opened;
      },
      numClosed()
      {
        var closed = 0;
        for (var prop in this.site_statuses)
        {
          if (this.site_statuses[prop] === 'N')
          {
            closed++;
          }
        }
        return closed;
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
