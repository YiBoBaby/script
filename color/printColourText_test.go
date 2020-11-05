package color

import (
	"fmt"
	"testing"
)

func TestBuildColorText(t *testing.T) {
	fmt.Println(BuildColorText("字体红色 背景青蓝色 带下划线", 31, 46, 4))
}
