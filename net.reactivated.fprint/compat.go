package fprint

import (
	godbus "github.com/godbus/dbus/v5"
	systemfprint "github.com/linuxdeepin/go-dbus-factory/system/net.reactivated.fprint"
)

type ObjectManager = systemfprint.ObjectManager
type Manager = systemfprint.Manager
type Device = systemfprint.Device
type MockObjectManager = systemfprint.MockObjectManager
type MockManager = systemfprint.MockManager
type MockInterfaceManager = systemfprint.MockInterfaceManager
type MockDevice = systemfprint.MockDevice
type MockInterfaceDevice = systemfprint.MockInterfaceDevice

func NewObjectManager(conn *godbus.Conn) ObjectManager {
	return systemfprint.NewObjectManager(conn)
}

func NewManager(conn *godbus.Conn) Manager {
	return systemfprint.NewManager(conn)
}

func NewDevice(conn *godbus.Conn, path godbus.ObjectPath) (Device, error) {
	return systemfprint.NewDevice(conn, path)
}
