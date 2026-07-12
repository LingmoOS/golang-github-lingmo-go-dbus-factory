package face

import (
	godbus "github.com/godbus/dbus/v5"
	systemface "github.com/LingmoOS/golang-github-lingmo-go-dbus-factory/system/com.lingmo.daemon.authenticate.face"
)

type Face = systemface.Face

func NewFace(conn *godbus.Conn, serviceName string, path godbus.ObjectPath) (Face, error) {
	return systemface.NewFace(conn, serviceName, path)
}
