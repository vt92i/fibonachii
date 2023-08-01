package producer

import (
	"fmt"
	"time"

	"github.com/vt92i/fibonachii/types"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/spf13/viper"
)

var (
	address   string = viper.GetString("KafkaAddress")
	topic     string = viper.GetString("KafkaTopic")
	partition int32  = viper.GetInt32("KafkaPartition")
)

func Produce(message types.Fibonacci) {
	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": address})
	if err != nil {
		panic(err)
	}
	defer p.Close()

	// messageStatus := true

	go func() {
		for e := range p.Events() {
			switch ev := e.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					// messageStatus = false
					fmt.Println(ev.TopicPartition.Error)
				}
			}
		}
	}()

	// if jsonData, err := json.Marshal(message); err != nil {
	// 	panic(err)
	// } else {
	// 	p.Produce(&kafka.Message{
	// 		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: partition},
	// 		Value:          []byte(string(jsonData)),
	// 	}, nil)
	// }

	// p.Flush(15 * 1000)

	// return messageStatus

	for {
		p.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: partition},
			Value:          []byte(string("message")),
		}, nil)

		time.Sleep(1000 * time.Millisecond)
		p.Flush(15 * 1000)
	}
}
