package kld

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// IStoreServiceGetStoreListInfo
// @id 261
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=获取门店列表)
func (client *Client) IStoreServiceGetStoreListInfo(request *IStoreServiceGetStoreListInfoRequest) (response *IStoreServiceGetStoreListInfoResponse, err error) {
	rpcResponse := CreateIStoreServiceGetStoreListInfoResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type IStoreServiceGetStoreListInfoRequest struct {
	*api.BaseRequest

	Data          IStoreServiceGetStoreListInfoRequestData `json:"data,omitempty"`
	CurrentUserId string                                   `json:"currentUserId,omitempty"`
	QueryName     string                                   `json:"queryName,omitempty"`
	PageIndex     int                                      `json:"pageIndex,omitempty"`
	PageSize      int                                      `json:"pageSize,omitempty"`
	Orderby       int                                      `json:"orderby,omitempty"`
}

type IStoreServiceGetStoreListInfoResponse struct {
}

func CreateIStoreServiceGetStoreListInfoRequest() (request *IStoreServiceGetStoreListInfoRequest) {
	request = &IStoreServiceGetStoreListInfoRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("kld", "1_0", "IStoreService/getStoreListInfo", "api")
	request.Method = api.POST
	return
}

func CreateIStoreServiceGetStoreListInfoResponse() (response *api.BaseResponse[IStoreServiceGetStoreListInfoResponse]) {
	response = api.CreateResponse[IStoreServiceGetStoreListInfoResponse](&IStoreServiceGetStoreListInfoResponse{})
	return
}

type IStoreServiceGetStoreListInfoRequestData struct {
}
