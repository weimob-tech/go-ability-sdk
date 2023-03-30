package weimob_guide

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// CustomerGuiderUnbind
// @id 1381
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1381?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=解除导购关系链)
func (client *Client) CustomerGuiderUnbind(request *CustomerGuiderUnbindRequest) (response *CustomerGuiderUnbindResponse, err error) {
	rpcResponse := CreateCustomerGuiderUnbindResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type CustomerGuiderUnbindRequest struct {
	*api.BaseRequest

	List             []CustomerGuiderUnbindRequestlist `json:"list,omitempty"`
	BindSceneType    int64                             `json:"bindSceneType,omitempty"`
	ModifyReasonType int64                             `json:"modifyReasonType,omitempty"`
}

type CustomerGuiderUnbindResponse struct {
	FailList []CustomerGuiderUnbindResponseFailList `json:"failList,omitempty"`
	Result   int64                                  `json:"result,omitempty"`
}

func CreateCustomerGuiderUnbindRequest() (request *CustomerGuiderUnbindRequest) {
	request = &CustomerGuiderUnbindRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_guide", "v2.0", "customer/guider/unbind", "apigw")
	request.Method = api.POST
	return
}

func CreateCustomerGuiderUnbindResponse() (response *api.BaseResponse[CustomerGuiderUnbindResponse]) {
	response = api.CreateResponse[CustomerGuiderUnbindResponse](&CustomerGuiderUnbindResponse{})
	return
}

type CustomerGuiderUnbindRequestlist struct {
	GuiderVid   int64 `json:"guiderVid,omitempty"`
	CustomerWid int64 `json:"customerWid,omitempty"`
}

type CustomerGuiderUnbindResponseFailList struct {
	CustomerWid int64  `json:"customerWid,omitempty"`
	Reason      string `json:"reason,omitempty"`
}
