package main

import (
	"flag"
	"fmt"
	"log"
	"strings"

	"goglgears/cmd/gear"
	"goglgears/pkg/gl"
	"goglgears/pkg/glfw"
)

var (
	fullscreen = false /* Create a single fullscreen window */
	stereo     = false /* Enable stereo.  */
	samples    = 0     /* Choose visual with at least N samples. */
	animate    = true  /* Animation */
	printInfo  = false
	winWidth   = 300
	winHeight  = 300
)

func main() {

	var geometry string

	flag.BoolVar(&stereo, "stereo", false, "run in stereo mode")
	flag.IntVar(&samples, "samples", 0, "run in multisample mode with at least N samples")
	flag.BoolVar(&printInfo, "info", false, "display OpenGL renderer info")
	flag.StringVar(&geometry, "geometry", "", "WxH window geometry")
	flag.Parse()

	if geometry != "" {
		_, err := fmt.Fscanf(strings.NewReader(geometry), "%dx%d", &winWidth, &winHeight)
		if err != nil {
			log.Fatal("error parse geometry: ", err)
		}
	}

	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {

	if !glfw.Init() {
		return glfw.GetError()
	}
	defer glfw.Terminate()

	// set window hints
	if stereo {
		glfw.WindowHint(glfw.GLFW_STEREO, glfw.GLFW_TRUE)
		if err := glfw.GetError(); err != nil {
			return err
		}
	}

	if samples > 0 {
		glfw.WindowHint(glfw.GLFW_SAMPLES, samples)
		if err := glfw.GetError(); err != nil {
			return err
		}
	}

	win := glfw.CreateWindow(winWidth, winHeight, "goglgears", nil, nil)
	if win == nil {
		return glfw.GetError()
	}
	defer glfw.DestroyWindow(win)

	glfw.MakeContextCurrent(win)
	glfw.SetWindowSizeCallback(win, onResize)
	glfw.SetKeyCallback(win, onKey)

	if printInfo {
		fmt.Printf("GL_RENDERER   = %s\n", gl.GetString(gl.GL_RENDERER))
		fmt.Printf("GL_VERSION    = %s\n", gl.GetString(gl.GL_VERSION))
		fmt.Printf("GL_VENDOR     = %s\n", gl.GetString(gl.GL_VENDOR))
		fmt.Printf("GL_EXTENSIONS = %s\n", gl.GetString(gl.GL_EXTENSIONS))
	}

	delGears := gear.GearInit(stereo)
	defer delGears()

	/* Set initial projection/viewing transformation.
	 * We can't be sure we'll get a ConfigureNotify event when the window
	 * first appears.
	 */
	gear.Reshape(winWidth, winHeight)

	for !glfw.WindowShouldClose(win) {

		drawFrame(win)
		glfw.PollEvents()

		if err := glfw.GetError(); err != nil {
			return err
		}
	}

	return glfw.GetError()
}

func onResize(win *glfw.Window, width, height int) {
	gear.Reshape(width, height)
}

var (
	view_rotx float32 = 20.0
	view_roty float32 = 30.0
	view_rotz float32 = 0.0
)

func onKey(win *glfw.Window, key, scancode, action, mods int) {
	if action == glfw.GLFW_RELEASE {
		switch key {
		case glfw.GLFW_KEY_A:
			animate = !animate
		}
	} else {
		switch key {
		case glfw.GLFW_KEY_LEFT:
			view_roty += 5.0
		case glfw.GLFW_KEY_RIGHT:
			view_roty -= 5.0
		case glfw.GLFW_KEY_UP:
			view_rotx += 5.0
		case glfw.GLFW_KEY_DOWN:
			view_rotx -= 5.0
		}
	}
}

var (
	frames         = 0
	tRot0  float64 = -1
	tRate0 float64 = -1
	angle  float32 = 0.0
)

/** Draw single frame, do SwapBuffers, compute FPS */
func drawFrame(win *glfw.Window) {

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

	gear.DrawGears(angle, view_rotx, view_roty, view_rotz)
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
