package entity

// OrderProcessor é responsável por processar uma transação
// atualizando as posições dos investidores e o status das ordens
type OrderProcessor struct {
	Transaction *Transaction
}

func NewOrderProcessor(transaction *Transaction) *OrderProcessor {
	return &OrderProcessor{
		Transaction: transaction,
	}
}

// Process executa todo o fluxo de processamento de uma transação:
// 1. Calcula a quantidade real de ações a serem negociadas
// 2. Atualiza as posições dos investidores
// 3. Atualiza o status das ordens
// 4. Calcula o valor total da transação
func (op *OrderProcessor) Process() {
	shares := op.calculateShares()
	op.updatePositions(shares)
	op.updateOrders(shares)
	op.Transaction.Total = float64(shares) * op.Transaction.Price
}

// calculateShares determina a quantidade real de ações que serão negociadas
// Usa o menor valor entre:
// - Quantidade solicitada na transação
// - Quantidade pendente na ordem de compra
// - Quantidade pendente na ordem de venda
func (op *OrderProcessor) calculateShares() int {
	availableShares := op.Transaction.Shares

	// Verifica se a ordem de compra tem ações pendentes suficientes
	if op.Transaction.BuyingOrder.PendingShares < availableShares {
		availableShares = op.Transaction.BuyingOrder.PendingShares
	}

	// Verifica se a ordem de venda tem ações pendentes suficientes
	if op.Transaction.SellingOrder.PendingShares < availableShares {
		availableShares = op.Transaction.SellingOrder.PendingShares
	}

	return availableShares
}

// updatePositions atualiza as posições dos investidores após a transação
// - Reduz a quantidade de ações do vendedor
// - Aumenta a quantidade de ações do comprador
func (op *OrderProcessor) updatePositions(shares int) {
	// Remove as ações da posição do vendedor (número negativo)
	op.Transaction.SellingOrder.Investor.AdjustAssetPosition(op.Transaction.SellingOrder.Asset.ID, -shares)
	// Adiciona as ações na posição do comprador (número positivo)
	op.Transaction.BuyingOrder.Investor.AdjustAssetPosition(op.Transaction.BuyingOrder.Asset.ID, shares)
}

// updateOrders atualiza o status das ordens envolvidas na transação
// Reduz a quantidade de ações pendentes em ambas as ordens
// Se uma ordem não tiver mais ações pendentes, seu status muda para "CLOSED"
func (op *OrderProcessor) updateOrders(shares int) {
	op.Transaction.BuyingOrder.ApplyTrade(shares)
	op.Transaction.SellingOrder.ApplyTrade(shares)
}
