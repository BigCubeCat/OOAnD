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
	ID            int `gorm:"primaryKey"`
	Owner         int
	Name          string
	BillPositions []BillPosition `gorm:"foreignKey:IDBill"`
}

type BillPosition struct {
	ID           int `gorm:"primaryKey"`
	IDBill       int
	WhoPaid      int
	FromWhomPaid int
	Amount       float64
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
