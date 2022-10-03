package main

import (
	"log"
	"time"

	"goglgears/pkg/gear"
	"goglgears/pkg/gl"
	"goglgears/pkg/glfw"
)

var (
	gear1 uint
	gear2 uint
	gear3 uint
)

var (
	view_rotx float32 = 20.0
	view_roty float32 = 30.0
	view_rotz float32 = 0.0
	angle     float32 = 0.0
)

var (
	fullscreen = false /* Create a single fullscreen window */
	stereo     = false /* Enable stereo.  */
	samples    = 0     /* Choose visual with at least N samples. */
	animate    = true  /* Animation */

	eyesep    float64 = 5.0  /* Eye separation. */
	fix_point float32 = 40.0 /* Fixation point distance.  */
)

var ( /* Stereo frustum params.  */
	left  float64
	right float64
	asp   float64
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
	gear1 = gear.NewGear(1.0, 4.0, 1.0, 20, 0.7, red)
	gear2 = gear.NewGear(0.5, 2.0, 2.0, 10, 0.7, green)
	gear3 = gear.NewGear(1.3, 2.0, 0.5, 10, 0.7, blue)

	gl.Enable(gl.GL_NORMALIZE)
}

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
	frames = 0
	tRot0 = -1.0 
	tRate0 = -1.0
)

/** Draw single frame, do SwapBuffers, compute FPS */
func draw_frame(win *glfw.Window)
{

   dt, t = time.Now().Unix();

   if (tRot0 < 0.0)
      tRot0 = t;
   dt = t - tRot0;
   tRot0 = t;

   if (animate) {
      /* advance rotation for next frame */
      angle += 70.0 * dt;  /* 70 degrees per second */
      if (angle > 3600.0)
         angle -= 3600.0;
   }

   draw_gears();
   glXSwapBuffers(dpy, win);

   frames++;
   
   if (tRate0 < 0.0)
      tRate0 = t;
   if (t - tRate0 >= 5.0) {
      GLfloat seconds = t - tRate0;
      GLfloat fps = frames / seconds;
      printf("%d frames in %3.1f seconds = %6.3f FPS\n", frames, seconds,
             fps);
      fflush(stdout);
      tRate0 = t;
      frames = 0;
   }
}