package cry

import "testing"

func TestAESDecrypt(t *testing.T) {
	v := "#EXT-X-KEY:METHOD=AES-128,URI=\"/20210106/a04nv6M6/2000kb/hls/key.key\""
	GetDecryptFunc(v)
}
