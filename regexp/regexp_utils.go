package utils

import (
	"fmt"
	"regexp"
	"strings"
)

// CheckTypeRange 检测是否属于有效类型范围
func CheckTypeRange(validTypeRange []string, checkType string) bool {
	validReg := fmt.Sprintf("^(%s)$", strings.Join(validTypeRange, "|"))
	validProtocols := regexp.MustCompile(validReg)
	checkStr := strings.TrimSpace(strings.ToUpper(checkType))
	if !validProtocols.MatchString(checkStr) {
		return false
	}
	return true
}
