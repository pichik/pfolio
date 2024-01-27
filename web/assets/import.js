  function importXTB() {
    // Get the file input element
    const fileInput = document.getElementById('fileInput');
  
    // Check if a file is selected
    if (fileInput.files.length > 0) {
      // Get the first selected file
      const file = fileInput.files[0];
  
      // Create a FormData object
      const formData = new FormData();
  
      // Append the file to the FormData object
      formData.append('file', file);
  
      // Perform the file upload using fetch or XMLHttpRequest
      // Example using fetch:
      fetch('/import-xtb', {
        method: 'POST',
        body: formData,
      })
      .then(response => response.json())
      .then(data => {
        // Handle the server response
        console.log('Server response:', data);

         // Update the chart data
      const chartData = {
        labels: Object.keys(data),
        datasets: [{
          data: Object.values(data),
          backgroundColor: ['#FF6384', '#36A2EB', '#FFCE56'],
        }],
      };

      // Use chart.js API to update the chart with new data
      updateChart(chartData);
      })
      .catch(error => {
        console.error('Error:', error);
      });
    } else {
      console.error('No file selected.');
    }
  }


  // Function to update the chart with new data
function updateChart(newData) {

  myPieChart.data = newData;
  myPieChart.update();

  myBarChart.data = newData;
  myBarChart.update();

  stockTicker.data = newData;
  stockTicker.update();
}

// Function to generate a random color
function getRandomColor() {
  const letters = '0123456789ABCDEF';
  let color = '#';
  for (let i = 0; i < 6; i++) {
    color += letters[Math.floor(Math.random() * 16)];
  }
  return color;
}