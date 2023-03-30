package cloud_api

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// DatacenterCustomerinfo
// @id 274
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=查询会员数据)
func (client *Client) DatacenterCustomerinfo(request *DatacenterCustomerinfoRequest) (response *DatacenterCustomerinfoResponse, err error) {
	rpcResponse := CreateDatacenterCustomerinfoResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type DatacenterCustomerinfoRequest struct {
	*api.BaseRequest

	BusinessType string `json:"business_type,omitempty"`
	PageNum      int    `json:"page_num,omitempty"`
	PageSize     int    `json:"page_size,omitempty"`
}

type DatacenterCustomerinfoResponse struct {
}

func CreateDatacenterCustomerinfoRequest() (request *DatacenterCustomerinfoRequest) {
	request = &DatacenterCustomerinfoRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("cloud_api", "1_0", "datacenter/customerinfo", "api")
	request.Method = api.POST
	return
}

func CreateDatacenterCustomerinfoResponse() (response *api.BaseResponse[DatacenterCustomerinfoResponse]) {
	response = api.CreateResponse[DatacenterCustomerinfoResponse](&DatacenterCustomerinfoResponse{})
	return
}
