package regex_utils

import (
	"regexp"
	"sync"
)

var cache sync.Map

func MustCompileAndCache(exp string) *regexp.Regexp {
	compiled, ok := cache.Load(exp)
	if !ok {
		compiled, _ = cache.LoadOrStore(exp, regexp.MustCompile(exp))
	}

	return compiled.(*regexp.Regexp)
}