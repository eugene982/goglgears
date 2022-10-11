package main

import (
	"fmt"
	"log"

	"goglgears/pkg/gear"
	"goglgears/pkg/gl"
	"goglgears/pkg/glfw"
)

var (
	fullscreen = false /* Create a single fullscreen window */
	stereo     = false /* Enable stereo.  */
	samples    = 0     /* Choose visual with at least N samples. */
	animate    = true  /* Animation */

	eyesep    float64 = 5.0  /* Eye separation. */
	fix_point float64 = 40.0 /* Fixation point distance.  */
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func onResize(win *glfw.Window, h, w int) {
	fmt.Println(win, h, w)
}

func run() error {

	var printInfo bool
	winWidth, winHeight := 800, 600

	if !glfw.Init() {
		return glfw.GetError()
	}

	defer glfw.Terminate()

	win := glfw.CreateWindow(winWidth, winHeight, "goglgears", nil, nil)
	if win == nil {
		return glfw.GetError()
	}
	glfw.MakeContextCurrent(win)
	glfw.SetWindowSizeCallback(win, onResize)

	if printInfo {
		fmt.Printf("GL_RENDERER   = %s\n", gl.GetString(gl.GL_RENDERER))
		fmt.Printf("GL_VERSION    = %s\n", gl.GetString(gl.GL_VERSION))
		fmt.Printf("GL_VENDOR     = %s\n", gl.GetString(gl.GL_VENDOR))
		fmt.Printf("GL_EXTENSIONS = %s\n", gl.GetString(gl.GL_EXTENSIONS))
	}

	glinit()

	/* Set initial projection/viewing transformation.
	 * We can't be sure we'll get a ConfigureNotify event when the window
	 * first appears.
	 */
	reshape(winWidth, winHeight)

	for !glfw.WindowShouldClose(win) {

		draw_frame(win)
		glfw.PollEvents()

		if err := glfw.GetError(); err != nil {
			break
		}
	}

	glfw.DestroyWindow(win)

	gl.DeleteLists(gear1, 1)
	gl.DeleteLists(gear2, 1)
	gl.DeleteLists(gear3, 1)

	return glfw.GetError()
}

var ( /* Stereo frustum params.  */
	left  float64
	right float64
	asp   float64
)

func reshape(width, height int) {
	gl.Viewport(0, 0, width, height)

	if stereo {
		var w float64

		asp = float64(height) / float64(width)
		w = fix_point * (1.0 / 5.0)

		left = -5.0 * ((w - 0.5*eyesep) / fix_point)
		right = 5.0 * ((w + 0.5*eyesep) / fix_point)
	} else {
		h := float64(height) / float64(width)

		gl.MatrixMode(gl.GL_PROJECTION)
		gl.LoadIdentity()
		gl.Frustum(-1.0, 1.0, -h, h, 5.0, 60.0)
	}

	gl.MatrixMode(gl.GL_MODELVIEW)
	gl.LoadIdentity()
	gl.Translatef(0.0, 0.0, -40.0)
}

var (
	gear1 uint
	gear2 uint
	gear3 uint
)

var (
	pos   = [4]float32{5.0, 5.0, 10.0, 0.0}
	red   = [4]float32{0.8, 0.1, 0.0, 1.0}
	green = [4]float32{0.0, 0.8, 0.2, 1.0}
	blue  = [4]float32{0.2, 0.2, 1.0, 1.0}
)

func glinit() {

	gl.Lightfv(gl.GL_LIGHT0, gl.GL_POSITION, &pos[0])
	gl.Enable(gl.GL_CULL_FACE)
	gl.Enable(gl.GL_LIGHTING)
	gl.Enable(gl.GL_LIGHT0)
	gl.Enable(gl.GL_DEPTH_TEST)

	/* make the gears */
	/* make the gears */
	gear1 = gl.GenLists(1)
	gl.NewList(gear1, gl.GL_COMPILE)
	gl.Materialfv(gl.GL_FRONT, gl.GL_AMBIENT_AND_DIFFUSE, &red[0])
	gear.Gear(1.0, 4.0, 1.0, 20, 0.7)
	gl.EndList()

	gear2 = gl.GenLists(1)
	gl.NewList(gear2, gl.GL_COMPILE)
	gl.Materialfv(gl.GL_FRONT, gl.GL_AMBIENT_AND_DIFFUSE, &green[0])
	gear.Gear(0.5, 2.0, 2.0, 10, 0.7)
	gl.EndList()

	gear3 = gl.GenLists(1)
	gl.NewList(gear3, gl.GL_COMPILE)
	gl.Materialfv(gl.GL_FRONT, gl.GL_AMBIENT_AND_DIFFUSE, &blue[0])
	gear.Gear(1.3, 2.0, 0.5, 10, 0.7)
	gl.EndList()

	gl.Enable(gl.GL_NORMALIZE)
}

var (
	view_rotx float32 = 20.0
	view_roty float32 = 30.0
	view_rotz float32 = 0.0
	angle     float32 = 0.0
)

func draw() {
	gl.Clear(gl.GL_COLOR_BUFFER_BIT | gl.GL_DEPTH_BUFFER_BIT)

	gl.PushMatrix()
	gl.Rotatef(view_rotx, 1.0, 0.0, 0.0)
	gl.Rotatef(view_roty, 0.0, 1.0, 0.0)
	gl.Rotatef(view_rotz, 0.0, 0.0, 1.0)

	gl.PushMatrix()
	gl.Translatef(-3.0, -2.0, 0.0)
	gl.Rotatef(angle, 0.0, 0.0, 1.0)
	gl.CallList(gear1)
	gl.PopMatrix()

	gl.PushMatrix()
	gl.Translatef(3.1, -2.0, 0.0)
	gl.Rotatef(-2.0*angle-9.0, 0.0, 0.0, 1.0)
	gl.CallList(gear2)
	gl.PopMatrix()

	gl.PushMatrix()
	gl.Translatef(-3.1, 4.2, 0.0)
	gl.Rotatef(-2.0*angle-25.0, 0.0, 0.0, 1.0)
	gl.CallList(gear3)
	gl.PopMatrix()

	gl.PopMatrix()
}

func draw_gears() {
	if stereo {
		/* First left eye.  */
		gl.DrawBuffer(gl.GL_BACK_LEFT)

		gl.MatrixMode(gl.GL_PROJECTION)
		gl.LoadIdentity()
		gl.Frustum(left, right, -asp, asp, 5.0, 60.0)

		gl.MatrixMode(gl.GL_MODELVIEW)

		gl.PushMatrix()
		gl.Translated(+0.5*eyesep, 0.0, 0.0)

		draw()

		gl.PopMatrix()

		/* Then right eye.  */
		gl.DrawBuffer(gl.GL_BACK_RIGHT)

		gl.MatrixMode(gl.GL_PROJECTION)
		gl.LoadIdentity()
		gl.Frustum(-right, -left, -asp, asp, 5.0, 60.0)

		gl.MatrixMode(gl.GL_MODELVIEW)

		gl.PushMatrix()
		gl.Translated(-0.5*eyesep, 0.0, 0.0)

		draw()

		gl.PopMatrix()

	} else {
		draw()
	}
}

var (
	frames         = 0
	tRot0  float64 = -1
	tRate0 float64 = -1
)

/** Draw single frame, do SwapBuffers, compute FPS */
func draw_frame(win *glfw.Window) {

	dt := glfw.GetTime()
	t := dt

	if tRot0 < 0.0 {
		tRot0 = t
	}
	dt = t - tRot0
	tRot0 = t

	if animate {
		/* advance rotation for next frame */
		angle += 70.0 * float32(dt) /* 70 degrees per second */
		if angle > 3600.0 {
			angle -= 3600.0
		}
	}

	draw_gears()
	glfw.SwapBuffers(win)

	frames++

	if tRate0 < 0.0 {
		tRate0 = t
	}
	if t-tRate0 >= 5.0 {
		seconds := t - tRate0
		fps := float64(frames) / seconds
		fmt.Printf("%d frames in %3.1f seconds = %6.3f FPS\n", frames, seconds,
			fps)
		//fflush(stdout)
		tRate0 = t
		frames = 0
	}
}
