package login1

import (
	godbus "github.com/godbus/dbus/v5"
	systemlogin1 "github.com/linuxdeepin/go-dbus-factory/system/org.freedesktop.login1"
)

type ScheduledShutdown = systemlogin1.ScheduledShutdown
type SessionInfo = systemlogin1.SessionInfo
type UserInfo = systemlogin1.UserInfo
type SeatInfo = systemlogin1.SeatInfo
type SessionDetail = systemlogin1.SessionDetail
type UserDetail = systemlogin1.UserDetail
type InhibitorInfo = systemlogin1.InhibitorInfo
type Manager = systemlogin1.Manager
type PropManagerScheduledShutdown = systemlogin1.PropManagerScheduledShutdown
type Seat = systemlogin1.Seat
type Session = systemlogin1.Session
type PropSessionUser = systemlogin1.PropSessionUser
type PropSessionSeat = systemlogin1.PropSessionSeat
type User = systemlogin1.User
type PropSessionInfo = systemlogin1.PropSessionInfo
type PropSessionInfoSlice = systemlogin1.PropSessionInfoSlice

func NewManager(conn *godbus.Conn) Manager {
	return systemlogin1.NewManager(conn)
}

func NewSeat(conn *godbus.Conn, path godbus.ObjectPath) (Seat, error) {
	return systemlogin1.NewSeat(conn, path)
}

func NewSession(conn *godbus.Conn, path godbus.ObjectPath) (Session, error) {
	return systemlogin1.NewSession(conn, path)
}

func NewUser(conn *godbus.Conn, path godbus.ObjectPath) (User, error) {
	return systemlogin1.NewUser(conn, path)
}
