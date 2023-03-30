package o2o

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// UsercenterGetWeimobUserInfo
// @id 45
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=授权对象（公众号/店铺）基本信息接口)
func (client *Client) UsercenterGetWeimobUserInfo(request *UsercenterGetWeimobUserInfoRequest) (response *UsercenterGetWeimobUserInfoResponse, err error) {
	rpcResponse := CreateUsercenterGetWeimobUserInfoResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type UsercenterGetWeimobUserInfoRequest struct {
	*api.BaseRequest
}

type UsercenterGetWeimobUserInfoResponse struct {
}

func CreateUsercenterGetWeimobUserInfoRequest() (request *UsercenterGetWeimobUserInfoRequest) {
	request = &UsercenterGetWeimobUserInfoRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("o2o", "1_0", "usercenter/getWeimobUserInfo", "api")
	request.Method = api.POST
	return
}

func CreateUsercenterGetWeimobUserInfoResponse() (response *api.BaseResponse[UsercenterGetWeimobUserInfoResponse]) {
	response = api.CreateResponse[UsercenterGetWeimobUserInfoResponse](&UsercenterGetWeimobUserInfoResponse{})
	return
}
