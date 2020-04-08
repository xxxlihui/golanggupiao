package crawler

import (
	"fmt"
	"testing"
)

func TestGetByCodes(t *testing.T) {
	codes := []string{"600031", "002513", "300215", "000114", "300015"}
	s, err := GetByCodes(codes)
	if err != nil {
		t.Error(err)
	}
	for _, e := range s {
		t.Log(fmt.Sprintf("%+v", e))
	}

}
