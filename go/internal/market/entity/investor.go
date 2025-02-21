package entity

// Investor representa um investidor no sistema
// Mantém o registro de todas as posições de ativos que o investidor possui
type Investor struct {
	ID            string                   // Identificador único do investidor
	Name          string                   // Nome do investidor
	AssetPosition []*InvestorAssetPosition // Lista de posições em ativos do investidor
}

// NewInvestor cria um novo investidor
// Parâmetro:
// - id: identificador único do investidor
func NewInvestor(id string) *Investor {
	return &Investor{
		ID:            id,
		AssetPosition: []*InvestorAssetPosition{},
	}
}

// AddAssetPosition adiciona uma nova posição de ativo ao portfólio do investidor
func (i *Investor) AddAssetPosition(assetPosition *InvestorAssetPosition) {
	i.AssetPosition = append(i.AssetPosition, assetPosition)
}

// AdjustAssetPosition ajusta a quantidade de ações de um ativo específico
// Se o investidor não possui o ativo, cria uma nova posição
// Se já possui, atualiza a quantidade de ações
// Parâmetros:
// - assetID: identificador do ativo
// - qtdShares: quantidade de ações a adicionar (ou subtrair se negativo)
func (i *Investor) AdjustAssetPosition(assetID string, qtdShares int) {
	assetPosition := i.GetAssetPosition(assetID)
	if assetPosition == nil {
		i.AssetPosition = append(i.AssetPosition, NewInvestorAssetPosition(assetID, qtdShares))
	} else {
		assetPosition.AddShares(qtdShares)
	}
}

// GetAssetPosition busca a posição de um ativo específico do investidor
// Retorna nil se o investidor não possui o ativo
// Parâmetro:
// - assetID: identificador do ativo
func (i *Investor) GetAssetPosition(assetID string) *InvestorAssetPosition {
	for _, assetPosition := range i.AssetPosition {
		if assetPosition.AssetID == assetID {
			return assetPosition
		}
	}
	return nil
}

// InvestorAssetPosition representa a posição (quantidade) de um ativo específico
// que um investidor possui em sua carteira
type InvestorAssetPosition struct {
	AssetID string // Identificador do ativo
	Shares  int    // Quantidade de ações possuídas
}

// NewInvestorAssetPosition cria uma nova posição de ativo
// Parâmetros:
// - assetID: identificador do ativo
// - shares: quantidade inicial de ações
func NewInvestorAssetPosition(assetID string, shares int) *InvestorAssetPosition {
	return &InvestorAssetPosition{
		AssetID: assetID,
		Shares:  shares,
	}
}

// AddShares adiciona (ou remove, se negativo) uma quantidade de ações à posição
// Parâmetro:
// - qtd: quantidade de ações a adicionar ou remover
func (iap *InvestorAssetPosition) AddShares(qtd int) {
	iap.Shares += qtd
}
