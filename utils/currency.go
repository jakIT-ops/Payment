package utils

const (
	MNT = "MNT"
	EUR = "EUR"
	USD = "USD"
	CNY = "CNY"
)

func IsSupportedCurrency(currency string) bool {
	switch currency {
	case USD, MNT, EUR, CNY:
		return true
	}
	return false
}
