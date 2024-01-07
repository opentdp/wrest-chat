package wcferry

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"regexp"
	"strings"

	"github.com/opentdp/go-helper/request"
)

// 获取网络文件
// param url string 文件URL或路径
// return string 失败则返回空字符串
func DownloadFile(url string) string {
	if strings.HasPrefix(url, "http://") || strings.HasPrefix(url, "https://") {
		if tmp, err := request.Download(url, "", false); err == nil {
			return tmp
		}
	}
	return ""
}

// 打印接收到的消息
// param msg *MsgPayload 消息
func MsgPrinter(msg *MsgPayload) {
	rs := "\n=== New Message ===\n"
	re := regexp.MustCompile(`(?m)^\s*|\n`)
	if msg.Id > 0 {
		rs += fmt.Sprintf("::Id:: %d\n", msg.Id)
	}
	if msg.Type > 0 {
		rs += fmt.Sprintf("::Type:: %d\n", msg.Type)
	}
	if msg.Roomid != "" {
		rs += fmt.Sprintf("::Roomid:: %s\n", msg.Roomid)
	}
	if msg.Sender != "" {
		rs += fmt.Sprintf("::Sender:: %v\n", msg.Sender)
	}
	if msg.Content != "" {
		rs += fmt.Sprintf("::Content:: %s\n", re.ReplaceAllString(msg.Content, ""))
	}
	if msg.ContentMap != nil {
		data, _ := json.Marshal(msg.ContentMap)
		rs += fmt.Sprintf("::ContentMap:: %s\n", string(data))
	}
	if msg.Xml != "" {
		rs += fmt.Sprintf("::Xml:: %s\n", re.ReplaceAllString(msg.Xml, ""))
	}
	if msg.XmlMap != nil {
		data, _ := json.Marshal(msg.XmlMap)
		rs += fmt.Sprintf("::XmlMap:: %s\n", string(data))
	}
	if msg.Extra != "" {
		rs += fmt.Sprintf("::Extra:: %s\n", re.ReplaceAllString(msg.Extra, ""))
	}
	fmt.Print(rs, "=== End Message ===\n")
}

// 解析 XML
// param data string XML数据
// return map[string]any 解析失败则返回空
func ParseXMLToMap(data string) (map[string]any, error) {
	var (
		decoder     = xml.NewDecoder(bytes.NewBufferString(data))
		nodeStack   []map[string]any
		currentNode map[string]any
	)

	for {
		token, err := decoder.Token()
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}

		switch tokenType := token.(type) {
		case xml.StartElement:
			// Create a new node and push it onto the stack
			newNode := make(map[string]any)
			for _, attr := range tokenType.Attr {
				newNode[attr.Name.Local] = attr.Value
			}
			if currentNode != nil {
				// Add this node as a child of the current node
				children, ok := currentNode[tokenType.Name.Local]
				if !ok {
					currentNode[tokenType.Name.Local] = []any{newNode}
				} else {
					currentNode[tokenType.Name.Local] = append(children.([]any), newNode)
				}
			}
			nodeStack = append(nodeStack, newNode)
			currentNode = newNode
		case xml.EndElement:
			// Pop the stack
			if len(nodeStack) > 1 {
				nodeStack = nodeStack[:len(nodeStack)-1]
				currentNode = nodeStack[len(nodeStack)-1]
			}
		case xml.CharData:
			charData := string(bytes.TrimSpace(tokenType))
			if charData != "" && len(nodeStack) > 0 {
				currentText, ok := currentNode["content"]
				if !ok {
					currentNode["content"] = charData
				} else {
					currentNode["content"] = currentText.(string) + charData
				}
			}
		}
	}

	if len(nodeStack) > 0 {
		return nodeStack[0], nil
	}
	return nil, fmt.Errorf("no root element")
}
