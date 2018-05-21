import Vue from 'vue';
import Vuex from 'vuex';
import moment from 'moment';

Vue.use(Vuex);

export default new Vuex.Store({
  state:{
    start: moment.utc().subtract(24, 'hours'),
    end: moment.utc(),
    range: '24 hours',
    site_code: 'coj'
  },
  getters: {
    start: state => {
      return state.start;
    },
    end: state => {
      return state.end;
    }, // from this: https://vuex.vuejs.org/en/getters.html it says we have to return a function if we want
    // to pass arguments to getters?
    site_status: (state) => (site_code) => {
      /* There should be an object in state that has something like:

        {
        .
        .
        .
          site_code: status
        }

      */
      //return state.site_code[status];

      function getSiteStatus()
      {
        return state.site_code[status];
      }

      return getSiteStatus();
    }
  },
  mutations:{
    setStart(state, time){
      state.start = time;
    },
    setEnd(state, time){
      state.end = time;
    },
    setRange(state, text){
      state.range = text;
    },
    setSiteStatus(state, site_code, status) {
      // Since we can't do: state.site_code.status, you have to use the Vue.set() function to set the property
      // for an object

      Vue.set(state, site_code, status);
    }
  }
});
