package recordOperation

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/alidns"
	"github.com/sirupsen/logrus"
	"github.com/sunliang711/aliddns/types"
)

type DeleteRecordResponse struct {
	RequestId string `json:"RequestId"`
	RecordId  string `json:"RecordId"`
}

func (o *Operator) DeleteRecord(recordId string) (string, error) {
	logrus.Infof("DeleteRecord(): recordId: %v", recordId)

	request := alidns.CreateDeleteDomainRecordRequest()

	request.RecordId = recordId

	response, err := o.client.DeleteDomainRecord(request)
	if err != nil {
		return "", fmt.Errorf("DeleteRecord error: %v", err)
	}
	if response.GetHttpStatus() != http.StatusOK {
		return "", fmt.Errorf("DeleteRecord error: %v", types.ErrHttpStatusNotOK)
	}

	var res DeleteRecordResponse
	err = json.Unmarshal(response.GetHttpContentBytes(), &res)
	if err != nil {
		return "", fmt.Errorf("DeleteRecord error: %v", err)
	}
	if res.RecordId != recordId {
		return "", fmt.Errorf("DeleteRecord error: %v", types.ErrResponseIdNotMatchRequestId)
	}
	return response.GetHttpContentString(), nil
}
