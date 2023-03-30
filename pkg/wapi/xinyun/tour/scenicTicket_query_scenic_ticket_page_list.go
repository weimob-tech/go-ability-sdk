package tour

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// ScenicTicketQueryScenicTicketPageList
// @id 1107
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=查询门票产品列表)
func (client *Client) ScenicTicketQueryScenicTicketPageList(request *ScenicTicketQueryScenicTicketPageListRequest) (response *ScenicTicketQueryScenicTicketPageListResponse, err error) {
	rpcResponse := CreateScenicTicketQueryScenicTicketPageListResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type ScenicTicketQueryScenicTicketPageListRequest struct {
	*api.BaseRequest

	ScenicCode string `json:"scenicCode,omitempty"`
	GoodsCode  string `json:"goodsCode,omitempty"`
	GoodsName  string `json:"goodsName,omitempty"`
	Status     int    `json:"status,omitempty"`
	PageIndex  int    `json:"pageIndex,omitempty"`
	PageSize   int    `json:"pageSize,omitempty"`
}

type ScenicTicketQueryScenicTicketPageListResponse struct {
}

func CreateScenicTicketQueryScenicTicketPageListRequest() (request *ScenicTicketQueryScenicTicketPageListRequest) {
	request = &ScenicTicketQueryScenicTicketPageListRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("tour", "1_0", "scenicTicket/queryScenicTicketPageList", "api")
	request.Method = api.POST
	return
}

func CreateScenicTicketQueryScenicTicketPageListResponse() (response *api.BaseResponse[ScenicTicketQueryScenicTicketPageListResponse]) {
	response = api.CreateResponse[ScenicTicketQueryScenicTicketPageListResponse](&ScenicTicketQueryScenicTicketPageListResponse{})
	return
}
