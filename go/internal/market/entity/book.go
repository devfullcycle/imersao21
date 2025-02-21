package entity

import "sync"

// Book representa o livro de ofertas do mercado financeiro
// Responsável por gerenciar e executar as ordens de compra e venda de ativos
type Book struct {
	Orders          []*Order        // Lista de todas as ordens
	Transactions    []*Transaction  // Lista de transações realizadas
	IncomingOrders  chan *Order     // Canal para receber novas ordens
	ProcessedOrders chan *Order     // Canal para enviar ordens processadas
	Wg              *sync.WaitGroup // WaitGroup para sincronização de goroutines
}

// NewBook cria uma nova instância do livro de ofertas
// Recebe canais para ordens de entrada e saída, além do WaitGroup para sincronização
func NewBook(incomingOrders chan *Order, processedOrders chan *Order, wg *sync.WaitGroup) *Book {
	return &Book{
		Orders:          []*Order{},
		Transactions:    []*Transaction{},
		IncomingOrders:  incomingOrders,
		ProcessedOrders: processedOrders,
		Wg:              wg,
	}
}

// orderQueue é uma fila de ordens implementada como slice de ponteiros para Order
type orderQueue []*Order

// Add adiciona uma nova ordem ao final da fila
func (oq *orderQueue) Add(order *Order) {
	*oq = append(*oq, order)
}

// GetNextOrder retorna e remove a próxima ordem da fila (FIFO)
// Retorna nil se a fila estiver vazia
func (oq *orderQueue) GetNextOrder() *Order {
	if len(*oq) == 0 {
		return nil
	}
	order := (*oq)[0]
	*oq = (*oq)[1:]
	return order
}

// Trade é o método principal que processa continuamente as ordens recebidas
// Mantém filas separadas para ordens de compra e venda por ativo
// Tenta realizar o matching das ordens assim que são recebidas
func (b *Book) Trade() {
	// Mapas para manter filas separadas de ordens de compra e venda por ativo
	// A chave é o ID do ativo e o valor é uma fila de ordens
	buyOrders := make(map[string]*orderQueue)
	sellOrders := make(map[string]*orderQueue)

	// Loop infinito que processa ordens conforme chegam no canal
	for order := range b.IncomingOrders {
		asset := order.Asset.ID

		// Inicializa as filas de compra e venda para o ativo se não existirem
		if buyOrders[asset] == nil {
			buyOrders[asset] = &orderQueue{}
		}

		if sellOrders[asset] == nil {
			sellOrders[asset] = &orderQueue{}
		}

		// Direciona a ordem para o matching apropriado:
		// - Se for compra, tenta matching com ordens de venda existentes
		// - Se for venda, tenta matching com ordens de compra existentes
		if order.OrderType == "BUY" {
			b.tryMatch(order, sellOrders[asset], buyOrders[asset])
		} else {
			b.tryMatch(order, buyOrders[asset], sellOrders[asset])
		}
	}
}

// tryMatch tenta realizar o matching de uma nova ordem com as ordens disponíveis
// Parâmetros:
// - newOrder: nova ordem a ser processada
// - availableOrders: fila de ordens disponíveis do lado oposto (compra/venda)
// - pendingOrders: fila onde a nova ordem será adicionada se não for totalmente executada
func (b *Book) tryMatch(newOrder *Order, availableOrders, pendingOrders *orderQueue) {
	// Loop continua enquanto houver ordens disponíveis para matching
	// e a nova ordem ainda não foi totalmente executada
	for {
		// Obtém a próxima ordem disponível para matching
		potentialMatch := availableOrders.GetNextOrder()
		if potentialMatch == nil {
			break // Não há mais ordens disponíveis para matching
		}

		// Verifica se os preços são compatíveis
		// Se não forem, devolve a ordem para a fila e encerra
		if !b.pricesMatch(newOrder, potentialMatch) {
			availableOrders.Add(potentialMatch)
			break
		}

		// Se ainda há ações pendentes para negociar
		if potentialMatch.PendingShares > 0 {
			// Cria e processa a transação entre as ordens
			matchedTransaction := b.createTransaction(newOrder, potentialMatch)
			b.processTransaction(matchedTransaction)

			// Se a ordem do matching ainda tem ações pendentes
			// coloca ela de volta na fila para futuros matchings
			if potentialMatch.PendingShares > 0 {
				availableOrders.Add(potentialMatch)
			}

			// Se a nova ordem foi totalmente executada, encerra o processo
			if newOrder.PendingShares == 0 {
				break
			}
		}
	}

	// Se a nova ordem ainda tem ações pendentes
	// adiciona ela na fila de ordens pendentes
	if newOrder.PendingShares > 0 {
		pendingOrders.Add(newOrder)
	}
}

// pricesMatch verifica se o preço entre duas ordens é compatível para execução
// Para ordens de compra: preço da ordem de venda deve ser menor ou igual ao preço de compra
// Para ordens de venda: preço da ordem de compra deve ser maior ou igual ao preço de venda
func (b *Book) pricesMatch(order, matchOrder *Order) bool {
	if order.OrderType == "BUY" {
		return matchOrder.Price <= order.Price
	}
	return matchOrder.Price >= order.Price
}

// createTransaction cria uma nova transação a partir de duas ordens compatíveis
// Determina a quantidade de ações a serem negociadas com base na menor quantidade disponível
// entre as duas ordens
func (b *Book) createTransaction(incomingOrder, matchedOrder *Order) *Transaction {
	var buyOrder, sellOrder *Order

	// Identifica qual é a ordem de compra e qual é a de venda
	// para criar a transação na ordem correta
	if incomingOrder.OrderType == "BUY" {
		buyOrder, sellOrder = incomingOrder, matchedOrder
	} else {
		buyOrder, sellOrder = matchedOrder, incomingOrder
	}

	// Determina a quantidade de ações a serem negociadas
	// Usa o menor valor entre as ações pendentes das duas ordens
	shares := incomingOrder.PendingShares
	if matchedOrder.PendingShares < shares {
		shares = matchedOrder.PendingShares
	}

	return NewTransaction(sellOrder, buyOrder, shares, matchedOrder.Price)
}

// recordTransaction registra uma transação concluída no livro e nas ordens envolvidas
func (b *Book) recordTransaction(transaction *Transaction) {
	b.Transactions = append(b.Transactions, transaction)
	transaction.BuyingOrder.Transactions = append(transaction.BuyingOrder.Transactions, transaction)
	transaction.SellingOrder.Transactions = append(transaction.SellingOrder.Transactions, transaction)
}

// processTransaction finaliza o processamento de uma transação
// - Processa a transação
// - Registra no livro
// - Envia as ordens processadas para o canal de saída
// - Marca a conclusão no WaitGroup
func (b *Book) processTransaction(transaction *Transaction) {
	defer b.Wg.Done()
	transaction.Process()
	b.recordTransaction(transaction)
	b.ProcessedOrders <- transaction.BuyingOrder
	b.ProcessedOrders <- transaction.SellingOrder
}
