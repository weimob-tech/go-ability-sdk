package ec

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// GuideUpdateGuider
// @id 1075
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=更新导购)
func (client *Client) GuideUpdateGuider(request *GuideUpdateGuiderRequest) (response *GuideUpdateGuiderResponse, err error) {
	rpcResponse := CreateGuideUpdateGuiderResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type GuideUpdateGuiderRequest struct {
	*api.BaseRequest

	GuiderPhone    string `json:"guiderPhone,omitempty"`
	GuiderName     string `json:"guiderName,omitempty"`
	IsUsed         int    `json:"isUsed,omitempty"`
	IsExclusiveCus int    `json:"isExclusiveCus,omitempty"`
	JobNumber      string `json:"jobNumber,omitempty"`
	GuiderWid      int64  `json:"guiderWid,omitempty"`
}

type GuideUpdateGuiderResponse struct {
}

func CreateGuideUpdateGuiderRequest() (request *GuideUpdateGuiderRequest) {
	request = &GuideUpdateGuiderRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("ec", "1_0", "guide/updateGuider", "api")
	request.Method = api.POST
	return
}

func CreateGuideUpdateGuiderResponse() (response *api.BaseResponse[GuideUpdateGuiderResponse]) {
	response = api.CreateResponse[GuideUpdateGuiderResponse](&GuideUpdateGuiderResponse{})
	return
}
