package gear

import (
	"math"

	"goglgears/pkg/gl"
)

func NewGear(innerRadius float32,
	outerRadius float32,
	width float32,
	teeth int,
	toothDepth float32,
	material [4]float32) uint {

	list := gl.GenLists(1)
	gl.NewList(list, gl.GL_COMPILE)
	gl.Materialfv(gl.GL_FRONT, gl.GL_AMBIENT_AND_DIFFUSE, &material[0])
	gear(innerRadius, outerRadius, width, teeth, toothDepth)
	gl.EndList()

	return list
}

func gear(innerRadius float32, outerRadius float32,
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

func sin(a float32) float32 {
	return float32(math.Sin(float64(a)))
}

func cos(a float32) float32 {
	return float32(math.Cos(float64(a)))
}

func sqrt(a float32) float32 {
	return float32(math.Sqrt(float64(a)))
}
