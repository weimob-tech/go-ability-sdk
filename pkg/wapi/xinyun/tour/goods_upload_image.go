package tour

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// GoodsUploadImage
// @id 999
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=上传商品图片信息)
func (client *Client) GoodsUploadImage(request *GoodsUploadImageRequest) (response *GoodsUploadImageResponse, err error) {
	rpcResponse := CreateGoodsUploadImageResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type GoodsUploadImageRequest struct {
	*api.BaseRequest

	File GoodsUploadImageRequestFile `json:"file,omitempty"`
}

type GoodsUploadImageResponse struct {
}

func CreateGoodsUploadImageRequest() (request *GoodsUploadImageRequest) {
	request = &GoodsUploadImageRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("tour", "1_0", "goods/uploadImage", "api")
	request.Method = api.POST
	return
}

func CreateGoodsUploadImageResponse() (response *api.BaseResponse[GoodsUploadImageResponse]) {
	response = api.CreateResponse[GoodsUploadImageResponse](&GoodsUploadImageResponse{})
	return
}
