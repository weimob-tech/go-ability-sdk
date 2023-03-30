package weimob_guide

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// CustomerGuiderChangeBind
// @id 1382
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1382?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=更换导购关系链)
func (client *Client) CustomerGuiderChangeBind(request *CustomerGuiderChangeBindRequest) (response *CustomerGuiderChangeBindResponse, err error) {
	rpcResponse := CreateCustomerGuiderChangeBindResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type CustomerGuiderChangeBindRequest struct {
	*api.BaseRequest

	GuiderWid        int64 `json:"guiderWid,omitempty"`
	GuiderVid        int64 `json:"guiderVid,omitempty"`
	CustomerWid      int64 `json:"customerWid,omitempty"`
	BindSceneType    int64 `json:"bindSceneType,omitempty"`
	ModifyReasonType int64 `json:"modifyReasonType,omitempty"`
}

type CustomerGuiderChangeBindResponse struct {
	Result bool `json:"result,omitempty"`
}

func CreateCustomerGuiderChangeBindRequest() (request *CustomerGuiderChangeBindRequest) {
	request = &CustomerGuiderChangeBindRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_guide", "v2.0", "customer/guider/changeBind", "apigw")
	request.Method = api.POST
	return
}

func CreateCustomerGuiderChangeBindResponse() (response *api.BaseResponse[CustomerGuiderChangeBindResponse]) {
	response = api.CreateResponse[CustomerGuiderChangeBindResponse](&CustomerGuiderChangeBindResponse{})
	return
}
