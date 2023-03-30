package o2o

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// GoodsGroupSyn
// @id 239
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=同步菜品分组)
func (client *Client) GoodsGroupSyn(request *GoodsGroupSynRequest) (response *GoodsGroupSynResponse, err error) {
	rpcResponse := CreateGoodsGroupSynResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type GoodsGroupSynRequest struct {
	*api.BaseRequest

	StoreId       int64   `json:"storeId,omitempty"`
	Groups        []int64 `json:"groups,omitempty"`
	GroupName     string  `json:"groupName,omitempty"`
	ThirdGroupId  string  `json:"thirdGroupId,omitempty"`
	SortNo        int64   `json:"sortNo,omitempty"`
	SuppOrderType int     `json:"suppOrderType,omitempty"`
}

type GoodsGroupSynResponse struct {
}

func CreateGoodsGroupSynRequest() (request *GoodsGroupSynRequest) {
	request = &GoodsGroupSynRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("o2o", "1_0", "goodsGroup/syn", "api")
	request.Method = api.POST
	return
}

func CreateGoodsGroupSynResponse() (response *api.BaseResponse[GoodsGroupSynResponse]) {
	response = api.CreateResponse[GoodsGroupSynResponse](&GoodsGroupSynResponse{})
	return
}
