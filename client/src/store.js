import Vue from 'vue';
import Vuex from 'vuex';
import moment from 'moment';

Vue.use(Vuex);

export default new Vuex.Store({
  state:{
    start: moment.utc().subtract(24, 'hours'),
    end: moment.utc(),
    range: '24 hours'
  },
  getters: {
    start: state => {
      return state.start;
    },
    end: state => {
      return state.end;
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
    }
  }
});
