Simple Internet connection sentinel.  
Works as http endpoint /ping

Environment vars:  
PINGER_CHAT - telegram chat id  
PINGER_TOKEN - telegram bot token  
PINGER_ADDR - addr to listen to. format '127.0.0.1:8080'.
PINGER_OK_MESSAGE - message is sent when Internet restored
PINGER_BAD_MESSAGE - message is sent when Internet is lost 
Best used with VPN due to no auth available.   
Maybe simple auth will be added later