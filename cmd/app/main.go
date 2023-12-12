package main

import (
	"database/sql"
	"encoding/json"
	"mensageria/internal/infra/akafka"
	"mensageria/internal/infra/repository"
	"mensageria/internal/infra/web"
	"mensageria/internal/usecase"
	"net/http"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/go-chi/chi/v5"
)

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(host.docker.internal:3306)/products")
	if err != nil {
		panic(err)
	}

	defer db.Close()

	repository := repository.NewProductRepository(db)
	createProductUsecase := usecase.NewCreateProductRepository(repository)
	listProductUseCase := usecase.NewListProductRepository(repository)

	productHandlers := web.NewProductHandlers(createProductUsecase, listProductUseCase)

	r := chi.NewRouter()
	r.Post("/products", productHandlers.CreateProductHandler)
	r.Get("/products", productHandlers.ListProducsHandler)

	go http.ListenAndServe(":8000", r)

	msgChan := make(chan *kafka.Message)
	go akafka.Consume([]string{"products"}, "host.docker.internal9094", msgChan)

	for msg := range msgChan {
		dto := usecase.CreateProductInputDto{}
		err := json.Unmarshal(msg.Value, &dto)
		if err != nil {

		}

		_, err = createProductUsecase.Execute(dto)
	}

}
