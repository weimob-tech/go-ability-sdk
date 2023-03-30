package cloud_api

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// DatacenterCustomergrowthtrend
// @id 275
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=查询一段时间内会员数量的变化)
func (client *Client) DatacenterCustomergrowthtrend(request *DatacenterCustomergrowthtrendRequest) (response *DatacenterCustomergrowthtrendResponse, err error) {
	rpcResponse := CreateDatacenterCustomergrowthtrendResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type DatacenterCustomergrowthtrendRequest struct {
	*api.BaseRequest

	BusinessType int64  `json:"business_type,omitempty"`
	Startdate    string `json:"startdate,omitempty"`
	Enddate      string `json:"enddate,omitempty"`
	PageNum      int    `json:"page_num,omitempty"`
	PageSize     int    `json:"page_size,omitempty"`
}

type DatacenterCustomergrowthtrendResponse struct {
}

func CreateDatacenterCustomergrowthtrendRequest() (request *DatacenterCustomergrowthtrendRequest) {
	request = &DatacenterCustomergrowthtrendRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("cloud_api", "1_0", "datacenter/customergrowthtrend", "api")
	request.Method = api.POST
	return
}

func CreateDatacenterCustomergrowthtrendResponse() (response *api.BaseResponse[DatacenterCustomergrowthtrendResponse]) {
	response = api.CreateResponse[DatacenterCustomergrowthtrendResponse](&DatacenterCustomergrowthtrendResponse{})
	return
}
