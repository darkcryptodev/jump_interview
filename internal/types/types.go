package types

type User struct {
	ID        int
	FirstName string
	LastName  string
	Balance   float32
}

type InvoiceInput struct {
	UserID int
	Amount float32
	Label  string
}

type Invoice struct {
	ID     int
	UserID int
	Amount int64
	Label  string
	Status string
}

type Transaction struct {
	InvoiceID int
	Amount    float32
	Reference string
}

type InvoiceStatus string

const (
	InvoiceStatusPending InvoiceStatus = "pending"
	InvoiceStatusPaid    InvoiceStatus = "paid"
)
