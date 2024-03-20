
const Sections = {
    Pfolio: 'Portfolio',
    Wlist: 'Watchlist',
    Months: 'Months'
  };


var currentSection  = Sections.Pfolio;


function SwitchSections(section){
    document.getElementById("searchTable").value="";
    //reset search for previous table
    regexSearch( new RegExp("", 'i'))

    // currentSection = currentSection == Sections.Pfolio ? Sections.Wlist : Sections.Pfolio;
    currentSection = section;

    // document.getElementById("sectionSwitcher").innerHTML = currentSection == Sections.Pfolio ? Sections.Pfolio : Sections.Wlist;
    if(currentSection == Sections.Pfolio){
        document.getElementById('pfolio').style.display = 'block';
        document.getElementById('watchlist').style.display = 'none';
        document.getElementById('months').style.display = 'none';
    }else if(currentSection == Sections.Wlist){
        document.getElementById('pfolio').style.display = 'none';
        document.getElementById('watchlist').style.display = 'block';
        document.getElementById('months').style.display = 'none';
    }else if(currentSection == Sections.Months){
        document.getElementById('pfolio').style.display = 'none';
        document.getElementById('watchlist').style.display = 'none';
        document.getElementById('months').style.display = 'block';
    }
    // document.getElementById(currentSection == Sections.Pfolio ? 'pfolio':'watchlist').style.display = 'block';
    // document.getElementById(currentSection == Sections.Pfolio ? 'watchlist':'pfolio').style.display = 'none';

    document.getElementById('wlist-inputs').style.visibility = currentSection == Sections.Wlist ? 'visible':'hidden';
    document.getElementById('searchTable').style.visibility = currentSection == Sections.Months ? 'hidden':'visible';
}




function GetDifference(from, to) {
    let arrow = '';
    let color = 'white';
    let sign = '';
    let value = parseFloat(to - from).toFixed(2);
    let percent = Math.abs(parseFloat(((to - from) / from) * 100).toFixed(2))+"%";

    if (value > 0) {
        arrow = '↑';
        sign = '+';
        color = 'green';
    } else if (value < 0) {
        arrow = '↓';
        color = 'red';
    } 

    return { value, percent, color, sign, arrow };
}