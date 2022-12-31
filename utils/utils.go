package utils

import (
	"fmt"
	"strings"
)

// Returns a log string with a chainId and tab as the prefix
// Ex:
//
//	| COSMOSHUB-4   |   string
func LogWithHostZone(chainId string, s string, a ...any) string {
	msg := fmt.Sprintf(s, a...)
	return fmt.Sprintf("|   %-13s |  %s", strings.ToUpper(chainId), msg)
}

// Returns a log string with a chain Id and callback as a prefix
// Ex:
//
//	| COSMOSHUB-4   |  DELEGATE CALLBACK  |  string
func LogCallbackWithHostZone(chainId string, callbackId string, s string, a ...any) string {
	msg := fmt.Sprintf(s, a...)
	return fmt.Sprintf("|   %-13s |  %s CALLBACK  |  %s", strings.ToUpper(chainId), strings.ToUpper(callbackId), msg)
}

// Returns a log header string with a dash padding on either side
// Ex:
//
//	------------------------------ string ------------------------------
func LogHeader(s string, a ...any) string {
	lineLength := 120
	header := fmt.Sprintf(s, a...)
	pad := strings.Repeat("-", (lineLength-len(header))/2)
	return fmt.Sprintf("%s %s %s", pad, header, pad)
}
