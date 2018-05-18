<template>
    <canvas :id="datumid" height="50"></canvas>
</template>
<script>
  import Chart from 'chart.js';
  import moment from 'moment';
  import '../timeline.js';
  // TODO: Use proper method documentation format for JS
// TODO: Remove unneeded props from this component
export default {
  name: 'Timeline',
  props: ['cdata', 'datumid', 'datumname','suntimes', 'timezone', 'fdata'],
  computed: {
    chartData(){
      /*
      const open = this.cdata.filter(item => item.ValueString === 'true').map(
        point => (moment.utc(point.TimeStamp, 'YYYY/MM/DD HH:mm:ss'))
      );
      const closed = this.cdata.filter(item => item.ValueString === 'false').map(
        point => (moment.utc(point.TimeStamp, 'YYYY/MM/DD HH:mm:ss'))
      );
      */

      const open = this.cdata.filter(item => item.ValueString === 'true').map(
        point => ({time: moment.utc(point.TimeStamp, 'YYYY/MM/DD HH:mm:ss'),
                  reason: ''})
      );

      const closed = this.cdata.filter(item => item.ValueString === 'false').map(
        point => ({time: moment.utc(point.TimeStamp, 'YYYY/MM/DD HH:mm:ss'),
                   reason: ''})
      );

      let intervals = [];
      for (let i = 0; i < open.length; i++) {
        const start = open[i];
        let end = null;
        for (let j = 0; j < closed.length; j++) {
          //if(closed[j].isAfter(start))
          if ((closed[j].time).isAfter(start.time))
          {
            end = closed[j];
            break;
          }
        }
        if(!end){
         //end = moment.utc();
          end = ({
            time: moment.utc(),
            reason: ''
          });
        }
        intervals.push([start, end]);
      }
      console.log("intervals from chartData(): ");
      console.log(intervals);
      return intervals;
    },
    failureData()
    { // TODO: Fix code duplication? Merge this function with chartData()

      // TODO: There shouldnt be a need for two functions here since one is the inverse of the other
      function failure_data_desired(error_json)
      {
        return ((error_json.ValueString !== 'Unknown') && (error_json.ValueString !== 'None'));
      }

      function failure_data_undesired(error_json)
      {
          return ((error_json.ValueString === 'Unknown') || (error_json.ValueString === 'None'));
      }

      const open = this.fdata.filter(failure_data_desired).map(
        point => ({time: moment.utc(point.TimeStamp, 'YYYY/MM/DD HH:mm:ss'),
                  reason: point.ValueString })
      );
      const closed = this.fdata.filter(failure_data_undesired).map(
        point => ({time: moment.utc(point.TimeStamp, 'YYYY/MM/DD HH:mm:ss'),
                  reason: point.ValueString })
      );

      let intervals = [];
      for (let i = 0; i < open.length; i++) {
        const start = open[i];
        let end = null;
        for (let j = 0; j < closed.length; j++) {
          //if(closed[j].isAfter(start))
          if ((closed[j].time).isAfter(start.time))
          {
            end = closed[j];
            break;
          }
        }
        if(!end){
          //end = moment.utc();
          end = ({
            time: moment.utc(),
            reason: 'End object'
          });
        }
        intervals.push([start, end]);
      }
      console.log("intervals from failureData(): ");
      console.log(intervals);
      return intervals;

    },
    sunrise(){
      // TODO: Fix code reuse here
      let last_sunrise_obj = moment(this.suntimes.slice(-1)[0].sunrise.valueOf()).utc();

      if (last_sunrise_obj.isAfter(moment().utc()))
      {
        let next_last_sunrise_obj = moment(this.suntimes.slice(-2)[0].sunrise.valueOf()).utc();
        return next_last_sunrise_obj;
      }

      return last_sunrise_obj;
    },
    sunset() { // return the last sunrise time, but if todays sunrise time has yet to occur,return yesterdays
      let last_sunset_obj = moment(this.suntimes.slice(-1)[0].sunset.valueOf()).utc();

      if (last_sunset_obj.isAfter(moment().utc()))
      {
        let next_last_sunset_obj = moment(this.suntimes.slice(-2)[0].sunset.valueOf()).utc();
        return next_last_sunset_obj;
      }

      return last_sunset_obj;
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
    fdata: function(){
      this.chart.data.datasets[1].data = this.failureData;
      this.chart.update();
    },
    '$route' (){
      this.chart.update();
    },
    chartMin(){
      this.chart.options.scales.xAxes[0].time.min = this.chartMin;
      this.chart.options.scales.xAxes[0].time.max = this.chartMax;
    },
    suntimes()
    {
      let suntimes_annotations = [];
      for (let suntime_index = 0; suntime_index < this.suntimes.length; suntime_index++)
      {
        // TODO: Fix code duplication in object creation, maybe use Object.assign()?
        // only show the annotation label if its the '24 hour option' since they take up too much space
        let sunrise_annotation = {
          type: "line",
          mode: "vertical",
          scaleID: "x-axis-0",
          value: (this.suntimes[suntime_index])['sunrise'],
          borderWidth: 5,
          borderColor: "yellow",
          label: {
            enabled: true,
            position: "top",
            content: (this.suntimes.length > 2) ? ("") : ("sunrise")
          }
        };

       let sunset_annotation = {
         type: "line",
         mode: "vertical",
         scaleID: "x-axis-0",
         value: (this.suntimes[suntime_index])['sunset'],
         borderWidth: 5,
         borderColor: "green",
         label: {
           enabled: true,
           position: "top",
           content: (this.suntimes.length > 2) ? (""): ("sunset")
         }
       };

        suntimes_annotations.push(sunrise_annotation, sunset_annotation);
      }

      this.chart.options.annotation.annotations = suntimes_annotations;
    }
  },
  mounted(){
    var that = this;
    const ctx = document.getElementById(this.datumid);
    this.chart = new Chart(ctx, {
      type: 'timeline',
      data:{
        // from this example: view-source:http://www.chartjs.org/samples/latest/scales/time/line.html it seems like you
        // can just assign a background color to each dataset object?
        labels: ['',''],
        datasets:[{
          data: that.chartData,
          backgroundColor: '#009ec366', // blue
          yAxisID: 'y-axis-0',
          label: 'open'
        },
          {
            data: that.failureData,
            backgroundColor: '#A9A9A9', // gray
            yAxisID: 'y-axis-0',
            label: 'closed'
          }]
      },
      options: {
        responsive: true, // resize chart when container element gets resized
        legend: {
          display: true,
        },
        annotation: {
          annotations: [{
            type: "line",
            mode: "vertical",
            scaleID: "x-axis-0",
            value: that.sunrise,
            borderWidth: 5,
            borderColor: "yellow",
            label: {
              content: "sunrise",
              enabled: true,
              position: "top"
            }
          },
            {
              type:'line',
              mode:'vertical',
              scaleID: 'x-axis-0',
              value: that.sunset,
              borderWidth: 5,
              borderColor: 'green',
              label: {
                content: "sunset",
                enabled: true,
                position: "top"
              }
            }]
        },
        scales: {
          yAxes: [{
            beginAtZero: false,
            display: true,
          }
          ],

          xAxes:[{
            display: true,
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
