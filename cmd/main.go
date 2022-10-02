package main

import (
	"log"

	"goglgears/pkg/gl"
	"goglgears/pkg/glfw"
)

var (
	gear1 Gear
	gear2 Gear
	gear3 Gear
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

	win := glfw.CreateWindow(600, 400, "goglgears", nil, nil)
	if win == nil {
		return glfw.GetError()
	}

	defer func() {
		glfw.DestroyWindow(win)
	}()

	var err error

	glinit()

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

func glinit() {
	var (
		pos   = [4]float32{5.0, 5.0, 10.0, 0.0}
		red   = [4]float32{0.8, 0.1, 0.0, 1.0}
		green = [4]float32{0.0, 0.8, 0.2, 1.0}
		blue  = [4]float32{0.2, 0.2, 1.0, 1.0}
	)
	gl.Lightfv(gl.GL_LIGHT0, gl.GL_POSITION, &pos[0])
	gl.Enable(gl.GL_CULL_FACE)
	gl.Enable(gl.GL_LIGHTING)
	gl.Enable(gl.GL_LIGHT0)
	gl.Enable(gl.GL_DEPTH_TEST)

	/* make the gears */
	gear1 = NewGear(1.0, 4.0, 1.0, 20, 0.7, red)
	gear2 = NewGear(0.5, 2.0, 2.0, 10, 0.7, green)
	gear3 = NewGear(1.3, 2.0, 0.5, 10, 0.7, blue)

	gl.Enable(gl.GL_NORMALIZE)

}

/*
func draw() {
	glClear(GL_COLOR_BUFFER_BIT | GL_DEPTH_BUFFER_BIT)

	glPushMatrix()
	glRotatef(view_rotx, 1.0, 0.0, 0.0)
	glRotatef(view_roty, 0.0, 1.0, 0.0)
	glRotatef(view_rotz, 0.0, 0.0, 1.0)

	glPushMatrix()
	glTranslatef(-3.0, -2.0, 0.0)
	glRotatef(angle, 0.0, 0.0, 1.0)
	glCallList(gear1)
	glPopMatrix()

	glPushMatrix()
	glTranslatef(3.1, -2.0, 0.0)
	glRotatef(-2.0*angle-9.0, 0.0, 0.0, 1.0)
	glCallList(gear2)
	glPopMatrix()

	glPushMatrix()
	glTranslatef(-3.1, 4.2, 0.0)
	glRotatef(-2.0*angle-25.0, 0.0, 0.0, 1.0)
	glCallList(gear3)
	glPopMatrix()

	glPopMatrix()
}
*/
