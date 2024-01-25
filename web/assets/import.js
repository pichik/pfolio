// function importXTB() {
//     // Get the input value
//     const inputData = document.getElementById('inputData').value;

//     // Send the data to your server (you can use fetch or XMLHttpRequest)
//     fetch('/import-xtb', {
//       method: 'POST',
//       headers: {
//         'Content-Type': 'application/json',
//       },
//       body: JSON.stringify({ data: inputData }),
//     })
//     .then(response => response.json())
//     .then(data => {
//       // Handle the server response
//       console.log('Server response:', data);
//     })
//     .catch(error => {
//       console.error('Error:', error);
//     });
//   }

  function importXTB() {
    // Get the file input element
    const fileInput = document.getElementById('fileInput');
  
    // Check if a file is selected
    if (fileInput.files.length > 0) {
      // Get the first selected file
      const file = fileInput.files[0];
  
      // Create a FormData object
      const formData = new FormData();
  
      // Append the file to the FormData object
      formData.append('file', file);
  
      // Perform the file upload using fetch or XMLHttpRequest
      // Example using fetch:
      fetch('/import-xtb', {
        method: 'POST',
        body: formData,
      })
      .then(response => response.json())
      .then(data => {
        // Handle the server response
        console.log('Server response:', data);
      })
      .catch(error => {
        console.error('Error:', error);
      });
    } else {
      console.error('No file selected.');
    }
  }