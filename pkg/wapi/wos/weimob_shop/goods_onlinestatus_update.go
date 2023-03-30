package weimob_shop

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// GoodsOnlinestatusUpdate
// @id 1639
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1639?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=更新上下架状态)
func (client *Client) GoodsOnlinestatusUpdate(request *GoodsOnlinestatusUpdateRequest) (response *GoodsOnlinestatusUpdateResponse, err error) {
	rpcResponse := CreateGoodsOnlinestatusUpdateResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type GoodsOnlinestatusUpdateRequest struct {
	*api.BaseRequest

	BasicInfo   GoodsOnlinestatusUpdateRequestBasicInfo `json:"basicInfo,omitempty"`
	GoodsIdList []int64                                 `json:"goodsIdList,omitempty"`
	IsOnline    bool                                    `json:"isOnline,omitempty"`
}

type GoodsOnlinestatusUpdateResponse struct {
	ReturnResult bool `json:"returnResult,omitempty"`
}

func CreateGoodsOnlinestatusUpdateRequest() (request *GoodsOnlinestatusUpdateRequest) {
	request = &GoodsOnlinestatusUpdateRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_shop", "v2.0", "goods/onlinestatus/update", "apigw")
	request.Method = api.POST
	return
}

func CreateGoodsOnlinestatusUpdateResponse() (response *api.BaseResponse[GoodsOnlinestatusUpdateResponse]) {
	response = api.CreateResponse[GoodsOnlinestatusUpdateResponse](&GoodsOnlinestatusUpdateResponse{})
	return
}

type GoodsOnlinestatusUpdateRequestBasicInfo struct {
	Vid int64 `json:"vid,omitempty"`
}
