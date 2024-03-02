
    function UpdateData(stocks){
        var investedTotal=0;
        var valueTotal=0;
        var dividendTotal=0;
        var stocksInvested = [];
        var stocksProfit = [];
        var stocksValue = new Map();



        stocks.forEach(stock => {
          const ticker=stock.Ticker;
          const buyPrice=parseFloat(stock.BuyPrice.toFixed(2));
          const dividend=parseFloat(stock.Dividend.toFixed(2));
          const lastPrice=parseFloat(stock.Stock.LastPrice.toFixed(2));
          const stockOwn=stock.BuyAmount;
          const invested=parseFloat((stock.BuyPrice*stock.BuyAmount).toFixed(2));
         


        //Calculate profit
        diffProfit = GetDifference(stock.BuyPrice,stock.Stock.LastPrice)
        const profit=parseFloat(((diffProfit.value)*stock.BuyAmount).toFixed(2));
        for (const element of document.getElementsByClassName(`colored-${ticker}`)) {
          element.style.color = diffProfit.color;
        }



        //Data for bar chart
        stocksValue.set(ticker,[invested+profit]);
        stocksInvested.push(invested);
        stocksProfit.push(profit);




        //Update table
        document.getElementById(`ticker-${ticker}`).innerHTML= ticker;
        document.getElementById(`lastprice-${ticker}`).innerHTML=`${lastPrice}`;
        document.getElementById(`buyprice-${ticker}`).innerHTML= buyPrice;
        document.getElementById(`stockown-${ticker}`).innerHTML=stockOwn;
        document.getElementById(`invested-${ticker}`).innerHTML=invested;
        document.getElementById(`profit-${ticker}`).innerHTML=profit;
        document.getElementById(`profitpercent-${ticker}`).innerHTML=diffProfit.percent;
        document.getElementById(`div-${ticker}`).innerHTML=dividend > 0 ? dividend : '';
        


        //Update Chart
        if (stock.Stock.HourlyPrices != null && Object.entries(stock.Stock.HourlyPrices).length > 0){
          priceProgress = Object.entries(stock.Stock.HourlyPrices)
            .map(([timestamp, price]) => ({ x: parseInt(timestamp), y: price }))
          //Calculate graph difference in last 10 days
          diffHourly = GetDifference(priceProgress[0].y,priceProgress[priceProgress.length - 1].y)
          document.getElementById(`difference-${ticker}`).innerHTML=`
          <div style="position: relative;">
          <span>${diffHourly.sign}${diffHourly.value}<br>(${diffHourly.percent})</span>
          <span style="position: absolute; top: -10px; right: -10px; font-size: 50px;">${diffHourly.arrow}</span>
          </div>`;
          document.getElementById(`difference-${ticker}`).style.color = diffHourly.color;
          document.getElementById(`difference-${ticker}`).style.position = 'relative';

          //Update line charts
          const chart = Chart.getChart(`chart-${ticker}`);
          chart.data.labels = createEmptyStringList(priceProgress.length);
          chart.data.datasets[0].data = priceProgress;
          chart.data.datasets[0].label = diffHourly.percent;
          chart.data.datasets[0].borderColor = diffHourly.color;
          chart.data.datasets[0].backgroundColor=tinycolor(diffHourly.color).setAlpha(0.2).toRgbString();
          chart.update();

        } 

        //Updating total value of portfolio
        investedTotal+=invested;
        dividendTotal+=dividend;
        valueTotal+=stockOwn*lastPrice;


      })

      //Update total values
      valueTotal=parseFloat(valueTotal.toFixed(2));
      investedTotal=parseFloat(investedTotal.toFixed(2));
      dividendTotal= parseFloat(dividendTotal.toFixed(2));
      diffTotal = GetDifference(investedTotal, valueTotal);



      document.getElementById(`portfolio-value`).innerHTML=`Portfolio Value: ${valueTotal}`;
      document.getElementById(`portfolio-invested`).innerHTML=`Invested: ${investedTotal}`;
      document.getElementById(`portfolio-profit`).innerHTML=`Profit: ${diffTotal.value} (${diffTotal.percent})`;
      document.getElementById(`portfolio-profit`).style.color= diffTotal.color;
      document.getElementById(`portfolio-dividend`).innerHTML=`Dividends: ${dividendTotal}`;

      if (dividendTotal>0){
      diffTotalWithDivs = GetDifference(investedTotal, valueTotal+dividendTotal);
      document.getElementById(`portfolio-profit-dividends`).innerHTML=`(${diffTotalWithDivs.value} ${diffTotalWithDivs.percent})`;
      document.getElementById(`portfolio-profit-dividends`).style.color= diffTotalWithDivs.color;
      }

      

      // Iterate throu all stocks and calulate their percentage value in portfolio
      stocksValue.forEach((value, key, map) => {
        value.push(((value[0] / valueTotal) * 100).toFixed(2));
        stocksValue.set(key, value);
      });
 
      //Update pie chart
      const piechart = Chart.getChart("piechart-total");
      piechart.data.labels = Array.from(stocksValue.keys());
      piechart.data.datasets[0].data = Array.from(stocksValue.values()).map(item => item[1]);
      piechart.update();



      


      // Update bar chart
      const barchart = Chart.getChart("barchart-total");
      barchart.data.labels = Array.from(stocksValue.keys());
      barchart.data.datasets[0].data =stocksInvested;
      barchart.data.datasets[1].data =stocksProfit;
      barchart.update();
 
      
    };





    function CreateTable(stocks) {
        
        const tableBody = document.getElementById('tableBody');

        stocks.forEach(stock => {
          const row = document.createElement('tr');

          const ticker=stock.Ticker
    
          row.innerHTML = `
              <td id="ticker-${ticker}"></td>
              <td><canvas id="chart-${ticker}" height="150" style="width:100%"></canvas></td>
              <td id="difference-${ticker}"></td>
              <td id="lastprice-${ticker}"></td>
              <td id="buyprice-${ticker}"></td>
              <td id="stockown-${ticker}"></td>
              <td id="invested-${ticker}"></td>
              <td id="profit-${ticker}" class="colored-${ticker}"></td>
              <td id="profitpercent-${ticker}" class="colored-${ticker}"></td>
              <td id="div-${ticker}" style="color:green;"></td>
          `;
          tableBody.appendChild(row);
    
          createLineChart(`chart-${ticker}`, [0])
        });
    }


