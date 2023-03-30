package o2o

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// CardTemplateUpdate
// @id 136
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=卡券模板更新)
func (client *Client) CardTemplateUpdate(request *CardTemplateUpdateRequest) (response *CardTemplateUpdateResponse, err error) {
	rpcResponse := CreateCardTemplateUpdateResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type CardTemplateUpdateRequest struct {
	*api.BaseRequest

	TextImages         []map[string]any `json:"textImages,omitempty"`
	CardTemplateId     int64            `json:"cardTemplateId,omitempty"`
	ColorType          int              `json:"colorType,omitempty"`
	Type               int              `json:"type,omitempty"`
	ExpDateStr         string           `json:"expDateStr,omitempty"`
	UsingTimeType      int              `json:"usingTimeType,omitempty"`
	WeekDays           []string         `json:"weekDays,omitempty"`
	StartTimeStr       string           `json:"startTimeStr,omitempty"`
	EndTimeStr         string           `json:"endTimeStr,omitempty"`
	UserTakeLimit      int              `json:"userTakeLimit,omitempty"`
	ApplyStoreType     int              `json:"applyStoreType,omitempty"`
	Store              []string         `json:"store,omitempty"`
	UserGuide          string           `json:"userGuide,omitempty"`
	UseNotice          string           `json:"useNotice,omitempty"`
	ServicePhoneNo     string           `json:"servicePhoneNo,omitempty"`
	IconUrl            string           `json:"iconUrl,omitempty"`
	ImageUrl           string           `json:"imageUrl,omitempty"`
	Text               string           `json:"text,omitempty"`
	CenterScene        int              `json:"centerScene,omitempty"`
	CenterTitle        string           `json:"centerTitle,omitempty"`
	CenterSubTitle     string           `json:"centerSubTitle,omitempty"`
	CenterUrl          string           `json:"centerUrl,omitempty"`
	UsingScene         int              `json:"usingScene,omitempty"`
	UsingSceneName     string           `json:"usingSceneName,omitempty"`
	UsingGuide         string           `json:"usingGuide,omitempty"`
	UsingLink          string           `json:"usingLink,omitempty"`
	MarketingScene     int              `json:"marketingScene,omitempty"`
	MarketingSceneName string           `json:"marketingSceneName,omitempty"`
	MarketingGuide     string           `json:"marketingGuide,omitempty"`
	MarketingLink      string           `json:"marketingLink,omitempty"`
	ServiceType        string           `json:"serviceType,omitempty"`
	Version            int              `json:"version,omitempty"`
}

type CardTemplateUpdateResponse struct {
}

func CreateCardTemplateUpdateRequest() (request *CardTemplateUpdateRequest) {
	request = &CardTemplateUpdateRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("o2o", "1_0", "cardTemplate/update", "api")
	request.Method = api.POST
	return
}

func CreateCardTemplateUpdateResponse() (response *api.BaseResponse[CardTemplateUpdateResponse]) {
	response = api.CreateResponse[CardTemplateUpdateResponse](&CardTemplateUpdateResponse{})
	return
}
