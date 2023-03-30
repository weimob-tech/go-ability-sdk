package weimob_guide

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// GuiderUpdate
// @id 1759
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1759?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=修改导购信息)
func (client *Client) GuiderUpdate(request *GuiderUpdateRequest) (response *GuiderUpdateResponse, err error) {
	rpcResponse := CreateGuiderUpdateResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type GuiderUpdateRequest struct {
	*api.BaseRequest

	GuiderName     string `json:"guiderName,omitempty"`
	JobNumber      string `json:"jobNumber,omitempty"`
	IsExclusiveCus int64  `json:"isExclusiveCus,omitempty"`
	Wid            int64  `json:"wid,omitempty"`
	IsUsed         int64  `json:"isUsed,omitempty"`
	GuiderPhone    string `json:"guiderPhone,omitempty"`
	GuiderWid      int64  `json:"guiderWid,omitempty"`
	GuiderVid      int64  `json:"guiderVid,omitempty"`
}

type GuiderUpdateResponse struct {
	Result bool `json:"result,omitempty"`
}

func CreateGuiderUpdateRequest() (request *GuiderUpdateRequest) {
	request = &GuiderUpdateRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_guide", "v2.0", "guider/update", "apigw")
	request.Method = api.POST
	return
}

func CreateGuiderUpdateResponse() (response *api.BaseResponse[GuiderUpdateResponse]) {
	response = api.CreateResponse[GuiderUpdateResponse](&GuiderUpdateResponse{})
	return
}
