package authenticate

import (
	godbus "github.com/godbus/dbus/v5"
	systemauth "github.com/linuxdeepin/go-dbus-factory/system/com.deepin.daemon.authenticate"
)

type CharaInfo = systemauth.CharaInfo
type DriverInfo = systemauth.DriverInfo
type Authenticate = systemauth.Authenticate
type Fingerprint = systemauth.Fingerprint
type UKey = systemauth.UKey
type Face = systemauth.Face
type CharaManger = systemauth.CharaManger

func NewAuthenticate(conn *godbus.Conn) Authenticate {
	return systemauth.NewAuthenticate(conn)
}

func NewFingerprint(conn *godbus.Conn) Fingerprint {
	return systemauth.NewFingerprint(conn)
}

func NewUKey(conn *godbus.Conn) UKey {
	return systemauth.NewUKey(conn)
}

func NewFace(conn *godbus.Conn) Face {
	return systemauth.NewFace(conn)
}

func NewCharaManger(conn *godbus.Conn) CharaManger {
	return systemauth.NewCharaManger(conn)
}
