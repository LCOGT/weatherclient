import Vue from 'vue';
import App from './App.vue';
import router from './router';
import store from './store';
import './assets/main.scss';

new Vue({
  el: '#app',
  router,
  store,
  components: { App },
  template: '<App/>'
});
