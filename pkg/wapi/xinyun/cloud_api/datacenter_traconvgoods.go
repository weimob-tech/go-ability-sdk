package cloud_api

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// DatacenterTraconvgoods
// @id 272
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=查询商品转化数据)
func (client *Client) DatacenterTraconvgoods(request *DatacenterTraconvgoodsRequest) (response *DatacenterTraconvgoodsResponse, err error) {
	rpcResponse := CreateDatacenterTraconvgoodsResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type DatacenterTraconvgoodsRequest struct {
	*api.BaseRequest

	BusinessType string `json:"business_type,omitempty"`
	Startdate    string `json:"startdate,omitempty"`
	Enddate      string `json:"enddate,omitempty"`
	PageNum      int    `json:"page_num,omitempty"`
	PageSize     int    `json:"page_size,omitempty"`
}

type DatacenterTraconvgoodsResponse struct {
}

func CreateDatacenterTraconvgoodsRequest() (request *DatacenterTraconvgoodsRequest) {
	request = &DatacenterTraconvgoodsRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("cloud_api", "1_0", "datacenter/traconvgoods", "api")
	request.Method = api.POST
	return
}

func CreateDatacenterTraconvgoodsResponse() (response *api.BaseResponse[DatacenterTraconvgoodsResponse]) {
	response = api.CreateResponse[DatacenterTraconvgoodsResponse](&DatacenterTraconvgoodsResponse{})
	return
}
