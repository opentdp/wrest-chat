package deliver

import (
	"errors"
	"strings"
	"time"

	"github.com/opentdp/go-helper/logman"
)

func Send(deliver, content string) error {

	content = strings.TrimSpace(content)
	delivers := strings.Split(deliver, "\n")

	for _, dr := range delivers {
		logman.Warn("deliver "+dr, "content", content)
		// 解析参数
		args := strings.Split(strings.TrimSpace(dr), ",")
		if len(args) < 2 {
			return errors.New("deliver is error")
		}
		// 分渠道投递
		switch args[0] {
		case "wechat":
			time.Sleep(1 * time.Second)
			wechatMessage(args[1:], content)
		}
	}

	return nil

}
