package tools

import (
	"strings"

	"os"

	"supermall/config"

	"fmt"
)

// ConfigToNoramlPath 将配置文件路径转换成真实路径
// eg:
// a|b|c -> a\b\c
func ConfigToNoramlPath(configPath string) string {
	str := strings.Split(configPath, "|")
	return strings.Join(str, string(os.PathSeparator))
}

// NormalPathToURL 文件路径转换为 URL
// eg:
// a\b\c -> http://$host:$port/a/b/c
func NormalPathToURL(normalPath string) string {
	pathSep := string(os.PathSeparator)
	str := strings.Split(normalPath, pathSep)
	url := fmt.Sprintf(
		"http://%s:%d/%s",
		config.ServerHost, config.ServerPort, strings.Join(str, "/"),
	)
	return url
}
