package util

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
-----BEGIN PRIVATE KEY-----
MIICdwIBADANBgkqhkiG9w0BAQEFAASCAmEwggJdAgEAAoGBAKPiHpijYX5av9rZ
/vkcJlEvomjRhULbemfd+Lssuu5OZhqfvRWRDNywmoB1v8VnZEp7WXWuPdNow5Zn
wrZQzg5WVVwbsxZFSyL6zyAhEHbFxVBmZ9seYte90YFLL2+5JVmPfAewoHHj9nHN
GIT+f62pXJL6pQM059V1bPc9s0ktAgMBAAECgYAp7tTHV569lvjaRcdQ9Fv1kAut
aFcByWjjpM9dDU/zZAoKK+9U0l2JJoMR9Y8RMqhdynwUeXYGXxsUpA4VWk2uwnrf
xvjbbq15imF92L8YLb6YXE7iUAfMUTmARBMK/wWP6RmQ3rUVNoWVnUjt6j0krapF
mlUd8cNk+jOch+tWwQJBANiN53oKdnXIka7cGTNp/oG4WibUzfnuz7ThNGZonGeN
kXvgDPxJ0BOohjWOhzNCTDH4jKzm0dM68Qx4iXayHtECQQDBvB/ExtR8flXh8JyM
olGJqxKw3JRXfJcn4tvK2Y5cKfUCqA4uFmxjRpsMk8I81czl44JVNmFGDbMXQYr1
VvOdAkEAhDVma8i1d8VCw/mV3SDKA9JUH77uHbeh0XFod1lIm6P/fRxVcTVzNn09
qrbgbff84sk2wVyOH6KthYqVigTG8QJBAIydzzE0X+Y8jHmB+x7YcfZKhTZ54/Hc
LJp2vrFtVzbt/TgAYspw3BrylHd8h+8//4icqWzQG6qNJwAqQoHwqsECQBxeb6/a
rHgmjSw0/PTeIk+tL402NTHlJyU7yV9fV5fttgxne/PiaIdAAq0TsR/aXmYq3t7U
mSbKRLFo2kE2YQI=
-----END PRIVATE KEY-----
`)

// 公钥
var publicKey = []byte(`
-----BEGIN PUBLIC KEY-----
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQCj4h6Yo2F+Wr/a2f75HCZRL6Jo
0YVC23pn3fi7LLruTmYan70VkQzcsJqAdb/FZ2RKe1l1rj3TaMOWZ8K2UM4OVlVc
G7MWRUsi+s8gIRB2xcVQZmfbHmLXvdGBSy9vuSVZj3wHsKBx4/ZxzRiE/n+tqVyS
+qUDNOfVdWz3PbNJLQIDAQAB
-----END PUBLIC KEY-----
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