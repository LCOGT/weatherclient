<template>
    <canvas :id="datumid" height="50"></canvas>
</template>
<script>
import Chart from 'chart.js';
import moment from 'moment';
export default {
  name: 'Timechart',
  props: ['cdata', 'datumid', 'datumname', 'unit'],
  computed: {
    chartData(){
      if(!this.cdata) return [];
      return this.cdata.map(point => ({t: moment.utc(point.TimeStamp, 'YYYY/MM/DD HH:mm:ss'), y: point.Value}));
    },
    chartMin(){
      return this.$store.getters.start;
    },
    chartMax(){
      return this.$store.getters.end;
    }
  },
  watch: {
    cdata(){
      this.chart.data.datasets[0].data = this.chartData;
      this.chart.update();
    },
    '$route' (){
      this.chart.update();
    },
    chartMin(){
      this.chart.options.scales.xAxes[0].time.min = this.chartMin;
      this.chart.options.scales.xAxes[0].time.max = this.chartMax;
    }
  },
  mounted(){
    var that = this;
    const ctx = document.getElementById(this.datumid);
    this.chart = new Chart(ctx, {
      type: 'line',
      data:{
        datasets:[{
          backgroundColor: '#009ec367',
          label: that.datumname,
          data: that.chartData,
        }]
      },
      options: {
        legend: {
          display: false
        },
        scales: {
          xAxes: [{
            time: {
              min: this.chartMin,
              max: this.chartMax
            },
            ticks: {
              callback: function(value, index, values){
                if (!values[index]) return;
                return moment.utc(values[index]['value']).format('HH:mm');
              }
            },
            type: 'time',
          }],
          yAxes: [{
            scaleLabel :{
              display: true,
              labelString: that.unit
            }
          }]
        }
      },
    });
  }
};
</script>
