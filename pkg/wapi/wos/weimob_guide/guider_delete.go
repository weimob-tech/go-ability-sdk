package weimob_guide

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// GuiderDelete
// @id 1757
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1757?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=删除导购)
func (client *Client) GuiderDelete(request *GuiderDeleteRequest) (response *GuiderDeleteResponse, err error) {
	rpcResponse := CreateGuiderDeleteResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type GuiderDeleteRequest struct {
	*api.BaseRequest

	GuiderWid   int64  `json:"guiderWid,omitempty"`
	GuiderVid   int64  `json:"guiderVid,omitempty"`
	Wid         int64  `json:"wid,omitempty"`
	JobNumber   string `json:"jobNumber,omitempty"`
	GuiderPhone string `json:"guiderPhone,omitempty"`
}

type GuiderDeleteResponse struct {
	Result bool `json:"result,omitempty"`
}

func CreateGuiderDeleteRequest() (request *GuiderDeleteRequest) {
	request = &GuiderDeleteRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_guide", "v2.0", "guider/delete", "apigw")
	request.Method = api.POST
	return
}

func CreateGuiderDeleteResponse() (response *api.BaseResponse[GuiderDeleteResponse]) {
	response = api.CreateResponse[GuiderDeleteResponse](&GuiderDeleteResponse{})
	return
}
