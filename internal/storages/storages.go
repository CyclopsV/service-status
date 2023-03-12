package storages

import "github.com/CyclopsV/service-status-skillbox/internal/billing"

type ResultSetT struct {
	SMS       []SMSStorage                 `json:"sms"`
	MMS       []MMSStorage                 `json:"mms"`
	VoiceCall VCStorage                    `json:"voice_call"`
	Email     map[string][]providerStorage `json:"email"`
	Billing   billing.BillingData          `json:"billing"`
	Support   []int                        `json:"support"`
	Incidents IncidentStorage              `json:"incident"`
}

type ResultT struct {
	Status bool        `json:"status"` // True, если все этапы сбора данных прошли успешно, False во всех остальных случаях
	Data   *ResultSetT `json:"data"`   // Заполнен, если все этапы сбора  данных прошли успешно, nil во всех остальных случаях
	Error  []string    `json:"errors"` // Пустая строка, если все этапы сбора данных прошли успешно, в случае ошибки заполнено текстом ошибки
}

func GetResultData() ResultT {
	data := ResultSetT{}
	status := ResultT{
		Status: true,
		Data:   &data,
		Error:  nil,
	}
	errs := []string{}
	smsData, err := smsDataF()
	if err != nil {
		errs = append(errs, err.Error())
	}
	data.SMS = smsData
	mmsData, err := mmsDataF()
	if err != nil {
		errs = append(errs, err.Error())
	}
	data.MMS = mmsData
	vcData, err := vcDataF()
	if err != nil {
		errs = append(errs, err.Error())
	}
	data.VoiceCall = vcData
	emailData, err := emailDataF()
	if err != nil {
		errs = append(errs, err.Error())
	}
	data.Email = emailData
	billingData, err := billingDataF()
	if err != nil {
		errs = append(errs, err.Error())
	}
	data.Billing = billingData
	supportData, err := supportDataF()
	if err != nil {
		errs = append(errs, err.Error())
	}
	data.Support = supportData
	incidentData, err := incidentDataF()
	if err != nil {
		errs = append(errs, err.Error())
	}
	data.Incidents = incidentData
	if len(errs) > 0 {
		status.Error = errs
		status.Status = false
	}
	return status
}

func smsDataF() ([]SMSStorage, error) {
	smsPath := "../skillbox-diploma/sms.data"
	smsStorage, err := NewSMSStorage(smsPath)
	if err != nil {
		return nil, err
	}
	sortedProvider := make(SMSStorage, len(*smsStorage))
	smsStorage.SortProvider()
	copy(sortedProvider, *smsStorage)
	smsStorage.SortCountry()
	sortedCountry := make(SMSStorage, len(*smsStorage))
	copy(sortedCountry, *smsStorage)
	return []SMSStorage{sortedProvider, sortedCountry}, nil
}

func mmsDataF() ([]MMSStorage, error) {
	mmsStorage, err := NewMMSStorage()
	if err != nil {
		return nil, err
	}
	sortedProvider := make(MMSStorage, len(*mmsStorage))
	mmsStorage.SortProvider()
	copy(sortedProvider, *mmsStorage)
	sortedCountry := make(MMSStorage, len(*mmsStorage))
	mmsStorage.SortCountry()
	copy(sortedCountry, *mmsStorage)
	return []MMSStorage{sortedProvider, sortedCountry}, nil
}

func vcDataF() (VCStorage, error) {
	vcPath := "../skillbox-diploma/voice.data"
	vcData, err := NewVCStorage(vcPath)
	return *vcData, err
}

func emailDataF() (map[string][]providerStorage, error) {
	emailPath := "../skillbox-diploma/email.data"
	emailStorage, err := NewEmailStorage(emailPath)
	if err != nil {
		return nil, err
	}
	catalogEmailByCountry := emailStorage.catalogingByCountry()
	result := map[string][]providerStorage{}
	for country, emails := range catalogEmailByCountry {
		providers := emails.createStatisticProviders()
		providers.sort()
		topsProviders := providers.BestAndWorst()
		result[country] = topsProviders
	}
	return result, nil
}

func billingDataF() (billing.BillingData, error) {
	billingPath := "../skillbox-diploma/billing.data"
	billingData, err := billing.New(billingPath)
	return *billingData, err
}

func supportDataF() ([]int, error) {
	supportData, err := NewSupportStorage()
	if err != nil {
		return nil, err
	}
	loadStatus, waitTime := supportData.CurrentLoad()
	return []int{loadStatus, waitTime}, nil
}

func incidentDataF() (IncidentStorage, error) {
	incidentData, err := NewIncidentStorage()
	if err != nil {
		return nil, err
	}
	return *incidentData, nil
}
