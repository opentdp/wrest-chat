package plugin

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"

	"github.com/opentdp/go-helper/logman"
	"github.com/opentdp/wechat-rest/dbase/cronjob"
)

type CronjobPlugin struct {
	Config *cronjob.CreateParam `json:"config"`
	Error  string               `json:"error"`
	Name   string               `json:"file"`
}

func CronjobPluginSetup() []*CronjobPlugin {

	configs := []*CronjobPlugin{}
	checker := NewCache("./plugin/cronjob.txt")

	filepath.Walk("./plugin/cronjob", func(rp string, info os.FileInfo, err error) error {
		// 忽略原则错误
		if err != nil || info.IsDir() {
			logman.Error("invalid cronjob plugin", "error", err)
			return nil
		}
		// 获取绝对路径
		fp, err := filepath.Abs(rp)
		if err != nil {
			logman.Error("invalid cronjob plugin", "error", err)
			return nil
		}
		// 提取插件参数
		config, err := CronjobPluginParser(fp)
		if err != nil {
			configs = append(configs, &CronjobPlugin{config, err.Error(), info.Name()})
			return nil
		}
		// 更新插件信息
		errstr := ""
		config.Rd = checker.Get(rp)
		if config.Rd == 0 {
			if rd, err := cronjob.Create(config); err == nil {
				checker.Put(rp, rd)
				config.Rd = rd
			} else {
				errstr = err.Error()
			}
		}
		configs = append(configs, &CronjobPlugin{config, errstr, info.Name()})
		return nil
	})

	return configs

}

func CronjobPluginParser(fp string) (*cronjob.CreateParam, error) {

	content, err := os.ReadFile(fp)
	if err != nil {
		return nil, err
	}

	// 提取插件参数
	re := regexp.MustCompile(`(?m)^(//|::|#)\s*@(Name|Second|Minute|Hour|DayOfMonth|Month|DayOfWeek|Timeout|Content|Deliver):\s*(.*)$`)
	matches := re.FindAllStringSubmatch(string(content), -1)
	if matches == nil {
		return nil, fmt.Errorf("cronjob config not found")
	}

	// 构造插件参数
	plugin := &cronjob.CreateParam{Type: "EXEC", Directory: "."}
	for _, match := range matches {
		match[3] = strings.TrimSpace(match[3])
		switch match[2] {
		case "Rd":
			n, _ := strconv.ParseInt(match[3], 10, 32)
			plugin.Rd = uint(n)
		case "Name":
			plugin.Name = match[3]
		case "Second":
			plugin.Second = match[3]
		case "Minute":
			plugin.Minute = match[3]
		case "Hour":
			plugin.Hour = match[3]
		case "DayOfMonth":
			plugin.DayOfMonth = match[3]
		case "Month":
			plugin.Month = match[3]
		case "DayOfWeek":
			plugin.DayOfWeek = match[3]
		case "Timeout":
			n, _ := strconv.ParseInt(match[3], 10, 32)
			plugin.Timeout = uint(n)
		case "Content":
			plugin.Content = match[3] + " " + fp
		case "Deliver":
			plugin.Deliver = match[3]
		}
	}

	return plugin, nil

}
