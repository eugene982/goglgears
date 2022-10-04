package glfw

// #cgo pkg-config: glfw3
// #include <GLFW/glfw3.h>
import "C"
import "unsafe"

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
