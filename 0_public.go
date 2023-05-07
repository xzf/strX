package strX

import "strings"

func SubBefore(source string, target string) (result string) {
	index := strings.Index(source, target)
	if index == -1 {
		return
	}
	return source[:index]
}

func SubAfter(source string, target string) (result string) {
	index := strings.Index(source, target)
	if index == -1 {
		return
	}
	return source[index+1:]
}

func SubBeforeLast(source string, target string) (result string) {
	index := strings.LastIndex(source, target)
	if index == -1 {
		return
	}
	return source[:index]
}

func SubAfterLast(source string, target string) (result string) {
	index := strings.LastIndex(source, target)
	if index == -1 {
		return
	}
	return source[index+1:]
}
