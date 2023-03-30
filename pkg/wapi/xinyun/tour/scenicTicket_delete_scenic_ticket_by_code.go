package tour

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// ScenicTicketDeleteScenicTicketByCode
// @id 1106
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=删除门票产品)
func (client *Client) ScenicTicketDeleteScenicTicketByCode(request *ScenicTicketDeleteScenicTicketByCodeRequest) (response *ScenicTicketDeleteScenicTicketByCodeResponse, err error) {
	rpcResponse := CreateScenicTicketDeleteScenicTicketByCodeResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type ScenicTicketDeleteScenicTicketByCodeRequest struct {
	*api.BaseRequest

	ScenicCode string `json:"scenicCode,omitempty"`
	GoodsCode  string `json:"goodsCode,omitempty"`
}

type ScenicTicketDeleteScenicTicketByCodeResponse struct {
}

func CreateScenicTicketDeleteScenicTicketByCodeRequest() (request *ScenicTicketDeleteScenicTicketByCodeRequest) {
	request = &ScenicTicketDeleteScenicTicketByCodeRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("tour", "1_0", "scenicTicket/deleteScenicTicketByCode", "api")
	request.Method = api.POST
	return
}

func CreateScenicTicketDeleteScenicTicketByCodeResponse() (response *api.BaseResponse[ScenicTicketDeleteScenicTicketByCodeResponse]) {
	response = api.CreateResponse[ScenicTicketDeleteScenicTicketByCodeResponse](&ScenicTicketDeleteScenicTicketByCodeResponse{})
	return
}
