package utils

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"ginapi/pkg/logging"
	"net/http"
	"regexp"
	"sort"
	"strings"
)

//ReadUserIP this is get user realip
func ReadUserIP(r *http.Request) string {
	IPAddress := r.Header.Get("X-Real-Ip")
	if IPAddress == "" {
		IPAddress = r.Header.Get("X-Forwarded-For")
	}
	if IPAddress == "" {
		IPAddress = r.RemoteAddr
	}
	return IPAddress
}

// ValidatorSign this is ValidatorSign
func ValidatorSign(params map[string]interface{}, keySign string) bool {
	var strBuilder strings.Builder
	var keys []string
	for key := range params {
		keys = append(keys, key)
	}
	signstr := fmt.Sprintf("%v", params["sign"])
	sort.Strings(keys)
	for _, varKey := range keys {
		if varKey == "sign" {
			continue
		}
		strBuilder.WriteString(varKey)
		strBuilder.WriteString("=")
		strBuilder.WriteString(fmt.Sprintf("%v", params[varKey]))
		strBuilder.WriteString("&")
	}
	strBuilder.WriteString("key=" + keySign)
	str := strBuilder.String()
	logging.Info(str[:len(str)-32])
	h := md5.New()
	h.Write([]byte(str))
	md5str := hex.EncodeToString(h.Sum(nil))
	return md5str == signstr
}

// ToSnakeCase this is ToSnakeCase
func ToSnakeCase(str string) string {
	var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")
	snake := matchAllCap.ReplaceAllString(str, "${1}_${2}")
	return strings.ToLower(snake)
}

// ToCamelCase this is ToCamelCase
func ToCamelCase(str string) string {
	var re = regexp.MustCompile("(_|-)([a-zA-Z]+)")
	camel := re.ReplaceAllString(str, " $2")
	camel = strings.Title(camel)
	camel = strings.Replace(camel, " ", "", -1)
	return camel
}

// MapKeyToSnakeCase this is MapKeyToSnakeCase
func MapKeyToSnakeCase(params map[string]interface{}) map[string]interface{} {
	for k, v := range params {
		delete(params, k)
		k1 := ToSnakeCase(k)
		params[k1] = v
	}
	return params
}
