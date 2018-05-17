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
  props: ['cdata', 'datumid', 'datumname','suntimes', 'sundown', 'sunup', 'timezone', 'fdata'],
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
    failureData() {
      // Same as chartData(), but with weather failure data
      // TODO: Fix function duplication
      // example on cool tricks for the filter function: https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Array/filter
      function retrieve_desired_values(error_object)
      {
        return ((error_object.ValueString !== 'None') && (error_object.ValueString !== 'Unknown'));
      }

      function remove_undesired_values(error_object)
      {
        return ((error_object.ValueString === 'None') || (error_object.ValueString === 'Unknown'));
      }

      const open = this.fdata.filter(retrieve_desired_values).map(
        point => moment.utc(point.TimeStamp, 'YYYY/MM/DD HH:mm:ss')
      );

      const closed = this.fdata.filter(remove_undesired_values).map(
        point => moment.utc(point.TimeStamp, 'YYYY/MM/DD HH:mm:ss')
      );

      let intervals = [];
      for (let i = 0; i < open.length; i++) {
        const start = open[i];
        let end = null;
        for (let j = 0; j < closed.length; j++) {
          if (closed[j].isAfter(start)) {
            end = closed[j];
            break;
          }
        }
        if (!end) {
          end = moment.utc();
        }
        intervals.push([start, end]);
      }

      return intervals;

    },
    sunrise(){
      // return the last sunrise time, but if todays sunrise time has yet to occur,return yesterdays

      // TODO: Fix code reuse here
      let last_sunrise_obj = moment(this.suntimes.slice(-1)[0].sunrise.valueOf()).utc();
      //console.log(last_sunrise_obj);

      if (last_sunrise_obj.isAfter(moment().utc()))
      {
        console.log("sunrise has yet to happen here");
        let next_last_sunrise_obj = moment(this.suntimes.slice(-2)[0].sunrise.valueOf()).utc();
        return next_last_sunrise_obj;
      }

      else
      {
        return last_sunrise_obj;
      }

      return last_sunrise_obj;
    },
    sunset() { // return the last sunrise time, but if todays sunrise time has yet to occur,return yesterdays
      //let last_suntime_obj = this.suntimes.slice(-1)[0];
      let last_sunset_obj = moment(this.suntimes.slice(-1)[0].sunset.valueOf()).utc();

      if (last_sunset_obj.isAfter(moment().utc()))
      {
        console.log("sunset has yet to happen here");
        //return last_sunset_obj.subtract(1, 'day');
        let next_last_sunset_obj = moment(this.suntimes.slice(-2)[0].sunset.valueOf()).utc();
        return next_last_sunset_obj;
      }

      else
      {
        return last_sunset_obj;
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
    { // create a new annotation object for every entry in the suntimes object
      /*
      this.chart.options.annotation.annotations[0].value = this.sunrise;
      this.chart.options.annotation.annotations[1].value = this.sunset;
      */

      let suntimes_annotations = [];
      //console.log("Creating annotations for " + this.suntimes.length + " days");
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
  /*
    console.log("suncalc time");
    console.log(typeof(that.suntimes.sunrise));

    console.log("sunrise time in UTC");
    // the 'X' here tells moment that its parsing a UNIX timestamp, since according to the documentation you're
    // supposed to do that for non-ISO 8601 timestamps
    console.log(moment(that.suntimes.sunrise.valueOf()).utc().format('YYYY/MM/DD HH:mm:ss'));

    console.log('sunrise time with the offset of: ' + that.timezone.valueOf());

    console.log(moment(that.suntimes.sunrise.valueOf()).utcOffset(that.timezone).format('YYYY/MM/DD HH:mm:ss'));

    console.log('sunset time in UTC, with yesterdays day');
    console.log(moment(that.suntimes.sunset.valueOf()).subtract(1, 'days').utc().format('YYYY/MM/DD HH:mm:ss'));

    console.log('sunset time with the offset of: ' + that.timezone.valueOf() + ' (and day -1)');
    console.log(moment(that.suntimes.sunset.valueOf()).subtract(1, 'days').utcOffset(that.timezone).format('YYYY/MM/DD HH:mm:ss'));

*/
    this.chart = new Chart(ctx, {
      type: 'timeline',
      data:{
        labels: [''],
        // from this example: view-source:http://www.chartjs.org/samples/latest/scales/time/line.html it seems like you
        // can just assign a background color to each dataset object
        datasets:[{
          data: that.chartData,
          backgroundColor: '#009ec366', // blue
          label: "open"
        },
          {
            data: that.failureData,
            backgroundColor: '#ff0000',
            label: "closed"
          }]
      },
      options: {
        legend: {
         // display: false ,
          display: true
        },
        annotation: {
          annotations: [{
            type: "line",
            mode: "vertical",
            scaleID: "x-axis-0",
            // value below doesnt work for SB sites, seems roughly 8 hours off?
            //value: moment(that.suntimes.sunrise, 'X').utc().format('YYYY/MM/DD HH:mm:ss'),

            //value: moment(that.suntimes.sunrise.valueOf()).utcOffset(that.timezone).format('YYYY/MM/DD HH:mm:ss'),
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
              // value below doesnt work for SB site, seems roughly 8 hours off?
              //value: moment(that.suntimes.sunset, 'X').subtract(1, 'days').utc().format('YYYY/MM/DD HH:mm:ss'),

              // value below seems to work for SB sites, but is horribly wrong for every other site
              //value: moment(that.suntimes.sunset.valueOf()).subtract(1, 'days').utcOffset(that.timezone).format('YYYY/MM/DD HH:mm:ss'),
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
