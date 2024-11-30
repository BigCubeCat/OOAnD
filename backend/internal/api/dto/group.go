package dto

type CreateGroupDTO struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Members     []int  `json:"members"`
}
