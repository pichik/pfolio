
function CreateHistory(stocks){
    CreateDividendTable(stocks)
    CreatePurchaseTable(stocks)

}


function CreateDividendTable(stocks){
    const accumulatedDividends = {};

    // Iterate over each stock
    stocks.forEach(stock => {
        if(stock.Dividends == null){
            return
        }
        // Iterate over each dividend of the stock
        stock.Dividends.forEach(dividend => {
            // Convert the dividend timestamp to milliseconds since the epoch
            const timestamp = dividend.Timestamp * 1000; // Convert seconds to milliseconds

            // Get the year and month of the dividend
            const date = new Date(timestamp);
            const year = date.getFullYear();
            const month = date.getMonth() + 1;

            // Initialize the object for the year if it doesn't exist
            if (!accumulatedDividends[year]) {
                accumulatedDividends[year] = {};
            }

            // Initialize the accumulated amount for the month if it doesn't exist
            if (typeof accumulatedDividends[year][month] === 'undefined' || accumulatedDividends[year][month] === null) {
                for (let m = 1; m <= 12; m++) {
                    accumulatedDividends[year][m] = {};
                    accumulatedDividends[year][m][0] = "";
                    accumulatedDividends[year][m][1] = 0;
                }
            }


            if (!accumulatedDividends[year]["all"]) {
                accumulatedDividends[year]["all"] = 0;
            }

            ticker = stock.Ticker.split('.')[0]

            const regex = new RegExp(`\\b${ticker}\\b`);

            if (!regex.test(accumulatedDividends[year][month][0])) {
                
                accumulatedDividends[year][month][0] += `${ticker}<br>`;
            }

            // Add the dividend amount to the accumulated dividends for the month
            accumulatedDividends[year][month][1] += dividend.TaxedPayout;
            accumulatedDividends[year]["all"] += dividend.TaxedPayout;
        });
    });

    createMonthlyDividendBarChart(`barchart-monthly-dividends`, accumulatedDividends)


    // Get the table body
    const tbody = document.getElementById('dividendMonthsTable');
    const tickertbody = document.getElementById('dividendMonthsTickerTable');

       // Initialize an object to store accumulated dividend amounts for each month


    // Iterate over each year in the accumulated dividends object
    Object.keys(accumulatedDividends).forEach(year => {
        // Create a new row for the current year
        let row = tickertbody.insertRow();
        let cell = row.insertCell();
        
        // Insert the year in the first cell and apply the yellow color directly
        cell.textContent = year;
        cell.style.color = 'yellow'; // Apply the yellow color directly to the cell

        // Iterate over each month (from January to December)
        for (let month = 1; month <= 12; month++) {
            // Get the accumulated dividend amount for the current month of the current year
            var ticker = accumulatedDividends[year][month][0] || "";
            if (ticker == "") {
                ticker = "-";
            }
            
            // Create a new cell and insert the dividend amount
            cell = row.insertCell();
            cell.innerHTML = ticker; // Convert the amount to a fixed decimal representation
        }




        row = tbody.insertRow();
        // Insert the year in the first cell and apply the yellow color directly
        cell = row.insertCell();
        cell.textContent = year;
        cell.style.color = 'yellow'; // Apply the yellow color directly to the cell
        
        // Iterate over each month (from January to December)
        for (let month = 1; month <= 12; month++) {
            // Get the accumulated dividend amount for the current month of the current year
            var amount = accumulatedDividends[year][month][1] || 0;
            if (amount == 0) {
                amount = "-";
            }else{
                amount = parseFloat(amount).toFixed(2)
            }
            
            // Create a new cell and insert the dividend amount
            cell = row.insertCell();
            cell.textContent = amount; // Convert the amount to a fixed decimal representation
        }

        var amount = accumulatedDividends[year]["all"] || 0;
            if (amount == 0) {
                amount = "-";
            }else{
                amount = parseFloat(amount).toFixed(2)
            }
            // Create a new cell and insert the dividend amount
            cell = row.insertCell();
            cell.textContent = amount; // Convert the amount to a fixed decimal representation
    });
}


function CreatePurchaseTable(stocks){
    const accumulatedPurchases = {};

    // Iterate over each stock
    stocks.forEach(stock => {
        if(stock.Purchases == null){
            return
        }
        // Iterate over each dividend of the stock
        stock.Purchases.forEach(purchase => {
            // Convert the dividend timestamp to milliseconds since the epoch
            const timestamp = purchase.Timestamp * 1000; // Convert seconds to milliseconds

            // Get the year and month of the dividend
            const date = new Date(timestamp);
            const year = date.getFullYear();
            const month = date.getMonth() + 1;

            // Initialize the object for the year if it doesn't exist
            if (!accumulatedPurchases[year]) {
                accumulatedPurchases[year] = {};
            }
            if (typeof accumulatedPurchases[year][month] === 'undefined' || accumulatedPurchases[year][month] === null) {
                    for (let m = 1; m <= 12; m++) {
                    accumulatedPurchases[year][m] = 0;
                }
            }

            if (!accumulatedPurchases[year]["all"]) {
                accumulatedPurchases[year]["all"] = 0;
            }

            // Add the dividend amount to the accumulated dividends for the month
            accumulatedPurchases[year][month] += purchase.Quantity * purchase.Price;
            accumulatedPurchases[year]["all"] += purchase.Quantity * purchase.Price;
        });
    });

    createMonthlyPurchaseBarChart(`barchart-monthly-purchases`, accumulatedPurchases)


    // Get the table body
    const tbody = document.getElementById('purchasesMonthsTable');

       // Initialize an object to store accumulated dividend amounts for each month
       // Clear all existing rows
    while (tbody.firstChild) {
        tbody.removeChild(tbody.firstChild);
    }
    // Iterate over each year in the accumulated dividends object
    Object.keys(accumulatedPurchases).forEach(year => {
        // Create a new row for the current year
        let row = tbody.insertRow();
        
        // Insert the year in the first cell and apply the yellow color directly
        let cell = row.insertCell();
        cell.textContent = year;
        cell.style.color = 'yellow'; // Apply the yellow color directly to the cell
        
        // Iterate over each month (from January to December)
        for (let month = 1; month <= 12; month++) {
            // Get the accumulated dividend amount for the current month of the current year
            var amount = accumulatedPurchases[year][month] || 0;
            if (amount == 0) {
                amount = "-";
            }else{
                amount = parseFloat(amount).toFixed(2)
            }
            
            // Create a new cell and insert the dividend amount
            cell = row.insertCell();
            cell.textContent = amount; // Convert the amount to a fixed decimal representation
        }

        var amount = accumulatedPurchases[year]["all"] || 0;
            if (amount == 0) {
                amount = "-";
            }else{
                amount = parseFloat(amount).toFixed(2)
            }
            // Create a new cell and insert the dividend amount
            cell = row.insertCell();
            cell.textContent = amount; // Convert the amount to a fixed decimal representation
    });
}