package weimob_shop

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// GoodsClassifyGet
// @id 1525
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1525?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=查询分组详情)
func (client *Client) GoodsClassifyGet(request *GoodsClassifyGetRequest) (response *GoodsClassifyGetResponse, err error) {
	rpcResponse := CreateGoodsClassifyGetResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type GoodsClassifyGetRequest struct {
	*api.BaseRequest

	BasicInfo  GoodsClassifyGetRequestBasicInfo `json:"basicInfo,omitempty"`
	ClassifyId int64                            `json:"classifyId,omitempty"`
}

type GoodsClassifyGetResponse struct {
	ClassifyList []GoodsClassifyGetResponseClassifyList `json:"classifyList,omitempty"`
}

func CreateGoodsClassifyGetRequest() (request *GoodsClassifyGetRequest) {
	request = &GoodsClassifyGetRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_shop", "v2.0", "goods/classify/get", "apigw")
	request.Method = api.POST
	return
}

func CreateGoodsClassifyGetResponse() (response *api.BaseResponse[GoodsClassifyGetResponse]) {
	response = api.CreateResponse[GoodsClassifyGetResponse](&GoodsClassifyGetResponse{})
	return
}

type GoodsClassifyGetRequestBasicInfo struct {
	Vid int64 `json:"vid,omitempty"`
}

type GoodsClassifyGetResponseClassifyList struct {
	ClassifyLevel int64  `json:"classifyLevel,omitempty"`
	ClassifyId    int64  `json:"classifyId,omitempty"`
	Name          string `json:"name,omitempty"`
	ParentId      int64  `json:"parentId,omitempty"`
	IsShow        bool   `json:"isShow,omitempty"`
}
