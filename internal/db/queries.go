package db

import (
	"fmt"
	"jump_interview/internal/types"
	"strconv"

	_ "github.com/lib/pq"
)

func GetUsers() ([]types.User, error) {
	// Query all users
	rows, err := DB.Query("SELECT id, first_name, last_name, balance::text FROM users LIMIT 50")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer rows.Close()

	// Iterate through users: checks
	users := []types.User{}
	for rows.Next() {
		user := types.User{}
		var balance string
		err := rows.Scan(&user.ID, &user.FirstName, &user.LastName, &balance)
		if err != nil {
			return nil, err
		}

		// Check balance
		value, err := strconv.ParseFloat(balance, 64)
		if err != nil {
			return nil, err
		}
		user.Balance = float32(value)

		users = append(users, user)
	}

	return users, nil
}

func CreateInvoice(invoice types.InvoiceInput) error {
	// Convert the float32 amount to an integer (assuming the database saves amounts as cents)
	amountAsInt := int64(100 * invoice.Amount)
	// Insert new invoice
	_, err := DB.Exec(
		"INSERT INTO invoices (user_id, label, amount) VALUES ($1, $2, $3::bigint)",
		invoice.UserID, invoice.Label, amountAsInt,
	)
	if err != nil {
		fmt.Println("Error creating invoice:", err)
		return err
	}
	return nil
}

func GetCorrespondingInvoice(transaction types.Transaction) (*types.Invoice, error) {
	var invoice types.Invoice
	err := DB.QueryRow(
		"SELECT id, user_id, status, label, amount FROM invoices WHERE id = $1",
		transaction.InvoiceID,
	).Scan(&invoice.ID, &invoice.UserID, &invoice.Status, &invoice.Label, &invoice.Amount)
	if err != nil {
		fmt.Println("Error corresponding invoice not found:", err)
		return nil, err
	}

	return &invoice, nil
}

func AcceptTransaction(transaction types.Transaction) error {
	// Must be in 1 commit: set invoice as paid if and only if balance is updated
	tx, err := DB.Begin()
	if err != nil {
		fmt.Println("Error starting transaction:", err)
		return err
	}
	defer func() {
		if err != nil {
			tx.Rollback()
			return
		}
		err = tx.Commit()
	}()

	// Increase the user's balance corresponding to the invoice
	transactionAmountAsInt := int64(100 * transaction.Amount)
	_, err = tx.Exec(
		"UPDATE users SET balance = balance + $1 WHERE id = (SELECT user_id FROM invoices WHERE id = $2)",
		transactionAmountAsInt, transaction.InvoiceID,
	)
	if err != nil {
		fmt.Println("Error updating user balance:", err)
		return err
	}

	// Update the invoice status to 'paid'
	_, err = tx.Exec(
		"UPDATE invoices SET status = 'paid' WHERE id = $1",
		transaction.InvoiceID,
	)
	if err != nil {
		fmt.Println("Error updating invoice status:", err)
		return err
	}

	// TO DO : Insert the transaction record
	/*
		_, err = DB.Exec(
			"INSERT INTO transactions (invoice_id, amount, reference) VALUES ($1, $2, $3)",
			transaction.InvoiceID, transaction.Amount, transaction.Reference,
		)
		if err != nil {
			fmt.Println("Error creating transaction:", err)
			return err
		}
	*/

	return nil
}
