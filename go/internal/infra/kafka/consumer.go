package kafka

import ckafka "github.com/confluentinc/confluent-kafka-go/kafka"

// Consumer encapsula a funcionalidade de consumo de mensagens do Kafka
// Permite consumir mensagens de múltiplos tópicos simultaneamente
type Consumer struct {
	ConfigMap *ckafka.ConfigMap // Configurações de conexão com o Kafka
	Topics    []string          // Lista de tópicos para consumir mensagens
}

// NewConsumer cria uma nova instância do consumidor Kafka
// Parâmetros:
// - configMap: mapa de configurações do Kafka (ex: broker, grupo de consumo, etc)
// - topics: lista de tópicos que serão consumidos
func NewConsumer(configMap *ckafka.ConfigMap, topics []string) *Consumer {
	return &Consumer{
		ConfigMap: configMap,
		Topics:    topics,
	}
}

// Consume inicia o consumo de mensagens dos tópicos configurados
// As mensagens são enviadas para o canal msgChan para processamento
// Este método bloqueia e roda indefinidamente até que ocorra um erro
// Parâmetro:
// - msgChan: canal onde as mensagens consumidas serão enviadas
func (c *Consumer) Consume(msgChan chan *ckafka.Message) error {
	// Cria uma nova instância do consumidor com as configurações
	consumer, err := ckafka.NewConsumer(c.ConfigMap)
	if err != nil {
		panic(err)
	}

	// Inscreve o consumidor nos tópicos especificados
	// O segundo parâmetro (nil) indica que não há callback específico para rebalanceamento
	err = consumer.SubscribeTopics(c.Topics, nil)
	if err != nil {
		panic(err)
	}

	// Loop infinito para consumo contínuo de mensagens
	// ReadMessage(-1) bloqueia até que uma nova mensagem chegue
	for {
		msg, err := consumer.ReadMessage(-1)
		if err == nil {
			msgChan <- msg
		}
	}
}
