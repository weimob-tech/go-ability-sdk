package xiaoke

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// NicheProductList
// @id 2654
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=查询商机下产品列表)
func (client *Client) NicheProductList(request *NicheProductListRequest) (response *NicheProductListResponse, err error) {
	rpcResponse := CreateNicheProductListResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type NicheProductListRequest struct {
	*api.BaseRequest

	UserWid  int64  `json:"userWid,omitempty"`
	Key      string `json:"key,omitempty"`
	PageNum  int    `json:"pageNum,omitempty"`
	PageSize int    `json:"pageSize,omitempty"`
}

type NicheProductListResponse struct {
}

func CreateNicheProductListRequest() (request *NicheProductListRequest) {
	request = &NicheProductListRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("xiaoke", "1_0", "niche/productList", "api")
	request.Method = api.POST
	return
}

func CreateNicheProductListResponse() (response *api.BaseResponse[NicheProductListResponse]) {
	response = api.CreateResponse[NicheProductListResponse](&NicheProductListResponse{})
	return
}
