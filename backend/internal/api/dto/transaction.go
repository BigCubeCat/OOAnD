package dto

type ClientTransactionRequestDTO struct {
	ID       int     `json:"id"`
	Receiver int     `json:"receiver_id"`
	Sender   int     `json:"sender_id"`
	Summary  float64 `json:"summary"`
	State    int
}
