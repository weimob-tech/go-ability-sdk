package ec

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// GuideGetGuiderCustomers
// @id 1132
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=查询导购绑定的客户)
func (client *Client) GuideGetGuiderCustomers(request *GuideGetGuiderCustomersRequest) (response *GuideGetGuiderCustomersResponse, err error) {
	rpcResponse := CreateGuideGetGuiderCustomersResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type GuideGetGuiderCustomersRequest struct {
	*api.BaseRequest

	GuiderPhone string `json:"guiderPhone,omitempty"`
	GuiderWid   int64  `json:"guiderWid,omitempty"`
	JobNumber   string `json:"jobNumber,omitempty"`
	StartId     int64  `json:"startId,omitempty"`
}

type GuideGetGuiderCustomersResponse struct {
}

func CreateGuideGetGuiderCustomersRequest() (request *GuideGetGuiderCustomersRequest) {
	request = &GuideGetGuiderCustomersRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("ec", "1_0", "guide/getGuiderCustomers", "api")
	request.Method = api.POST
	return
}

func CreateGuideGetGuiderCustomersResponse() (response *api.BaseResponse[GuideGetGuiderCustomersResponse]) {
	response = api.CreateResponse[GuideGetGuiderCustomersResponse](&GuideGetGuiderCustomersResponse{})
	return
}
