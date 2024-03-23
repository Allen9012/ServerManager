package buserr

import (
	ginI18n "github.com/gin-contrib/i18n"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

func GetErrMsg(key string, maps map[string]interface{}) string {
	content := ""
	if maps == nil {
		content = ginI18n.MustGetMessage(&i18n.LocalizeConfig{
			MessageID: key,
		})
	} else {
		content = ginI18n.MustGetMessage(&i18n.LocalizeConfig{
			MessageID:    key,
			TemplateData: maps,
		})
	}
	return content
}
