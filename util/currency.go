package util

const (
	USD = "USD"
	EUR = "EUR"
	CAD = "CAD"
	GBP = "GBP"
	JPY = "JPY"
	CHF = "CHF"
	RUB = "RUB"
	CNY = "CNY"
)

func IsSupportedCurrency(currency string) bool {
	switch currency {
	case USD, EUR, CAD, GBP, JPY, CHF, RUB, CNY:
		return true
	}
	return false
}
