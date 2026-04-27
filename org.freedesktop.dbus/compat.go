package dbus

import (
	godbus "github.com/godbus/dbus/v5"
	systemdbus "github.com/linuxdeepin/go-dbus-factory/system/org.freedesktop.dbus"
)

type DBus = systemdbus.DBus
type MockDBus = systemdbus.MockDBus
type MockInterfaceDbusIfc = systemdbus.MockInterfaceDbusIfc

func NewDBus(conn *godbus.Conn) DBus {
	return systemdbus.NewDBus(conn)
}
