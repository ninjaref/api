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

/* Ninja Rating Chart */

var ratingDataOverall = {
  labels: ['Speed', 'Consistency', 'Success', 'Rating'],
  datasets: [
    {
      label: 'Average',
      backgroundColor: '#D6E685',
      hoverBackgroundColor: '#D6E685',
      borderColor: '#D6E685',
      data: rDataProfile
    },
    {
      label: 'Profile',
      backgroundColor: '#8CC665',
      hoverBackgroundColor: '#8CC665',
      borderColor: '#8CC665',
      data: rDataProfile
    }
  ]
}

var ratingDataFilter = {
  labels: ['Speed', 'Consistency', 'Success', 'Rating'],
  datasets: [
    {
      label: 'Average',
      backgroundColor: 'rgba(205, 220, 57, 0.85)',
      data: rDataProfile
    },
    {
      label: 'Profile',
      backgroundColor: 'rgba(66, 133, 234, .2)',
      data: rDataProfile
    }
  ]
}

var ratingCtx = document.getElementById('ninja-rating').getContext('2d')
var ratingChart = new Chart(ratingCtx, {
  type: 'bar',
  data: ratingDataOverall,
  options: {}
})
