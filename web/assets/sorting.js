


document.getElementById("searchTable").addEventListener("input", function() {
    var filter = this.value.toLowerCase(); // Convert input value to lowercase for case-insensitive matching
    var tableRows = document.getElementById(currentSection == Sections.Pfolio ? 'pfolioTable':'wlistTable').getElementsByTagName("tr");
  
    // Loop through all table rows
    for (var i = 0; i < tableRows.length; i++) {
      var row = tableRows[i];
      var firstColumn = row.getElementsByTagName("td")[0]; // Get the first column of the current row
  
      if (firstColumn) {
        var textContent = firstColumn.textContent.toLowerCase(); // Get text content of the first column
  
        // If the text content contains the search query, show the row, otherwise hide it
        if (textContent.indexOf(filter) > -1) {
          row.style.display = "";
        } else {
          row.style.display = "none";
        }
      }
    }
  });
  


var sortDirection = {};

function sortTable(tableId, columnIndex) {
  var table = document.getElementById(tableId);
  var rows = Array.from(table.rows).slice(1); // Convert HTMLCollection to array and exclude the header row
  
  // Initialize sorting direction for this column if not already set
  if (!sortDirection[columnIndex]) {
    sortDirection[columnIndex] = 'asc';
  } else {
    sortDirection[columnIndex] = sortDirection[columnIndex] === 'asc' ? 'desc' : 'asc';
  }

  rows.sort(function(a, b) {
    var x = a.getElementsByTagName("td")[columnIndex].textContent.trim();
    var y = b.getElementsByTagName("td")[columnIndex].textContent.trim();

    // Attempt to parse as float
    var xFloat = parseFloat(x);
    var yFloat = parseFloat(y);

    // Check if both values are valid floats
    var isFloat = !isNaN(xFloat) && !isNaN(yFloat);

    if (isFloat) {
      x = xFloat;
      y = yFloat;
    }

    // Compare based on data type
    if (sortDirection[columnIndex] === 'asc') {
      if (isFloat) {
        return x - y;
      } else {
        return x.localeCompare(y);
      }
    } else {
      if (isFloat) {
        return y - x;
      } else {
        return y.localeCompare(x);
      }
    }
  });

  // Clear the table body
  while (table.rows.length > 1) {
    table.deleteRow(1);
  }

  // Re-add sorted rows to the table
  rows.forEach(function(row) {
    table.appendChild(row);
  });
}




function sortPercentTable(tableId, columnIndex) {
  var table = document.getElementById(tableId);
  var rows = Array.from(table.rows).slice(1); // Exclude the header row
  
  // Initialize sorting direction for this column if not already set
  if (!sortDirection[columnIndex]) {
    sortDirection[columnIndex] = 'asc';
  } else {
    sortDirection[columnIndex] = sortDirection[columnIndex] === 'asc' ? 'desc' : 'asc';
  }

  rows.sort(function(a, b) {
    var x = getPercentageValue(a.getElementsByTagName("td")[columnIndex]);
    var y = getPercentageValue(b.getElementsByTagName("td")[columnIndex]);

    if (sortDirection[columnIndex] === 'asc') {
      return x - y;
    } else {
      return y - x;
    }
  });

  // Clear the table body
  while (table.rows.length > 1) {
    table.deleteRow(1);
  }

  // Re-add sorted rows to the table
  rows.forEach(function(row) {
    table.appendChild(row);
  });
}

function getPercentageValue(cell) {
  var textContent = cell.textContent.trim(); // Get the text content of the cell
  var percentageMatch = textContent.match(/(\d+(\.\d+)?%)/)[0]; // Match the percentage value

  if (percentageMatch) {
    return parseFloat(percentageMatch[0]); // Parse the matched percentage value as float
  } else {
    return NaN; // Return NaN if no percentage value found
  }
}



  