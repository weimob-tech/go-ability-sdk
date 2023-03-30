package weimob_shop

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// GoodsGoodtemplateGetList
// @id 1854
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1854?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=查询商详模板列表)
func (client *Client) GoodsGoodtemplateGetList(request *GoodsGoodtemplateGetListRequest) (response *GoodsGoodtemplateGetListResponse, err error) {
	rpcResponse := CreateGoodsGoodtemplateGetListResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type GoodsGoodtemplateGetListRequest struct {
	*api.BaseRequest

	BasicInfo               GoodsGoodtemplateGetListRequestBasicInfo `json:"basicInfo,omitempty"`
	PageNum                 int64                                    `json:"pageNum,omitempty"`
	PageSize                int64                                    `json:"pageSize,omitempty"`
	TargetProductInstanceId int64                                    `json:"targetProductInstanceId,omitempty"`
	MerchantId              int64                                    `json:"merchantId,omitempty"`
	BosId                   int64                                    `json:"bosId,omitempty"`
}

type GoodsGoodtemplateGetListResponse struct {
	PageList   []GoodsGoodtemplateGetListResponsePageList `json:"pageList,omitempty"`
	PageNum    int64                                      `json:"pageNum,omitempty"`
	PageSize   int64                                      `json:"pageSize,omitempty"`
	TotalCount int64                                      `json:"totalCount,omitempty"`
}

func CreateGoodsGoodtemplateGetListRequest() (request *GoodsGoodtemplateGetListRequest) {
	request = &GoodsGoodtemplateGetListRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_shop", "v2.0", "goods/goodtemplate/getList", "apigw")
	request.Method = api.POST
	return
}

func CreateGoodsGoodtemplateGetListResponse() (response *api.BaseResponse[GoodsGoodtemplateGetListResponse]) {
	response = api.CreateResponse[GoodsGoodtemplateGetListResponse](&GoodsGoodtemplateGetListResponse{})
	return
}

type GoodsGoodtemplateGetListRequestBasicInfo struct {
	Vid int64 `json:"vid,omitempty"`
}

type GoodsGoodtemplateGetListResponsePageList struct {
	Id    int64  `json:"id,omitempty"`
	Title string `json:"title,omitempty"`
}
