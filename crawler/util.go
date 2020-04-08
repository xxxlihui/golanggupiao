package crawler

import (
	"strconv"
	"strings"
)

func intsJoin(i []int) string {
	s := make([]string, 0, len(i))
	for k, e := range i {
		s[k] = strconv.FormatInt(int64(e), 10)
	}
	return strings.Join(s, ",")
}
