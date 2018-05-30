<template>
    <canvas :id="datumid" height="50"></canvas>
</template>
<script>
import Chart from 'chart.js';
import 'chartjs-plugin-annotation';
import moment from 'moment';
export default {
  name: 'Timechart',
  props: ['cdata', 'datumid', 'datumname', 'unit', 'limit', 'max', 'limit_direction'],
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
      switch(this.unit){
      case 'deg':
        return 360;
      case '%':
        return 100;
      default:
        return this.max;
      }
    },
    scaleMin(){
      switch(this.unit){
        case ('%' || 'm/s' || 'deg' || 'mag/arcsec^2'):
          return 0;
      }
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
          cubicInterpolationMode: 'monotone',
          spanGaps: true
        }]
      },
      options: {
        responsive: true,
        annotation:{
          events:['mouseover', 'mouseout', 'mouseleave', 'mouseenter', 'click'],
          annotations:[{
            label:
              {
                enabled: false,
                //position: 'left',
                position: 'center',
                content: (this.limit_direction === 'min') ? 'Minimum' : 'Maximum',
                yAdjust: (this.limit_direction === 'min') ? -10: 10, // prevent labels from showing up off screen
              },
            type: 'line',
            borderWidth: 5,
            mode: 'horizontal',
            scaleID: 'y-axis-0',
            value: this.limit,
            borderColor: (this.limit_direction === 'min') ? 'green': 'red',
            onClick: function(e) {
              let chart_label = this.chartInstance.chart.options.annotation.annotations[0].label;
              if (chart_label.enabled == true) {
                chart_label.enabled = false;
                that.chart.update();
              }

              else {
                chart_label.enabled = true;
                that.chart.update();
              }
            }
          }]
        },
        legend: {
          display: false,
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
              max: this.scaleMax,
              min: this.scaleMin
            }
          }]
        }
      },
    });
  }
};
</script>
