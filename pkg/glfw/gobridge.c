#include "gobridge.h" 

void
goBridgeSetWindowSizeCallback(GLFWwindow *win) {
	glfwSetWindowSizeCallback(win, goBrigdeWindowsizefun);
}