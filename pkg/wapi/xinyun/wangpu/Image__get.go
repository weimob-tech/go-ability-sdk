package wangpu

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// ImageGet
// @id 38
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=获取商品主图)
func (client *Client) ImageGet(request *ImageGetRequest) (response *ImageGetResponse, err error) {
	rpcResponse := CreateImageGetResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type ImageGetRequest struct {
	*api.BaseRequest

	SpuCode string `json:"spu_code,omitempty"`
}

type ImageGetResponse struct {
}

func CreateImageGetRequest() (request *ImageGetRequest) {
	request = &ImageGetRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("wangpu", "1_0", "Image/Get", "api")
	request.Method = api.POST
	return
}

func CreateImageGetResponse() (response *api.BaseResponse[ImageGetResponse]) {
	response = api.CreateResponse[ImageGetResponse](&ImageGetResponse{})
	return
}
