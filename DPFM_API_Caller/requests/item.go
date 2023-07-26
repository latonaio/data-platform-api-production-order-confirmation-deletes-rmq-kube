package requests

type Item struct {
	ProductionOrderConfirmation      int   `json:"ProductionOrderConfirmation"`
	ProductionOrderConfirmationItem  int   `json:"ProductionOrderConfirmationItem"`
	IsMarkedForDeletion				 *bool `json:"IsMarkedForDeletion"`
}
