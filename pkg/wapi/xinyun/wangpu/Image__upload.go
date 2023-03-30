package wangpu

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// ImageUpload
// @id 20
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=上传商品图片)
func (client *Client) ImageUpload(request *ImageUploadRequest) (response *ImageUploadResponse, err error) {
	rpcResponse := CreateImageUploadResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type ImageUploadRequest struct {
	*api.BaseRequest

	SpuCode   string `json:"spu_code,omitempty"`
	IsMainImg bool   `json:"is_main_img,omitempty"`
	ImgName   string `json:"img_name,omitempty"`
	ImgData   string `json:"img_data,omitempty"`
}

type ImageUploadResponse struct {
}

func CreateImageUploadRequest() (request *ImageUploadRequest) {
	request = &ImageUploadRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("wangpu", "1_0", "Image/Upload", "api")
	request.Method = api.POST
	return
}

func CreateImageUploadResponse() (response *api.BaseResponse[ImageUploadResponse]) {
	response = api.CreateResponse[ImageUploadResponse](&ImageUploadResponse{})
	return
}
