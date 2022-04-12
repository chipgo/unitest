package compapy

import "time"

type (
	Company struct {
		DocumentID string

		BackgroundColor string
		Color           string

		Name        string
		Symbol      string
		ClosePrices []ClosePrice
	}

	ClosePrice struct {
		Price    float64
		Timeline *time.Time
	}
)
