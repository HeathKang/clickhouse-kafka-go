package main

import (
	"context"
	"fmt"
	"log"
    "encoding/json"
	// "time"
	kafka "github.com/segmentio/kafka-go"
    "github.com/heathkang/clickhouse-kafka/internal/data"

)

func main() {
	fmt.Println("helo word")
    r := kafka.NewReader(kafka.ReaderConfig{
        Brokers: []string{"kuka-connect-kafka-0.kuka-connect-kafka.kuka-connect-dev.svc.digital-dev.kukaplus.com:9092"},
        Topic: "operational_data_batch",
        GroupID: "ning-test",
    })
    
    messageChan := make(chan kafka.Message)
    // kafka get message
    go kafkaConsumeChannel(r, messageChan)
    // consume kafka message
    go consumeKafkaMessage(messageChan)
    // main loop
    for {

    }
}

func kafkaConsumeChannel(r *kafka.Reader, ch chan kafka.Message) {
    for{
        message, err := r.ReadMessage(context.Background())
        if err != nil {
            log.Fatal(err)
            break
        }
        ch <- message
    }
}

func consumeKafkaMessage(ch chan kafka.Message) {
    for {
        select {
        // 1. 1 second interval; 
        case message := <- ch:
            msg := data.OperationalData{}
            err := json.Unmarshal(message.Value, &msg)
            if err != nil {
                log.Fatal("Wrong data format for json")
            }
            fmt.Printf("message at offset %d: %s = %s \n", message.Offset, string(message.Key), msg.BaseName)
        // 2. 1000 number data;
        }
    }
} 

