// +build go1.11,!go1.12

package stdlib

// Code generated by 'goexports log/syslog'. DO NOT EDIT.

import (
	"log/syslog"
	"reflect"
)

func init() {
	Symbols["log/syslog"] = map[string]reflect.Value{
		// function, constant and variable definitions
		"Dial":         reflect.ValueOf(syslog.Dial),
		"LOG_ALERT":    reflect.ValueOf(syslog.LOG_ALERT),
		"LOG_AUTH":     reflect.ValueOf(syslog.LOG_AUTH),
		"LOG_AUTHPRIV": reflect.ValueOf(syslog.LOG_AUTHPRIV),
		"LOG_CRIT":     reflect.ValueOf(syslog.LOG_CRIT),
		"LOG_CRON":     reflect.ValueOf(syslog.LOG_CRON),
		"LOG_DAEMON":   reflect.ValueOf(syslog.LOG_DAEMON),
		"LOG_DEBUG":    reflect.ValueOf(syslog.LOG_DEBUG),
		"LOG_EMERG":    reflect.ValueOf(syslog.LOG_EMERG),
		"LOG_ERR":      reflect.ValueOf(syslog.LOG_ERR),
		"LOG_FTP":      reflect.ValueOf(syslog.LOG_FTP),
		"LOG_INFO":     reflect.ValueOf(syslog.LOG_INFO),
		"LOG_KERN":     reflect.ValueOf(syslog.LOG_KERN),
		"LOG_LOCAL0":   reflect.ValueOf(syslog.LOG_LOCAL0),
		"LOG_LOCAL1":   reflect.ValueOf(syslog.LOG_LOCAL1),
		"LOG_LOCAL2":   reflect.ValueOf(syslog.LOG_LOCAL2),
		"LOG_LOCAL3":   reflect.ValueOf(syslog.LOG_LOCAL3),
		"LOG_LOCAL4":   reflect.ValueOf(syslog.LOG_LOCAL4),
		"LOG_LOCAL5":   reflect.ValueOf(syslog.LOG_LOCAL5),
		"LOG_LOCAL6":   reflect.ValueOf(syslog.LOG_LOCAL6),
		"LOG_LOCAL7":   reflect.ValueOf(syslog.LOG_LOCAL7),
		"LOG_LPR":      reflect.ValueOf(syslog.LOG_LPR),
		"LOG_MAIL":     reflect.ValueOf(syslog.LOG_MAIL),
		"LOG_NEWS":     reflect.ValueOf(syslog.LOG_NEWS),
		"LOG_NOTICE":   reflect.ValueOf(syslog.LOG_NOTICE),
		"LOG_SYSLOG":   reflect.ValueOf(syslog.LOG_SYSLOG),
		"LOG_USER":     reflect.ValueOf(syslog.LOG_USER),
		"LOG_UUCP":     reflect.ValueOf(syslog.LOG_UUCP),
		"LOG_WARNING":  reflect.ValueOf(syslog.LOG_WARNING),
		"New":          reflect.ValueOf(syslog.New),
		"NewLogger":    reflect.ValueOf(syslog.NewLogger),

		// type definitions
		"Priority": reflect.ValueOf((*syslog.Priority)(nil)),
		"Writer":   reflect.ValueOf((*syslog.Writer)(nil)),

		// interface wrapper definitions

	}
}
