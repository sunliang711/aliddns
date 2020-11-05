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

type UpdateRecordResponse struct {
	RequestId string `json:"RequestId"`
	RecordId  string `json:"RecordId"`
}

func (o *Operator) UpdateRecord(recordId, RR, Type, Value, TTL string) (string, error) {
	logrus.Infof("UpdateRecord(): recordId:%v, RR:%v, Type:%v, Value:%v, TTL:%v", recordId, RR, Type, Value, TTL)

	request := alidns.CreateUpdateDomainRecordRequest()

	request.RecordId = recordId
	request.RR = RR
	request.Type = Type
	request.Value = Value
	request.TTL = requests.Integer(TTL)

	response, err := o.client.UpdateDomainRecord(request)
	if err != nil {
		return "", fmt.Errorf("UpdateDomainRecord error: %v", err)
	}

	if response.GetHttpStatus() != http.StatusOK {
		return "", fmt.Errorf("UpdateDomainRecord error: %v", types.ErrHttpStatusNotOK)
	}
	var res UpdateRecordResponse
	err = json.Unmarshal(response.GetHttpContentBytes(), &res)
	if err != nil {
		return "", fmt.Errorf("UpdateDomainRecord error: %v", err)
	}

	if res.RecordId != recordId {
		return "", fmt.Errorf("UpdateDomainRecord error: %v", types.ErrResponseIdNotMatchRequestId)
	}

	return response.GetHttpContentString(), nil
}
