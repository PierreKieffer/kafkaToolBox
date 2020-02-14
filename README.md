# kafkaToolBox 

kafkaToolBox provides
- producer-api : Simple API allowing data ingestion and pushing messages to Kafka 
- consumer :  Simple Kafka consumer 

## Build and deploy 
- Set the config.json file with your own settings 
- Build an application : 
	
	`go build . `
	
- producer-api : 
	- Deploy : 
	
	`./producer-api /path/config.json`
	
	- Post a message :   	
	
	`curl -i http://localhost:8080/producer -X POST -d '{"message":"Here we go !"}' -H  "Authorization: Bearer token"`
	
- consumer : 
	- Deploy : 
	
	`./consumer /path/config.json`



