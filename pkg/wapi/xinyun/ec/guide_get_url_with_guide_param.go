package ec

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// GuideGetUrlWithGuideParam
// @id 1144
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=查询带导购参数链接/码)
func (client *Client) GuideGetUrlWithGuideParam(request *GuideGetUrlWithGuideParamRequest) (response *GuideGetUrlWithGuideParamResponse, err error) {
	rpcResponse := CreateGuideGetUrlWithGuideParamResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type GuideGetUrlWithGuideParamRequest struct {
	*api.BaseRequest

	GuiderWid   int64  `json:"guiderWid,omitempty"`
	Type        string `json:"type,omitempty"`
	Url         string `json:"url,omitempty"`
	Id          int64  `json:"id,omitempty"`
	GuiderPhone string `json:"guiderPhone,omitempty"`
	JobNumber   string `json:"jobNumber,omitempty"`
	Path        string `json:"path,omitempty"`
}

type GuideGetUrlWithGuideParamResponse struct {
}

func CreateGuideGetUrlWithGuideParamRequest() (request *GuideGetUrlWithGuideParamRequest) {
	request = &GuideGetUrlWithGuideParamRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("ec", "1_0", "guide/getUrlWithGuideParam", "api")
	request.Method = api.POST
	return
}

func CreateGuideGetUrlWithGuideParamResponse() (response *api.BaseResponse[GuideGetUrlWithGuideParamResponse]) {
	response = api.CreateResponse[GuideGetUrlWithGuideParamResponse](&GuideGetUrlWithGuideParamResponse{})
	return
}
