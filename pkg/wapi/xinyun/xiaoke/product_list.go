package xiaoke

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// ProductList
// @id 2669
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=获取产品列表)
func (client *Client) ProductList(request *ProductListRequest) (response *ProductListResponse, err error) {
	rpcResponse := CreateProductListResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type ProductListRequest struct {
	*api.BaseRequest

	UserWid            int64   `json:"userWid,omitempty"`
	ProductUniqueNo    string  `json:"productUniqueNo,omitempty"`
	Version            int     `json:"version,omitempty"`
	ProductName        string  `json:"productName,omitempty"`
	ProductCode        string  `json:"productCode,omitempty"`
	ClassificationId   int64   `json:"classificationId,omitempty"`
	StandardPriceStart float64 `json:"standardPriceStart,omitempty"`
	StandardPriceEnd   float64 `json:"standardPriceEnd,omitempty"`
	State              int     `json:"state,omitempty"`
	OrderByApiName     string  `json:"orderByApiName,omitempty"`
	OrderByRule        int     `json:"orderByRule,omitempty"`
	PageNum            int     `json:"pageNum,omitempty"`
	PageSize           int     `json:"pageSize,omitempty"`
}

type ProductListResponse struct {
	PageNum    int64   `json:"pageNum,omitempty"`
	PageSize   int64   `json:"pageSize,omitempty"`
	TotalCount int64   `json:"totalCount,omitempty"`
	List       []int64 `json:"list,omitempty"`
}

func CreateProductListRequest() (request *ProductListRequest) {
	request = &ProductListRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("xiaoke", "1_0", "product/list", "api")
	request.Method = api.POST
	return
}

func CreateProductListResponse() (response *api.BaseResponse[ProductListResponse]) {
	response = api.CreateResponse[ProductListResponse](&ProductListResponse{})
	return
}
