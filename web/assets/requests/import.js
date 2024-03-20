  function importXTB() {
    // Get the file input element
    const fileInput = document.getElementById('xtb-file');

    // Get the selected option from the dropdown menu
    const dropdown = document.getElementById('xtb-currency');
    const selectedOption = dropdown.value;
  
    // Check if a file is selected
    if (fileInput.files.length > 0) {
      // Get the first selected file
      const file = fileInput.files[0];
  
      // Create a FormData object
      const formData = new FormData();

      // Append the selected option to the FormData object
      formData.append('xtb-currency', selectedOption);  
  
      // Append the file to the FormData object
      formData.append('xtb-file', file);
  
      // Perform the file upload using fetch or XMLHttpRequest
      // Example using fetch:
      fetch('/import-xtb', {
        method: 'POST',
        body: formData,
      })
      .then(response => {
        // Reload the page after the response is received
        location.reload();
      })
      .catch(error => {
        console.error('Error:', error);
      });
    } else {
      console.error('No file selected.');
    }
  }

  function importWlistStocks(){
    // Get the file input element
    const stockInput = document.getElementById('wlist-stock').value;


    // Check if a file is selected
    if (stockInput != "") {
      // Create a FormData object
      const formData = new FormData();

      // Append the selected option to the FormData object
      formData.append('tickers', stockInput);  

      fetch('/add-to-wlist', {
        method: 'POST',
        body: formData,
      })
    } 
    document.getElementById("wlist-stock").value = "";
  }

