package fingerprint

import (
	godbus "github.com/godbus/dbus/v5"
	systemfp "github.com/linuxdeepin/go-dbus-factory/system/com.deepin.daemon.authenticate.fingerprint"
)

type Device = systemfp.Device
type CommonDevice = systemfp.CommonDevice
type MockDevice = systemfp.MockDevice
type MockInterfaceDevice = systemfp.MockInterfaceDevice
type MockCommonDevice = systemfp.MockCommonDevice
type MockInterfaceCommonDevice = systemfp.MockInterfaceCommonDevice

func NewDevice(conn *godbus.Conn, serviceName string, path godbus.ObjectPath) (Device, error) {
	return systemfp.NewDevice(conn, serviceName, path)
}

func NewCommonDevice(conn *godbus.Conn, serviceName string, path godbus.ObjectPath) (CommonDevice, error) {
	return systemfp.NewCommonDevice(conn, serviceName, path)
}
