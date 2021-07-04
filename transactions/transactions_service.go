package transactions

import (
	"fmt"
	"regexp"
	"strconv"
)

func (b *Bookings) GetByRemittanceInfo(remittanceInfo string) (*Bookings, error) {

	var fb []Booking
	var amount float64 = 0
	var debiting float64 = 0
	exp := regexp.MustCompile(remittanceInfo)

	for _, sb := range b.Values {
		if match := exp.MatchString(sb.RemittanceInfo); match {
			fb = append(fb, sb)
			val, err := strconv.ParseFloat(sb.Amount.Value, 64)
			if err != nil {
				return nil, err
			}
			if val < 0 {
				debiting = debiting + val
			} else {
				amount = amount + val
			}
		}
	}

	filtered := Bookings{
		RemittanceInfo: remittanceInfo,
		Values:         fb,
		Count:          len(fb),
		Amount:         fmt.Sprintf("%.2f", amount),
		Debiting:       fmt.Sprintf("%.2f", debiting),
	}

	return &filtered, nil
}
