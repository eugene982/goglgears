package glfw

// #cgo pkg-config: glfw3
// #include <GLFW/glfw3.h>
import "C"

const (
	GLFW_TRUE  = C.GLFW_TRUE
	GLFW_FALSE = C.GLFW_FALSE
)

func Init() bool {
	return C.glfwInit() == GLFW_TRUE
}

func Terminate() {
	C.glfwTerminate()
}

func PollEvents() {
	C.glfwPollEvents()
}
