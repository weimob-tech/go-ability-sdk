package weimob_guide

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// CustomerGuiderGet
// @id 1392
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1392?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=查询客户的导购)
func (client *Client) CustomerGuiderGet(request *CustomerGuiderGetRequest) (response *CustomerGuiderGetResponse, err error) {
	rpcResponse := CreateCustomerGuiderGetResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type CustomerGuiderGetRequest struct {
	*api.BaseRequest

	Wid      int64 `json:"wid,omitempty"`
	BizVid   int64 `json:"bizVid,omitempty"`
	BindType int64 `json:"bindType,omitempty"`
}

type CustomerGuiderGetResponse struct {
	GuiderName        string `json:"guiderName,omitempty"`
	GuiderWid         int64  `json:"guiderWid,omitempty"`
	GuiderId          string `json:"guiderId,omitempty"`
	GuiderVid         int64  `json:"guiderVid,omitempty"`
	GuiderPhone       string `json:"guiderPhone,omitempty"`
	GuiderJobNum      string `json:"guiderJobNum,omitempty"`
	CustomerWid       int64  `json:"customerWid,omitempty"`
	BindTime          string `json:"bindTime,omitempty"`
	BindSceneTypeDesc string `json:"bindSceneTypeDesc,omitempty"`
}

func CreateCustomerGuiderGetRequest() (request *CustomerGuiderGetRequest) {
	request = &CustomerGuiderGetRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_guide", "v2.0", "customer/guider/get", "apigw")
	request.Method = api.POST
	return
}

func CreateCustomerGuiderGetResponse() (response *api.BaseResponse[CustomerGuiderGetResponse]) {
	response = api.CreateResponse[CustomerGuiderGetResponse](&CustomerGuiderGetResponse{})
	return
}
