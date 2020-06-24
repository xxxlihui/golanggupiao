package service

import (
	"nn/data"
	"testing"
)

func TestLoadDaysToCache(t *testing.T) {
	data.InitDb(
		"server2",
		"postgres",
		"123",
		"gupiao",
		30002,
	)
	//LoadDaysToCache()

}
