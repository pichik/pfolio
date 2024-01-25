function importXTB() {
    // Get the input value
    const inputData = document.getElementById('inputData').value;

    // Send the data to your server (you can use fetch or XMLHttpRequest)
    fetch('/import-xtb', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({ data: inputData }),
    })
    .then(response => response.json())
    .then(data => {
      // Handle the server response
      console.log('Server response:', data);
    })
    .catch(error => {
      console.error('Error:', error);
    });
  }