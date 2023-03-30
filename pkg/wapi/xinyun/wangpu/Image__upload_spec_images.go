package wangpu

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// ImageUploadSpecImages
// @id 263
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=上传规格图片)
func (client *Client) ImageUploadSpecImages(request *ImageUploadSpecImagesRequest) (response *ImageUploadSpecImagesResponse, err error) {
	rpcResponse := CreateImageUploadSpecImagesResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type ImageUploadSpecImagesRequest struct {
	*api.BaseRequest

	SpuCode string `json:"spu_code,omitempty"`
	SkuCode string `json:"sku_code,omitempty"`
	ImgName string `json:"img_name,omitempty"`
	ImgData string `json:"img_data,omitempty"`
}

type ImageUploadSpecImagesResponse struct {
}

func CreateImageUploadSpecImagesRequest() (request *ImageUploadSpecImagesRequest) {
	request = &ImageUploadSpecImagesRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("wangpu", "1_0", "Image/UploadSpecImages", "api")
	request.Method = api.POST
	return
}

func CreateImageUploadSpecImagesResponse() (response *api.BaseResponse[ImageUploadSpecImagesResponse]) {
	response = api.CreateResponse[ImageUploadSpecImagesResponse](&ImageUploadSpecImagesResponse{})
	return
}
