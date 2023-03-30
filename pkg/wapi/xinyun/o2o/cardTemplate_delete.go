package o2o

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// CardTemplateDelete
// @id 128
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=删除卡模板)
func (client *Client) CardTemplateDelete(request *CardTemplateDeleteRequest) (response *CardTemplateDeleteResponse, err error) {
	rpcResponse := CreateCardTemplateDeleteResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type CardTemplateDeleteRequest struct {
	*api.BaseRequest
}

type CardTemplateDeleteResponse struct {
}

func CreateCardTemplateDeleteRequest() (request *CardTemplateDeleteRequest) {
	request = &CardTemplateDeleteRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("o2o", "1_0", "cardTemplate/delete", "api")
	request.Method = api.POST
	return
}

func CreateCardTemplateDeleteResponse() (response *api.BaseResponse[CardTemplateDeleteResponse]) {
	response = api.CreateResponse[CardTemplateDeleteResponse](&CardTemplateDeleteResponse{})
	return
}
