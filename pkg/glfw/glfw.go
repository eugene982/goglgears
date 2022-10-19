package glfw

// #cgo pkg-config: glfw3
// #include <GLFW/glfw3.h>
import "C"

const (
	GLFW_TRUE  = C.GLFW_TRUE
	GLFW_FALSE = C.GLFW_FALSE

	GLFW_PRESS   = C.GLFW_PRESS
	GLFW_RELEASE = C.GLFW_RELEASE
	GLFW_REPEAT  = C.GLFW_REPEAT
)

const ( // hints
	GLFW_DOUBLEBUFFER = C.GLFW_DOUBLEBUFFER
	GLFW_STEREO       = C.GLFW_STEREO
	GLFW_SAMPLES      = C.GLFW_SAMPLES
)

const ( /* Printable keys */
	GLFW_KEY_UNKNOWN = C.GLFW_KEY_UNKNOWN

	GLFW_KEY_A = C.GLFW_KEY_A

	GLFW_KEY_RIGHT = C.GLFW_KEY_RIGHT
	GLFW_KEY_LEFT  = C.GLFW_KEY_LEFT
	GLFW_KEY_DOWN  = C.GLFW_KEY_DOWN
	GLFW_KEY_UP    = C.GLFW_KEY_UP
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

func GetTime() float64 {
	return float64(C.glfwGetTime())
}

func WindowHint(hint C.int, value int) {
	C.glfwWindowHint(hint, C.int(value))
}

func DefaultWindowHints() {
	C.glfwDefaultWindowHints()
}
