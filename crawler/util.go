package crawler

import (
	"strings"
)

func intsJoin(i []string) string {
	f := i
	for k, e := range f {
		if strings.HasPrefix(e, "00") || strings.HasPrefix(e, "300") {
			i[k] = "sz" + e
		} else {
			i[k] = "sh" + e
		}
	}
	return strings.Join(i, ",")
}
