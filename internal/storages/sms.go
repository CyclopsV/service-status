package storages

import (
	"github.com/CyclopsV/service-status-skillbox/internal/sms"
	"github.com/CyclopsV/service-status-skillbox/pkg/pars"
	"sort"
)

type SMSStorage []*sms.SMS

func (ss *SMSStorage) Add(obj *sms.SMS) {
	*ss = append(*ss, obj)
}

func NewSMSStorage(filename string) (*SMSStorage, error) {
	smsStr, err := pars.FileToStr(filename)
	if err != nil {
		return nil, err
	}
	ss := SMSStorage{}
	for _, s := range smsStr {
		res := sms.FromSTR(s)
		if res == nil {
			continue
		}
		ss.Add(res)
	}
	return &ss, nil
}

func (ss SMSStorage) SortCountry() {
	sortF := func(i, j int) bool {
		return ss[i].Country < ss[j].Country
	}
	sort.SliceStable(ss, sortF)
}

func (ss SMSStorage) SortProvider() {
	sortF := func(i, j int) bool {
		return ss[i].Provider < ss[j].Provider
	}
	sort.SliceStable(ss, sortF)
}
