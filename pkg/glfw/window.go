package glfw

//#cgo pkg-config: glfw3
//#include "gobridge.h"
import "C"
import (
	"unsafe"
)

type Window = C.GLFWwindow

type Monitor = C.GLFWmonitor

func CreateWindow(width int, height int, title string,
	monitor *Monitor, share *Window) *Window {

	pt := append([]byte(title), 0)

	return C.glfwCreateWindow(C.int(width), C.int(height),
		(*C.char)(unsafe.Pointer(&pt[0])),
		monitor, share)
}

func DestroyWindow(win *Window) {
	C.glfwDestroyWindow(win)
}

func WindowShouldClose(win *Window) bool {
	return C.glfwWindowShouldClose(win) != 0
}

func SwapBuffers(win *Window) {
	C.glfwSwapBuffers(win)
}

func MakeContextCurrent(win *Window) {
	C.glfwMakeContextCurrent(win)
}

////////////////////////////////////
// typedef void (* GLFWwindowsizefun)(GLFWwindow* window, int width, int height);
// GLFWAPI GLFWwindowsizefun glfwSetWindowSizeCallback(GLFWwindow* window,
//	GLFWwindowsizefun callback);

type Windowsizefun func(*Window, int, int)

var bindsWindowsizefun = make(map[*Window]Windowsizefun)

func SetWindowSizeCallback(win *Window, callback Windowsizefun) (prev Windowsizefun) {
	prev = bindsWindowsizefun[win]
	bindsWindowsizefun[win] = callback

	C.goBridgeSetWindowSizeCallback(win)
	return prev
}

//export goBrigdeWindowsizefun
func goBrigdeWindowsizefun(gwin *C.GLFWwindow, width C.int, height C.int) {
	callback, ok := bindsWindowsizefun[gwin]
	if ok {
		callback(gwin, int(width), int(height))
	}
}

//////////////////////////////////////////////
//typedef void (* GLFWkeyfun)(GLFWwindow*,int,int,int,int);
//GLFWAPI GLFWkeyfun glfwSetKeyCallback(GLFWwindow* window,
//	GLFWkeyfun callback)

const ( //action

)

type Keyfun func(*Window, int, int, int, int)

var bindsKeyfun = make(map[*Window]Keyfun)

func SetKeyCallback(win *Window, callback Keyfun) (prev Keyfun) {
	prev = bindsKeyfun[win]
	bindsKeyfun[win] = callback

	C.goBridgeSetKeyCallback(win)
	return prev
}

//export goBrigdeKeyfun
func goBrigdeKeyfun(gwin *C.GLFWwindow, key, scancode, action, mods C.int) {
	callback, ok := bindsKeyfun[gwin]
	if ok {
		callback(gwin, int(key), int(scancode), int(action), int(mods))
	}
}
