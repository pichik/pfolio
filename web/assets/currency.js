
var UsdToEur;
var EurToUsd;

const Currency={
    USD:"USD",
    EUR:"EUR"
}

var currentCurrency=Currency.USD;

function ChangeCurrency(){
    currentCurrency = currentCurrency == Currency.USD ? Currency.EUR : Currency.USD;
    document.getElementById('currency-btn').innerHTML = currentCurrency == Currency.USD ? '$' : '€';
    UpdateStockData()
}


function ChangePfolioCurrency(stocks) {
    if (currentCurrency === Currency.USD){
        stocks.forEach(stock => {
            var strPrice = stock.Stock.LastPriceStr;
            if (strPrice.includes("€")){
                stock.BuyPrice *= EurToUsd;
                stock.Dividend *= EurToUsd;

                stock.Stock.LastPrice *= EurToUsd;

                if (stock.Stock.HourlyPrices) {
                    for (const timestamp in stock.Stock.HourlyPrices) {
                        if (Object.hasOwnProperty.call(stock.Stock.HourlyPrices, timestamp)) {
                            stock.Stock.HourlyPrices[timestamp] *= EurToUsd;
                        }
                    }
                }
            }

        });
    }else if (currentCurrency === Currency.EUR){
        stocks.forEach(stock => {
            var strPrice = stock.Stock.LastPriceStr;
            if (strPrice.includes("$")){
                stock.BuyPrice *= UsdToEur;
                stock.Dividend *= UsdToEur;

                stock.Stock.LastPrice *= UsdToEur;

                if (stock.Stock.HourlyPrices) {
                    for (const timestamp in stock.Stock.HourlyPrices) {
                        if (Object.hasOwnProperty.call(stock.Stock.HourlyPrices, timestamp)) {
                            stock.Stock.HourlyPrices[timestamp] *= UsdToEur;
                        }
                    }
                }
            }
        });
    }
    return stocks
}


function SynchroniseEurPfolio(stocks){
    stocks.forEach(stock => {

        var strPrice = stock.Stock.LastPriceStr;

        if (strPrice.includes("€")){

        }
        if (strPrice.includes("$")){
            stock.Stock.LastPrice *= UsdToEur;

            if (stock.Stock.HourlyPrices) {
                for (const timestamp in stock.Stock.HourlyPrices) {
                    if (Object.hasOwnProperty.call(stock.Stock.HourlyPrices, timestamp)) {
                        stock.Stock.HourlyPrices[timestamp] *= UsdToEur;
                    }
                }
            }
        }
        
      });
      return stocks
}





