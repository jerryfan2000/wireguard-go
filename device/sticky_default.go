//go:build !linux

package device

import (
	"github.com/jerryfan2000/wireguard-go/conn"
	"github.com/jerryfan2000/wireguard-go/rwcancel"
)

func (device *Device) startRouteListener(bind conn.Bind) (*rwcancel.RWCancel, error) {
	return nil, nil
}
