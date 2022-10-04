package gl

// #cgo pkg-config: opengl
// #include <GL/gl.h>
import "C"
import "unsafe"

const (
	GL_FLAT = C.GL_FLAT

	GL_COMPILE = C.GL_COMPILE

	GL_QUAD_STRIP = C.GL_QUAD_STRIP
	GL_QUADS      = C.GL_QUADS
	GL_SMOOTH     = C.GL_SMOOTH

	GL_CULL_FACE  = C.GL_CULL_FACE
	GL_DEPTH_TEST = C.GL_DEPTH_TEST
	GL_FRONT      = C.GL_FRONT
	GL_NORMALIZE  = C.GL_NORMALIZE

	GL_AMBIENT_AND_DIFFUSE = C.GL_AMBIENT_AND_DIFFUSE

	GL_COLOR_BUFFER_BIT = C.GL_COLOR_BUFFER_BIT
	GL_DEPTH_BUFFER_BIT = C.GL_DEPTH_BUFFER_BIT

	GL_BACK_LEFT  = C.GL_BACK_LEFT
	GL_BACK_RIGHT = C.GL_BACK_RIGHT

	GL_PROJECTION = C.GL_PROJECTION
	GL_MODELVIEW  = C.GL_MODELVIEW
)

const ( // light
	GL_LIGHTING = C.GL_LIGHTING
	GL_LIGHT0   = C.GL_LIGHT0

	GL_POSITION = C.GL_POSITION
)

const (
	GL_RENDERER   = C.GL_RENDERER
	GL_VERSION    = C.GL_VERSION
	GL_VENDOR     = C.GL_VENDOR
	GL_EXTENSIONS = C.GL_EXTENSIONS
)

func Begin(mode C.uint) {
	C.glBegin(mode)
}

func End() {
	C.glEnd()
}

func Enable(mode C.uint) {
	C.glEnable(mode)
}

func Disable(mode C.uint) {
	C.glDisable(mode)
}

func ShadeModel(mode C.uint) {
	C.glShadeModel(mode)
}

func Normal3f(nx, ny, nz float32) {
	C.glNormal3f(C.float(nx), C.float(ny), C.float(nz))
}

func Vertex3f(x, y, z float32) {
	C.glVertex3f(C.float(x), C.float(y), C.float(z))
}

func Lightfv(light uint, pname uint, params *float32) {
	C.glLightfv(C.uint(light), C.uint(pname), (*C.float)(params))
}

func GenLists(lrange int) uint {
	return uint(C.glGenLists(C.int(lrange)))
}

func NewList(list uint, mode uint) {
	C.glNewList(C.uint(list), C.uint(mode))
}

func Materialfv(face uint, pname uint, params *float32) {
	C.glMaterialfv(C.uint(face), C.uint(pname), (*C.float)(params))
}

func EndList() {
	C.glEndList()
}

func CallList(list uint) {
	C.glCallList(C.uint(list))
}

func DeleteLists(list uint, lrange int) {
	C.glDeleteLists(C.uint(list), C.int(lrange))
}

func Clear(mask C.uint) {
	C.glClear(mask)
}

func DrawBuffer(mode C.uint) {
	C.glDrawBuffer(mode)
}

func MatrixMode(mode C.uint) {
	C.glMatrixMode(mode)
}

func PushMatrix() {
	C.glPushMatrix()
}

func PopMatrix() {
	C.glPopMatrix()
}

func Rotatef(angle, x, y, z float32) {
	C.glRotatef(C.float(angle), C.float(x), C.float(y), C.float(z))
}

func Translatef(x, y, z float32) {
	C.glTranslatef(C.float(x), C.float(y), C.float(z))
}

func Translated(x, y, z float64) {
	C.glTranslated(C.double(x), C.double(y), C.double(z))
}

func LoadIdentity() {
	C.glLoadIdentity()
}

func Frustum(left, right, bottom, top, zNear, zFar float64) {
	C.glFrustum(C.double(left), C.double(right),
		C.double(bottom), C.double(top),
		C.double(zNear), C.double(zFar))
}

func GetString(name uint) string {
	var ch = C.glGetString(C.uint(name))
	return C.GoString((*C.char)(unsafe.Pointer(ch)))
}

func Viewport(x, y, width, height int) {
	C.glViewport(C.int(x), C.int(y), C.int(width), C.int(height))
}
