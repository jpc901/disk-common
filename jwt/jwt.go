package jwt

import (
	"fmt"

	"github.com/jpc901/disk-common/util"
)

// CreateToken 根据uuid和有效期截止时间生成token
func CreateToken(uuid, timeStamp int64) (string, error) {
	return util.RsaEncrypt(fmt.Sprintf("%d:%d", uuid, timeStamp))
}

func ParseToken(token string) (uuid, timeStamp int64, err error) {
	parseToken, err := util.RsaDecrypt(token)
	if err != nil {
		return
	}

	_, err = fmt.Sscanf(parseToken, "%d:%d", &uuid, &timeStamp)
	return
}