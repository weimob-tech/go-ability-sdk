package ec

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// BuyerShowGetBuyerShowGroupPage
// @id 2120
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=查询买家秀分组)
func (client *Client) BuyerShowGetBuyerShowGroupPage(request *BuyerShowGetBuyerShowGroupPageRequest) (response *BuyerShowGetBuyerShowGroupPageResponse, err error) {
	rpcResponse := CreateBuyerShowGetBuyerShowGroupPageResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type BuyerShowGetBuyerShowGroupPageRequest struct {
	*api.BaseRequest
}

type BuyerShowGetBuyerShowGroupPageResponse struct {
	RelateGoodsGroupList []BuyerShowGetBuyerShowGroupPageResponseRelateGoodsGroupList `json:"relateGoodsGroupList,omitempty"`
	Name                 string                                                       `json:"name,omitempty"`
	Id                   int64                                                        `json:"id,omitempty"`
	Type                 int                                                          `json:"type,omitempty"`
	BuyerShowCount       int                                                          `json:"buyerShowCount,omitempty"`
}

func CreateBuyerShowGetBuyerShowGroupPageRequest() (request *BuyerShowGetBuyerShowGroupPageRequest) {
	request = &BuyerShowGetBuyerShowGroupPageRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("ec", "1_0", "buyerShow/getBuyerShowGroupPage", "api")
	request.Method = api.POST
	return
}

func CreateBuyerShowGetBuyerShowGroupPageResponse() (response *api.BaseResponse[BuyerShowGetBuyerShowGroupPageResponse]) {
	response = api.CreateResponse[BuyerShowGetBuyerShowGroupPageResponse](&BuyerShowGetBuyerShowGroupPageResponse{})
	return
}

type BuyerShowGetBuyerShowGroupPageResponseRelateGoodsGroupList struct {
	GroupId   int64  `json:"groupId,omitempty"`
	GroupName string `json:"groupName,omitempty"`
}
