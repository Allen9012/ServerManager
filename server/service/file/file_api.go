package file

import (
	"fmt"
	"github.com/Allen9012/ServerManager/server/global"
	"github.com/Allen9012/ServerManager/server/model/file"
	"regexp"
	"strings"
)

func isRegexValid(FP *file.FilePermission) error {
	// 针对 \转义字符解决问题
	input := FP.Regexp
	nregstr := strings.ReplaceAll(input, "\\", "\\\\")
	FP.Regexp = nregstr
	_, err := regexp.Compile(FP.Regexp)
	global.GVA_LOG.Debug(fmt.Sprintf("isRegixValid Regexp: %s", FP.Regexp))
	return err
}
