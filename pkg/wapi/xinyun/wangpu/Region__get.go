package wangpu

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// RegionGet
// @id 12
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=查询行政区域信息)
func (client *Client) RegionGet(request *RegionGetRequest) (response *RegionGetResponse, err error) {
	rpcResponse := CreateRegionGetResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type RegionGetRequest struct {
	*api.BaseRequest

	RegionPid int64 `json:"region_pid,omitempty"`
}

type RegionGetResponse struct {
}

func CreateRegionGetRequest() (request *RegionGetRequest) {
	request = &RegionGetRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("wangpu", "1_0", "Region/Get", "api")
	request.Method = api.POST
	return
}

func CreateRegionGetResponse() (response *api.BaseResponse[RegionGetResponse]) {
	response = api.CreateResponse[RegionGetResponse](&RegionGetResponse{})
	return
}
