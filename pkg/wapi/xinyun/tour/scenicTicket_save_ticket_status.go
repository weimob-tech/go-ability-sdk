package tour

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// ScenicTicketSaveTicketStatus
// @id 1109
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=上下架门票产品)
func (client *Client) ScenicTicketSaveTicketStatus(request *ScenicTicketSaveTicketStatusRequest) (response *ScenicTicketSaveTicketStatusResponse, err error) {
	rpcResponse := CreateScenicTicketSaveTicketStatusResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type ScenicTicketSaveTicketStatusRequest struct {
	*api.BaseRequest

	GoodsCode string `json:"goodsCode,omitempty"`
	Status    int    `json:"status,omitempty"`
}

type ScenicTicketSaveTicketStatusResponse struct {
}

func CreateScenicTicketSaveTicketStatusRequest() (request *ScenicTicketSaveTicketStatusRequest) {
	request = &ScenicTicketSaveTicketStatusRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("tour", "1_0", "scenicTicket/saveTicketStatus", "api")
	request.Method = api.POST
	return
}

func CreateScenicTicketSaveTicketStatusResponse() (response *api.BaseResponse[ScenicTicketSaveTicketStatusResponse]) {
	response = api.CreateResponse[ScenicTicketSaveTicketStatusResponse](&ScenicTicketSaveTicketStatusResponse{})
	return
}
