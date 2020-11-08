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
	对一个I/O操作的reader（通常为文件）进行数据读取，对上传的数据进行hash计算，返回md5

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

/**
	对一个io操作的reader（通常是文件）进行数据读取，并计算hash，返回sha256哈希值
 */
func SHA256HashReader(reader io.Reader) (string, error) {
	bytes, err := ioutil.ReadAll(reader)
	if err != nil {
		return  "",err
	}
	sha256Hash := sha256.New()
	sha256Hash.Write(bytes)
	return hex.EncodeToString(sha256Hash.Sum(nil)),nil
}


//对区块数据进行hash计算

func SHA256Hash(data []byte) ([]byte) {

	//对数据进行sha256
	sha256Hash := sha256.New()
	sha256Hash.Write(data)
	return sha256Hash.Sum(nil)
}
