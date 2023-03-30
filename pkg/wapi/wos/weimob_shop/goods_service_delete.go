package weimob_shop

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// GoodsServiceDelete
// @id 1863
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1863?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=删除商品服务)
func (client *Client) GoodsServiceDelete(request *GoodsServiceDeleteRequest) (response *GoodsServiceDeleteResponse, err error) {
	rpcResponse := CreateGoodsServiceDeleteResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type GoodsServiceDeleteRequest struct {
	*api.BaseRequest

	BasicInfo GoodsServiceDeleteRequestBasicInfo `json:"basicInfo,omitempty"`
	PropId    int64                              `json:"propId,omitempty"`
}

type GoodsServiceDeleteResponse struct {
	ReturnResult bool `json:"returnResult,omitempty"`
}

func CreateGoodsServiceDeleteRequest() (request *GoodsServiceDeleteRequest) {
	request = &GoodsServiceDeleteRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_shop", "v2.0", "goods/service/delete", "apigw")
	request.Method = api.POST
	return
}

func CreateGoodsServiceDeleteResponse() (response *api.BaseResponse[GoodsServiceDeleteResponse]) {
	response = api.CreateResponse[GoodsServiceDeleteResponse](&GoodsServiceDeleteResponse{})
	return
}

type GoodsServiceDeleteRequestBasicInfo struct {
	Vid int64 `json:"vid,omitempty"`
}
