package recordOperation

import (
	"fmt"
	"regexp"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/alidns"
	"github.com/sunliang711/aliddns/config"
)

type Operator struct {
	Config        *config.Config
	client        *alidns.Client
	filterIpRegex *regexp.Regexp
}

func NewOperator(cfg *config.Config) (*Operator, error) {
	o := &Operator{
		Config: cfg,
	}
	var err error
	o.client, err = alidns.NewClientWithAccessKey(o.Config.RegionId, o.Config.AccessKey, o.Config.AccessSecret)
	if err != nil {
		return nil, err
	}
	if len(o.Config.NewIpCommand) == 0 {
		//compile regex
		re, err := regexp.Compile(o.Config.FilterIpRegex)
		if err != nil {
			msg := fmt.Sprintf("NewipCommand is null and compile filterIpRegex failed:%v", err)
			return nil, fmt.Errorf(msg)
		}
		o.filterIpRegex = re
	}
	return o, nil
}
