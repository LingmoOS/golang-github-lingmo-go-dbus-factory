package fingerprint

import (
	godbus "github.com/godbus/dbus/v5"
	systemhuawei "github.com/linuxdeepin/go-dbus-factory/system/com.huawei.fingerprint"
)

type Fingerprint = systemhuawei.Fingerprint
type MockFingerprint = systemhuawei.MockFingerprint
type MockInterfaceFingerprint = systemhuawei.MockInterfaceFingerprint

func NewFingerprint(conn *godbus.Conn) Fingerprint {
	return systemhuawei.NewFingerprint(conn)
}
