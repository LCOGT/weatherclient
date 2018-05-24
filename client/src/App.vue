<template>
  <div id="app">
    <section class="section">
      <div class="container is-fluid">
        <div class="columns">
          <div class="column is-one-quarter">
            <div class="fixed">
              <WeatherNav/>
            </div>
          </div>
          <div class="column">
            <router-view/>
          </div>
        </div>
      </div>
    </section>
  </div>
</template>
<script>
  import WeatherNav from './components/WeatherNav';
  export default {
    name: 'app',
    components: {WeatherNav}
  };
</script>
<style lang="scss">
  @media screen and (min-width: 1224px){
  .fixed {
      position: fixed;
      width: 19rem;
    }
  }
</style>
