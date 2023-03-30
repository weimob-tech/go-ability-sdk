package weimob_shop

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// GoodsAssign
// @id 1640
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1640?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=分配商品)
func (client *Client) GoodsAssign(request *GoodsAssignRequest) (response *GoodsAssignResponse, err error) {
	rpcResponse := CreateGoodsAssignResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type GoodsAssignRequest struct {
	*api.BaseRequest

	BasicInfo   GoodsAssignRequestBasicInfo `json:"basicInfo,omitempty"`
	AssignType  int64                       `json:"assignType,omitempty"`
	IsPutAway   int64                       `json:"isPutAway,omitempty"`
	GoodsIdList []int64                     `json:"goodsIdList,omitempty"`
	VidList     []int64                     `json:"vidList,omitempty"`
}

type GoodsAssignResponse struct {
	ReturnResult bool `json:"returnResult,omitempty"`
}

func CreateGoodsAssignRequest() (request *GoodsAssignRequest) {
	request = &GoodsAssignRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_shop", "v2.0", "goods/assign", "apigw")
	request.Method = api.POST
	return
}

func CreateGoodsAssignResponse() (response *api.BaseResponse[GoodsAssignResponse]) {
	response = api.CreateResponse[GoodsAssignResponse](&GoodsAssignResponse{})
	return
}

type GoodsAssignRequestBasicInfo struct {
	Vid int64 `json:"vid,omitempty"`
}
