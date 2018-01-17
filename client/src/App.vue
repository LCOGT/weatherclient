<template>
  <div id="app">
    <section class="section">
      <div class="container">
        <div class="columns">
          <div class="column is-one-quarter">
            <WeatherNav/>
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
