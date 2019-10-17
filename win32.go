/**
 * Auth :   liubo
 * Date :   2019/10/17 9:09
 * Comment:
 */

package main

import (
	"fmt"
	"golang.org/x/sys/windows"
	"syscall"
	"unsafe"
)

type (
	DWORD uint32
	WPARAM uintptr
	LPARAM uintptr
	LRESULT uintptr
	HANDLE uintptr
	HINSTANCE HANDLE
	HHOOK HANDLE
	HWND HANDLE
)

var (
	sendMessage                 *windows.LazyProc
	sendMessageTimeout          *windows.LazyProc
)
const (
	HWND_BROADCAST = HWND(0xFFFF)
	WM_SETTINGCHANGE          = 26
	SMTO_ABORTIFHUNG = 0x0002
)

func init() {
	libuser32 := windows.NewLazySystemDLL("user32.dll")
	sendMessage = libuser32.NewProc("SendMessageW")
	sendMessageTimeout = libuser32.NewProc("SendMessageTimeoutW")
}

func SendMessage(hWnd HWND, msg uint32, wParam, lParam uintptr) uintptr {
	ret, _, _ := syscall.Syscall6(sendMessage.Addr(), 4,
		uintptr(hWnd),
		uintptr(msg),
		wParam,
		lParam,
		0,
		0)

	return ret
}
func SendMessageTimeout(hWnd HWND, msg uint32, wParam, lParam uintptr, fuFlags, uTimeout uint, lpdwResult uintptr) uintptr {
	ret, _, err := syscall.Syscall9(sendMessageTimeout.Addr(), 7,
		uintptr(hWnd),
		uintptr(msg),
		wParam,
		lParam,
		uintptr(fuFlags),
		uintptr(uTimeout),
		lpdwResult,
		0,
		0)

	if err != syscall.Errno(0) {
		fmt.Println(err.Error())
	}

	return ret
}

func RefreshRegister() {
	var dwReturnValue uint
	str, _ := syscall.UTF16PtrFromString( "Environment")
	SendMessageTimeout(HWND_BROADCAST, WM_SETTINGCHANGE,
		uintptr(0), uintptr(unsafe.Pointer(str)),
		SMTO_ABORTIFHUNG, 5000, uintptr(unsafe.Pointer(&dwReturnValue)))
}