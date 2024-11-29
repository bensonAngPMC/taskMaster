package util

import (
	"fmt"
	"strings"
)
func ParsePopulate(populate []string) map[string][]string {
	fmt.Println("Input populate:", populate)
	populateMap := make(map[string][]string)
	for _, param := range populate {
		if !strings.Contains(param, "[") && !strings.Contains(param, "=") {
			populateMap[param] = []string{}
		} else if strings.Contains(param, "[") && strings.HasPrefix(param, "populate") {
			innerKey := param[strings.Index(param, "[")+1 : strings.Index(param, "]")]
			if strings.Contains(param, "=") {
				value := strings.Split(param[strings.Index(param, "=")+1:], ",")
				populateMap[innerKey] = value
			} else {
				populateMap[innerKey] = []string{}
			}
		} else {
			populateMap[param] = []string{}
		}
	}

	fmt.Println("Parsed populate map:", populateMap)

	return populateMap
}
func CapitalizeFirstLetter(s string) string {
	if len(s) == 0 {
		return s // 如果字符串为空，直接返回
	}
	return strings.ToUpper(string(s[0])) + s[1:]
}