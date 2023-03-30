package cloud_api

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// DatacenterOperpagemoni
// @id 270
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=查询被监测页面的流量数据)
func (client *Client) DatacenterOperpagemoni(request *DatacenterOperpagemoniRequest) (response *DatacenterOperpagemoniResponse, err error) {
	rpcResponse := CreateDatacenterOperpagemoniResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type DatacenterOperpagemoniRequest struct {
	*api.BaseRequest

	BusinessType string `json:"business_type,omitempty"`
	Startdate    string `json:"startdate,omitempty"`
	Enddate      string `json:"enddate,omitempty"`
	PageNum      int    `json:"page_num,omitempty"`
	PageSize     int    `json:"page_size,omitempty"`
}

type DatacenterOperpagemoniResponse struct {
}

func CreateDatacenterOperpagemoniRequest() (request *DatacenterOperpagemoniRequest) {
	request = &DatacenterOperpagemoniRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("cloud_api", "1_0", "datacenter/operpagemoni", "api")
	request.Method = api.POST
	return
}

func CreateDatacenterOperpagemoniResponse() (response *api.BaseResponse[DatacenterOperpagemoniResponse]) {
	response = api.CreateResponse[DatacenterOperpagemoniResponse](&DatacenterOperpagemoniResponse{})
	return
}
