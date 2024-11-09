package db

type User struct {
	SerialID                int    `gorm:"primaryKey"            json:"id"`
	TelegramID              int    `                             json:"tg"`
	Handle                  string `                             json:"handle"`
	Username                string `                             json:"username"`
	Email                   string `                             json:"email"`
	Token                   string
	PriorityPaymentMethodID int
	Avatar                  string          `                             json:"avatar"`
	PaymentMethods          []PaymentMethod `gorm:"foreignKey:UserID"`
	ResultsAsReceiver       []Result        `gorm:"foreignKey:Receiver"`
	ResultsAsSender         []Result        `gorm:"foreignKey:Sender"`
	Bills                   []Bill          `gorm:"foreignKey:Owner"`
	BillPositions           []BillPosition  `gorm:"foreignKey:WhoPaid"`
	Groups                  []Group         `gorm:"many2many:user_groups"`
}

type PaymentMethod struct {
	ID         int `gorm:"primaryKey"`
	UserID     int
	Name       string
	Requisites string
	Type       string
}

type Bill struct {
	ID            int            `gorm:"primaryKey"`
	Owner         int            `                         json:"owner"`
	Name          string         `                         json:"name"`
	BillPositions []BillPosition `gorm:"foreignKey:IDBill" json:"positions"`
}

type BillPosition struct {
	ID           int `gorm:"primaryKey"`
	IDBill       int
	Name         string  `json: "name"`
	WhoPaid      int     `json:"who_paid"`
	FromWhomPaid int     `json:"from_whom_paid"`
	Amount       float64 `json:"amount"`
}

type Result struct {
	ID       int `gorm:"primaryKey"`
	Receiver int
	Sender   int
	Summary  float64
	State    int
}

type Group struct {
	ID          int `gorm:"primaryKey"`
	Name        string
	Members     int
	Description string
}
