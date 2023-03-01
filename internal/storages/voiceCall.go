package storages

import (
	"github.com/CyclopsV/service-status-skillbox/internal/voiceCall"
	"github.com/CyclopsV/service-status-skillbox/pkg/pars"
)

type VCStorage []*voiceCall.VoiceCall

func (vcs *VCStorage) Add(obj *voiceCall.VoiceCall) {
	*vcs = append(*vcs, obj)
}

func NewVCStorage(filename string) *VCStorage {
	smsStr := pars.ParseFile(filename)
	ss := VCStorage{}
	for _, s := range smsStr {
		res := voiceCall.FromSTR(s)
		if res == nil {
			continue
		}
		ss.Add(res)
	}
	return &ss
}
