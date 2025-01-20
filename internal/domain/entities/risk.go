package entities

import (
	"time"
)

type Risk struct {
	CardNumber string
	Reason     RiskReason
	Timestamp  time.Time
}

type RiskReason = string

const (
	RiskHighAmount  RiskReason = "high amount"
	RiskNotStandard RiskReason = "not standard"
)
