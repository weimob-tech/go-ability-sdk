package ec

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// DiscountUpdateDiscountGoods
// @id 2731
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=限时折扣商品更新)
func (client *Client) DiscountUpdateDiscountGoodsV1_1(request *DiscountUpdateDiscountGoodsRequestV1_1) (response *DiscountUpdateDiscountGoodsResponseV1_1, err error) {
	rpcResponse := CreateDiscountUpdateDiscountGoodsResponseV1_1()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type DiscountUpdateDiscountGoodsRequestV1_1 struct {
	*api.BaseRequest
}

type DiscountUpdateDiscountGoodsResponseV1_1 struct {
}

func CreateDiscountUpdateDiscountGoodsRequestV1_1() (request *DiscountUpdateDiscountGoodsRequestV1_1) {
	request = &DiscountUpdateDiscountGoodsRequestV1_1{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("ec", "1_1", "discount/updateDiscountGoods", "api")
	request.Method = api.POST
	return
}

func CreateDiscountUpdateDiscountGoodsResponseV1_1() (response *api.BaseResponse[DiscountUpdateDiscountGoodsResponseV1_1]) {
	response = api.CreateResponse[DiscountUpdateDiscountGoodsResponseV1_1](&DiscountUpdateDiscountGoodsResponseV1_1{})
	return
}
