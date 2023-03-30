package o2o

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// DishesGoodsDishesQueryList
// @id 2095
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=雅座查询菜品列表)
func (client *Client) DishesGoodsDishesQueryList(request *DishesGoodsDishesQueryListRequest) (response *DishesGoodsDishesQueryListResponse, err error) {
	rpcResponse := CreateDishesGoodsDishesQueryListResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type DishesGoodsDishesQueryListRequest struct {
	*api.BaseRequest

	BrandId  int64 `json:"brandId,omitempty"`
	StoreId  int64 `json:"storeId,omitempty"`
	BizType  int   `json:"bizType,omitempty"`
	PageNum  int   `json:"pageNum,omitempty"`
	PageSize int   `json:"pageSize,omitempty"`
}

type DishesGoodsDishesQueryListResponse struct {
}

func CreateDishesGoodsDishesQueryListRequest() (request *DishesGoodsDishesQueryListRequest) {
	request = &DishesGoodsDishesQueryListRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("o2o", "1_0", "dishesGoods/dishesQueryList", "api")
	request.Method = api.POST
	return
}

func CreateDishesGoodsDishesQueryListResponse() (response *api.BaseResponse[DishesGoodsDishesQueryListResponse]) {
	response = api.CreateResponse[DishesGoodsDishesQueryListResponse](&DishesGoodsDishesQueryListResponse{})
	return
}
