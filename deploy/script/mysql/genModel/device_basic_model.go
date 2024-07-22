package genModel

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ DeviceBasicModel = (*customDeviceBasicModel)(nil)

type (
	// DeviceBasicModel is an interface to be customized, add more methods here,
	// and implement the added methods in customDeviceBasicModel.
	DeviceBasicModel interface {
		deviceBasicModel
	}

	customDeviceBasicModel struct {
		*defaultDeviceBasicModel
	}
)

// NewDeviceBasicModel returns a model for the database table.
func NewDeviceBasicModel(conn sqlx.SqlConn, c cache.CacheConf) DeviceBasicModel {
	return &customDeviceBasicModel{
		defaultDeviceBasicModel: newDeviceBasicModel(conn, c),
	}
}
