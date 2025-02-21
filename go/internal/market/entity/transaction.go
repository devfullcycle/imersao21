package entity

import (
	"time"

	"github.com/google/uuid"
)

// Transaction representa uma transação realizada entre duas ordens (compra e venda)
// Registra todos os detalhes da negociação, incluindo quantidade, preço e valor total
type Transaction struct {
	ID           string    // Identificador único da transação
	SellingOrder *Order    // Ordem de venda
	BuyingOrder  *Order    // Ordem de compra
	Shares       int       // Quantidade de ações negociadas
	Price        float64   // Preço por ação da negociação
	Total        float64   // Valor total da transação (Shares * Price)
	DateTime     time.Time // Data e hora da transação
}

// NewTransaction cria uma nova transação entre uma ordem de venda e uma de compra
// Parâmetros:
// - sellingOrder: ordem de venda
// - buyingOrder: ordem de compra
// - shares: quantidade de ações a serem negociadas
// - price: preço por ação acordado
// Retorna uma nova transação com ID único e timestamp atual
func NewTransaction(sellingOrder *Order, buyingOrder *Order, shares int, price float64) *Transaction {
	// Cria uma nova transação com:
	// - ID único gerado automaticamente
	// - Total inicializado como 0 (será calculado durante o processamento)
	// - Timestamp atual como momento da criação
	return &Transaction{
		ID:           uuid.New().String(),
		SellingOrder: sellingOrder,
		BuyingOrder:  buyingOrder,
		Shares:       shares,
		Price:        price,
		Total:        0, // Total será calculado durante o processamento
		DateTime:     time.Now(),
	}
}

// Process executa o processamento da transação
// Utiliza um OrderProcessor para atualizar as ordens envolvidas
// e registrar a transação em ambas as ordens
func (t *Transaction) Process() {
	// Cria um processador específico para esta transação
	// que irá atualizar as posições dos investidores e
	// o status das ordens envolvidas
	processor := NewOrderProcessor(t)
	processor.Process()
}
