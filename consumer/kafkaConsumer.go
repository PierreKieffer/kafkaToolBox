package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/segmentio/kafka-go"
	"github.com/tkanos/gonfig"
)

type Configuration struct {
	KafkaBrokers         string
	KafkaTopic           string
	KafkaConsumerGroupId string
}

func initKafkaReader(kafkaBrokers, kafkaTopic, kafkaConsumerGroupId string) *kafka.Reader {
	brokers := strings.Split(kafkaBrokers, ",")

	return kafka.NewReader(kafka.ReaderConfig{
		Brokers:  brokers,
		GroupID:  kafkaConsumerGroupId,
		Topic:    kafkaTopic,
		MinBytes: 10e3,
		MaxBytes: 10e6,
		MaxWait:  time.Second,
	})
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Missing config.json file path argument")
		return
	}

	configFilePath := os.Args[1]
	configuration := Configuration{}
	err := gonfig.GetConf(configFilePath, &configuration)
	if err != nil {
		log.Fatal(err)
		return
	}

	kafkaReader := initKafkaReader(configuration.KafkaBrokers, configuration.KafkaTopic, configuration.KafkaConsumerGroupId)

	defer kafkaReader.Close()

	log.Println("Starting kafka consumer")

	for {
		m, err := kafkaReader.ReadMessage(context.Background())
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(m.Value))
	}
}
