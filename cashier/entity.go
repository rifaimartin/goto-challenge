package cashier

import "time"

// Cashier struct
type Cashier struct {
	ID             int
	Name           string
	Passcode       string
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
