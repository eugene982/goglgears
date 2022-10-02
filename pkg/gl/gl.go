package gl

const (
	GL_FLAT = iota

	GL_QUAD_STRIP
	GL_QUADS
)

func Begin(_ int) {
	panic("glBegin")
}

func End() {
	panic("glEnd")
}

func ShadeModel(m int) {
	panic("glShadeModel")
}

func Normal3f(_, _, _ float32) {
	panic("glNormal3f")
}

func Vertex3f(_, _, _ float32) {
	panic("glVertex3f")
}
