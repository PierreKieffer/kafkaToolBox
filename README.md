# kafkaToolBox 

kafkaToolBox is based on segmentio/kafka-go package and provides : 
- producer-api : Simple API allowing data ingestion and pushing messages to Kafka 
- consumer :  Simple Kafka consumer 

## Build and deploy 
- To quickly deploy a Kafka cluster : https://github.com/PierreKieffer/docker-kafka-cluster
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



