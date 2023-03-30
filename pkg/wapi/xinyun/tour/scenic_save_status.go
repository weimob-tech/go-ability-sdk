package tour

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// ScenicSaveStatus
// @id 1105
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=景区商品上/下架)
func (client *Client) ScenicSaveStatus(request *ScenicSaveStatusRequest) (response *ScenicSaveStatusResponse, err error) {
	rpcResponse := CreateScenicSaveStatusResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type ScenicSaveStatusRequest struct {
	*api.BaseRequest

	GoodsCode string `json:"goodsCode,omitempty"`
	Status    int    `json:"status,omitempty"`
}

type ScenicSaveStatusResponse struct {
}

func CreateScenicSaveStatusRequest() (request *ScenicSaveStatusRequest) {
	request = &ScenicSaveStatusRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("tour", "1_0", "scenic/saveStatus", "api")
	request.Method = api.POST
	return
}

func CreateScenicSaveStatusResponse() (response *api.BaseResponse[ScenicSaveStatusResponse]) {
	response = api.CreateResponse[ScenicSaveStatusResponse](&ScenicSaveStatusResponse{})
	return
}
