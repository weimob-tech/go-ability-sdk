package ec

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// NavigationPageUrlWithExtendParam
// @id 1848
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=生成导购链接)
func (client *Client) NavigationPageUrlWithExtendParam(request *NavigationPageUrlWithExtendParamRequest) (response *NavigationPageUrlWithExtendParamResponse, err error) {
	rpcResponse := CreateNavigationPageUrlWithExtendParamResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type NavigationPageUrlWithExtendParamRequest struct {
	*api.BaseRequest

	SysParam      NavigationPageUrlWithExtendParamRequestSysParam    `json:"sysParam,omitempty"`
	ExtendParam   NavigationPageUrlWithExtendParamRequestExtendParam `json:"extendParam,omitempty"`
	StoreId       string                                             `json:"storeId,omitempty"`
	Sign          string                                             `json:"sign,omitempty"`
	GuiderWid     int64                                              `json:"guiderWid,omitempty"`
	ChannelSource string                                             `json:"channelSource,omitempty"`
}

type NavigationPageUrlWithExtendParamResponse struct {
	GlobalTicket   string `json:"globalTicket,omitempty"`
	MonitorTrackId string `json:"monitorTrackId,omitempty"`
}

func CreateNavigationPageUrlWithExtendParamRequest() (request *NavigationPageUrlWithExtendParamRequest) {
	request = &NavigationPageUrlWithExtendParamRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("ec", "1_0", "navigation/pageUrlWithExtendParam", "api")
	request.Method = api.POST
	return
}

func CreateNavigationPageUrlWithExtendParamResponse() (response *api.BaseResponse[NavigationPageUrlWithExtendParamResponse]) {
	response = api.CreateResponse[NavigationPageUrlWithExtendParamResponse](&NavigationPageUrlWithExtendParamResponse{})
	return
}

type NavigationPageUrlWithExtendParamRequestSysParam struct {
}

type NavigationPageUrlWithExtendParamRequestExtendParam struct {
}
