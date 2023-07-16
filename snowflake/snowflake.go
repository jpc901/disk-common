package snowflake

import (
	"github.com/GUAIK-ORG/go-snowflake/snowflake"
)

// GetId 通过两个number生成一个唯一的id
func GetId(datacenterid, workerid int64) (int64, error) {
	s, err := snowflake.NewSnowflake(datacenterid, workerid)
	if err != nil {
		return -1, err
	}
	return s.NextVal(), nil
}