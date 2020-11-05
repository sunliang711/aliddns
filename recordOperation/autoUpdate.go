package recordOperation

import (
	"fmt"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/sunliang711/aliddns/types"
)

func (o *Operator) AutomaticUpdate() {
	o.automaticUpdate()
	if o.Config.UpdateInterval > 0 {
		tick := time.NewTicker(time.Duration(o.Config.UpdateInterval) * time.Second)
		for {
			select {
			case <-tick.C:
				o.automaticUpdate()
			}
		}
	}
}

func (o *Operator) automaticUpdate() {
	logrus.Infof("Begin update DDNS...")
	newValue, err := o.GetNewIP()
	if err != nil {
		logrus.Errorf("Get new ip error: %v", err)
		return
	}
	newValue = strings.TrimSpace(newValue)
	logrus.Infof("Got new ip: %s", newValue)
	//1. getRecordId
	subDomain := fmt.Sprintf("%v.%v", o.Config.RR, o.Config.DomainName)
	recordId, currentIp, err := o.GetRecordId(subDomain)
	if err == types.ErrNoSubDomain {
		recordId, err = o.AddRecord(o.Config.DomainName, o.Config.Type, o.Config.RR, newValue, o.Config.TTL)
		return
	} else if err != nil {
		logrus.Errorf(">>Exist such subDomain,but cann't get recordId")
		return
	}
	currentIp = strings.TrimSpace(currentIp)
	logrus.Infof("Current ip( in record id ): %s", currentIp)
	if currentIp != newValue {
		//2. update
		res, err := o.UpdateRecord(recordId, o.Config.RR, o.Config.Type, newValue, o.Config.TTL)
		if err != nil {
			return
		}
		logrus.Infof(">>Update OK:%v", res)
	} else {
		logrus.Infof("Do nothing: currentIp == new ip")
	}
}
