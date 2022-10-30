package main

import (
	"flag"
	"fmt"
	"log"
	"runtime"
	"strings"

	"goglgears/cmd/gear"
	"goglgears/pkg/gl"
	"goglgears/pkg/glfw"
)

var (
	fullscreenFlag = false /* Create a single fullscreen window */
	stereoFlag     = false /* Enable stereo.  */
	samplesFlag    = 0     /* Choose visual with at least N samples. */
	animateFlag    = true  /* Animation */
	printInfoFlag  = false
	winWidthFlag   = 300
	winHeightFlag  = 300
)

func init() {
	// Иначе падает через несколько секунд.
	runtime.LockOSThread()
}

func main() {

	var geometry string
	flag.BoolVar(&fullscreenFlag, "fullscreen", false, "run in fullscreen mode")
	flag.BoolVar(&stereoFlag, "stereo", false, "run in stereo mode")
	flag.IntVar(&samplesFlag, "samples", 0, "run in multisample mode with at least N samples")
	flag.BoolVar(&printInfoFlag, "info", false, "display OpenGL renderer info")
	flag.StringVar(&geometry, "geometry", "", "WxH window geometry")
	flag.Parse()

	if geometry != "" {
		_, err := fmt.Fscanf(strings.NewReader(geometry), "%dx%d", &winWidthFlag, &winHeightFlag)
		if err != nil {
			log.Fatal("error parse geometry: ", err)
		}
	}

	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	err := glfw.Init()
	if err != nil {
		return err
	}
	defer glfw.Terminate()

	win, err := createWindow("goglgears")
	if err != nil {
		return err
	}
	defer glfw.DestroyWindow(win)

	if printInfoFlag {
		printInfo()
	}

	delGears := gear.Init(stereoFlag)
	defer delGears()

	/* Set initial projection/viewing transformation.
	 * We can't be sure we'll get a ConfigureNotify event when the window
	 * first appears.
	 */
	gear.Reshape(winWidthFlag, winHeightFlag)

	// main loop
	for !glfw.WindowShouldClose(win) {
		drawFrame(win)
		glfw.PollEvents()
	}

	return nil
}

func createWindow(title string) (*glfw.Window, error) {

	// set window hints
	if stereoFlag {
		err := glfw.WindowHint(glfw.STEREO, glfw.TRUE)
		if err != nil {
			return nil, err
		}
	}

	if samplesFlag > 0 {
		err := glfw.WindowHint(glfw.SAMPLES, samplesFlag)
		if err != nil {
			return nil, err
		}
	}

	var pmon *glfw.Monitor
	if fullscreenFlag {
		pmon = glfw.GetPrimaryMonitor()
	}

	win, err := glfw.CreateWindow(winWidthFlag, winHeightFlag, title,
		pmon, nil)
	if err != nil {
		return nil, err
	}

	glfw.MakeContextCurrent(win)
	glfw.SetWindowSizeCallback(win, onResize)
	glfw.SetKeyCallback(win, onKey)
	return win, nil
}

func printInfo() {
	fmt.Println("GL_RENDERER   = ", gl.GetString(gl.GL_RENDERER))
	fmt.Println("GL_VERSION    = ", gl.GetString(gl.GL_VERSION))
	fmt.Println("GL_VENDOR     = ", gl.GetString(gl.GL_VENDOR))
	fmt.Println("GL_EXTENSIONS = ", gl.GetString(gl.GL_EXTENSIONS))
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
	if action == glfw.RELEASE {
		switch key {
		case glfw.KEY_A:
			animateFlag = !animateFlag
		case glfw.KEY_ESCAPE:
			glfw.SetWindowShouldClose(win)
		}
		// Если был запуск в русской расскладке
		// клавиши не подхватываюится
	} else if key == glfw.KEY_UNKNOWN {
		switch scancode {
		case 113: // LEFT
			view_roty += 5.0
		case 114: // RIGHT
			view_roty -= 5.0
		case 111: // UP
			view_rotx += 5.0
		case 116: // DOWN
			view_rotx -= 5.0
		}
	} else {
		switch key {
		case glfw.KEY_LEFT:
			view_roty += 5.0
		case glfw.KEY_RIGHT:
			view_roty -= 5.0
		case glfw.KEY_UP:
			view_rotx += 5.0
		case glfw.KEY_DOWN:
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

	if animateFlag {
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
		tRate0 = t
		frames = 0
	}
}
