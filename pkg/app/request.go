package app

import (
	"github.com/astaxie/beego/validation"

	"github.com/kainhuck/gold/pkg/log"
)

// MarkErrors logs error logs
func MarkErrors(errors []*validation.Error) {
	for _, err := range errors {
		log.SugarLogger.Info(err.Key, err.Message)
	}

	return
}
