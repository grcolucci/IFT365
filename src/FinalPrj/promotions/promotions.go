package promotions

import (
	"fmt"
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

	fmt.Println("Days", daysPrior)
	fmt.Println("type", pType)

	for _, cust := range CustomersList {
		fmt.Println("OC ", cust.LastOCPromo.PromoDate)
		fmt.Println("CW ", cust.LastCWPromo.PromoDate)

		if pType == services.OILCHANGE {
			if cust.LastOCPromo.PromoDate == "" {
				promoList = append(promoList, cust)
				//			cust.LastCWPromo.PromoDate = time.Now().AddDate(0, 0, -300).Format("2006-01-02")
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
