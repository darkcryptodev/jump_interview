package types

type User struct {
	ID        int
	FirstName string
	LastName  string
	Balance   float32
}

type Invoice struct {
	UserID int
	Amount float32
	Label  string
}

type Transaction struct {
	InvoiceID int
	Amount    float32
	Reference string
}
