// ////////////////////////////////////////////////////////////////////////////////////////
//
// Filename: promotions.go
// File type: Go
// Author: Glenn Colucci
// Class: IFT 365
// Date: December 2, 2022
//
// Description: This file deals with the promotion records and csv file
// All structures and operations pertaining to Promotional records should be
// handled by this package.
package promotions

import (
	"strconv"
	"time"

	"github.com/IFT365/src/FinalPrj/customers"
	"github.com/IFT365/src/FinalPrj/services"
)

func PromoSelections(daysPrior string, pType string, CustomersList map[string]customers.Customer) ([]customers.Customer, error) {

	DaysPast, err := strconv.Atoi(daysPrior)
	if err != nil {
		return nil, err
	}
	DaysPast = DaysPast * -1

	now := time.Now().AddDate(0, 0, DaysPast) //

	var promoList []customers.Customer

	var cDate time.Time

	for _, cust := range CustomersList {
		if pType == services.OILCHANGE {
			if cust.LastOCPromo.PromoDate == "" {
				promoList = append(promoList, cust)
			} else {
				cDate, _ = time.Parse("2006-01-02", cust.LastOCPromo.PromoDate)
				if cDate.Unix() < now.Unix() {
					promoList = append(promoList, cust)
				}
			}
		} else {
			if cust.LastCWPromo.PromoDate == "" {
				promoList = append(promoList, cust)
			} else {
				cDate, _ = time.Parse("2006-01-02", cust.LastCWPromo.PromoDate)
				if cDate.Unix() < now.Unix() {
					promoList = append(promoList, cust)
				}
			}
		}
	}

	return promoList, nil
}
