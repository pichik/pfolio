
    function UpdateWatchlist(stocks){
    stocks.forEach(stock => {
        const ticker=stock.Ticker;
        const lastPrice=parseFloat(stock.LastPrice.toFixed(2));

        const sectionId=`watchlist-${ticker}`

        

        //Update Chart
        if (stock.HourlyPrices != null && Object.entries(stock.HourlyPrices).length > 0){
            priceProgress = Object.entries(stock.HourlyPrices)
            .map(([timestamp, price]) => ({ x: parseInt(timestamp), y: price }))
            //Calculate graph difference in last 10 days
            diffHourly = GetDifference(priceProgress[0].y,priceProgress[priceProgress.length - 1].y)
            
            document.getElementById(`difference-${sectionId}`).innerHTML=`
            <div style="position: relative;">
            <span>${diffHourly.sign}${diffHourly.value}<br>(${diffHourly.percent})</span>
            <span style="position: absolute; top: -10px; right: 10px; font-size: 50px;">${diffHourly.arrow}</span>
            </div>`;
            document.getElementById(`difference-${sectionId}`).style.color = diffHourly.color;
            document.getElementById(`difference-${sectionId}`).style.position = 'relative';

            //Update line charts
            const chart = Chart.getChart(`chart-${sectionId}`);
            chart.data.labels = createEmptyStringList(priceProgress.length);
            chart.data.datasets[0].data = priceProgress;
            chart.data.datasets[0].label = diffHourly.percent;
            chart.data.datasets[0].borderColor = diffHourly.color;
            chart.data.datasets[0].backgroundColor=tinycolor(diffHourly.color).setAlpha(0.2).toRgbString();
            chart.update();
        } 
    


        //Update table
        document.getElementById(`ticker-${sectionId}`).innerHTML= ticker;
        document.getElementById(`lastprice-${sectionId}`).innerHTML=`${lastPrice}`;

  })

    
};





function CreateWatchlistTable(stocks) {
    const tableBody = document.getElementById('watchlistTable');

    stocks.forEach(stock => {
    const row = document.createElement('tr');

    const sectionId=`watchlist-${stock.Ticker}`

    row.innerHTML = `
        <td id="ticker-${sectionId}"></td>
        <td><canvas id="chart-${sectionId}" height="150" style="width:100%"></canvas></td>
        <td id="difference-${sectionId}"></td>
        <td id="lastprice-${sectionId}"></td>
    `;
    tableBody.appendChild(row);

    createLineChart(`chart-${sectionId}`, [0])
    });
    //Hide by default, need to be used here, after data loads, otherwise it weirdly stretches
    document.getElementById('watchlist').style.display = 'none';
  }