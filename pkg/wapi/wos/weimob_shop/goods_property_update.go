package weimob_shop

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// GoodsPropertyUpdate
// @id 1631
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1631?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=更新商品属性)
func (client *Client) GoodsPropertyUpdate(request *GoodsPropertyUpdateRequest) (response *GoodsPropertyUpdateResponse, err error) {
	rpcResponse := CreateGoodsPropertyUpdateResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type GoodsPropertyUpdateRequest struct {
	*api.BaseRequest

	BasicInfo    GoodsPropertyUpdateRequestBasicInfo `json:"basicInfo,omitempty"`
	PropId       int64                               `json:"propId,omitempty"`
	PropName     string                              `json:"propName,omitempty"`
	IsSearchable bool                                `json:"isSearchable,omitempty"`
	PropType     int64                               `json:"propType,omitempty"`
}

type GoodsPropertyUpdateResponse struct {
	ReturnResult bool `json:"returnResult,omitempty"`
}

func CreateGoodsPropertyUpdateRequest() (request *GoodsPropertyUpdateRequest) {
	request = &GoodsPropertyUpdateRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_shop", "v2.0", "goods/property/update", "apigw")
	request.Method = api.POST
	return
}

func CreateGoodsPropertyUpdateResponse() (response *api.BaseResponse[GoodsPropertyUpdateResponse]) {
	response = api.CreateResponse[GoodsPropertyUpdateResponse](&GoodsPropertyUpdateResponse{})
	return
}

type GoodsPropertyUpdateRequestBasicInfo struct {
	Vid int64 `json:"vid,omitempty"`
}
