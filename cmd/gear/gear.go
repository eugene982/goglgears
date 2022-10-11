package gear

import (
	"math"

	"goglgears/pkg/gl"
)

func Gear(innerRadius float32, outerRadius float32,
	width float32, teeth int, toothDepth float32) {
	var (
		i          int
		r0, r1, r2 float32
		angle, da  float32
		u, v, len  float32
	)

	r0 = innerRadius
	r1 = outerRadius - toothDepth/2.0
	r2 = outerRadius + toothDepth/2.0

	da = 2.0 * math.Pi / float32(teeth) / 4.0

	gl.ShadeModel(gl.GL_FLAT)

	gl.Normal3f(0.0, 0.0, 1.0)

	/* draw front face */
	gl.Begin(gl.GL_QUAD_STRIP)
	for i = 0; i <= teeth; i++ {
		angle = float32(i) * 2.0 * math.Pi / float32(teeth)
		gl.Vertex3f(r0*cos(angle), r0*sin(angle), width*0.5)
		gl.Vertex3f(r1*cos(angle), r1*sin(angle), width*0.5)
		if i < teeth {
			gl.Vertex3f(r0*cos(angle), r0*sin(angle), width*0.5)
			gl.Vertex3f(r1*cos(angle+3*da), r1*sin(angle+3*da), width*0.5)
		}
	}
	gl.End()

	/* draw front sides of teeth */
	gl.Begin(gl.GL_QUADS)
	da = 2.0 * math.Pi / float32(teeth) / 4.0
	for i = 0; i < teeth; i++ {
		angle = float32(i) * 2.0 * math.Pi / float32(teeth)

		gl.Vertex3f(r1*cos(angle), r1*sin(angle), width*0.5)
		gl.Vertex3f(r2*cos(angle+da), r2*sin(angle+da), width*0.5)
		gl.Vertex3f(r2*cos(angle+2*da), r2*sin(angle+2*da),
			width*0.5)
		gl.Vertex3f(r1*cos(angle+3*da), r1*sin(angle+3*da),
			width*0.5)
	}
	gl.End()

	gl.Normal3f(0.0, 0.0, -1.0)

	/* draw back face */
	gl.Begin(gl.GL_QUAD_STRIP)
	for i = 0; i <= teeth; i++ {
		angle = float32(i) * 2.0 * math.Pi / float32(teeth)
		gl.Vertex3f(r1*cos(angle), r1*sin(angle), -width*0.5)
		gl.Vertex3f(r0*cos(angle), r0*sin(angle), -width*0.5)
		if i < teeth {
			gl.Vertex3f(r1*cos(angle+3*da), r1*sin(angle+3*da), -width*0.5)
			gl.Vertex3f(r0*cos(angle), r0*sin(angle), -width*0.5)
		}
	}
	gl.End()

	/* draw back sides of teeth */
	gl.Begin(gl.GL_QUADS)
	da = 2.0 * math.Pi / float32(teeth) / 4.0
	for i = 0; i < teeth; i++ {
		angle = float32(i) * 2.0 * math.Pi / float32(teeth)

		gl.Vertex3f(r1*cos(angle+3*da), r1*sin(angle+3*da), -width*0.5)
		gl.Vertex3f(r2*cos(angle+2*da), r2*sin(angle+2*da), -width*0.5)
		gl.Vertex3f(r2*cos(angle+da), r2*sin(angle+da), -width*0.5)
		gl.Vertex3f(r1*cos(angle), r1*sin(angle), -width*0.5)
	}
	gl.End()

	/* draw outward faces of teeth */
	gl.Begin(gl.GL_QUAD_STRIP)
	for i = 0; i < teeth; i++ {
		angle = float32(i) * 2.0 * math.Pi / float32(teeth)

		gl.Vertex3f(r1*cos(angle), r1*sin(angle), width*0.5)
		gl.Vertex3f(r1*cos(angle), r1*sin(angle), -width*0.5)
		u = r2*cos(angle+da) - r1*cos(angle)
		v = r2*sin(angle+da) - r1*sin(angle)
		len = sqrt(u*u + v*v)
		u /= len
		v /= len
		gl.Normal3f(v, -u, 0.0)
		gl.Vertex3f(r2*cos(angle+da), r2*sin(angle+da), width*0.5)
		gl.Vertex3f(r2*cos(angle+da), r2*sin(angle+da), -width*0.5)
		gl.Normal3f(cos(angle), sin(angle), 0.0)
		gl.Vertex3f(r2*cos(angle+2*da), r2*sin(angle+2*da), width*0.5)
		gl.Vertex3f(r2*cos(angle+2*da), r2*sin(angle+2*da), -width*0.5)
		u = r1*cos(angle+3*da) - r2*cos(angle+2*da)
		v = r1*sin(angle+3*da) - r2*sin(angle+2*da)
		gl.Normal3f(v, -u, 0.0)
		gl.Vertex3f(r1*cos(angle+3*da), r1*sin(angle+3*da), width*0.5)
		gl.Vertex3f(r1*cos(angle+3*da), r1*sin(angle+3*da), -width*0.5)
		gl.Normal3f(cos(angle), sin(angle), 0.0)
	}

	gl.Vertex3f(r1*cos(0), r1*sin(0), width*0.5)
	gl.Vertex3f(r1*cos(0), r1*sin(0), -width*0.5)

	gl.End()

	gl.ShadeModel(gl.GL_SMOOTH)

	/* draw inside radius cylinder */
	gl.Begin(gl.GL_QUAD_STRIP)
	for i = 0; i <= teeth; i++ {
		angle = float32(i) * 2.0 * math.Pi / float32(teeth)
		gl.Normal3f(-cos(angle), -sin(angle), 0.0)
		gl.Vertex3f(r0*cos(angle), r0*sin(angle), -width*0.5)
		gl.Vertex3f(r0*cos(angle), r0*sin(angle), width*0.5)
	}
	gl.End()
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

var (
	stereo bool
)

func GearInit(stereo_ bool) (delete func()) {

	stereo = stereo_

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
	Gear(1.0, 4.0, 1.0, 20, 0.7)
	gl.EndList()

	gear2 = gl.GenLists(1)
	gl.NewList(gear2, gl.GL_COMPILE)
	gl.Materialfv(gl.GL_FRONT, gl.GL_AMBIENT_AND_DIFFUSE, &green[0])
	Gear(0.5, 2.0, 2.0, 10, 0.7)
	gl.EndList()

	gear3 = gl.GenLists(1)
	gl.NewList(gear3, gl.GL_COMPILE)
	gl.Materialfv(gl.GL_FRONT, gl.GL_AMBIENT_AND_DIFFUSE, &blue[0])
	Gear(1.3, 2.0, 0.5, 10, 0.7)
	gl.EndList()

	gl.Enable(gl.GL_NORMALIZE)

	return func() {
		gl.DeleteLists(gear1, 1)
		gl.DeleteLists(gear2, 1)
		gl.DeleteLists(gear3, 1)
	}
}

var (
	eyesep    float64 = 5.0  /* Eye separation. */
	fix_point float64 = 40.0 /* Fixation point distance.  */
)

var ( /* Stereo frustum params.  */
	left  float64
	right float64
	asp   float64
)

func DrawGears(angle, view_rotx, view_roty, view_rotz float32) {
	if stereo {
		/* First left eye.  */
		gl.DrawBuffer(gl.GL_BACK_LEFT)

		gl.MatrixMode(gl.GL_PROJECTION)
		gl.LoadIdentity()
		gl.Frustum(left, right, -asp, asp, 5.0, 60.0)

		gl.MatrixMode(gl.GL_MODELVIEW)

		gl.PushMatrix()
		gl.Translated(+0.5*eyesep, 0.0, 0.0)

		draw(angle, view_rotx, view_roty, view_rotz)

		gl.PopMatrix()

		/* Then right eye.  */
		gl.DrawBuffer(gl.GL_BACK_RIGHT)

		gl.MatrixMode(gl.GL_PROJECTION)
		gl.LoadIdentity()
		gl.Frustum(-right, -left, -asp, asp, 5.0, 60.0)

		gl.MatrixMode(gl.GL_MODELVIEW)

		gl.PushMatrix()
		gl.Translated(-0.5*eyesep, 0.0, 0.0)

		draw(angle, view_rotx, view_roty, view_rotz)

		gl.PopMatrix()

	} else {
		draw(angle, view_rotx, view_roty, view_rotz)
	}
}

func Reshape(width, height int) {
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

func draw(angle, view_rotx, view_roty, view_rotz float32) {
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

func sin(a float32) float32 {
	return float32(math.Sin(float64(a)))
}

func cos(a float32) float32 {
	return float32(math.Cos(float64(a)))
}

func sqrt(a float32) float32 {
	return float32(math.Sqrt(float64(a)))
}
