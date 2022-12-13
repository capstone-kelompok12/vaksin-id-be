package response

type VaccinesResponse struct {
	ID    string
	Name  string
	Dose  int
	Stock int
}
type VaccinesStockResponse struct {
	Name  string
	Stock int
}
