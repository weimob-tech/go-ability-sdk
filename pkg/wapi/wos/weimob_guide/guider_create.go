package weimob_guide

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// GuiderCreate
// @id 1756
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1756?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=新增导购)
func (client *Client) GuiderCreate(request *GuiderCreateRequest) (response *GuiderCreateResponse, err error) {
	rpcResponse := CreateGuiderCreateResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type GuiderCreateRequest struct {
	*api.BaseRequest

	GuiderVid      int64  `json:"guiderVid,omitempty"`
	GuiderName     string `json:"guiderName,omitempty"`
	GuiderPhone    string `json:"guiderPhone,omitempty"`
	JobNumber      string `json:"jobNumber,omitempty"`
	IsExclusiveCus int64  `json:"isExclusiveCus,omitempty"`
	IsUsed         int64  `json:"isUsed,omitempty"`
	Zone           string `json:"zone,omitempty"`
}

type GuiderCreateResponse struct {
	Result bool  `json:"result,omitempty"`
	Wid    int64 `json:"wid,omitempty"`
}

func CreateGuiderCreateRequest() (request *GuiderCreateRequest) {
	request = &GuiderCreateRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_guide", "v2.0", "guider/create", "apigw")
	request.Method = api.POST
	return
}

func CreateGuiderCreateResponse() (response *api.BaseResponse[GuiderCreateResponse]) {
	response = api.CreateResponse[GuiderCreateResponse](&GuiderCreateResponse{})
	return
}
