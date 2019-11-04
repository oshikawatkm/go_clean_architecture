package views

type Response struct {
	Code int         `json:"code"`
	Body interface{} `json:"Body`
}

type User struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	Description string `json:"description"`
	Created_at  string `json:"created_at"`
	Updated_at  string `json:"updated_at"`
}

type Setting struct {
	ID         string `json:"id"`
	User       string `json:"user"`
	Created_at string `json:"created_at"`
	Updated_at string `json:"updated_at"`
}

type Transaction struct {
	ID           string  `json:"id"`
	Amount       string  `json:"amount"`
	Sender_id    *User   `json:"sender_id"`
	Recepient_id string  `json:"recepient_id"`
	Wallet_id    *Wallet `json:"wallet"`
	Created_at   string  `json:"created_at"`
}

type Wallet struct {
	ID         string `json:"id"`
	Balance    string `json:"balance"`
	Created_at string `json:"created_at"`
	Updated_at string `json:"updated_at"`
}
