package weimob_shop

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// GoodsSellstatusUpdate
// @id 2035
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/2035?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=更新商品可售状态)
func (client *Client) GoodsSellstatusUpdate(request *GoodsSellstatusUpdateRequest) (response *GoodsSellstatusUpdateResponse, err error) {
	rpcResponse := CreateGoodsSellstatusUpdateResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type GoodsSellstatusUpdateRequest struct {
	*api.BaseRequest

	BasicInfo   GoodsSellstatusUpdateRequestBasicInfo `json:"basicInfo,omitempty"`
	IsCanSell   bool                                  `json:"isCanSell,omitempty"`
	GoodsIdList []int64                               `json:"goodsIdList,omitempty"`
}

type GoodsSellstatusUpdateResponse struct {
	ReturnResult bool `json:"returnResult,omitempty"`
}

func CreateGoodsSellstatusUpdateRequest() (request *GoodsSellstatusUpdateRequest) {
	request = &GoodsSellstatusUpdateRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_shop", "v2.0", "goods/sellstatus/update", "apigw")
	request.Method = api.POST
	return
}

func CreateGoodsSellstatusUpdateResponse() (response *api.BaseResponse[GoodsSellstatusUpdateResponse]) {
	response = api.CreateResponse[GoodsSellstatusUpdateResponse](&GoodsSellstatusUpdateResponse{})
	return
}

type GoodsSellstatusUpdateRequestBasicInfo struct {
	Vid int64 `json:"vid,omitempty"`
}
