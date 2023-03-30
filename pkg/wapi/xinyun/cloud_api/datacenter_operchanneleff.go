package cloud_api

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// DatacenterOperchanneleff
// @id 269
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=查询渠道带来的流量数据)
func (client *Client) DatacenterOperchanneleff(request *DatacenterOperchanneleffRequest) (response *DatacenterOperchanneleffResponse, err error) {
	rpcResponse := CreateDatacenterOperchanneleffResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type DatacenterOperchanneleffRequest struct {
	*api.BaseRequest

	BusinessType string `json:"business_type,omitempty"`
	Startdate    string `json:"startdate,omitempty"`
	Enddate      string `json:"enddate,omitempty"`
	PageNum      int    `json:"page_num,omitempty"`
	PageSize     int    `json:"page_size,omitempty"`
}

type DatacenterOperchanneleffResponse struct {
}

func CreateDatacenterOperchanneleffRequest() (request *DatacenterOperchanneleffRequest) {
	request = &DatacenterOperchanneleffRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("cloud_api", "1_0", "datacenter/operchanneleff", "api")
	request.Method = api.POST
	return
}

func CreateDatacenterOperchanneleffResponse() (response *api.BaseResponse[DatacenterOperchanneleffResponse]) {
	response = api.CreateResponse[DatacenterOperchanneleffResponse](&DatacenterOperchanneleffResponse{})
	return
}
