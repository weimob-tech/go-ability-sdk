package weimob_shop

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// GoodsTagInsert
// @id 1853
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1853?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=添加标签)
func (client *Client) GoodsTagInsert(request *GoodsTagInsertRequest) (response *GoodsTagInsertResponse, err error) {
	rpcResponse := CreateGoodsTagInsertResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type GoodsTagInsertRequest struct {
	*api.BaseRequest

	BasicInfo GoodsTagInsertRequestBasicInfo `json:"basicInfo,omitempty"`
	Title     string                         `json:"title,omitempty"`
}

type GoodsTagInsertResponse struct {
	TagId int64 `json:"tagId,omitempty"`
}

func CreateGoodsTagInsertRequest() (request *GoodsTagInsertRequest) {
	request = &GoodsTagInsertRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_shop", "v2.0", "goods/tag/insert", "apigw")
	request.Method = api.POST
	return
}

func CreateGoodsTagInsertResponse() (response *api.BaseResponse[GoodsTagInsertResponse]) {
	response = api.CreateResponse[GoodsTagInsertResponse](&GoodsTagInsertResponse{})
	return
}

type GoodsTagInsertRequestBasicInfo struct {
	Vid int64 `json:"vid,omitempty"`
}
