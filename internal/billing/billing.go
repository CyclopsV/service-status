package billing

import (
	"github.com/CyclopsV/service-status-skillbox/pkg/pars"
	"log"
)

type BillingData struct {
	createCustomer bool
	purchase       bool
	payout         bool
	recurring      bool
	fraudControl   bool
	checkoutPage   bool
}

func New(path string) *BillingData {
	content := pars.ReadFile(path)
	if len(content) != 6 {
		log.Fatalf("Не верная длина битовой маски")
	}
	return &BillingData{
		createCustomer: check(content[5]),
		purchase:       check(content[4]),
		payout:         check(content[3]),
		recurring:      check(content[2]),
		fraudControl:   check(content[1]),
		checkoutPage:   check(content[0]),
	}
}

func check(status byte) bool {
	return status == byte('1')
}
