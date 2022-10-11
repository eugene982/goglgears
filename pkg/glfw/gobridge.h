#ifndef _GOBRIDGE_H_
#define _GOBRIDGE_H_

#include <GLFW/glfw3.h>

// glfwSetWindowSizeCallback

void extern
goBrigdeWindowsizefun(GLFWwindow* window, int width, int height);

void extern
goBridgeSetWindowSizeCallback(GLFWwindow *window);


// glfwSetKeyCallback

void extern
goBrigdeKeyfun(GLFWwindow* window, int key, int scancode, int action, int mods); 

void extern
goBridgeSetKeyCallback(GLFWwindow *window);

#endif