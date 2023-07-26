package requests

type Header struct {
	ProductionOrderConfirmation     int   `json:"ProductionOrderConfirmation"`
	IsMarkedForDeletion 			*bool `json:"IsMarkedForDeletion"`
}
