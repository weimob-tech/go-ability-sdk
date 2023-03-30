package o2o

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// ComboQueryPage
// @id 256
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=分页查询菜品套餐)
func (client *Client) ComboQueryPage(request *ComboQueryPageRequest) (response *ComboQueryPageResponse, err error) {
	rpcResponse := CreateComboQueryPageResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type ComboQueryPageRequest struct {
	*api.BaseRequest

	PageNum  int   `json:"pageNum,omitempty"`
	PageSize int   `json:"pageSize,omitempty"`
	StoreId  int64 `json:"storeId,omitempty"`
}

type ComboQueryPageResponse struct {
}

func CreateComboQueryPageRequest() (request *ComboQueryPageRequest) {
	request = &ComboQueryPageRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("o2o", "1_0", "combo/queryPage", "api")
	request.Method = api.POST
	return
}

func CreateComboQueryPageResponse() (response *api.BaseResponse[ComboQueryPageResponse]) {
	response = api.CreateResponse[ComboQueryPageResponse](&ComboQueryPageResponse{})
	return
}
