package entity

// Asset representa um ativo financeiro negociável no mercado
// Por exemplo: ações de uma empresa, títulos, etc.
type Asset struct {
	ID           string // Identificador único do ativo (ex: código da ação - PETR4)
	Name         string // Nome do ativo (ex: Petrobras PN)
	MarketVolume int    // Volume total de ações disponíveis no mercado
}

// NewAsset cria uma nova instância de um ativo financeiro
// Parâmetros:
// - id: código identificador do ativo
// - name: nome do ativo
// - marketVolume: volume total disponível no mercado
func NewAsset(id, name string, marketVolume int) *Asset {
	return &Asset{
		ID:           id,
		Name:         name,
		MarketVolume: marketVolume,
	}
}
