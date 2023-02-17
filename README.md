# WebWatcher
Golang server for testing blind XSS.   
Special endpoint for testing requests and file extensions.  

## Setup


Setup .wwconfig and add it to your home directory.  
Make sure that `assets` and `templates` are in the directory specified in config.  

## Usage
Login to your website with /login?token=[token]

**Application have 2 parts**
1. Collecting data from blind XSS



2. Collecting simple request data from path you specified in .wwconfig  
This endpoint will contain response specified in `assets/extensions.json` file  
