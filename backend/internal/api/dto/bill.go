package dto

// DTO for bill position
type CreateBillPositionDTO struct {
	WhoPaid      int     `json:"who_paid"`
	FromWhomPaid int     `json:"from_whom_paid"`
	Amount       float64 `json:"amount"`
}

type CreateBillDTO struct {
	Name       string                  `json:"name"`
	Postitions []CreateBillPositionDTO `json:"postitions"`
}
