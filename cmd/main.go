package main

import (
	"context"
	"fmt"
	"log"
	"time"

	kafka "github.com/segmentio/kafka-go"
)

func main() {
	fmt.Println("helo word")
	topic := "operational_data_batch"
	partion := 0
	context, cancel := context.WithCancel(context.Background());
	conn, err := kafka.DialLeader(context, 
		"tcp",
		"kuka-connect-kafka-0.kuka-connect-kafka.kuka-connect-dev.svc.digital-dev.kukaplus.com:9092",
		topic,
		partion)
	if err != nil {
		log.Fatal("failed to dial leader:", err)
	}

	conn.SetReadDeadline(time.Now().Add(10*time.Second))
	batch := conn.ReadBatch(10e3, 1e6)
	

	b := make([]byte, 10e3)
	go cancelProcess(cancel)
	for {
		_, err := batch.Read(b)
		if err != nil {
			break
		}
		fmt.Println(string(b))
	}

	if err := batch.Close(); err != nil {
		log.Fatal(" failed to close batch:", err)
	}

	if err := conn.Close(); err != nil {
		log.Fatal("failed to close connection:", err)
	}

}

func cancelProcess(cancel context.CancelFunc) {
	time.Sleep(1*time.Second)
	// cancel()
	fmt.Println("*********************")
	fmt.Println("*********************")
	fmt.Println("Cancel kafka")
	fmt.Println("*********************")
	fmt.Println("*********************")
}