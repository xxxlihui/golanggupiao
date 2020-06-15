package service

import (
	"testing"
)

func TestLoadDaysToCache(t *testing.T) {
	InitDb(
		"server2",
		"postgres",
		"123",
		"gupiao",
		30002,
	)
	//LoadDaysToCache()

}
