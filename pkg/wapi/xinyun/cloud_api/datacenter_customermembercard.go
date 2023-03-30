package cloud_api

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// DatacenterCustomermembercard
// @id 276
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=查询一段时间内不同等级会员卡数据)
func (client *Client) DatacenterCustomermembercard(request *DatacenterCustomermembercardRequest) (response *DatacenterCustomermembercardResponse, err error) {
	rpcResponse := CreateDatacenterCustomermembercardResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type DatacenterCustomermembercardRequest struct {
	*api.BaseRequest

	BusinessType string `json:"business_type,omitempty"`
	PageNum      int    `json:"page_num,omitempty"`
	PageSize     int    `json:"page_size,omitempty"`
}

type DatacenterCustomermembercardResponse struct {
}

func CreateDatacenterCustomermembercardRequest() (request *DatacenterCustomermembercardRequest) {
	request = &DatacenterCustomermembercardRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("cloud_api", "1_0", "datacenter/customermembercard", "api")
	request.Method = api.POST
	return
}

func CreateDatacenterCustomermembercardResponse() (response *api.BaseResponse[DatacenterCustomermembercardResponse]) {
	response = api.CreateResponse[DatacenterCustomermembercardResponse](&DatacenterCustomermembercardResponse{})
	return
}
