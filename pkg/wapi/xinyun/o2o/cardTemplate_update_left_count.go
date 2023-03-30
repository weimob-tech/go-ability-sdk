package o2o

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// CardTemplateUpdateLeftCount
// @id 127
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=修改库存)
func (client *Client) CardTemplateUpdateLeftCount(request *CardTemplateUpdateLeftCountRequest) (response *CardTemplateUpdateLeftCountResponse, err error) {
	rpcResponse := CreateCardTemplateUpdateLeftCountResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type CardTemplateUpdateLeftCountRequest struct {
	*api.BaseRequest

	CardTemplateId int64 `json:"cardTemplateId,omitempty"`
	UpdateTime     int64 `json:"updateTime,omitempty"`
	InventoryType  int   `json:"inventoryType,omitempty"`
	ChangeNum      int   `json:"changeNum,omitempty"`
}

type CardTemplateUpdateLeftCountResponse struct {
}

func CreateCardTemplateUpdateLeftCountRequest() (request *CardTemplateUpdateLeftCountRequest) {
	request = &CardTemplateUpdateLeftCountRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("o2o", "1_0", "cardTemplate/updateLeftCount", "api")
	request.Method = api.POST
	return
}

func CreateCardTemplateUpdateLeftCountResponse() (response *api.BaseResponse[CardTemplateUpdateLeftCountResponse]) {
	response = api.CreateResponse[CardTemplateUpdateLeftCountResponse](&CardTemplateUpdateLeftCountResponse{})
	return
}
