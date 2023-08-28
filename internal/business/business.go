package business

import (
	"jump_interview/internal/db"
	"jump_interview/internal/types"
	"net/http"
)

func GetUsers() ([]types.User, int) {
	// Retrieve users from database
	users, err := db.GetUsers()
	if err != nil {
		return nil, http.StatusInternalServerError
	}

	return users, http.StatusOK
}

func CreateInvoice(invoice types.InvoiceInput) int {
	// Insert invoice into database
	err := db.CreateInvoice(invoice)
	if err != nil {
		return http.StatusInternalServerError
	}

	return http.StatusNoContent
}

func CreateTransaction(transaction types.Transaction) int {
	// Check if the invoice with the given ID exists
	invoice, err := db.GetCorrespondingInvoice(transaction)
	if err != nil {
		return http.StatusNotFound
	}
	// Check if the amount is correct
	// Convert the float32 amount to an integer (assuming the database saves amounts as cents)
	transactionAmountAsInt := int64(100 * transaction.Amount)
	if invoice.Amount != transactionAmountAsInt {
		return http.StatusBadRequest
	}
	// Check if the invoice has already been paid
	if invoice.Status == string(types.InvoiceStatusPaid) {
		return http.StatusUnprocessableEntity
	}
	// Success: Update status of invoice and increase the user's balance
	err = db.AcceptTransaction(transaction)
	if err != nil {
		return http.StatusUnprocessableEntity
	}
	return http.StatusNoContent
}
