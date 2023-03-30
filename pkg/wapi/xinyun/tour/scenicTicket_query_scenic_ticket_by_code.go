package tour

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// ScenicTicketQueryScenicTicketByCode
// @id 1108
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=查询门票产品详情t)
func (client *Client) ScenicTicketQueryScenicTicketByCode(request *ScenicTicketQueryScenicTicketByCodeRequest) (response *ScenicTicketQueryScenicTicketByCodeResponse, err error) {
	rpcResponse := CreateScenicTicketQueryScenicTicketByCodeResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type ScenicTicketQueryScenicTicketByCodeRequest struct {
	*api.BaseRequest

	GoodsCode string `json:"goodsCode,omitempty"`
}

type ScenicTicketQueryScenicTicketByCodeResponse struct {
}

func CreateScenicTicketQueryScenicTicketByCodeRequest() (request *ScenicTicketQueryScenicTicketByCodeRequest) {
	request = &ScenicTicketQueryScenicTicketByCodeRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("tour", "1_0", "scenicTicket/queryScenicTicketByCode", "api")
	request.Method = api.POST
	return
}

func CreateScenicTicketQueryScenicTicketByCodeResponse() (response *api.BaseResponse[ScenicTicketQueryScenicTicketByCodeResponse]) {
	response = api.CreateResponse[ScenicTicketQueryScenicTicketByCodeResponse](&ScenicTicketQueryScenicTicketByCodeResponse{})
	return
}
