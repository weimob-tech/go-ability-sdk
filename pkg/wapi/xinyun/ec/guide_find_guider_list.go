package ec

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// GuideFindGuiderList
// @id 1070
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=查询导购列表)
func (client *Client) GuideFindGuiderList(request *GuideFindGuiderListRequest) (response *GuideFindGuiderListResponse, err error) {
	rpcResponse := CreateGuideFindGuiderListResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type GuideFindGuiderListRequest struct {
	*api.BaseRequest

	PageNum  int   `json:"pageNum,omitempty"`
	PageSize int   `json:"pageSize,omitempty"`
	StoreId  int64 `json:"storeId,omitempty"`
}

type GuideFindGuiderListResponse struct {
}

func CreateGuideFindGuiderListRequest() (request *GuideFindGuiderListRequest) {
	request = &GuideFindGuiderListRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("ec", "1_0", "guide/findGuiderList", "api")
	request.Method = api.POST
	return
}

func CreateGuideFindGuiderListResponse() (response *api.BaseResponse[GuideFindGuiderListResponse]) {
	response = api.CreateResponse[GuideFindGuiderListResponse](&GuideFindGuiderListResponse{})
	return
}
