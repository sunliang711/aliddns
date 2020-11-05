package recordOperation

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/alidns"
	"github.com/sirupsen/logrus"
	"github.com/sunliang711/aliddns/types"
)

//REF https://help.aliyun.com/document_detail/29772.html?spm=a2c4g.11186623.6.640.2c6f3192OeaBJ0
type addRecordResponse struct {
	RecordId  string `json:"RecordId"`
	RequestId string `json:"RequestId"`
}

//return RecordId,error
func (o *Operator) AddRecord(DomainName, Type, RR, Value, TTL string) (string, error) {
	logrus.Infof("AddRecord(): DomainName:%v, Type:%v, RR:%v, Value:%v, TTL:%v", DomainName, Type, RR, Value, TTL)

	request := alidns.CreateAddDomainRecordRequest()

	request.DomainName = DomainName
	request.Type = Type
	request.Value = Value
	request.RR = RR
	request.TTL = requests.Integer(TTL)

	response, err := o.client.AddDomainRecord(request)
	if err != nil {
		return "", fmt.Errorf("AddRecord error: %v", err)
	}
	if response.GetHttpStatus() != http.StatusOK {
		return "", fmt.Errorf("AddRecord error: %v", types.ErrHttpStatusNotOK)
	}
	var res addRecordResponse
	err = json.Unmarshal(response.GetHttpContentBytes(), &res)
	if err != nil {
		return "", fmt.Errorf("AddRecord error: %v", err)
	}
	return res.RecordId, nil
}
