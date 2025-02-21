package entity

// Order representa uma ordem de compra ou venda de ações no mercado
// Contém todas as informações necessárias para processar uma negociação
type Order struct {
	ID            string         // Identificador único da ordem
	Investor      *Investor      // Investidor que criou a ordem
	Asset         *Asset         // Ativo sendo negociado
	Shares        int            // Quantidade total de ações da ordem
	PendingShares int            // Quantidade de ações ainda não executadas
	Price         float64        // Preço por ação
	OrderType     string         // Tipo da ordem: "BUY" (compra) ou "SELL" (venda)
	Status        string         // Status da ordem: "OPEN" (aberta) ou "CLOSED" (fechada)
	Transactions  []*Transaction // Lista de transações realizadas com esta ordem
}

// NewOrder cria uma nova ordem de negociação
// Parâmetros:
// - orderID: identificador único da ordem
// - investor: investidor que está criando a ordem
// - asset: ativo a ser negociado
// - shares: quantidade de ações
// - price: preço por ação
// - orderType: tipo da ordem ("BUY" ou "SELL")
func NewOrder(orderID string, investor *Investor, asset *Asset, shares int, price float64, orderType string) *Order {
	// Cria uma nova ordem com status inicial "OPEN"
	// PendingShares é inicializado com o total de ações da ordem
	// pois nenhuma ação foi negociada ainda
	return &Order{
		ID:            orderID,
		Investor:      investor,
		Asset:         asset,
		Shares:        shares,
		PendingShares: shares, // Inicialmente, todas as ações estão pendentes
		Price:         price,
		OrderType:     orderType,
		Status:        "OPEN", // Ordem começa com status aberto
		Transactions:  []*Transaction{},
	}
}

// ApplyTrade executa uma negociação parcial ou total da ordem
// Atualiza a quantidade de ações pendentes e o status da ordem
// Se todas as ações forem negociadas, a ordem é fechada
func (o *Order) ApplyTrade(tradedShares int) {
	// Garante que não serão negociadas mais ações do que o disponível
	if tradedShares > o.PendingShares {
		tradedShares = o.PendingShares
	}

	// Reduz a quantidade de ações pendentes
	o.PendingShares -= tradedShares

	// Se não houver mais ações pendentes, a ordem é fechada
	if o.PendingShares == 0 {
		o.Status = "CLOSED"
	}
}

// AddTransaction adiciona uma nova transação à lista de transações da ordem
func (o *Order) AddTransaction(t *Transaction) {
	o.Transactions = append(o.Transactions, t)
}
