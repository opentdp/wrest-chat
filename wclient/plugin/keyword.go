package plugin

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"

	"github.com/opentdp/wechat-rest/dbase/keyword"
)

type KeywordPlugin struct {
	Config *keyword.CreateParam `json:"config"`
	Error  string               `json:"error"`
	Name   string               `json:"file"`
}

func KeywordPluginSetup() ([]*KeywordPlugin, error) {

	configs := []*KeywordPlugin{}
	checker := NewCache("./plugin/keyword.txt")

	err := filepath.Walk("./plugin/keyword", func(rp string, info os.FileInfo, err error) error {
		if err != nil || info.IsDir() {
			return err
		}
		// 获取绝对路径
		fp, err := filepath.Abs(rp)
		if err != nil {
			return err
		}
		// 提取插件参数
		config, err := KeywordPluginParser(fp)
		if err != nil {
			return err
		}
		// 更新插件信息
		errstr := ""
		config.Rd = checker.Get(rp)
		if config.Rd == 0 {
			if rd, err := keyword.Create(config); err == nil {
				checker.Put(rp, rd)
				config.Rd = rd
			} else {
				errstr = err.Error()
			}
		}
		configs = append(configs, &KeywordPlugin{
			config, errstr, info.Name(),
		})
		return err
	})

	return configs, err

}

func KeywordPluginParser(fp string) (*keyword.CreateParam, error) {

	content, err := os.ReadFile(fp)
	if err != nil {
		return nil, err
	}

	// 提取插件参数
	re := regexp.MustCompile(`(?m)^(//|::)\s*@(Roomid|Phrase|Level|Target|Remark):\s*(.*)$`)
	matches := re.FindAllStringSubmatch(string(content), -1)
	if matches == nil {
		return nil, fmt.Errorf("command config not found")
	}

	// 构造插件参数
	plugin := &keyword.CreateParam{Group: "command"}
	for _, match := range matches {
		match[3] = strings.TrimSpace(match[3])
		switch match[2] {
		case "Rd":
			n, _ := strconv.ParseInt(match[3], 10, 32)
			plugin.Rd = uint(n)
		case "Roomid":
			plugin.Roomid = match[3]
		case "Phrase":
			plugin.Phrase = match[3]
		case "Level":
			n, _ := strconv.ParseInt(match[3], 10, 32)
			plugin.Level = int32(n)
		case "Target":
			plugin.Target = match[3] + " " + fp
		case "Remark":
			plugin.Remark = match[3]
		}
	}

	return plugin, nil

}
