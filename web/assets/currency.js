
var UsdToEur;
var EurToUsd;

const Currency={
    USD:{Name:"USD",Symbol:"$"},
    EUR:{Name:"EUR",Symbol:"€"}
}

var currentCurrency=Currency.USD;

function ChangeCurrency(){
    currentCurrency = currentCurrency == Currency.USD ? Currency.EUR : Currency.USD;
    document.getElementById('currency-btn').innerHTML = currentCurrency.Symbol;
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
                if (stock.Stock.DailyPrices) {
                    for (const timestamp in stock.Stock.DailyPrices) {
                        if (Object.hasOwnProperty.call(stock.Stock.DailyPrices, timestamp)) {
                            stock.Stock.DailyPrices[timestamp] *= EurToUsd;
                        }
                    }
                }
                if (stock.Purchases) {
                    for (const k in stock.Purchases) {
                        stock.Purchases[k].Price *= EurToUsd;
                    }
                }
                if (stock.Dividends) {
                    for (const k in stock.Dividends) {
                        stock.Dividends[k].TaxedPayout *= EurToUsd;
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
                if (stock.Stock.DailyPrices) {
                    for (const timestamp in stock.Stock.DailyPrices) {
                        if (Object.hasOwnProperty.call(stock.Stock.DailyPrices, timestamp)) {
                            stock.Stock.DailyPrices[timestamp] *= UsdToEur;
                        }
                    }
                }
                if (stock.Purchases) {
                    for (const k in stock.Purchases) {
                        stock.Purchases[k].Price *= UsdToEur;
                    }
                }
                if (stock.Dividends) {
                    for (const k in stock.Dividends) {
                        stock.Dividends[k].TaxedPayout *= UsdToEur;
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
            if (stock.Stock.DailyPrices) {
                for (const timestamp in stock.Stock.DailyPrices) {
                    if (Object.hasOwnProperty.call(stock.Stock.DailyPrices, timestamp)) {
                        stock.Stock.DailyPrices[timestamp] *= UsdToEur;
                    }
                }
            }
        }
        
      });
      return stocks
}





