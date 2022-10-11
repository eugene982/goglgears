#include "gobridge.h" 

void
goBridgeSetWindowSizeCallback(GLFWwindow *window) {
	glfwSetWindowSizeCallback(window, goBrigdeWindowsizefun);
}

void
goBridgeSetKeyCallback(GLFWwindow *window) {
	glfwSetKeyCallback(window, goBrigdeKeyfun);
}