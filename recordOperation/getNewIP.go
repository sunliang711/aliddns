package recordOperation

import (
	"github.com/sunliang711/aliddns/types"
	"bytes"
	"io"
	"log"
	"net/http"
	"os/exec"
	"regexp"
)

func (o *Operator) GetNewIP() (string, error) {
	if len(o.Config.NewIpCommand) > 0 {
		result, err := exec.Command("sh", "-c", o.Config.NewIpCommand).Output()
		if err != nil {
			return "", types.ErrCannotGetIpFromIpCommnad
		}
		return string(result), nil
	} else {
		return getIpByRegex(o.Config.NewIpSource, o.filterIpRegex, o.Config.MasterIndex, o.Config.SlaveIndex)
	}
}

func getIpByRegex(url string, re *regexp.Regexp, masterIndex, slaveIndex uint) (string, error) {
	log.Printf("getIpByRegex(): url:%s,masterIndex: %d,slaveIndex: %d", url, masterIndex, slaveIndex)

	res, err := http.Get(url)
	if err != nil {
		log.Printf(">>http.Get error: %s", err)
		return "", err
	}
	var buf bytes.Buffer
	_, err = io.Copy(&buf, res.Body)
	if err != nil {
		log.Printf(">>io.Copy error: %s", err)
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
