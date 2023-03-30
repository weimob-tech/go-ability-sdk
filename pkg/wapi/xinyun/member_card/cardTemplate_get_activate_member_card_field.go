package member_card

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// CardTemplateGetActivateMemberCardField
// @id 980
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=查询会员卡开卡字段)
func (client *Client) CardTemplateGetActivateMemberCardField(request *CardTemplateGetActivateFieldRequest) (response *CardTemplateGetActivateFieldResponse, err error) {
	rpcResponse := CreateCardTemplateGetActivateFieldResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type CardTemplateGetActivateFieldRequest struct {
	*api.BaseRequest

	Type    int `json:"type,omitempty"`
	Channel int `json:"channel,omitempty"`
}

type CardTemplateGetActivateFieldResponse struct {
}

func CreateCardTemplateGetActivateFieldRequest() (request *CardTemplateGetActivateFieldRequest) {
	request = &CardTemplateGetActivateFieldRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("member_card", "1_0", "cardTemplate/getActivateMemberCardField", "api")
	request.Method = api.POST
	return
}

func CreateCardTemplateGetActivateFieldResponse() (response *api.BaseResponse[CardTemplateGetActivateFieldResponse]) {
	response = api.CreateResponse[CardTemplateGetActivateFieldResponse](&CardTemplateGetActivateFieldResponse{})
	return
}
