<template>
    <canvas :id="datumid" height="50"></canvas>
</template>
<script>
import Chart from 'chart.js';
import moment from 'moment';
export default {
  name: 'Timechart',
  props: ['cdata', 'datumid', 'datumname'],
  computed: {
    chartData(){
      return this.cdata.map(point => ({t: moment.utc(point.TimeStamp, 'YYYY/MM/DD HH:mm:ss'), y: point.Value}));
    }
  },
  watch: {
    cdata(){
      this.chart.data.datasets[0].data = this.chartData;
      this.chart.update();
    },
    '$route' (){
      this.chart.update();
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
            ticks: {
              callback: function(value, index, values){
                if (!values[index]) return;
                return moment.utc(values[index]['value']).format('HH:mm');
              }
            },
            type: 'time',
          }]
        }
      },
    });
  }
};
</script>
