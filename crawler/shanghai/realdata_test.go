package shanghai

import "testing"

func TestGetRealData(t *testing.T) {
	str,err:=GetRealData()
	if err!=nil{
		t.Error(err)
	}
	t.Log(str)
}