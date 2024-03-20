

const apiUrl = '/stock-update'; 

function GetData(){
    return fetch(apiUrl)
        .then(response => response.json())
        .then(data => {
            UsdToEur = data.UsdToEur;
            EurToUsd = data.EurToUsd;

            wList = data.wList;


            eurPfolio = data.xtb_eur || [];
            usdPfolio = data.xtb_usd || [];

            eurPfolio = SynchroniseEurPfolio(eurPfolio)
            allPfolio = [...eurPfolio,...usdPfolio];

            if (eurPfolio.length > 0 && usdPfolio.length > 0){
                allPfolio = ChangePfolioCurrency(allPfolio)
                document.getElementById('currency-btn').style.visibility = 'visible';
            }else if (eurPfolio.length > 0){
                currentCurrency=Currency.EUR
            }else if (usdPfolio.length > 0){
                currentCurrency=Currency.USD
            }

            return {allPfolio,wList}
    })
    .catch(error => {
        console.error('Error fetching stock data:', error);
        throw error;
    });
}

async function CreateStockTables() {
    try {
        const data = await GetData();

        const allPfolio = data.allPfolio;
        const wList = data.wList;

        CreateTable(allPfolio)
        UpdateData(allPfolio)
        CreateHistory(allPfolio)
        //Hide by default, need to be used here, after data loads, otherwise it weirdly stretches
        document.getElementById('months').style.display = 'none';

        if(wList != null){
            CreateWatchlistTable(wList)
            UpdateWatchlist(wList)
        }else{
            document.getElementById('watchlist').style.display = 'none';
        }
    } catch (error) {
        console.error('Error adding data:', error);
    }
}

async function UpdateStockData() {
    
    try {
        const data = await GetData();

        const allPfolio = data.allPfolio;
        const wList = data.wList;

        UpdateData(allPfolio)
        CreateHistory(allPfolio)

        if(wList != null){
            UpdateWatchlist(wList)
        }

    } catch (error) {
        console.error('Error fetching stock data:', error);
      }
}