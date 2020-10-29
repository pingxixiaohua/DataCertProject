package util

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"io"
	"io/ioutil"
)

/*
	对一个字符串进行MD5哈希计算，并返回
*/

func MD5HashString(data string) string {
	md5Hash := md5.New()
	md5Hash.Write([]byte(data))
	passwordBytes := md5Hash.Sum(nil)
	return hex.EncodeToString(passwordBytes)
}

/*
	对上传的数据进行hash计算
*/

func MD5HashReader(reader io.Reader) (string, error) {
	bytes, err := ioutil.ReadAll(reader)
	if err != nil {
		return  "",err
	}
	md5Hash := md5.New()
	md5Hash.Write(bytes)
	hashBytes := md5Hash.Sum(nil)
	return hex.EncodeToString(hashBytes),nil
}

//对区块数据进行hash计算

func SHA256Hash(data []byte) ([]byte) {

	//对数据进行sha256
	sha256Hash := sha256.New()
	sha256Hash.Write(data)
	return sha256Hash.Sum(nil)
}
