package BiometricsDriver

import (
	godbus "github.com/godbus/dbus/v5"
	systembiometrics "github.com/LingmoOS/golang-github-lingmo-go-dbus-factory/system/com.lingmo.daemon.authenticate.BiometricsDriver"
)

type BiometricsDriver = systembiometrics.BiometricsDriver

func NewBiometricsDriver(conn *godbus.Conn, serviceName string, path godbus.ObjectPath) (BiometricsDriver, error) {
	return systembiometrics.NewBiometricsDriver(conn, serviceName, path)
}
