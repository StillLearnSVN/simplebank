package util

const (
	USD = "USD"
	EUR = "EUR"
	IDR = "IDR"
)



// isSupportedCurrency returns true if the currency is supported
func IsSupportedCurrency(currency string) bool{
	switch currency {
	case USD, EUR, IDR:
		return true
	}
	return false
}