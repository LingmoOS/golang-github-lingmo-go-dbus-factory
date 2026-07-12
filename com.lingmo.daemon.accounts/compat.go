package accounts

import (
	godbus "github.com/godbus/dbus/v5"
	systemaccounts "github.com/LingmoOS/golang-github-lingmo-go-dbus-factory/system/org.lingmo.accounts1"
)

type LoginUtmpx = systemaccounts.LoginUtmpx
type LoginReminderInfo = systemaccounts.LoginReminderInfo
type Accounts = systemaccounts.Accounts
type User = systemaccounts.User

func NewAccounts(conn *godbus.Conn) Accounts {
	return systemaccounts.NewAccounts(conn)
}

func NewUser(conn *godbus.Conn, path godbus.ObjectPath) (User, error) {
	return systemaccounts.NewUser(conn, path)
}
