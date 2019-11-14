import Vue from 'vue';
import Router from 'vue-router';
import Site from './components/Site';

Vue.use(Router);

export default new Router({
  routes: [
    {
      path: '/:sitecode',
      name: 'site',
      component: Site,
      props: true
    },{
      path: '/',
      redirect: '/bpl'
    }
  ]
});

