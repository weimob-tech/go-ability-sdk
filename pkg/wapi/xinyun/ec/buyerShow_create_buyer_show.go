package ec

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// BuyerShowCreateBuyerShow
// @id 2119
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=新建买家秀)
func (client *Client) BuyerShowCreateBuyerShow(request *BuyerShowCreateBuyerShowRequest) (response *BuyerShowCreateBuyerShowResponse, err error) {
	rpcResponse := CreateBuyerShowCreateBuyerShowResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type BuyerShowCreateBuyerShowRequest struct {
	*api.BaseRequest

	RelateGoods      []BuyerShowCreateBuyerShowRequestRelateGoods      `json:"relateGoods,omitempty"`
	VideoContentList []BuyerShowCreateBuyerShowRequestVideoContentList `json:"videoContentList,omitempty"`
	ProfilePicture   string                                            `json:"profilePicture,omitempty"`
	UserName         string                                            `json:"userName,omitempty"`
	BuyerShowGroups  []int64                                           `json:"buyerShowGroups,omitempty"`
	ImageContent     []string                                          `json:"imageContent,omitempty"`
	TextContent      string                                            `json:"textContent,omitempty"`
	ApprovalStatus   int                                               `json:"approvalStatus,omitempty"`
}

type BuyerShowCreateBuyerShowResponse struct {
	ActivityId int64 `json:"activityId,omitempty"`
}

func CreateBuyerShowCreateBuyerShowRequest() (request *BuyerShowCreateBuyerShowRequest) {
	request = &BuyerShowCreateBuyerShowRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("ec", "1_0", "buyerShow/createBuyerShow", "api")
	request.Method = api.POST
	return
}

func CreateBuyerShowCreateBuyerShowResponse() (response *api.BaseResponse[BuyerShowCreateBuyerShowResponse]) {
	response = api.CreateResponse[BuyerShowCreateBuyerShowResponse](&BuyerShowCreateBuyerShowResponse{})
	return
}

type BuyerShowCreateBuyerShowRequestRelateGoods struct {
	GoodsId int64 `json:"goodsId,omitempty"`
}

type BuyerShowCreateBuyerShowRequestVideoContentList struct {
	ImageUrl string `json:"imageUrl,omitempty"`
	MediaUrl string `json:"mediaUrl,omitempty"`
}
