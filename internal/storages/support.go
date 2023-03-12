package storages

import (
	"fmt"
	"github.com/CyclopsV/service-status-skillbox/internal/support"
	"github.com/CyclopsV/service-status-skillbox/pkg/apiRequest"
	"github.com/CyclopsV/service-status-skillbox/pkg/pars"
)

type SupportStorage []*support.Support

func NewSupportStorage() (*SupportStorage, error) {
	resp := apiRequest.Get("http://127.0.0.1:8383/support")
	ss := SupportStorage{}
	if resp == nil {
		return &ss, fmt.Errorf("ошибка получения данных")
	}
	if err := pars.JSON(&ss, resp.Body); err != nil {
		return &ss, err
	}
	deleteErrData(ss)
	return &ss, nil
}
