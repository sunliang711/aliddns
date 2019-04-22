package recordOperation

import (
	"github.com/sunliang711/aliddns/types"
	"encoding/json"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/alidns"
	"log"
	"net/http"
	"strings"
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
	log.Printf("GetRecordId(): subDomain:%v", subDomain)

	request := alidns.CreateDescribeSubDomainRecordsRequest()

	request.PageSize = requests.Integer(o.Config.PageSize)
	request.PageNumber = requests.Integer(o.Config.PageNumber)
	request.SubDomain = subDomain

	response, err := o.client.DescribeSubDomainRecords(request)
	if err != nil {
		log.Printf(">>DescribeSubDomainRecords error:%s", err)
		return "", "", err
	}
	if response.GetHttpStatus() != http.StatusOK {
		log.Printf(">>%v", types.ErrHttpStatusNotOK)
		return "", "", types.ErrHttpStatusNotOK
	}
	log.Println(">>response Content: ", response.GetHttpContentString())
	var res SubDomainRecordResponse
	err = json.Unmarshal(response.GetHttpContentBytes(), &res)
	if err != nil {
		log.Printf(">>json.Unmarshal error:%v", err)
		return "", "", err
	}
	if res.TotalCount == 0 {
		log.Printf(">>%v", types.ErrNoSubDomain)
		return "", "", types.ErrNoSubDomain
	}
	//RR.DomainName === subDomain
	if strings.Compare(res.DomainRecords.Record[0].RR+"."+res.DomainRecords.Record[0].DomainName, subDomain) != 0 {
		log.Printf(">>%v", types.ErrSubDomainNotMatch)
		return "", "", types.ErrSubDomainNotMatch
	}
	return res.DomainRecords.Record[0].RecordId, res.DomainRecords.Record[0].Value, nil
}
