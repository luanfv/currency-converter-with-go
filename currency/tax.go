package currency

type taxJson struct {
	Rates map[string]float64 `json:"rates"`
}