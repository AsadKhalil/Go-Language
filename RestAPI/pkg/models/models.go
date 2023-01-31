package models

// """This code defines a struct called "Product"
// The struct has five fields: ID, Name, Description, Price and Quantity.
// Each field is of a specific type: int, string, float64,
// and each field is also tagged with a json:"fieldname" which is used in encoding/decoding of json.
type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Quantity    int     `json:"quantity"`
}

// """This code defines a struct called "USER"
// The struct has four fields: ID, Name, Email, Password
// Each field is of a specific type: int, string
// and each field is also tagged with a json:"fieldname" which is used in encoding/decoding of json.
type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// """This code defines a struct called "CreditCard"
// The struct has four fields: ID, UserID, CardNumber, ExpiryMonth, ExpiryYear ,CVV and NameONCard
// Each field is of a specific type: int string
// and each field is also tagged with a json:"fieldname" which is used in encoding/decoding of json.
type CreditCard struct {
	ID          int    `json:"id"`
	UserID      int    `json:"user_id"`
	CardNumber  string `json:"card_number"`
	ExpiryMonth string `json:"expiry_month"`
	ExpiryYear  string `json:"expiry_year"`
	CVV         string `json:"cvv"`
	NameOnCard  string `json:"name_on_card"`
}
