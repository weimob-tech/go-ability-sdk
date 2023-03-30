package weimob_guide

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// CustomerGuiderBind
// @id 1383
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1383?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=新建客户导购关系)
func (client *Client) CustomerGuiderBind(request *CustomerGuiderBindRequest) (response *CustomerGuiderBindResponse, err error) {
	rpcResponse := CreateCustomerGuiderBindResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type CustomerGuiderBindRequest struct {
	*api.BaseRequest

	GuiderWid        int64 `json:"guiderWid,omitempty"`
	BindSceneType    int64 `json:"bindSceneType,omitempty"`
	ModifyReasonType int64 `json:"modifyReasonType,omitempty"`
	Wid              int64 `json:"wid,omitempty"`
	BizVid           int64 `json:"bizVid,omitempty"`
}

type CustomerGuiderBindResponse struct {
	Result bool `json:"result,omitempty"`
}

func CreateCustomerGuiderBindRequest() (request *CustomerGuiderBindRequest) {
	request = &CustomerGuiderBindRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_guide", "v2.0", "customer/guider/bind", "apigw")
	request.Method = api.POST
	return
}

func CreateCustomerGuiderBindResponse() (response *api.BaseResponse[CustomerGuiderBindResponse]) {
	response = api.CreateResponse[CustomerGuiderBindResponse](&CustomerGuiderBindResponse{})
	return
}
