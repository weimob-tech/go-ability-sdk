package ec

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// GoodsQueryOuterGoodsCodeByCateWithCursor
// @id 1684
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=根据分组或类目滚动查询商品信息)
func (client *Client) GoodsQueryOuterGoodsCodeByCateWithCursor(request *GoodsQueryOuterGoodsCodeByCateWithCursorRequest) (response *GoodsQueryOuterGoodsCodeByCateWithCursorResponse, err error) {
	rpcResponse := CreateGoodsQueryOuterGoodsCodeByCateWithCursorResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type GoodsQueryOuterGoodsCodeByCateWithCursorRequest struct {
	*api.BaseRequest

	Cursor              int64 `json:"cursor,omitempty"`
	PageSize            int   `json:"pageSize,omitempty"`
	CategoryId          int64 `json:"categoryId,omitempty"`
	ClassifyId          int64 `json:"classifyId,omitempty"`
	FilterNullGoodsCode bool  `json:"filterNullGoodsCode,omitempty"`
}

type GoodsQueryOuterGoodsCodeByCateWithCursorResponse struct {
}

func CreateGoodsQueryOuterGoodsCodeByCateWithCursorRequest() (request *GoodsQueryOuterGoodsCodeByCateWithCursorRequest) {
	request = &GoodsQueryOuterGoodsCodeByCateWithCursorRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("ec", "1_0", "goods/queryOuterGoodsCodeByCateWithCursor", "api")
	request.Method = api.POST
	return
}

func CreateGoodsQueryOuterGoodsCodeByCateWithCursorResponse() (response *api.BaseResponse[GoodsQueryOuterGoodsCodeByCateWithCursorResponse]) {
	response = api.CreateResponse[GoodsQueryOuterGoodsCodeByCateWithCursorResponse](&GoodsQueryOuterGoodsCodeByCateWithCursorResponse{})
	return
}
