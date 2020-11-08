package models

import (
	"bytes"
	"encoding/gob"
)

/**
 *数据上链的认证结构
 */
type CertRecord struct {
	CertHash       []byte //认证文件的sha256 Hash值
	CertHashStr    string
	CertId         []byte //Id
	CertIdStr      string
	CertAuthor     string //认证人
	Phone          string //电话
	AuthorCard     string //身份证号
	FileName       string //文件名
	FileSize       int64  //文件大小
	CertTime       int64  //认证时间
	CertTimeFormat string
}

/**
认证数据记录的序列化
*/
func (c CertRecord) SerializeRecord() ([]byte, error) {
	buff := new(bytes.Buffer)
	err := gob.NewEncoder(buff).Encode(c)
	return buff.Bytes(), err
}

/**
认证数据记录的反序列化
*/
func DeSerializeRecord(data []byte) (*CertRecord, error) {
	var certRceord *CertRecord
	err := gob.NewDecoder(bytes.NewReader(data)).Decode(&certRceord)
	return certRceord, err
}
