package store

import (
	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
	"fmt"
    "encoding/json"
    "store/config"
)

var producer *kafka.Producer
var consumer *kafka.Consumer
var topic string

func InitProducer(platform string){
	var err error
    var options = config.Reader(platform)
	producer, err = kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": options["bootstrap_servers"]})
	if err != nil {
		panic(err)
	}
	topic = options["destination"]
}

func InitConsumer(platform string){
	var err error
    var options = config.Reader(platform)
	consumer, err = kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": options["bootstrap_servers"],
		"group.id":          options["group_id"],
		"broker.address.family": "v4",
		"session.timeout.ms":    6000,
		"auto.offset.reset": "earliest",
	})
	topic = options["destination"]

	if err != nil {
		panic(err)
	}
	defer consumer.Close()
	KafkaConsumer()
	
}

func KafkaProducer() (*kafka.Producer, string){
	
	return producer, topic
}

func KafkaConsumer(){
    
	
    consumer.SubscribeTopics([]string{topic}, nil)

	var dat map[string]interface{}
    for {
        msg, err := consumer.ReadMessage(-1)
        if err == nil {
			if err := json.Unmarshal(msg.Value, &dat); err != nil {
				panic(err)
			}
            if dat["eventType"] == "Paid"{
                wheneverPaid_주문목록에추가(dat)
            }
            if dat["eventType"] == "OrderCanceled"{
                wheneverOrderCanceled_주문취소알림(dat)
            }
			if dat["eventType"] == "OrderPlaced"{
				whenOrderPlaced_then_CREATE_1(dat)
			}
			if dat["eventType"] == "OrderPlaced"{
				whenOrderPlaced_then_UPDATE_1(dat)
			}

			
        } else {
            // The client will automatically try to recover from all errors.
            fmt.Printf("Consumer error: %v (%v)\n", err, msg)
        }
    }
}

func Streamhandler(message string){
	producer, topic := KafkaProducer()
	
	producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          []byte(message),
	}, nil)

}