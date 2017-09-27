var trendCtx = document.getElementById('trend').getContext('2d')
var trendData = {
  labels: [4, 5, 6, 7, 8, 9],
  datasets: [
    {
      backgroundColor: 'rgba(139,195,74,0.4)',
      data: [1, 3, 2, 2, 3, 1]
    }
  ]
}
var INT_2_TYPE = {
  '1': 'Qualifying',
  '2': 'City Finals',
  '3': 'Stage 1',
  '4': 'Stage 2',
  '5': 'Stage 3',
  '6': 'Stage 4'
}
var trendChart = new Chart(trendCtx, {
  type: 'line',
  data: trendData,
  options: {
    legend: {
      display: false
    },
    scales: {
      xAxes: [{
        scaleLabel: {
          display: true,
          labelString: 'season'
        }
      }],
      yAxes: [{
        ticks: {
          min: 1,
          max: 6,
          callback: function (value, index, values) {
            return INT_2_TYPE[value.toString()]
          }
        }
      }]
    },
    tooltips: {
        callbacks: {
            title: function(item, data) {
              return 'Season ' + item[0].xLabel
            },
            label: function(tooltipItem, data) {
                return INT_2_TYPE[tooltipItem.yLabel.toString()]
            }
        }
    }
  }
})
