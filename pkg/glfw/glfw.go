package glfw

// #cgo pkg-config: glfw3
// #include <GLFW/glfw3.h>
import "C"

const (
	TRUE  = C.GLFW_TRUE
	FALSE = C.GLFW_FALSE

	PRESS   = C.GLFW_PRESS
	RELEASE = C.GLFW_RELEASE
	REPEAT  = C.GLFW_REPEAT
)

const ( // hints
	DOUBLEBUFFER = C.GLFW_DOUBLEBUFFER
	STEREO       = C.GLFW_STEREO
	SAMPLES      = C.GLFW_SAMPLES
)

const ( /* Printable keys */
	KEY_UNKNOWN = C.GLFW_KEY_UNKNOWN

	KEY_ESCAPE = C.GLFW_KEY_ESCAPE
	KEY_A      = C.GLFW_KEY_A

	KEY_RIGHT = C.GLFW_KEY_RIGHT
	KEY_LEFT  = C.GLFW_KEY_LEFT
	KEY_DOWN  = C.GLFW_KEY_DOWN
	KEY_UP    = C.GLFW_KEY_UP
)

type Monitor = C.GLFWmonitor

func Init() error {
	if C.glfwInit() == TRUE {
		return nil
	}
	return GetError()
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

func WindowHint(hint C.int, value int) error {
	C.glfwWindowHint(hint, C.int(value))
	return GetError()
}

func DefaultWindowHints() {
	C.glfwDefaultWindowHints()
}

func GetPrimaryMonitor() *Monitor {
	return C.glfwGetPrimaryMonitor()
}
