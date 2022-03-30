package cashier

import "time"

//cashierFormatter struct
type cashierFormatter struct {
	ID         		int    `json:"id"`
	Name       	    string `json:"name"`
}

type createCashierFormatter struct {
	ID         		int    `json:"id"`
	Name       	    string `json:"name"`
	Passcode        string `json:"passcode"`
	CreatedAt       string `json:"createdAt"`
	UpdatedAt       string `json:"updatedAt"`
}

//cashierFormatter struct
type listCashierFormatter struct {
	Cashiers        interface{}    `json:"cashiers"`
	Meta interface{} `json:"meta"`
}

// Meta Strucut
type Meta struct {
	Total int64 `json:"total"`
	Limit    int    `json:"limit"`
	Skip  int `json:"skip" `
}

//Formatcashier function for fomating data
func Formatcashier(cashier Cashier) cashierFormatter {
	formatter := cashierFormatter{
		ID:         	cashier.ID,
		Name:       	cashier.Name,
	}

	return formatter
}

//Formatcashier function for fomating data
func FormatCreateCashier(cashier Cashier) createCashierFormatter {
	formatter := createCashierFormatter{
		ID:         	cashier.ID,
		Name:       	cashier.Name,
		Passcode:       cashier.Passcode,
		CreatedAt:		time.Now().Format(time.RFC3339),
		UpdatedAt:		time.Now().Format(time.RFC3339),
	}

	return formatter
}

//Formatcashier function for fomating data
func ListFormatCashier(cashiers []Cashier, limit int, skip int, count int64) listCashierFormatter {

	cashiersFormatter := []cashierFormatter{}
	//
	for _, cashier := range cashiers {
		cashierFormatter := Formatcashier(cashier)
		cashiersFormatter = append(cashiersFormatter, cashierFormatter)
	}

	meta := Meta{
		Total: count,
		Limit: limit,
		Skip:  skip,
	}

	formatter := listCashierFormatter{
		Cashiers:         	cashiersFormatter,
		Meta: meta,
	}

	return formatter
}
