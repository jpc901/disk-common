package utils

import (
	"crypto/md5"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/x509"
	"encoding/hex"
	"encoding/pem"
	"errors"
	"hash"
	"io"
	"os"
	"path/filepath"
)

type Sha1Stream struct {
	_sha1 hash.Hash
}

func (obj *Sha1Stream) Update(data []byte) {
	if obj._sha1 == nil {
		obj._sha1 = sha1.New()
	}
	obj._sha1.Write(data)
}

func (obj *Sha1Stream) Sum() string {
	return hex.EncodeToString(obj._sha1.Sum([]byte("")))
}

func Sha1(data []byte) string {
	_sha1 := sha1.New()
	_sha1.Write(data)
	return hex.EncodeToString(_sha1.Sum([]byte("")))
}

func FileSha1(file *os.File) string {
	_sha1 := sha1.New()
	io.Copy(_sha1, file)
	return hex.EncodeToString(_sha1.Sum(nil))
}

func MD5(data []byte) string {
	_md5 := md5.New()
	_md5.Write(data)
	return hex.EncodeToString(_md5.Sum([]byte("")))
}

func FileMD5(file *os.File) string {
	_md5 := md5.New()
	io.Copy(_md5, file)
	return hex.EncodeToString(_md5.Sum(nil))
}

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func GetFileSize(filename string) int64 {
	var result int64
	filepath.Walk(filename, func(path string, f os.FileInfo, err error) error {
		result = f.Size()
		return nil
	})
	return result
}

var privateKey = []byte(`
-----BEGIN RSA PRIVATE KEY-----
MIICWwIBAAKBgQCvknv51fEkQjjD1k27HuqZU12H2mvgtsQ2vLvAHG3c7xdYqrLg
YvZXh7bKV87n5sZEh4BwuMswxqzoxQqX3DdMPxcwrTMEnv0CYoAtyYLFH9Bv/6yD
i0v1JnCVUVOSTAPa9Fc2HGq+Bf7rrd7qPhRLxz6o5mi6fJ89+PEpszuGhQIDAQAB
AoGAb5RWufZfZD25uawOhmcljP/QJzCG8q70kBzt7S+kGo7QdYc2WyhJimMzpfCc
DAE76/15gSnP6FW7OctP6icH9Jzd9gMroxEBd2PXxrOdlq0wAVVAfmaMs/rpnZTM
dMHNUwAZARpJqeNWcRa4IxcUsBKmaL4HxazaQjf4OKWakoECQQDjMmB8tdN4x5t4
BF/faiKezFfO+U7fBG5ipAMk1hiFsFM3sdLIHJtcvtI6lvyxEXiOZSayrOflhAJ/
iu1pfY4RAkEAxdSkhJyTCKc1GNVBrjpaOqjQ/oV6bcmn7vlswGkb6q0K1QXa/jBJ
gCc0IY/6VxujLFH4e1lbsHeR9BeeVmJNNQJBAJCWwukbHlZDUiHzRsB8X0QIb+l8
qEEuJMIJ9yY+SqTqLkvHk4lfC1De8BPxeyeFIuAcZ6BWgc6DUMOyupzkFsECPxA6
YUR/k5AiJzjiRYEFSGGHd51pVaGr6RqxWzptZNzbVQgkctJnI6Bflucp6F885SW9
k6SKr/rJ1C8xwMtVRQJADutuk6vhhD+uLVWxOdO6lBPwtAvvMBbCx6Sc4SmeJCDP
Szs2SS2VPLyfrkR/Fp2T3y7t8s9xHQ3Kko4qVwHizQ==
-----END RSA PRIVATE KEY-----
`)

// 公钥
var publicKey = []byte(`
-----BEGIN RSA PUBLIC KEY-----
MIGJAoGBAK+Se/nV8SRCOMPWTbse6plTXYfaa+C2xDa8u8AcbdzvF1iqsuBi9leH
tspXzufmxkSHgHC4yzDGrOjFCpfcN0w/FzCtMwSe/QJigC3JgsUf0G//rIOLS/Um
cJVRU5JMA9r0VzYcar4F/uut3uo+FEvHPqjmaLp8nz348SmzO4aFAgMBAAE=
-----END RSA PUBLIC KEY-----
`)

// 加密
func RsaEncrypt(str string) (string, error) {
	block, _ := pem.Decode(publicKey) //将密钥解析成公钥实例
	if block == nil {
		return "", errors.New("public key error")
	}
	pubInterface, err := x509.ParsePKCS1PublicKey(block.Bytes) //解析pem.Decode（）返回的Block指针实例
	if err != nil {
		return "", err
	}
	encStrBytes, err := rsa.EncryptPKCS1v15(rand.Reader, pubInterface, []byte(str)) //RSA算法加密
	return hex.EncodeToString(encStrBytes), err
}

func RsaDecrypt(encStr string) (string, error) {
	block, _ := pem.Decode(privateKey) //将密钥解析成私钥实例
	if block == nil {
		return "", errors.New("private key error!")
	}
	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes) //解析pem.Decode（）返回的Block指针实例
	if err != nil {
		return "", err
	}

	encStrBytes, err := hex.DecodeString(encStr)
	decStr, err := rsa.DecryptPKCS1v15(rand.Reader, priv, encStrBytes) //RSA算法解密
	return string(decStr), err
}