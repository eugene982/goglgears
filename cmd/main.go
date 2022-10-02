package main

import (
	"log"

	"goglgears/pkg/glfw"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {

	if !glfw.Init() {
		return glfw.GetError()
	}

	defer glfw.Terminate()

	win := glfw.CreateWindow(300, 300, "goglgears", nil, nil)
	if win == nil {
		return glfw.GetError()
	}

	defer func() {
		glfw.DestroyWindow(win)
	}()

	var err error
	for !glfw.WindowShouldClose(win) {

		glfw.SwapBuffers(win)
		glfw.PollEvents()

		err = glfw.GetError()
		if err != nil {
			return err
		}
	}

	return glfw.GetError()
}
