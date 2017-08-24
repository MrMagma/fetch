package parser

import (
    "regexp"
    "sort"
)

type ArgType int

const (
    Single ArgType = iota
    Multi
)

type ArgData struct {
    Type ArgType
    Name string
}

type PatternArgs []ArgData

func (p PatternArgs) Len() int {
    return len(p)
}

func (p PatternArgs) Swap(i, j int) {
    p[i], p[j] = p[j], p[i]
}

func (p PatternArgs) Less(i, j int) bool {
    return stringLt(p[i].Name, p[j].Name)
}

var escapeRegexp = regexp.MustCompile("([\\-\\.\\_\\~\\:\\/\\?\\[\\]\\@\\!\\$\\&\\'\\\"\\(\\)\\*\\+\\,\\;\\=\\%])")
var singleRegexp = regexp.MustCompile("\\<.+?\\>")
var multiRegexp = regexp.MustCompile("\\<\\<.+?\\>\\>")

/*
Takes in a URL pattern represented by a string and outputs a lists of arguments found in that pattern, along with a regular expression formatted version of that pattern.
*/
func ParsePattern(patternStr string) (args PatternArgs, pattern string) {
    patternStr = escapeRegexp.ReplaceAllString(patternStr, "\\$1")
    patternStr = singleRegexp.ReplaceAllStringFunc(patternStr, func (match string) string {
        args = append(args, ArgData{Type: Single, Name: match[1:len(match) - 1]})
        return "(?P<" + match[1:len(match) - 1] + ">[^\\/]+?)"
    })
    pattern = multiRegexp.ReplaceAllStringFunc(patternStr, func (match string) string {
        args = append(args, ArgData{Type: Multi, Name: match[2:len(match) - 2]})
        return "(?P<" + match[2:len(match) - 2] + ">[^]+?)"
    })
    sort.Sort(args)
    return
}

