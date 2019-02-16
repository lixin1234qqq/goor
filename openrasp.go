package openrasp

import (
	"log"
	"path/filepath"

	"github.com/baidu/openrasp/common"
	"github.com/baidu/openrasp/config"
)

const (
	Version = "1.1"
)

var workSpace *common.WorkSpace
var basic *config.BasicConfig
var general *config.GeneralConfig

func init() {
	workSpace = common.NewWorkSpace()
	workSpace.Init()
	basic = config.NewBasicConfig()
	if workSpace.Active() {
		confDir, err := workSpace.GetDir(common.Conf)
		if err != nil {
			log.Printf("%v", err)
		} else {
			path := filepath.Join(confDir, "openrasp.yml")
			err := basic.LoadProperties(path)
			if err != nil {
				log.Printf("%v", err)
			}
		}
	}
	general = config.NewGeneralConfig()
	if !basic.GetBool("cloud.enable") {
		workSpace.StartWatch(common.Conf)
		workSpace.RegisterListener(common.Conf, general)
	}
}

func GetWorkSpace() *common.WorkSpace {
	return workSpace
}

func GetBasic() *config.BasicConfig {
	return basic
}

func GetGeneral() *config.GeneralConfig {
	return general
}
