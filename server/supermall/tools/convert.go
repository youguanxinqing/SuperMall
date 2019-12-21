package tools

import (
	"reflect"
	"strings"

	"os"

	"supermall/config"

	"fmt"
)

// ConfigToNoramlPath 将配置文件路径转换成真实路径
// eg:
// a|b|c => a\b\c
func ConfigToNoramlPath(configPath string) string {
	str := strings.Split(configPath, "|")
	return strings.Join(str, string(os.PathSeparator))
}

// NormalPathToURL 文件路径转换为 URL
// eg:
// a\b\c => http://$host:$port/a/b/c
func NormalPathToURL(normalPath string) string {
	pathSep := string(os.PathSeparator)
	str := strings.Split(normalPath, pathSep)
	url := fmt.Sprintf(
		"http://%s:%d/%s",
		config.ServerHost, config.ServerPort, strings.Join(str, "/"),
	)
	return url
}

// CompleteURLPath ...
// eg:
// a/b/c => http://$hos:$port/a/b/c
func CompleteURLPath(urlPath string) string {
	return fmt.Sprintf(
		"http://%s:%d%s",
		config.ServerHost, config.ServerPort, urlPath,
	)
}

// StructToMap 结构体转字典
// eg:
// struct{...} => map[string]interface{}{...}
func StructToMap(stru interface{}) map[string]interface{} {
	t := reflect.TypeOf(stru)
	v := reflect.ValueOf(stru)

	mapping := make(map[string]interface{})
	for i := 0; i < t.NumField(); i++ {
		if tag, err := ExtractStructTag(t.Field(i).Tag); err != nil {
			fmt.Println(err)
			mapping[t.Field(i).Name] = v.Field(i).Interface()
		} else {
			mapping[tag] = v.Field(i).Interface()
		}
	}
	return mapping
}
