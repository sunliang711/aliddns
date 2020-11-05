package recordOperation

import (
	"bytes"
	"io"
	"net/http"
	"os/exec"
	"regexp"

	"github.com/sirupsen/logrus"
	"github.com/sunliang711/aliddns/types"
)

func (o *Operator) GetNewIP() (string, error) {
	if len(o.Config.NewIpCommand) > 0 {
		logrus.Debugf("Get new ip by command: %v", o.Config.NewIpCommand)
		result, err := exec.Command("sh", "-c", o.Config.NewIpCommand).Output()
		if err != nil {
			return "", types.ErrCannotGetIpFromIpCommnad
		}
		return string(result), nil
	} else {
		logrus.Debugf("Get new ip by ip source")
		return getIpByRegex(o.Config.NewIpSource, o.filterIpRegex, o.Config.MasterIndex, o.Config.SlaveIndex)
	}
}

func getIpByRegex(url string, re *regexp.Regexp, masterIndex, slaveIndex uint) (string, error) {
	logrus.Debugf("getIpByRegex(): url:%s,masterIndex: %d,slaveIndex: %d", url, masterIndex, slaveIndex)

	res, err := http.Get(url)
	if err != nil {
		logrus.Errorf(">>http.Get error: %s", err)
		return "", err
	}
	var buf bytes.Buffer
	_, err = io.Copy(&buf, res.Body)
	if err != nil {
		logrus.Errorf(">>io.Copy error: %s", err)
		return "", err
	}
	result := re.FindAllStringSubmatch(buf.String(), -1)
	if uint(len(result)) <= masterIndex {
		return "", types.ErrCannotGetIpFromIpSource
	}
	if uint(len(result[masterIndex])) <= slaveIndex {
		return "", types.ErrCannotGetIpFromIpSource
	}
	return result[masterIndex][slaveIndex], nil
}
