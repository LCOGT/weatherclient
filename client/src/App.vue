<template>
  <div id="app">
    <section class="section">
      <div class="container is-fluid is-marginless">
        <div class="columns">
          <div class="column is-one-fifth-desktop is-one-half-tablet">
            <div class="fixed">
              <WeatherNav/>
            </div>
          </div>
          <div class="column is-three-fourths-desktop">
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

    width: 95%;
    padding-right: 5px;
    resize: both
    }
  }
</style>
