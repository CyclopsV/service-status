package storages

import (
	"github.com/CyclopsV/service-status-skillbox/internal/incident"
	"github.com/CyclopsV/service-status-skillbox/pkg/apiRequest"
	"github.com/CyclopsV/service-status-skillbox/pkg/pars"
)

type IncidentStorage []*incident.Incident

func NewIncidentStorage() *IncidentStorage {
	resp := apiRequest.Get("http://127.0.0.1:8383/accendent")
	is := IncidentStorage{}
	if resp == nil {
		return &is
	}
	if pars.JSON(&is, resp.Body) {
		return &is
	}
	deleteErrData(is)
	return &is
}
