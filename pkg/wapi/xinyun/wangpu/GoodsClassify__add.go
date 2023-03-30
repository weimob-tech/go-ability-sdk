package wangpu

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// GoodsClassifyAdd
// @id 5
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=同步商品分组)
func (client *Client) GoodsClassifyAdd(request *GoodsClassifyAddRequest) (response *GoodsClassifyAddResponse, err error) {
	rpcResponse := CreateGoodsClassifyAddResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type GoodsClassifyAddRequest struct {
	*api.BaseRequest

	ClassifyName string `json:"classify_name,omitempty"`
	ClassifyPid  int64  `json:"classify_pid,omitempty"`
	Sort         int    `json:"sort,omitempty"`
	Imgurl       string `json:"imgurl,omitempty"`
}

type GoodsClassifyAddResponse struct {
}

func CreateGoodsClassifyAddRequest() (request *GoodsClassifyAddRequest) {
	request = &GoodsClassifyAddRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("wangpu", "1_0", "GoodsClassify/Add", "api")
	request.Method = api.POST
	return
}

func CreateGoodsClassifyAddResponse() (response *api.BaseResponse[GoodsClassifyAddResponse]) {
	response = api.CreateResponse[GoodsClassifyAddResponse](&GoodsClassifyAddResponse{})
	return
}
