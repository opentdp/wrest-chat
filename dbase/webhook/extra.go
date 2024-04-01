package webhook

import (
	"crypto/rand"
	"encoding/hex"
	"io"
)

// 生成 GUID

func generateGUID() string {

	// 创建一个16字节的切片用于存储随机数据
	byteGUID := make([]byte, 16)
	// 使用加密的随机数源生成随机数填充切片
	_, err := io.ReadFull(rand.Reader, byteGUID)
	if err != nil {
		panic("无法生成GUID")
	}

	// 设置版本号和变体以符合GUID版本4的标准
	byteGUID[8] = byteGUID[8]&^0xc0 | 0x80
	byteGUID[6] = byteGUID[6]&^0xf0 | 0x40

	// 将16字节的字节切片转换为32字节的字节切片
	guidBytes := append(byteGUID[0:4], byteGUID[4:6]...)
	guidBytes = append(guidBytes, byteGUID[6:8]...)
	guidBytes = append(guidBytes, byteGUID[8:10]...)
	guidBytes = append(guidBytes, byteGUID[10:16]...)

	// 使用十六进制编码32字节的字节切片
	return hex.EncodeToString(guidBytes)

}
