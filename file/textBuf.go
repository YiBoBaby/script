package main

import (
	"bytes"
	"fmt"
)

func main() {

}

const (
	TABLE = byte(9)
	WRAP  = byte(10)
)

// 文本输入缓冲
type TextBuf struct {
	*bytes.Buffer
}

// 初始化
func (textBuf *TextBuf) Init() {
	textBuf.Buffer = new(bytes.Buffer)
}

// 写入文本，在本文前加tab个制表符，在文本后加wrap个空格
func (textBuf *TextBuf) WriteText(ctx string, tab, wrap int) {
	for i := 0; i < tab; i++ {
		textBuf.writeTab()
	}
	if _, err := textBuf.WriteString(ctx); err != nil {
		fmt.Println(fmt.Sprintf("write text: %s, tab: %d, wrap: %d fail", ctx, tab, wrap))
	}
	for i := 0; i < wrap; i++ {
		textBuf.writeByte(WRAP)
	}
}

func (textBuf *TextBuf) writeTab() {
	textBuf.writeByte(TABLE)
}

func (textBuf *TextBuf) writeByte(b byte) {
	_ = textBuf.WriteByte(b)
}
