package o2o

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// CardTemplateDetails
// @id 137
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=查询卡模板详情)
func (client *Client) CardTemplateDetails(request *CardTemplateDetailsRequest) (response *CardTemplateDetailsResponse, err error) {
	rpcResponse := CreateCardTemplateDetailsResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type CardTemplateDetailsRequest struct {
	*api.BaseRequest

	CardTemplateIds []int `json:"cardTemplateIds,omitempty"`
}

type CardTemplateDetailsResponse struct {
}

func CreateCardTemplateDetailsRequest() (request *CardTemplateDetailsRequest) {
	request = &CardTemplateDetailsRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("o2o", "1_0", "cardTemplate/details", "api")
	request.Method = api.POST
	return
}

func CreateCardTemplateDetailsResponse() (response *api.BaseResponse[CardTemplateDetailsResponse]) {
	response = api.CreateResponse[CardTemplateDetailsResponse](&CardTemplateDetailsResponse{})
	return
}
