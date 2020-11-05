package config

import (
	"bytes"
	"fmt"
	"io/ioutil"

	"github.com/pelletier/go-toml"
	"github.com/sirupsen/logrus"
)

type Config struct {
	RegionId     string `toml:"region_id"`
	PageSize     string `toml:"page_size"`
	PageNumber   string `toml:"page_number"`
	AccessKey    string `toml:"access_key"`
	AccessSecret string `toml:"access_secret"`

	Loglevel string `toml:"log_level"`
	//0 表示只更新一次
	UpdateInterval uint   `toml:"update_interval"`
	Type           string `toml:"type"`
	RR             string `toml:"rr"`
	DomainName     string `toml:"domain_name"`
	TTL            string `toml:"ttl"`
	//获取新ip的命令
	// curl -s --interface ppp0 whatismyip.akamai.com
	// curl -s myip.ipip.net|grep -o '[0-9.]\+'
	NewIpCommand string `toml:"new_ip_command"`
	//当NewIpCommand为空的时候
	//或者直接http.get(NewIpSource)然后对结果使用正则表达式FilterIpRegex来匹配,最后在匹配到的结果中选择第Index个作为新ip
	NewIpSource   string `toml:"new_ip_source"`
	FilterIpRegex string `toml:"filter_ip_regex"`
	MasterIndex   uint   `toml:"master_index"`
	SlaveIndex    uint   `toml:"slave_index"`
}

func (cfg *Config) String() string {
	return fmt.Sprintf("%+v", cfg)
}

func NewConfig(filename string) (*Config, error) {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	//配置文件中正则表达式需要两个反斜杠来转移,比较麻烦,在这里,把所有'\'替换成'\\'
	bs = bytes.Replace(bs, []byte("\\"), []byte("\\\\"), -1)
	var cfg Config
	err = toml.Unmarshal(bs, &cfg)
	if err != nil {
		return nil, err
	}
	logrus.Infof("Config: %+v", cfg)

	return &cfg, nil
}
