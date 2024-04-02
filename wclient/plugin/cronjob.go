package plugin

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"

	"github.com/opentdp/go-helper/filer"
	"github.com/opentdp/go-helper/logman"
	"github.com/opentdp/wrest-chat/dbase/cronjob"
)

type CronjobPlugin struct {
	Config *cronjob.CreateParam `json:"config"`
	Error  string               `json:"error"`
	Name   string               `json:"file"`
}

func CronjobPluginSetup() []*CronjobPlugin {

	dir := "./plugin/cronjob"
	if !filer.Exists(dir) {
		return nil
	}

	files, err := os.ReadDir(dir)
	if err != nil {
		return nil
	}

	configs := []*CronjobPlugin{}
	checker := NewCache(dir + ".txt")

	for _, info := range files {
		name := info.Name()
		// 忽略目录和隐藏文件
		if info.IsDir() || strings.HasPrefix(name, ".") {
			logman.Error("invalid cronjob plugin", "name", name, "error", err)
			continue
		}
		// 获取绝对路径
		rp := filepath.Join(dir, name)
		fp, err := filepath.Abs(rp)
		if err != nil {
			logman.Error("invalid cronjob plugin", "name", name, "error", err)
			continue
		}
		// 提取插件参数
		config, err := CronjobPluginParser(fp)
		if err != nil {
			configs = append(configs, &CronjobPlugin{config, err.Error(), name})
			continue
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
		configs = append(configs, &CronjobPlugin{config, errstr, name})
	}

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
