package ec

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// GuideGetCustomerGuider
// @id 1131
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=查询客户的专属导购)
func (client *Client) GuideGetCustomerGuider(request *GuideGetCustomerGuiderRequest) (response *GuideGetCustomerGuiderResponse, err error) {
	rpcResponse := CreateGuideGetCustomerGuiderResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type GuideGetCustomerGuiderRequest struct {
	*api.BaseRequest

	CustomerPhone string `json:"customerPhone,omitempty"`
	CustomerWid   int64  `json:"customerWid,omitempty"`
}

type GuideGetCustomerGuiderResponse struct {
}

func CreateGuideGetCustomerGuiderRequest() (request *GuideGetCustomerGuiderRequest) {
	request = &GuideGetCustomerGuiderRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("ec", "1_0", "guide/getCustomerGuider", "api")
	request.Method = api.POST
	return
}

func CreateGuideGetCustomerGuiderResponse() (response *api.BaseResponse[GuideGetCustomerGuiderResponse]) {
	response = api.CreateResponse[GuideGetCustomerGuiderResponse](&GuideGetCustomerGuiderResponse{})
	return
}
