    const apiUrl = '/your/stock/api'; // Replace this with your stock API endpoint

    // Initial chart data
    const initialData = {
      labels: [],
      datasets: [{
        label: 'Stock Price',
        borderColor: 'blue',
        borderWidth: 2,
        fill: false,
        data: [],
      }],
    };



    // Function to update chart data from the server
    const updateStockData = async () => {
      try {
        const response = await fetch(apiUrl);
        const newData = await response.json();

        // Update chart labels and data
        stockTicker.data.labels.push(newData.timestamp);
        stockTicker.data.datasets[0].data.push(newData.price);

        // Limit the number of data points (e.g., last 20 points)
        const maxDataPoints = 20;
        if (stockTicker.data.labels.length > maxDataPoints) {
          stockTicker.data.labels.shift();
          stockTicker.data.datasets[0].data.shift();
        }

        // Update the chart
        stockTicker.update();
      } catch (error) {
        console.error('Error fetching stock data:', error);
      }
    };

    // Update the stock data every second
    // setInterval(updateStockData, 1000);


