package cloud_api

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// DatacenterTraconvorder
// @id 271
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=查询订单数据)
func (client *Client) DatacenterTraconvorder(request *DatacenterTraconvorderRequest) (response *DatacenterTraconvorderResponse, err error) {
	rpcResponse := CreateDatacenterTraconvorderResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type DatacenterTraconvorderRequest struct {
	*api.BaseRequest

	BusinessType string `json:"business_type,omitempty"`
	Startdate    string `json:"startdate,omitempty"`
	Enddate      string `json:"enddate,omitempty"`
	PageNum      int    `json:"page_num,omitempty"`
	PageSize     int    `json:"page_size,omitempty"`
}

type DatacenterTraconvorderResponse struct {
}

func CreateDatacenterTraconvorderRequest() (request *DatacenterTraconvorderRequest) {
	request = &DatacenterTraconvorderRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("cloud_api", "1_0", "datacenter/traconvorder", "api")
	request.Method = api.POST
	return
}

func CreateDatacenterTraconvorderResponse() (response *api.BaseResponse[DatacenterTraconvorderResponse]) {
	response = api.CreateResponse[DatacenterTraconvorderResponse](&DatacenterTraconvorderResponse{})
	return
}
