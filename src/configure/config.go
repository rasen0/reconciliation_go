package configure

import (
	"encoding/xml"
	"io/ioutil"
	"os"
	"path/filepath"
)

const (
	configureFile = "conf.xml"
	configureDir  = "conf"
	logFileDir    = "log"
)

var defaultConfigure string = `
<config>
	<logpath>conf/log</logpath>
	<logprefix>log_</logprefix>
	<excelfile>核对账.xlsx</excelfile>
</config>
`

type Configure struct {
	Config    xml.Name `xml:"config"`
	LogPath   string   `xml:"logpath"`
	LogPrefix string   `xml:"logprefix"`
	ExcelFile string   `xml:"excelfile"`
}

func New() Configure {
	absPath, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	cfgFilePath := filepath.Join(absPath, configureDir, configureFile)
	_, err := os.Stat(cfgFilePath)
	if os.IsNotExist(err) {
		os.MkdirAll(configureDir, 755)
		var fs *os.File
		fs, _ = os.Create(filepath.Join(absPath, configureDir, configureFile))
		fs.WriteString(defaultConfigure)
		fs.Close()
	}
	bs, _ := ioutil.ReadFile(cfgFilePath)
	cfg := Configure{}
	xml.Unmarshal(bs, &cfg)
	return cfg
}
