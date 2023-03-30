package xiaoke

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// ProductClassificationList
// @id 2655
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=获取产品分类树)
func (client *Client) ProductClassificationList(request *ProductClassificationListRequest) (response *ProductClassificationListResponse, err error) {
	rpcResponse := CreateProductClassificationListResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type ProductClassificationListRequest struct {
	*api.BaseRequest

	UserWid int64 `json:"userWid,omitempty"`
}

type ProductClassificationListResponse struct {
	Id               int64   `json:"id,omitempty"`
	ObjectKey        string  `json:"objectKey,omitempty"`
	ParentObjectKey  string  `json:"parentObjectKey,omitempty"`
	ClassifyLevel    int     `json:"classifyLevel,omitempty"`
	MaxClassifyLevel int     `json:"maxClassifyLevel,omitempty"`
	ClassifyName     string  `json:"classifyName,omitempty"`
	Child            []int64 `json:"child,omitempty"`
}

func CreateProductClassificationListRequest() (request *ProductClassificationListRequest) {
	request = &ProductClassificationListRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("xiaoke", "1_0", "product/classificationList", "api")
	request.Method = api.POST
	return
}

func CreateProductClassificationListResponse() (response *api.BaseResponse[ProductClassificationListResponse]) {
	response = api.CreateResponse[ProductClassificationListResponse](&ProductClassificationListResponse{})
	return
}
