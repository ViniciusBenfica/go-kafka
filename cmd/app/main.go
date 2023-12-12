package main

import (
	"database/sql"
	"encoding/json"
	"mensageria/internal/infra/akafka"
	"mensageria/internal/infra/repository"
	"mensageria/internal/usecase"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(host.docker.internal:3306/products)")
	if err != nil {
		panic(err)
	}

	defer db.Close()

	msgChan := make(chan *kafka.Message)
	go akafka.Consume([]string{"products"}, "host.docker.internal9094", msgChan)

	repository := repository.NewProductRepository(db)
	createProductUsecase := usecase.NewCreateProductRepository(repository)

	for msg := range msgChan {
		dto := usecase.CreateProductInputDto{}
		err := json.Unmarshal(msg.Value, &dto)

		_, err = createProductUsecase.Execute(dto)
	}

}
