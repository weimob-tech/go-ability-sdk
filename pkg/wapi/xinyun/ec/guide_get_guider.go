package ec

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// GuideGetGuider
// @id 1073
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=查询导购)
func (client *Client) GuideGetGuider(request *GuideGetGuiderRequest) (response *GuideGetGuiderResponse, err error) {
	rpcResponse := CreateGuideGetGuiderResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type GuideGetGuiderRequest struct {
	*api.BaseRequest

	GuiderPhone string `json:"guiderPhone,omitempty"`
	GuiderWid   int64  `json:"guiderWid,omitempty"`
	JobNumber   string `json:"jobNumber,omitempty"`
}

type GuideGetGuiderResponse struct {
}

func CreateGuideGetGuiderRequest() (request *GuideGetGuiderRequest) {
	request = &GuideGetGuiderRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("ec", "1_0", "guide/getGuider", "api")
	request.Method = api.POST
	return
}

func CreateGuideGetGuiderResponse() (response *api.BaseResponse[GuideGetGuiderResponse]) {
	response = api.CreateResponse[GuideGetGuiderResponse](&GuideGetGuiderResponse{})
	return
}
