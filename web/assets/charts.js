// Function to create and initialize a chart for a ticker
function createLineChart(canvasId, data) {
    const ctx = document.getElementById(canvasId).getContext('2d');

    return new Chart(ctx, {
      type: 'line',
      data: {
        labels: createEmptyStringList(data.length),
        datasets: [{
          label: 'Stock Price',
          data: data, 
          borderColor: 'red',
          tension: 0.1,
          pointRadius:10,
          pointBackgroundColor:'rgba(0,0,0,0)',
          pointBorderColor:'rgba(0,0,0,0)',
          fill: true,
          backgroundColor:'rgba(255,0,0,0.2)',
          pointHoverBackgroundColor:'green',
          pointHoverRadius:3,
          borderWidth:1
        }]
      },
      options: {
        plugins:{
          legend: {
            display: false
          },
          tooltip: {
            boxWidth: 0,
            mode: 'nearest', // Set the tooltip mode to 'nearest'
            position: 'nearest', // Set the tooltip position to 'nearest'
            intersect: false, // Disable tooltip intersection with items
            yAlign: 'bottom',
            callbacks: {
              label: function(context) {
                  label = '$' + context.raw.y + ' '; // Add price to label
                  label += '[ '+moment(context.raw.x).format('DD.MM HH:mm')+' ]';
                return label;
              }
            }
          }
        },
        responsive: false, // Prevent chart from resizing
      }
    });
  }


  function createEmptyStringList(length) {
    return Array.from({ length: length }, () => '');
}


function createPieChart(canvasId, data){
  const ctx = document.getElementById(canvasId).getContext('2d');
  // Create the pie chart
  const myPieChart = new Chart(ctx, {
      type: 'pie',
      data: {
        labels: createEmptyStringList(data.length),
        datasets: [{
          label: '',
          data: data, 
          pointRadius:10,
          backgroundColor: [
            'rgba(205, 86, 86, 0.5)',   // Red
            'rgba(86, 205, 86, 0.5)',   // Green
            'rgba(86, 86, 205, 0.5)',   // Blue
            'rgba(205, 205, 86, 0.5)',  // Yellow
            'rgba(205, 86, 205, 0.5)',  // Purple
            'rgba(86, 205, 205, 0.5)',  // Cyan
            'rgba(205, 154, 86, 0.5)',  // Orange
            'rgba(154, 86, 205, 0.5)',  // Violet
            'rgba(86, 154, 205, 0.5)',  // Sky Blue
            'rgba(154, 205, 86, 0.5)'   // Lime Green
        ]
        }]
      },
      options: {
        plugins:{
          tooltip:{
            displayColors:false,
            callbacks: {
              label: function(context) {
                  let label = context.raw || '';
                  if (context.parsed) {
                      label += `%`;
                  }
                  return label;
              }
            }
          },
          legend: {
            display: false
          }
        },
      responsive: false // Prevent chart from resizing
      },
  });
}


function createBarChart(canvasId, data){
  const ctx = document.getElementById(canvasId).getContext('2d');

  // Create the bar chart
  const myBarChart = new Chart(ctx, {
    type: 'bar',
    data: {
      labels: createEmptyStringList(data.length),
      datasets: [{
        label: 'Invested',
        backgroundColor: 'rgba(0,150,150,0.5)',
        data:data
      },
      {
        label: 'Profit',
        backgroundColor: 'rgba(0,255,0,0.2)',
        data:data
      }]
    },
    options: {
      responsive: false, // Prevent chart from resizing
      plugins: {
        tooltip:{
          displayColors:false
        },
        title: {
          display: true,
          text: 'Invested / Profit'
        },
      },
      scales: {
        x: {
          stacked: true,
        },
        y: {
          stacked: true
        }
      }
    },
  });
}


