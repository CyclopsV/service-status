package storages

import (
	"github.com/CyclopsV/service-status-skillbox/internal/sms"
	"github.com/CyclopsV/service-status-skillbox/pkg/pars"
)

type SMSStorage []*sms.SMS

func (ss *SMSStorage) Add(obj *sms.SMS) {
	*ss = append(*ss, obj)
}

func NewSMSStorage(filename string) *SMSStorage {
	smsStr := pars.ParseFile(filename)
	ss := SMSStorage{}
	for _, s := range smsStr {
		res := sms.StrToSMS(s)
		if res == nil {
			continue
		}
		ss.Add(res)
	}
	return &ss
}
