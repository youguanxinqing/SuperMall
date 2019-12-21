package tools

import (
	"fmt"
	"reflect"
	"regexp"
)

// TagRegxp struct tag 提取规则
var TagRegxp = regexp.MustCompile(`json:"(\w+)"`)

// SubGruopIndex 子组索引值
const SubGruopIndex = 1

// ExtractStructTag 提取 StructTag 中的 tag 信息
func ExtractStructTag(structTag reflect.StructTag) (string, error) {
	str := string([]byte(structTag))
	res := TagRegxp.FindStringSubmatch(str)
	if len(res) == 0 {
		return "", fmt.Errorf("no tag")
	} else {
		return res[SubGruopIndex], nil
	}
}
