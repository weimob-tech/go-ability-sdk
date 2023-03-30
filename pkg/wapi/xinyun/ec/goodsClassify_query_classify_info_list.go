package ec

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// GoodsClassifyQueryClassifyInfoList
// @id 357
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=获取分组列表)
func (client *Client) GoodsClassifyQueryClassifyInfoList(request *GoodsClassifyQueryClassifyInfoListRequest) (response *GoodsClassifyQueryClassifyInfoListResponse, err error) {
	rpcResponse := CreateGoodsClassifyQueryClassifyInfoListResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type GoodsClassifyQueryClassifyInfoListRequest struct {
	*api.BaseRequest
}

type GoodsClassifyQueryClassifyInfoListResponse struct {
	GoodsClassifyList []int64 `json:"goodsClassifyList,omitempty"`
}

func CreateGoodsClassifyQueryClassifyInfoListRequest() (request *GoodsClassifyQueryClassifyInfoListRequest) {
	request = &GoodsClassifyQueryClassifyInfoListRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("ec", "1_0", "goodsClassify/queryClassifyInfoList", "api")
	request.Method = api.POST
	return
}

func CreateGoodsClassifyQueryClassifyInfoListResponse() (response *api.BaseResponse[GoodsClassifyQueryClassifyInfoListResponse]) {
	response = api.CreateResponse[GoodsClassifyQueryClassifyInfoListResponse](&GoodsClassifyQueryClassifyInfoListResponse{})
	return
}
