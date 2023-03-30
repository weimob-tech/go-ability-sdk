package o2o

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// CardTemplateList
// @id 126
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=查询卡券列表（分页）)
func (client *Client) CardTemplateList(request *CardTemplateListRequest) (response *CardTemplateListResponse, err error) {
	rpcResponse := CreateCardTemplateListResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type CardTemplateListRequest struct {
	*api.BaseRequest

	PageNum  int `json:"pageNum,omitempty"`
	PageSize int `json:"pageSize,omitempty"`
}

type CardTemplateListResponse struct {
}

func CreateCardTemplateListRequest() (request *CardTemplateListRequest) {
	request = &CardTemplateListRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("o2o", "1_0", "cardTemplate/list", "api")
	request.Method = api.POST
	return
}

func CreateCardTemplateListResponse() (response *api.BaseResponse[CardTemplateListResponse]) {
	response = api.CreateResponse[CardTemplateListResponse](&CardTemplateListResponse{})
	return
}
