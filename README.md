# Portfolio Tracker
Track your investing portfolio

## Setup  
1. Install  
```
go install github.com/pichik/pfolio@latest
```
or
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
3. Fill `.pfconfig` and add it to your home directory.  
4. Move web directory to a place specified in config.  
5. Remove the rest.

All files in `assets/` are publicly accessible.  
If you are running https, directory with certificates should be generated.  



## Usage

