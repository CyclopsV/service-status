package storages

import (
	"encoding/json"
	"github.com/CyclopsV/service-status-skillbox/internal/support"
	"io"
	"net/http"
)

type SupportStorage []*support.Support

func NewSupportStorage() *SupportStorage {
	url := "http://127.0.0.1:8383/support"
	resp, err := http.Get(url)
	ss := SupportStorage{}
	if err != nil || resp.StatusCode != http.StatusOK {
		return &ss
	}
	content, err := io.ReadAll(resp.Body)
	if err = json.Unmarshal(content, &ss); err != nil {
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
