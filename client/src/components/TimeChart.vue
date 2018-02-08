<template>
    <canvas :id="datumid" height="50"></canvas>
</template>
<script>
import Chart from 'chart.js';
import 'chartjs-plugin-annotation';
import moment from 'moment';
export default {
  name: 'Timechart',
  props: ['cdata', 'datumid', 'datumname', 'unit', 'limit'],
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
    },
    scaleMax(){
      return this.unit === 'deg' ? 360 : undefined;
    }
  },
  watch: {
    cdata (){
      this.chart.data.datasets[0].data = this.chartData;
      this.chart.update();
    },
    '$route' (){
      this.chart.update();
    },
    chartMin(){
      this.chart.options.scales.xAxes[0].time.min = this.chartMin;
      this.chart.options.scales.xAxes[0].time.max = this.chartMax;
    },
    limit(){
      this.chart.options.annotation.annotations[0].value = this.limit;
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
        annotation:{
          annotations:[{
            type: 'line',
            mode: 'horizontal',
            scaleID: 'y-axis-0',
            value: this.limit,
            borderColor: 'red'
          }]
        },
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
            },
            ticks:{
              max: this.scaleMax
            }
          }]
        }
      },
    });
  }
};
</script>
