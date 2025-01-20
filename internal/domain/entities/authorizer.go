package entities

import "time"

type Authorizer struct {
	CardNumber string
	Amount     float64
	Currency   string
	Merchant   string
	Timestamp  time.Time
}
