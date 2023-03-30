package member_card

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// CardTemplateGetMemberTemplate
// @id 976
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=查询会员卡基本信息)
func (client *Client) CardTemplateGetMemberTemplate(request *CardTemplateGetMemberTemplateRequest) (response *CardTemplateGetMemberTemplateResponse, err error) {
	rpcResponse := CreateCardTemplateGetMemberTemplateResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type CardTemplateGetMemberTemplateRequest struct {
	*api.BaseRequest

	Type int `json:"type,omitempty"`
}

type CardTemplateGetMemberTemplateResponse struct {
}

func CreateCardTemplateGetMemberTemplateRequest() (request *CardTemplateGetMemberTemplateRequest) {
	request = &CardTemplateGetMemberTemplateRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("member_card", "1_0", "cardTemplate/getMemberTemplate", "api")
	request.Method = api.POST
	return
}

func CreateCardTemplateGetMemberTemplateResponse() (response *api.BaseResponse[CardTemplateGetMemberTemplateResponse]) {
	response = api.CreateResponse[CardTemplateGetMemberTemplateResponse](&CardTemplateGetMemberTemplateResponse{})
	return
}
