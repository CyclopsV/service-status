package storages

import (
	"encoding/json"
	"github.com/CyclopsV/service-status-skillbox/internal/mms"
	"io"
	"log"
	"net/http"
)

type MMSStorage []*mms.MMS

func (ms *MMSStorage) Add(obj *mms.MMS) {
	*ms = append(*ms, obj)
}

func NewMMSStorage() (*MMSStorage, error) {
	url := "http://127.0.0.1:8383/mms"
	resp, err := http.Get(url)
	if err != nil || resp.StatusCode != http.StatusOK {
		log.Fatalf("Ошибка получения данных MMS:\n%v", err)
		return nil, err
	}
	content, err := io.ReadAll(resp.Body)
	var buf []map[string]interface{}
	if err = json.Unmarshal(content, &buf); err != nil {
		log.Fatalf("Ошибка чтения данных MMS:\n%v", err)
		return nil, err
	}
	ms := MMSStorage{}
	for _, el := range buf {
		m := mms.New(el["country"].(string), el["provider"].(string), el["bandwidth"].(string), el["response_time"].(string))
		if m != nil {
			ms = append(ms, m)
		}
	}

	return &ms, nil
}
