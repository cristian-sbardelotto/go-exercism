package parsinglogfiles

import (
	"fmt"
	"regexp"
)

func IsValidLine(text string) bool {
	regex, err := regexp.Compile(`^(\[TRC\]|\[DBG\]|\[INF\]|\[WRN\]|\[ERR\]|\[FTL\])`)

	if err != nil {
		return false
	}

	return regex.MatchString(text)
}

func SplitLogLine(text string) []string {
	regex, err := regexp.Compile(`<[-~*=]*>`)

	if err != nil {
		return []string{}
	}

	return regex.Split(text, -1)
}

func CountQuotedPasswords(lines []string) int {
	regex, err := regexp.Compile(`(?i)"[^"]*password[^"]*"`)
	var count int

	if err != nil {
		return 0
	}

	for _, line := range lines {
		if regex.FindStringSubmatch(line) != nil {
			count += len(regex.FindStringSubmatch(line))
		}
	}

	return count
}

func RemoveEndOfLineText(text string) string {
	regex, err := regexp.Compile(`end-of-line\d+`)

	if err != nil {
		return ""
	}

	return regex.ReplaceAllString(text, "")
}

func TagWithUserName(lines []string) []string {
	regex, err := regexp.Compile(`User\s+(\S+)`)

	if err != nil {
		return []string{}
	}

	for i, line := range lines {
		matches := regex.FindStringSubmatch(line)

		if matches != nil {
			userName := matches[1]
			lines[i] = fmt.Sprintf("[USR] %s %s", userName, line)
		}
	}

	return lines
}
