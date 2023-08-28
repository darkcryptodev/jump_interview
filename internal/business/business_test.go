package business

import (
	"jump_interview/internal/db"
	"jump_interview/internal/types"
	"log"
	"reflect"
	"testing"

	_ "github.com/lib/pq"
)

func TestGetUsers(t *testing.T) {
	var err error
	// Connect to DB
	db.DB, err = db.ConnectDB()
	if err != nil {
		log.Fatal("Error connecting to database:", err)
	}
	tests := []struct {
		name     string
		want     []types.User
		wantCode int
	}{
		{
			"all users",
			[]types.User{
				{
					ID:        1,
					FirstName: "Bob",
					LastName:  "Loco",
					Balance:   241817,
				},
				{
					ID:        2,
					FirstName: "Kevin",
					LastName:  "Findus",
					Balance:   49297,
				},
				{
					ID:        3,
					FirstName: "Lynne",
					LastName:  "Gwafranca",
					Balance:   82540,
				},
				{
					ID:        4,
					FirstName: "Art",
					LastName:  "Decco",
					Balance:   402758,
				},
				{
					ID:        5,
					FirstName: "Lynne",
					LastName:  "Gwistic",
					Balance:   226777,
				},
				{
					ID:        6,
					FirstName: "Polly",
					LastName:  "Ester Undawair",
					Balance:   144970,
				},
				{
					ID:        7,
					FirstName: "Oscar",
					LastName:  "Nommanee",
					Balance:   205387,
				},
				{
					ID:        8,
					FirstName: "Laura",
					LastName:  "Biding",
					Balance:   520060,
				},
				{
					ID:        9,
					FirstName: "Laura",
					LastName:  "Norda",
					Balance:   565074,
				},
				{
					ID:        10,
					FirstName: "Des",
					LastName:  "Ignayshun",
					Balance:   436180,
				},
				{
					ID:        11,
					FirstName: "Mike",
					LastName:  "Rowe-Soft",
					Balance:   818313,
				},
				{
					ID:        12,
					FirstName: "Anne",
					LastName:  "Kwayted",
					Balance:   189588,
				},
				{
					ID:        13,
					FirstName: "Wayde",
					LastName:  "Thabalanz",
					Balance:   97005,
				},
				{
					ID:        14,
					FirstName: "Dee",
					LastName:  "Mandingboss",
					Balance:   276296,
				},
				{
					ID:        15,
					FirstName: "Sly",
					LastName:  "Meedentalfloss",
					Balance:   932505,
				},
				{
					ID:        16,
					FirstName: "Stanley",
					LastName:  "Knife",
					Balance:   500691,
				},
				{
					ID:        17,
					FirstName: "Wynn",
					LastName:  "Dozeaplikayshun",
					Balance:   478333,
				},
			},
			200,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, httpCode := GetUsers()
			if httpCode != tt.wantCode {
				t.Errorf("GetUsers() http code = %v, want %v", httpCode, tt.wantCode)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetUsers() on test \"%s\" got http code %v, wanted %v", tt.name, got, tt.want)
			}
		})
	}
}

func TestCreateInvoice(t *testing.T) {
	var err error
	// Connect to DB
	db.DB, err = db.ConnectDB()
	if err != nil {
		log.Fatal("Error connecting to database:", err)
	}
	type args struct {
		invoice types.InvoiceInput
	}
	tests := []struct {
		name     string
		args     args
		wantCode int
	}{
		{
			"bad invoice",
			args{
				invoice: types.InvoiceInput{
					UserID: 199,
					Amount: 163.22,
					Label:  "Work for May",
				},
			},
			500,
		},
		{
			"new invoice",
			args{
				invoice: types.InvoiceInput{
					UserID: 1,
					Amount: 113.75,
					Label:  "Work for April",
				},
			},
			204,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if httpCode := CreateInvoice(tt.args.invoice); httpCode != tt.wantCode {
				t.Errorf("CreateInvoice() on test \"%s\" http code = %v, wanted http code %v", tt.name, httpCode, tt.wantCode)
			}
		})
	}
}

func TestCreateTransaction(t *testing.T) {
	var err error
	// Connect to DB
	db.DB, err = db.ConnectDB()
	if err != nil {
		log.Fatal("Error connecting to database:", err)
	}
	type args struct {
		transaction types.Transaction
	}
	tests := []struct {
		name     string
		args     args
		wantCode int
	}{
		{
			"no invoice found",
			args{
				transaction: types.Transaction{
					InvoiceID: 199,
					Amount:    113.75,
					Reference: "JMPINV200220117",
				},
			},
			404,
		},
		{
			"not matching amounts",
			args{
				transaction: types.Transaction{
					InvoiceID: 2,
					Amount:    13.75,
					Reference: "JMPINV200220117",
				},
			},
			400,
		},
		{
			"success",
			args{
				transaction: types.Transaction{
					InvoiceID: 2,
					Amount:    113.75,
					Reference: "JMPINV200220117",
				},
			},
			204,
		},
		{
			"invoice already paid",
			args{
				transaction: types.Transaction{
					InvoiceID: 2,
					Amount:    113.75,
					Reference: "JMPINV200220117",
				},
			},
			422,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if httpCode := CreateTransaction(tt.args.transaction); httpCode != tt.wantCode {
				t.Errorf("CreateTransaction() on test \"%s\" got an http code: %v, wanted code: %v", tt.name, httpCode, tt.wantCode)
			}
		})
	}
}
