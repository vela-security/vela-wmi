//go:build windows
// +build windows

package wmi

import (
	"github.com/go-ole/go-ole"
	"github.com/go-ole/go-ole/oleutil"
	"github.com/vela-security/vela-public/assert"
	"github.com/vela-security/vela-public/export"
	"github.com/vela-security/vela-public/lua"
	"runtime"
)

var xEnv assert.Environment

func callback(L *lua.LState, src *ole.IDispatch, fn *lua.LFunction) error {
	count, err := oleInt64(src, "Count")
	if err != nil {
		return err
	}

	co := xEnv.Clone(L)
	defer xEnv.Free(co)
	cp := xEnv.P(fn)

	err = oleutil.ForEach(src, func(v *ole.VARIANT) error {
		defer v.Clear()

		a := newAnt(v)
		e := co.CallByParam(cp, a.toLV(co), lua.LNumber(count))
		if e != nil {
			return e
		}

		if a.err != nil {
			return err
		}

		return nil
	})
	return err
}

func LQuery(L *lua.LState, query string, fn *lua.LFunction, connectServerArgs ...interface{}) error {
	lock.Lock()
	defer lock.Unlock()

	runtime.LockOSThread()
	defer runtime.UnlockOSThread()

	service, cleanup, err := DefaultClient.coinitService(connectServerArgs...)
	if err != nil {
		return err
	}
	defer cleanup()

	// result is a SWBemObjectSet
	resultRaw, err := oleutil.CallMethod(service, "ExecQuery", query)
	if err != nil {
		return err
	}
	defer resultRaw.Clear()

	result := resultRaw.ToIDispatch()

	return callback(L, result, fn)
}

func wmiLuaQuery(L *lua.LState) int {
	query := L.CheckString(1)
	fn := L.CheckFunction(2)
	err := LQuery(L, query, fn)
	if err != nil {
		L.Pushf("%v", err)
	} else {
		L.Push(lua.LNil)
	}
	return 1
}

func WithEnv(env assert.Environment) {
	xEnv = env
	kv := lua.NewUserKV()
	kv.Set("query", lua.NewFunction(wmiLuaQuery))
	xEnv.Set("wmi", export.New("vela.wmi.export", export.WithTable(kv)))
}
