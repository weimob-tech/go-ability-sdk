package niepan

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// RightRightsusingRetrieveSoftwareRights
// @id 1946
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1946?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=权益中台-查询软件权益列表)
func (client *Client) RightRightsusingRetrieveSoftwareRights(request *RightRightsusingRetrieveSoftwareRightsRequest) (response *RightRightsusingRetrieveSoftwareRightsResponse, err error) {
	rpcResponse := CreateRightRightsusingRetrieveSoftwareRightsResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type RightRightsusingRetrieveSoftwareRightsRequest struct {
	*api.BaseRequest

	BosId             int64   `json:"bosId,omitempty"`
	FulfillItemIdList []int64 `json:"fulfillItemIdList,omitempty"`
	ProductId         int64   `json:"productId,omitempty"`
	ProductInstanceId int64   `json:"productInstanceId,omitempty"`
	GoodsId           int64   `json:"goodsId,omitempty"`
	GoodsType         []int64 `json:"goodsType,omitempty"`
	UsingStatusList   []int64 `json:"usingStatusList,omitempty"`
	ChannelCode       string  `json:"channelCode,omitempty"`
}

type RightRightsusingRetrieveSoftwareRightsResponse struct {
	RightsUsingId      int64  `json:"rightsUsingId,omitempty"`
	RightsName         string `json:"rightsName,omitempty"`
	ParentRightsId     int64  `json:"parentRightsId,omitempty"`
	LastFulfillItemId  int64  `json:"lastFulfillItemId,omitempty"`
	CustomerId         int64  `json:"customerId,omitempty"`
	MerchantId         int64  `json:"merchantId,omitempty"`
	BosId              int64  `json:"bosId,omitempty"`
	BosName            string `json:"bosName,omitempty"`
	ProductId          int64  `json:"productId,omitempty"`
	ProductVersionId   int64  `json:"productVersionId,omitempty"`
	ProductVersionName string `json:"productVersionName,omitempty"`
	ProductInstanceId  int64  `json:"productInstanceId,omitempty"`
	IsFree             int64  `json:"isFree,omitempty"`
	GoodsId            int64  `json:"goodsId,omitempty"`
	GoodsType          int64  `json:"goodsType,omitempty"`
	SkuId              int64  `json:"skuId,omitempty"`
	UsingStatus        int64  `json:"usingStatus,omitempty"`
	StartTime          string `json:"startTime,omitempty"`
	EndTime            string `json:"endTime,omitempty"`
	CreateTime         string `json:"createTime,omitempty"`
}

func CreateRightRightsusingRetrieveSoftwareRightsRequest() (request *RightRightsusingRetrieveSoftwareRightsRequest) {
	request = &RightRightsusingRetrieveSoftwareRightsRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("niepan", "v2.0", "right/rightsusing/retrieveSoftwareRights", "apigw")
	request.Method = api.POST
	return
}

func CreateRightRightsusingRetrieveSoftwareRightsResponse() (response *api.BaseResponse[RightRightsusingRetrieveSoftwareRightsResponse]) {
	response = api.CreateResponse[RightRightsusingRetrieveSoftwareRightsResponse](&RightRightsusingRetrieveSoftwareRightsResponse{})
	return
}
