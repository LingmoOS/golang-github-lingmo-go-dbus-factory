package ukey

import (
	godbus "github.com/godbus/dbus/v5"
	systemukey "github.com/linuxdeepin/go-dbus-factory/system/com.deepin.daemon.authenticate.ukey"
)

type UKey = systemukey.UKey

func NewUKey(conn *godbus.Conn, serviceName string, path godbus.ObjectPath) (UKey, error) {
	return systemukey.NewUKey(conn, serviceName, path)
}
