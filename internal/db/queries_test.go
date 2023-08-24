package db

import (
	"jump/jump_interview/internal/types"
	"log"
	"reflect"
	"testing"

	_ "github.com/lib/pq"
)

func TestGetUsers(t *testing.T) {
	var err error
	DB, err = ConnectDB()
	if err != nil {
		log.Fatal("Error connecting to database:", err)
	}
	tests := []struct {
		name    string
		want    []types.User
		wantErr bool
	}{
		{
			"first user",
			[]types.User{
				{
					ID:        1,
					FirstName: "Bob",
					LastName:  "Loco",
					Balance:   241817,
				},
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetUsers()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetUsers() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got[0], tt.want[0]) {
				t.Errorf("GetUsers() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCreateInvoice(t *testing.T) {
	var err error
	DB, err = ConnectDB()
	if err != nil {
		log.Fatal("Error connecting to database:", err)
	}
	type args struct {
		invoice types.Invoice
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"new invoice",
			args{
				invoice: types.Invoice{
					UserID: 1,
					Amount: 113.75,
					Label:  "Work for April",
				},
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := CreateInvoice(tt.args.invoice); (err != nil) != tt.wantErr {
				t.Errorf("CreateInvoice() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCreateTransaction(t *testing.T) {
	var err error
	DB, err = ConnectDB()
	if err != nil {
		log.Fatal("Error connecting to database:", err)
	}
	type args struct {
		transaction types.Transaction
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"new invoice",
			args{
				transaction: types.Transaction{
					InvoiceID: 2,
					Amount:    113.75,
					Reference: "JMPINV200220117",
				},
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := CreateTransaction(tt.args.transaction); (err != nil) != tt.wantErr {
				t.Errorf("CreateTransaction() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
