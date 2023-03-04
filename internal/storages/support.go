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
	if !pars.JSON(&ss, resp.Body) {
		return &ss
	}
	deleteErrData(ss)
	return &ss
}
