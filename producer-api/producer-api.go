package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/segmentio/kafka-go"
	"github.com/tkanos/gonfig"
)

type Configuration struct {
	KafkaBrokers         string
	KafkaTopic           string
	KafkaConsumerGroupId string
	KafkaClientId        string
	Secret               string
}

func initKafkaWriter(kafkaBrokers, kafkaTopic string) *kafka.Writer {
	return kafka.NewWriter(kafka.WriterConfig{
		Brokers:  []string{kafkaBrokers},
		Topic:    kafkaTopic,
		Balancer: &kafka.LeastBytes{},
	})

}

func getHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	w.Write([]byte(`{"message": "log"}`))
}

func checkToken(r *http.Request, secret string) bool {
	headerToken := r.Header.Get("Authorization")
	bearerToken := strings.Split(headerToken, "Bearer ")[1]

	if bearerToken != secret {
		return false
	} else {
		return true
	}
}

func producerHandler(kafkaWriter *kafka.Writer, secret string) func(http.ResponseWriter, *http.Request) {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		auth := checkToken(r, secret)
		if auth == false {
			log.Println("401 Unauthorized : Message not pushed to kafka")
			w.Write([]byte("401 Unauthorized : Message not pushed to kafka"))
			return
		}

		payload, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Fatal(err)
		}
		log.Println("Try to push message to kafka : " + string(payload))
		msg := kafka.Message{Value: payload}
		err = kafkaWriter.WriteMessages(r.Context(), msg)
		if err != nil {
			w.Write([]byte(err.Error()))
			log.Fatal(err)
		}
		log.Println("Message successfully pushed to kafka")
		w.Write([]byte("Message successfully pushed to kafka"))

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

	kafkaWriter := initKafkaWriter(configuration.KafkaBrokers, configuration.KafkaTopic)

	defer kafkaWriter.Close()

	log.Println("Listening on http://localhost:8080")

	http.HandleFunc("/producer", producerHandler(kafkaWriter, configuration.Secret))
	http.ListenAndServe(":8080", nil)

}
