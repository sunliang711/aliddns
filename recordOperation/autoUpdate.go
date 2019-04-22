package recordOperation

import (
	"fmt"
	"github.com/sunliang711/aliddns/types"
	"log"
	"strings"
	"time"
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
	log.Printf("**************************automaticUpdate()**************************")
	defer func() {
		log.Printf("**************************Leave automaticUpdate()**************************")
	}()
	newValue, err := o.GetNewIP()
	if err != nil {
		log.Printf(">>GetNewIp error:%s", err)
		return
	}
	newValue = strings.TrimSpace(newValue)
	log.Printf("New ip: %s", newValue)
	//1. getRecordId
	subDomain := fmt.Sprintf("%v.%v", o.Config.RR, o.Config.DomainName)
	recordId, currentIp, err := o.GetRecordId(subDomain)
	if err == types.ErrNoSubDomain {
		recordId, err = o.AddRecord(o.Config.DomainName, o.Config.Type, o.Config.RR, newValue, o.Config.TTL)
	} else if err != nil {
		log.Printf(">>Exist such subDomain,but cann't get recordId")
		return
	}
	currentIp = strings.TrimSpace(currentIp)
	log.Printf("Current ip: %s", currentIp)
	if currentIp != newValue {
		//2. update
		res, err := o.UpdateRecord(recordId, o.Config.RR, o.Config.Type, newValue, o.Config.TTL)
		if err != nil {
			return
		}
		log.Printf(">>update OK:%v", res)
	} else {
		log.Printf("currentIp == new ip,do nothing.")
	}
}
