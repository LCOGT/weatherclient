<template>
    <canvas :id="datumid" height="50"></canvas>
</template>
<script>
import Chart from 'chart.js';
import moment from 'moment';
import '../timeline.js';
export default {
  name: 'Timeline',
  props: ['cdata', 'datumid', 'datumname','suntimes', 'sundown', 'sunup'],
  computed: {
    chartData(){
      const open = this.cdata.filter(item => item.ValueString === 'true').map(
        point => moment.utc(point.TimeStamp, 'YYYY/MM/DD HH:mm:ss')
      );
      const closed = this.cdata.filter(item => item.ValueString === 'false').map(
        point => moment.utc(point.TimeStamp, 'YYYY/MM/DD HH:mm:ss')
      );
      let intervals = [];
      for (let i = 0; i < open.length; i++) {
        const start = open[i];
        let end = null;
        for (let j = 0; j < closed.length; j++) {
          if(closed[j].isAfter(start)){
            end = closed[j];
            break;
          }
        }
        if(!end){
          end = moment.utc();
        }
        intervals.push([start, end]);
      }
      return intervals;
    },
    chartMin(){
      return this.$store.getters.start;
    },
    chartMax(){
      return this.$store.getters.end;
    }
  },
  watch: {
    cdata: function(){
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
    console.log('sunrise time:');
    //console.log(moment.utc().subtract(1, 'day').format('YYYY/MM/DD ') + this.sunup.valueOf() + ':00');
    //console.log((moment.utc((this.suntimes.sunrise)).format('YYYY/MM/DD HH:mm:ss')).valueOf());

    console.log('sundown time');
    //console.log(moment.utc().subtract(1, 'day').format('YYYY/MM/DD ') + this.sundown.valueOf() + ':00');
    //console.log((moment.utc((this.suntimes.sundown)).format('YYYY/MM/DD HH:mm:ss')).valueOf());
    this.chart = new Chart(ctx, {
      type: 'timeline',
      data:{
        labels: [''],
        datasets:[{
          data: that.chartData,
        }]
      },
      options: {
        legend: {
          display: false
        },
        annotation: {
          annotations: [{
            type: "line",
            mode: "vertical",
            scaleID: "x-axis-0",
           //value: (moment.utc((that.suntimes.sunrise)).format('YYYY/MM/DD HH:mm:ss')).valueOf(),
            //value: '2018/05/12 11:20:00',
            value: moment.utc().subtract(1, 'day').format('YYYY/MM/DD ') + this.sunup.valueOf() + ':00',
            borderWidth: 5,
            borderColor: "yellow",
            label: {
              content: "SUNRISE",
              enabled: true,
              position: "top"
            }
          },
            {
              type:'line',
              mode:'vertical',
              scaleID: 'x-axis-0',
              //value: (moment.utc((that.suntimes.sunset)).format('YYYY/MM/DD HH:mm:ss')).valueOf(),
              value: moment.utc().subtract(1, 'day').format('YYYY/MM/DD ') + this.sundown.valueOf() + ':00',
              borderColor: 'green',
              borderWidth: 5,
              label: {
                content: "SUNSET",
                enabled: true,
                position: "bottom"
              }
            }]
        },
        scales: {
          yAxes: [{
            beginAtZero: false
          }],

          xAxes:[{
            time: {
              min: this.chartMin,
              max: this.chartMax
            },
            ticks: {
              callback: function(value, index, values){
                if (!values[index]) return;
                return moment.utc(values[index]['value']).format('HH:mm');
              }
            }
          }]
        }
      }
    });
  }
};
</script>
