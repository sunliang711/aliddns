package recordOperation

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/alidns"
	"github.com/sirupsen/logrus"
	"github.com/sunliang711/aliddns/types"
)

type SubDomainRecordResponse struct {
	PageNumber    int    `json:"PageNumber"`
	TotalCount    int    `json:"TotalCount"`
	PageSize      int    `json:"PageSize"`
	RequestId     string `json:"RequestId"`
	DomainRecords struct {
		Record []struct {
			RR         string `json:"RR"`
			Status     string `json:"Status"`
			Value      string `json:"Value"`
			Weight     int    `json:"Weight"`
			RecordId   string `json:"RecordId"`
			Type       string `json:"Type"`
			DomainName string `json:"DomainName"`
			Locked     bool   `json:"Locked"`
			Line       string `json:"Line"`
			TTL        int    `json:"TTL"`
		} `json:"Record"`
	} `json:"DomainRecords"`
}

//查询子域名记录Id
//return : id,current ip,err
func (o *Operator) GetRecordId(subDomain string) (string, string, error) {
	logrus.Infof("GetRecordId...")

	request := alidns.CreateDescribeSubDomainRecordsRequest()

	request.PageSize = requests.Integer(o.Config.PageSize)
	request.PageNumber = requests.Integer(o.Config.PageNumber)
	request.SubDomain = subDomain

	response, err := o.client.DescribeSubDomainRecords(request)
	if err != nil {
		return "", "", fmt.Errorf("GetRecordId error: %v", err)
	}
	if response.GetHttpStatus() != http.StatusOK {
		return "", "", fmt.Errorf("GetRecordId error: %v", types.ErrHttpStatusNotOK)
	}
	logrus.Infof("Record id info: %v", response.GetHttpContentString())
	var res SubDomainRecordResponse
	err = json.Unmarshal(response.GetHttpContentBytes(), &res)
	if err != nil {
		return "", "", fmt.Errorf("GetRecordId error: %v", err)
	}
	if res.TotalCount == 0 {
		return "", "", fmt.Errorf("GetRecordId error: %v", types.ErrNoSubDomain)
	}
	//RR.DomainName === subDomain
	if strings.Compare(res.DomainRecords.Record[0].RR+"."+res.DomainRecords.Record[0].DomainName, subDomain) != 0 {
		return "", "", fmt.Errorf("GetRecordId error: %v", types.ErrSubDomainNotMatch)
	}
	return res.DomainRecords.Record[0].RecordId, res.DomainRecords.Record[0].Value, nil
}
