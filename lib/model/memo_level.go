package model

import (
	"regexp"
)

// MemoLevel is
type MemoLevel struct{}

// GetLevel3 is
func getLevel3(line string) (string, string) {
	re := regexp.MustCompile(`^###\s\[.+\]\s(.*)`)
	switch {
	case regexp.MustCompile(`^###\s\[WORK\]`).Match([]byte(line)):
		return re.FindStringSubmatch(line)[1], "WORK"
	case regexp.MustCompile(`^###\s\[NECESSARY\]`).Match([]byte(line)):
		return "", "NECESSARY"
	case regexp.MustCompile(`^###\s\[FEASIBILITY\]`).Match([]byte(line)):
		return "", "FEASIBILITY"
	case regexp.MustCompile(`^###\s\[REFS\]`).Match([]byte(line)):
		return "", "REFS"
	case regexp.MustCompile(`^###\s\[WHAT\]`).Match([]byte(line)):
		return "", "WHAT"
	case regexp.MustCompile(`^###\s\[WHY\]`).Match([]byte(line)):
		return "", "WHY"
	case regexp.MustCompile(`^###\s\[HOW\]`).Match([]byte(line)):
		return "", "HOW"
	default:
		return "", ""
	}
}

// GetLevel2 is
func getLevel2(line string) (string, string) {
	re := regexp.MustCompile(`^##\s\[.+\]\s(.*)`)
	switch {
	case regexp.MustCompile(`^##\s\[STOP\]`).Match([]byte(line)):
		return re.FindStringSubmatch(line)[1], "STOP"
	case regexp.MustCompile(`^##\s\[THINK\]`).Match([]byte(line)):
		return re.FindStringSubmatch(line)[1], "THINK"
	case regexp.MustCompile(`^##\s\[REVIEW\]`).Match([]byte(line)):
		return re.FindStringSubmatch(line)[1], "REVIEW"
	case regexp.MustCompile(`^##\s\[WIP\]`).Match([]byte(line)):
		return re.FindStringSubmatch(line)[1], "WIP"
	case regexp.MustCompile(`^##\s\[WAIT\]`).Match([]byte(line)):
		return re.FindStringSubmatch(line)[1], "WAIT"
	case regexp.MustCompile(`^##\s\[DONE\]`).Match([]byte(line)):
		return re.FindStringSubmatch(line)[1], "DONE"
	default:
		return "", ""
	}
}

// GetLevel1 is
func getLevel1(line string) (string, string) {
	re := regexp.MustCompile(`^#\s\[.+\]\s(.*)`)
	switch {
	case regexp.MustCompile(`^#\s\[TOP\]`).Match([]byte(line)):
		return re.FindStringSubmatch(line)[1], "TOP"
	default:
		return "", ""
	}
}
