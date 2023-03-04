package storages

import (
	"github.com/CyclopsV/service-status-skillbox/internal/support"
	"github.com/CyclopsV/service-status-skillbox/pkg/apiRequest"
	"github.com/CyclopsV/service-status-skillbox/pkg/pars"
)

type SupportStorage []*support.Support

func NewSupportStorage() *SupportStorage {
	resp := apiRequest.Get("http://127.0.0.1:8383/support")
	ss := SupportStorage{}
	if resp == nil {
		return &ss
	}
	if pars.JSON(&ss, resp.Body) {
		return &ss
	}
	ss.deleteErrData()
	return &ss
}

func (ss *SupportStorage) deleteErrData() {
	for i, el := range *ss {
		if !el.Check() {
			ss.drop(i)
		}
	}
}

func (ss *SupportStorage) drop(i int) {
	(*ss)[i] = (*ss)[len(*ss)-1]
	*ss = (*ss)[:len(*ss)-1]
}
