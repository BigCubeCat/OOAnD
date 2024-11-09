package db

type User struct {
	SerialID                int                        `gorm:"primaryKey"            json:"id"`
	TelegramID              int                        `                             json:"tg"`
	Handle                  string                     `                             json:"handle"`
	Username                string                     `                             json:"username"`
	Email                   string                     `                             json:"email"`
	Token                   string                     `                             json:"-"`
	PriorityPaymentMethodID int                        `                             json:"method_id"`
	Avatar                  string                     `                             json:"avatar"`
	PaymentMethods          []PaymentMethod            `gorm:"foreignKey:UserID"     json:"payment_method"`
	TransactionsAsReceiver  []ClientTransactionRequest `gorm:"foreignKey:Receiver"`
	TransactionsAsSender    []ClientTransactionRequest `gorm:"foreignKey:Sender"`
	Bills                   []Bill                     `gorm:"foreignKey:Owner"`
	Groups                  []Group                    `gorm:"many2many:user_groups"`
}

type PaymentMethod struct {
	ID         int `gorm:"primaryKey"`
	UserID     int
	Name       string
	Requisites string
	Type       string
}

type Bill struct {
	ID            int            `gorm:"primaryKey"        json:"id"`
	Owner         int            `                         json:"owner"`
	Name          string         `                         json:"name"`
	BillPositions []BillPosition `gorm:"foreignKey:IDBill" json:"positions"`
}

type BillPosition struct {
	ID           int     `gorm:"primaryKey"    json:"position_id"`
	IDBill       int     `gorm:"foreignKey:ID" json:"bill_id"`
	Name         string  `                     json:"position_name"`
	WhoPaid      int     `                     json:"who_paid"`
	FromWhomPaid int     `                     json:"from_whom_paid"`
	Amount       float64 `                     json:"amount"`
}

type ClientTransactionRequest struct {
	ID       int     `gorm:"primaryKey" json:"id"`
	Receiver int     `                  json:"receiver_id"`
	Sender   int     `                  json:"sender_id"`
	Summary  float64 `                  json:"summary"`
	State    string  `                  json:"state"`
}

type Group struct {
	ID          int    `gorm:"primaryKey" json:"group_id"`
	Name        string `                  json:"group_name"`
	Members     int    `                  json:"member"`
	Description string `                  json:"description"`
}
