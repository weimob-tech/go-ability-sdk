package xiaoke

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// LogsNiche
// @id 1727
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=新建商机跟进记录)
func (client *Client) LogsNiche(request *LogsNicheRequest) (response *LogsNicheResponse, err error) {
	rpcResponse := CreateLogsNicheResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type LogsNicheRequest struct {
	*api.BaseRequest

	Key      string `json:"key,omitempty"`
	Wid      int64  `json:"wid,omitempty"`
	PageNum  int    `json:"pageNum,omitempty"`
	PageSize int    `json:"pageSize,omitempty"`
}

type LogsNicheResponse struct {
	TotalPage        int     `json:"totalPage,omitempty"`
	PageSize         int     `json:"pageSize,omitempty"`
	Total            int64   `json:"total,omitempty"`
	PageNum          int64   `json:"pageNum,omitempty"`
	OpUserWid        string  `json:"opUserWid,omitempty"`
	OpUserName       string  `json:"opUserName,omitempty"`
	OpType           int     `json:"opType,omitempty"`
	OpContentText    string  `json:"opContentText,omitempty"`
	OpContent        string  `json:"opContent,omitempty"`
	Remark           string  `json:"remark,omitempty"`
	CreateTime       int     `json:"createTime,omitempty"`
	CreateUserWid    int64   `json:"createUserWid,omitempty"`
	UpdateTime       int     `json:"updateTime,omitempty"`
	UpdateUserWid    int64   `json:"updateUserWid,omitempty"`
	ImgAddress       string  `json:"imgAddress,omitempty"`
	VoiceAddress     string  `json:"voiceAddress,omitempty"`
	VoiceDuration    string  `json:"voiceDuration,omitempty"`
	CheckInAddress   string  `json:"checkInAddress,omitempty"`
	Longitude        float64 `json:"longitude,omitempty"`
	Latitude         float64 `json:"latitude,omitempty"`
	FileAddress      string  `json:"fileAddress,omitempty"`
	NicheDescription string  `json:"nicheDescription,omitempty"`
	VoiceText        string  `json:"voiceText,omitempty"`
}

func CreateLogsNicheRequest() (request *LogsNicheRequest) {
	request = &LogsNicheRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("xiaoke", "1_0", "logs/niche", "api")
	request.Method = api.POST
	return
}

func CreateLogsNicheResponse() (response *api.BaseResponse[LogsNicheResponse]) {
	response = api.CreateResponse[LogsNicheResponse](&LogsNicheResponse{})
	return
}
