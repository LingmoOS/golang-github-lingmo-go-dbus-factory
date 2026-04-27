package BiometricsDriver

import (
	godbus "github.com/godbus/dbus/v5"
	systembiometrics "github.com/linuxdeepin/go-dbus-factory/system/com.deepin.daemon.authenticate.BiometricsDriver"
)

type BiometricsDriver = systembiometrics.BiometricsDriver

func NewBiometricsDriver(conn *godbus.Conn, serviceName string, path godbus.ObjectPath) (BiometricsDriver, error) {
	return systembiometrics.NewBiometricsDriver(conn, serviceName, path)
}
