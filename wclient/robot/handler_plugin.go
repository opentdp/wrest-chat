package robot

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"

	"github.com/opentdp/go-helper/command"
	"github.com/opentdp/go-helper/logman"

	"github.com/opentdp/wechat-rest/wcferry"
)

func pluginHandler() []*Handler {

	cmds := []*Handler{}

	err := filepath.Walk("./plugins", func(rp string, info os.FileInfo, err error) error {
		if err != nil || info.IsDir() {
			return err
		}
		fp, err := filepath.Abs(rp)
		if err != nil {
			return err
		}
		// 提取指令参数
		v, err := pluginParser(fp)
		if err != nil {
			return err
		}
		// 生成插件指令
		cmds = append(cmds, &Handler{
			Level:    v.Level,
			Order:    700 + v.Order,
			Roomid:   v.Roomid,
			Command:  v.Command,
			Describe: v.Describe,
			Callback: func(msg *wcferry.WxMsg) string {
				exec := v.Binary + " " + fp
				output, err := command.Exec(&command.ExecPayload{
					Name:        "Handler:" + v.Command,
					CommandType: "EXEC",
					Content:     exec,
				})
				if err != nil {
					logman.Error("cmd: "+v.Command, "error", err)
				}
				return output
			},
		})
		return nil
	})

	if err != nil {
		logman.Error("scan plugin dir", "error", err)
	}

	return cmds

}

type pluginConfig struct {
	Level    int32
	Order    int32
	Roomid   string
	Command  string
	Describe string
	Binary   string
}

func pluginParser(fp string) (*pluginConfig, error) {

	content, err := os.ReadFile(fp)
	if err != nil {
		return nil, err
	}

	re := regexp.MustCompile(`(?m)^(//|::)\s*(Binary|Level|Order|Roomid|Command|Describe):\s*(.*)$`)
	matches := re.FindAllStringSubmatch(string(content), -1)
	if matches == nil {
		return nil, fmt.Errorf("no matching comment found")
	}

	h := &pluginConfig{}
	for _, match := range matches {
		match[3] = strings.TrimSpace(match[3])
		switch match[2] {
		case "Level":
			level, _ := strconv.ParseInt(match[3], 10, 32)
			h.Level = int32(level)
		case "Order":
			order, _ := strconv.ParseInt(match[3], 10, 32)
			h.Order = int32(order)
		case "Roomid":
			h.Roomid = match[3]
		case "Command":
			h.Command = match[3]
		case "Describe":
			h.Describe = match[3]
		case "Binary":
			h.Binary = match[3]
		}
	}

	return h, nil

}
