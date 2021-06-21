package cry

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"log"
	"net/http"
	"net/url"
	"nn/spider"
	"strings"
)

//解密
func AESDecrypt(crypted, key []byte) []byte {
	block, _ := aes.NewCipher(key)
	blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, key[:blockSize])
	origData := make([]byte, len(crypted))
	blockMode.CryptBlocks(origData, crypted)
	origData = PKCS7UnPadding(origData)
	return origData
}

//去补码
func PKCS7UnPadding(origData []byte) []byte {
	length := len(origData)
	if length == 0 {
		return origData
	}
	unpadding := int(origData[length-1])
	return origData[:length-unpadding]
}

//加密
func AESEncrypt(origData, key []byte) []byte {
	//获取block块
	block, _ := aes.NewCipher(key)
	//补码
	origData = PKCS7Padding(origData, block.BlockSize())
	//加密模式，
	blockMode := cipher.NewCBCEncrypter(block, key[:block.BlockSize()]) //创建明文长度的数组
	crypted := make([]byte, len(origData))
	//加密明文
	blockMode.CryptBlocks(crypted, origData)
	return crypted
}

//补码
func PKCS7Padding(origData []byte, blockSize int) []byte {
	//计算需要补几位数
	padding := blockSize - len(origData)%blockSize
	//在切片后面追加char数量的byte(char)
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(origData, padtext...)
}
func GetDecryptFunc(fromurl, v string) (func(bys []byte) []byte, error) {
	content := v[len("#EXT-X-KEY:"):]
	ps := strings.Split(content, ",")
	pp := make(map[string]string)
	for _, pv := range ps {
		ks := strings.Split(pv, "=")
		pp[strings.ToUpper(ks[0])] = ks[1]
	}
	if pp["METHOD"] == "AES-128" {
		if _, ok := pp["URI"]; ok {
			uri := pp["URI"]
			uri = strings.Trim(uri, "\"")
			u, _ := url.Parse(fromurl)
			u2, _ := u.Parse(uri)
			client := spider.NewClient(spider.RandomUserAgent())
			str, _, _, err1 := spider.GetResponseString(nil, func() (*http.Response, error) {
				return client.Get(u2.String(), "", nil, nil)
			})

			if err1 != nil {
				return nil, err1
			}

			log.Printf("密钥:%s\n", str)
			key := []byte(str)

			return func(bys []byte) []byte {
				return AESDecrypt(bys, key)
			}, nil
		}

	}
	return nil, nil
}
