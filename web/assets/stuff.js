
const Sections = {
    Pfolio: 'Portfolio',
    Wlist: 'Watchlist'
  };


var currentSection  = Sections.Pfolio;


function SwitchSections(){
    document.getElementById("searchTable").value="";

    currentSection = currentSection == Sections.Pfolio ? Sections.Wlist : Sections.Pfolio;

    document.getElementById("sectionSwitcher").innerHTML = currentSection == Sections.Pfolio ? Sections.Pfolio : Sections.Wlist;
    document.getElementById(currentSection == Sections.Pfolio ? 'pfolio':'watchlist').style.display = 'block';
    document.getElementById(currentSection == Sections.Pfolio ? 'watchlist':'pfolio').style.display = 'none';

    document.getElementById('wlist-inputs').style.visibility = currentSection == Sections.Pfolio ? 'hidden':'visible';

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