<template>
<aside class="menu">
  <ul class="menu-list">
    <li>
      <div class="navheader">
        <div class="columns">
          <div class="column">
            <a href="https://lco.global">
              <img src="https://cdn.lco.global/mainstyle/img/LCO-logo-web.jpg" alt="Las Cumbres Observatory Weather" class="image logo">
            </a>
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
              <option value="14:days">14 Days</option>
            </select>
          </div>
        </div>
        <div class="columns level-left mini-info">
          <div class="column status is-one-fifth heading">Open?</div>
          <div class="column place is-two-thirds heading">Location</div>
          <div class="column time is-one-fifth  heading">UTC</div>
          <div class="column mini-weather is-pulled-right heading">Night</div>
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
      <span class="tag is-success is-large"> {{numOpenedandClosed[0]}}  open </span>
    </div>
    <div class="column is-narrow">
      <span class="tag is-warning is-large"> {{numOpenedandClosed[1]}} closed </span>
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
        /**
         * Returns a letter (Y or N or ?) corresponding to the site status.
         * @param site_code : The 3-letter side code.
         * @param start: Query start, determined by start property
         * @param end: query end, determined by end property
         * @param callback: Handles the response received from Weather API
         */
        {
          let request = new XMLHttpRequest();
          let url = 'https://weather-api.lco.global/query?site=' + site_code + '&datumname=' + 'Weather Ok To Open' +
            '&agg=False' + '&start=' + start.format() + '&end=' + end.format();

          request.open('GET', url, false);
          request.onload = () => {
            if (request.status >= 200 && request.status < 400)
            {
              callback(JSON.parse(request.responseText));

            }
            else
            {
              console.log("error, couldn't receive response");
            }
          };

          request.onerror = function()
          {
            console.log("A connection error occurred.");
          };

          request.send();
        },

        status(site_code)
        {
          /**
           * Calls the internal fetch_site_status function.
           */
          var my_status;
          this.fetch_site_status(site_code, this.start, this.$store.getters.end, (resp) =>
          {
            if (!resp || resp.length < 1)
            {
              this.site_statuses[site_code] = '?';
              my_status = '?';
            }
            else {
              var last_val = resp[resp.length - 1].ValueString;
              var last_letter = this.statusMsgToLetter(last_val);
              var last_letter = this.statusMsgToLetter(last_val);

              this.site_statuses[site_code] = last_letter;
              my_status = last_letter;
            }
          });
          return my_status;
        },

        statusMsgToLetter(last_val)
        {
          var letter =  (last_val === 'true' | last_val === 'Unknown') ? 'Y' : 'N';
          return letter;
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
      numOpenedandClosed()
      {
        let opened = 0;
        let closed = 0;
        let unknown = 0;
        for (var property in this.site_statuses)
        {
          if (this.site_statuses[property] === 'Y')
          {
            opened++;
          }
          else if (this.site_statuses[property] === 'N')
          {
            closed++;
          }
          else if (this.site_statuses[property] === '?')
          {
            unknown++;
          }
        }

        return ([opened, closed, unknown]);

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

.menu-list {
  padding-bottom: 10px;
}


</style>
