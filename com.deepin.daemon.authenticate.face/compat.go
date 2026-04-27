package face

import (
	godbus "github.com/godbus/dbus/v5"
	systemface "github.com/linuxdeepin/go-dbus-factory/system/com.deepin.daemon.authenticate.face"
)

type Face = systemface.Face

func NewFace(conn *godbus.Conn, serviceName string, path godbus.ObjectPath) (Face, error) {
	return systemface.NewFace(conn, serviceName, path)
}
