# kafkaToolBox 

kafkaToolBox provides
- producer-api : Simple API allowing data ingestion and pushing messages to Kafka 
- consumer :  Simple Kafka consumer 

## Build and deploy 
- set the config file with your own settings 
- For each application : go build . 
- ./producer-api /path/config.json
	`curl -i http://localhost:8080/producer -X POST -d '{"message":"Here we go !"}' -H  "Authorization: Bearer token"`
- ./consumer /path/config.json



