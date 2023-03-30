package ec

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// GuideDeleteGuider
// @id 1072
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=删除导购)
func (client *Client) GuideDeleteGuider(request *GuideDeleteGuiderRequest) (response *GuideDeleteGuiderResponse, err error) {
	rpcResponse := CreateGuideDeleteGuiderResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type GuideDeleteGuiderRequest struct {
	*api.BaseRequest

	GuiderPhone string `json:"guiderPhone,omitempty"`
	GuiderWid   int64  `json:"guiderWid,omitempty"`
	JobNumber   string `json:"jobNumber,omitempty"`
}

type GuideDeleteGuiderResponse struct {
}

func CreateGuideDeleteGuiderRequest() (request *GuideDeleteGuiderRequest) {
	request = &GuideDeleteGuiderRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("ec", "1_0", "guide/deleteGuider", "api")
	request.Method = api.POST
	return
}

func CreateGuideDeleteGuiderResponse() (response *api.BaseResponse[GuideDeleteGuiderResponse]) {
	response = api.CreateResponse[GuideDeleteGuiderResponse](&GuideDeleteGuiderResponse{})
	return
}
