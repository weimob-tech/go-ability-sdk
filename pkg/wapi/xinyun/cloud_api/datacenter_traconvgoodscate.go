package cloud_api

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// DatacenterTraconvgoodscate
// @id 273
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=查询商品类目数据)
func (client *Client) DatacenterTraconvgoodscate(request *DatacenterTraconvgoodscateRequest) (response *DatacenterTraconvgoodscateResponse, err error) {
	rpcResponse := CreateDatacenterTraconvgoodscateResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type DatacenterTraconvgoodscateRequest struct {
	*api.BaseRequest

	BusinessType string `json:"business_type,omitempty"`
	Startdate    string `json:"startdate,omitempty"`
	Enddate      string `json:"enddate,omitempty"`
	PageNum      int    `json:"page_num,omitempty"`
	PageSize     int    `json:"page_size,omitempty"`
}

type DatacenterTraconvgoodscateResponse struct {
}

func CreateDatacenterTraconvgoodscateRequest() (request *DatacenterTraconvgoodscateRequest) {
	request = &DatacenterTraconvgoodscateRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("cloud_api", "1_0", "datacenter/traconvgoodscate", "api")
	request.Method = api.POST
	return
}

func CreateDatacenterTraconvgoodscateResponse() (response *api.BaseResponse[DatacenterTraconvgoodscateResponse]) {
	response = api.CreateResponse[DatacenterTraconvgoodscateResponse](&DatacenterTraconvgoodscateResponse{})
	return
}
