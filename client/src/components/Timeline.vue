<template>
    <canvas :id="datumid" height="50"></canvas>
</template>
<script>
import Chart from 'chart.js';
import moment from 'moment';
import '../timeline.js';
export default {
  name: 'Timeline',
  props: ['cdata', 'datumid', 'datumname','suntimes', 'sundown', 'sunup', 'timezone'],
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

    console.log("sunrise time in UTC");
    // the 'X' here tells moment that its parsing a UNIX timestamp, since according to the documentation you're
    // supposed to do that for non-ISO 8601 timestamps
    console.log(moment(that.suntimes.sunrise, 'X').utc().format('YYYY/MM/DD HH:mm:ss'));

    console.log('sunrise time with the offset of: ' + that.timezone.valueOf());

    console.log(moment(that.suntimes.sunrise, 'X').utcOffset(that.timezone).format('YYYY/MM/DD HH:mm:ss'));

    console.log('sunset time in UTC, with yesterdays day');
    console.log(moment(that.suntimes.sunset, 'X').subtract(1, 'days').utc().format('YYYY/MM/DD HH:mm:ss'));

    console.log('sunset time with the offset of: ' + that.timezone.valueOf() + ' (and day -1)');
    console.log(moment(that.suntimes.sunset, 'X').subtract(1, 'days').utcOffset(that.timezone).format('YYYY/MM/DD HH:mm:ss'));

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
            // value below doesnt work for SB sites, seems roughly 8 hours off?
            value: moment(that.suntimes.sunrise, 'X').utc().format('YYYY/MM/DD HH:mm:ss'),

            //value: moment(that.suntimes.sunrise, 'X').utcOffset(that.timezone).format('YYYY/MM/DD HH:mm:ss'),

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
              // value below doesnt work for SB site, seems roughly 8 hours off?
              value: moment(that.suntimes.sunset, 'X').subtract(1, 'days').utc().format('YYYY/MM/DD HH:mm:ss'),

              // value below seems to work for SB sites, but is horribly wrong for every other site
              //value: moment(that.suntimes.sunset, 'X').subtract(1, 'days').utcOffset(that.timezone).format('YYYY/MM/DD HH:mm:ss'),
              borderWidth: 5,
              borderColor: 'yellow',
              label: {
                content: "SUNSET",
                enabled: true,
                position: "top"
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
