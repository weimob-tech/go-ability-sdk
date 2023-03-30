package wangpu

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// ImageGetSpecImages
// @id 203
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=获取规格图片)
func (client *Client) ImageGetSpecImages(request *ImageGetSpecImagesRequest) (response *ImageGetSpecImagesResponse, err error) {
	rpcResponse := CreateImageGetSpecImagesResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type ImageGetSpecImagesRequest struct {
	*api.BaseRequest

	SpuId int64 `json:"spu_id,omitempty"`
}

type ImageGetSpecImagesResponse struct {
}

func CreateImageGetSpecImagesRequest() (request *ImageGetSpecImagesRequest) {
	request = &ImageGetSpecImagesRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("wangpu", "1_0", "Image/GetSpecImages", "api")
	request.Method = api.POST
	return
}

func CreateImageGetSpecImagesResponse() (response *api.BaseResponse[ImageGetSpecImagesResponse]) {
	response = api.CreateResponse[ImageGetSpecImagesResponse](&ImageGetSpecImagesResponse{})
	return
}
