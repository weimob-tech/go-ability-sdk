package weimob_shop

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// GoodsPropertyGet
// @id 1634
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1634?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=查询商品属性详情)
func (client *Client) GoodsPropertyGet(request *GoodsPropertyGetRequest) (response *GoodsPropertyGetResponse, err error) {
	rpcResponse := CreateGoodsPropertyGetResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type GoodsPropertyGetRequest struct {
	*api.BaseRequest

	BasicInfo GoodsPropertyGetRequestBasicInfo `json:"basicInfo,omitempty"`
	PropId    int64                            `json:"propId,omitempty"`
	PropType  int64                            `json:"propType,omitempty"`
}

type GoodsPropertyGetResponse struct {
	PropId       int64  `json:"propId,omitempty"`
	Sort         int64  `json:"sort,omitempty"`
	PropName     string `json:"propName,omitempty"`
	IsSearchable bool   `json:"isSearchable,omitempty"`
}

func CreateGoodsPropertyGetRequest() (request *GoodsPropertyGetRequest) {
	request = &GoodsPropertyGetRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_shop", "v2.0", "goods/property/get", "apigw")
	request.Method = api.POST
	return
}

func CreateGoodsPropertyGetResponse() (response *api.BaseResponse[GoodsPropertyGetResponse]) {
	response = api.CreateResponse[GoodsPropertyGetResponse](&GoodsPropertyGetResponse{})
	return
}

type GoodsPropertyGetRequestBasicInfo struct {
	Vid int64 `json:"vid,omitempty"`
}
