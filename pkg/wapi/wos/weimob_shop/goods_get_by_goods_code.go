package weimob_shop

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// GoodsGetByGoodsCode
// @id 2036
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/2036?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=通过商品编码查询商品id)
func (client *Client) GoodsGetByGoodsCode(request *GoodsGetByGoodsCodeRequest) (response *GoodsGetByGoodsCodeResponse, err error) {
	rpcResponse := CreateGoodsGetByGoodsCodeResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type GoodsGetByGoodsCodeRequest struct {
	*api.BaseRequest

	BasicInfo          GoodsGetByGoodsCodeRequestBasicInfo `json:"basicInfo,omitempty"`
	OuterGoodsCodeList []int64                             `json:"outerGoodsCodeList,omitempty"`
}

type GoodsGetByGoodsCodeResponse struct {
	GoodsInfoList []GoodsGetByGoodsCodeResponseGoodsInfoList `json:"goodsInfoList,omitempty"`
}

func CreateGoodsGetByGoodsCodeRequest() (request *GoodsGetByGoodsCodeRequest) {
	request = &GoodsGetByGoodsCodeRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_shop", "v2.0", "goods/getByGoodsCode", "apigw")
	request.Method = api.POST
	return
}

func CreateGoodsGetByGoodsCodeResponse() (response *api.BaseResponse[GoodsGetByGoodsCodeResponse]) {
	response = api.CreateResponse[GoodsGetByGoodsCodeResponse](&GoodsGetByGoodsCodeResponse{})
	return
}

type GoodsGetByGoodsCodeRequestBasicInfo struct {
	Vid int64 `json:"vid,omitempty"`
}

type GoodsGetByGoodsCodeResponseGoodsInfoList struct {
	GoodsId        int64 `json:"goodsId,omitempty"`
	OuterGoodsCode int64 `json:"outerGoodsCode,omitempty"`
}
