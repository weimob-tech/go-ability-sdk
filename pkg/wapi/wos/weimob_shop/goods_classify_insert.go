package weimob_shop

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// GoodsClassifyInsert
// @id 1651
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1651?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=添加商品分组)
func (client *Client) GoodsClassifyInsert(request *GoodsClassifyInsertRequest) (response *GoodsClassifyInsertResponse, err error) {
	rpcResponse := CreateGoodsClassifyInsertResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type GoodsClassifyInsertRequest struct {
	*api.BaseRequest

	BasicInfo GoodsClassifyInsertRequestBasicInfo `json:"basicInfo,omitempty"`
	Name      string                              `json:"name,omitempty"`
	ParentId  int64                               `json:"parentId,omitempty"`
	ImageUrl  string                              `json:"imageUrl,omitempty"`
	Sort      int64                               `json:"sort,omitempty"`
	IsShow    bool                                `json:"isShow,omitempty"`
	IsHot     bool                                `json:"isHot,omitempty"`
}

type GoodsClassifyInsertResponse struct {
	Id int64 `json:"id,omitempty"`
}

func CreateGoodsClassifyInsertRequest() (request *GoodsClassifyInsertRequest) {
	request = &GoodsClassifyInsertRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_shop", "v2.0", "goods/classify/insert", "apigw")
	request.Method = api.POST
	return
}

func CreateGoodsClassifyInsertResponse() (response *api.BaseResponse[GoodsClassifyInsertResponse]) {
	response = api.CreateResponse[GoodsClassifyInsertResponse](&GoodsClassifyInsertResponse{})
	return
}

type GoodsClassifyInsertRequestBasicInfo struct {
	Vid int64 `json:"vid,omitempty"`
}
