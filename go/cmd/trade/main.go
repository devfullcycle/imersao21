package main

import (
	"encoding/json"
	"fmt"
	"sync"

	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/devfullcycle/imersao21/go/internal/infra/kafka"
	"github.com/devfullcycle/imersao21/go/internal/market/dto"
	"github.com/devfullcycle/imersao21/go/internal/market/entity"
	"github.com/devfullcycle/imersao21/go/internal/market/transformer"
)

func main() {
	// Canais para comunicação entre componentes
	ordersIn := make(chan *entity.Order)  // Canal para novas ordens
	ordersOut := make(chan *entity.Order) // Canal para ordens processadas
	wg := &sync.WaitGroup{}               // WaitGroup para sincronização de goroutines
	defer wg.Wait()                       // Aguarda todas as goroutines terminarem

	// Canal para receber mensagens do Kafka
	kafkaMsgChan := make(chan *ckafka.Message)

	// Configuração do consumidor Kafka
	// - bootstrap.servers: endereço do broker Kafka
	// - group.id: identificador do grupo de consumo
	// - auto.offset.reset: de onde começar a ler as mensagens
	consumerConfig := &ckafka.ConfigMap{
		"bootstrap.servers": "host.docker.internal:9094",
		"group.id":          "trade",
		"auto.offset.reset": "latest",
	}

	// Configuração do produtor Kafka
	// - bootstrap.servers: endereço do broker Kafka
	producerConfig := &ckafka.ConfigMap{
		"bootstrap.servers": "host.docker.internal:9094",
	}

	// Inicializa produtor e consumidor Kafka
	producer := kafka.NewKafkaProducer(producerConfig)
	consumer := kafka.NewConsumer(consumerConfig, []string{"orders"})

	// Inicia o consumo de mensagens em uma goroutine separada
	go consumer.Consume(kafkaMsgChan)

	// Cria e inicia o livro de ofertas em uma goroutine separada
	book := entity.NewBook(ordersIn, ordersOut, wg)
	go book.Trade()

	// Goroutine para processar mensagens recebidas do Kafka
	go func() {
		for msg := range kafkaMsgChan {
			wg.Add(1) // Incrementa o contador do WaitGroup para cada ordem
			fmt.Println(string(msg.Value))

			// Converte a mensagem JSON recebida para struct
			tradeInput := dto.TradeInput{}
			err := json.Unmarshal(msg.Value, &tradeInput)
			if err != nil {
				panic(err)
			}

			// Transforma o input em uma ordem e envia para processamento
			order := transformer.TransformInput(tradeInput)
			ordersIn <- order
		}
	}()

	// Loop principal para processar ordens executadas
	for res := range ordersOut {
		// Transforma a ordem processada em output JSON
		output := transformer.TransformOutput(res)
		jsonOutput, err := json.MarshalIndent(output, "", "  ")
		if err != nil {
			panic(err)
		}

		// Imprime e publica o resultado no Kafka
		fmt.Println(string(jsonOutput))
		producer.Publish(jsonOutput, []byte("processed_orders"), "processed_orders")
	}

	// Observação: O loop acima mantem o programa rodando indefinidamente
	// Nesse caso a utilizacao do waitgroup nao é necessaria, mas implementamos
	// para que possamos utilizar no futuro em caso de necessidade de processar
	// outras mensagens de outros canais ou lotes de mensagens de entrada
	// e saída de forma assíncrona antes de fechar o programa
}
