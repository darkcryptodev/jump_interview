package db

import (
	"fmt"
	"jump/jump_interview/internal/types"
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

func CreateInvoice(invoice types.Invoice) error {
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

func CreateTransaction(transaction types.Transaction) error {
	// Check if the invoice with the given ID and amount exists
	var invoiceID int
	var invoiceAmount float64
	// Convert the float32 amount to an integer (assuming the database saves amounts as cents)
	transactionAmountAsInt := int64(100 * transaction.Amount)
	err := DB.QueryRow(
		"SELECT id, amount FROM invoices WHERE id = $1 AND amount = $2::bigint",
		transaction.InvoiceID, transactionAmountAsInt,
	).Scan(&invoiceID, &invoiceAmount)
	if err != nil {
		fmt.Println("Error corresponding invoice not found:", err)
		return err
	}

	// Increase the user's balance corresponding to the invoice
	_, err = DB.Exec(
		"UPDATE users SET balance = balance + $1::bigint WHERE id = (SELECT user_id FROM invoices WHERE id = $2)",
		transactionAmountAsInt, transaction.InvoiceID,
	)
	if err != nil {
		fmt.Println("Error updating user balance:", err)
		return err
	}

	// Insert the transaction record
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
