package glfw

// #cgo pkg-config: glfw3
// #include <GLFW/glfw3.h>
import "C"
import "fmt"

const (
	GLFW_NO_ERROR = C.GLFW_NO_ERROR
)

type GlfwError struct {
	code int
	desc string
}

func (e GlfwError) Error() string {
	return fmt.Sprintf("glfv error (0x%08x): %s", e.code, e.desc)
}

func GetError() error {
	var (
		code C.int
		desc *C.char
	)

	code = C.glfwGetError(&desc)
	if code == GLFW_NO_ERROR {
		return nil
	}
	return &GlfwError{
		code: int(code),
		desc: C.GoString(desc),
	}
}
