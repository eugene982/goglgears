package main

import (
	"goglgears/pkg/gl"
	"math"
)

type Gear struct {
	innnerRadius float32
	outerRadius  float32
	width        float32
	teeth        int
	toothDepth   float32
}

func NewGear(innnerRadius float32,
	outerRadius float32,
	width float32,
	teeth int,
	toothDepth float32) Gear {

	return Gear{
		innnerRadius: innnerRadius,
		outerRadius:  outerRadius,
		width:        width,
		teeth:        teeth,
		toothDepth:   toothDepth,
	}
}

func (g Gear) Draw() {
	var (
		i          int
		r0, r1, r2 float32
		angle, da  float32
		u, v, len  float32
	)

	r0 = g.innnerRadius
	r1 = g.outerRadius - g.toothDepth/2.
	r2 = g.outerRadius + g.toothDepth/2.

	da = 2.0 * math.Pi / float32(g.teeth) / 4.0

	gl.ShadeModel(gl.GL_FLAT)

	gl.Normal3f(0.0, 0.0, 1.0)

	/* draw front face */
	gl.Begin(gl.GL_QUAD_STRIP)
	for i = 0; i <= g.teeth; i++ {
		angle = float32(i) * 2.0 * math.Pi / float32(g.teeth)
		gl.Vertex3f(r0*cos(angle), r0*sin(angle), g.width*0.5)
		gl.Vertex3f(r1*cos(angle), r1*sin(angle), g.width*0.5)
		if i < g.teeth {
			gl.Vertex3f(r0*cos(angle), r0*sin(angle), g.width*0.5)
			gl.Vertex3f(r1*cos(angle+3*da), r1*sin(angle+3*da), g.width*0.5)
		}
	}
	gl.End()

	/* draw front sides of teeth */
	gl.Begin(gl.GL_QUADS)
	da = 2.0 * math.Pi / float32(g.teeth) / 4.0
	for i = 0; i < g.teeth; i++ {
		angle = float32(i) * 2.0 * math.Pi / float32(g.teeth)

		gl.Vertex3f(r1*cos(angle), r1*sin(angle), g.width*0.5)
		gl.Vertex3f(r2*cos(angle+da), r2*sin(angle+da), g.width*0.5)
		gl.Vertex3f(r2*cos(angle+2*da), r2*sin(angle+2*da),
			g.width*0.5)
		gl.Vertex3f(r1*cos(angle+3*da), r1*sin(angle+3*da),
			g.width*0.5)
	}
	gl.End()

	glNormal3f(0.0, 0.0, -1.0);

   /* draw back face */
   glBegin(GL_QUAD_STRIP);
   for (i = 0; i <= teeth; i++) {
      angle = i * 2.0 * M_PI / teeth;
      glVertex3f(r1 * cos(angle), r1 * sin(angle), -width * 0.5);
      glVertex3f(r0 * cos(angle), r0 * sin(angle), -width * 0.5);
      if (i < teeth) {
	 glVertex3f(r1 * cos(angle + 3 * da), r1 * sin(angle + 3 * da),
		    -width * 0.5);
	 glVertex3f(r0 * cos(angle), r0 * sin(angle), -width * 0.5);
      }
   }
   glEnd();

   /* draw back sides of teeth */
   glBegin(GL_QUADS);
   da = 2.0 * M_PI / teeth / 4.0;
   for (i = 0; i < teeth; i++) {
      angle = i * 2.0 * M_PI / teeth;

      glVertex3f(r1 * cos(angle + 3 * da), r1 * sin(angle + 3 * da),
		 -width * 0.5);
      glVertex3f(r2 * cos(angle + 2 * da), r2 * sin(angle + 2 * da),
		 -width * 0.5);
      glVertex3f(r2 * cos(angle + da), r2 * sin(angle + da), -width * 0.5);
      glVertex3f(r1 * cos(angle), r1 * sin(angle), -width * 0.5);
   }
   glEnd();

   /* draw outward faces of teeth */
   glBegin(GL_QUAD_STRIP);
   for (i = 0; i < teeth; i++) {
      angle = i * 2.0 * M_PI / teeth;

      glVertex3f(r1 * cos(angle), r1 * sin(angle), width * 0.5);
      glVertex3f(r1 * cos(angle), r1 * sin(angle), -width * 0.5);
      u = r2 * cos(angle + da) - r1 * cos(angle);
      v = r2 * sin(angle + da) - r1 * sin(angle);
      len = sqrt(u * u + v * v);
      u /= len;
      v /= len;
      glNormal3f(v, -u, 0.0);
      glVertex3f(r2 * cos(angle + da), r2 * sin(angle + da), width * 0.5);
      glVertex3f(r2 * cos(angle + da), r2 * sin(angle + da), -width * 0.5);
      glNormal3f(cos(angle), sin(angle), 0.0);
      glVertex3f(r2 * cos(angle + 2 * da), r2 * sin(angle + 2 * da),
		 width * 0.5);
      glVertex3f(r2 * cos(angle + 2 * da), r2 * sin(angle + 2 * da),
		 -width * 0.5);
      u = r1 * cos(angle + 3 * da) - r2 * cos(angle + 2 * da);
      v = r1 * sin(angle + 3 * da) - r2 * sin(angle + 2 * da);
      glNormal3f(v, -u, 0.0);
      glVertex3f(r1 * cos(angle + 3 * da), r1 * sin(angle + 3 * da),
		 width * 0.5);
      glVertex3f(r1 * cos(angle + 3 * da), r1 * sin(angle + 3 * da),
		 -width * 0.5);
      glNormal3f(cos(angle), sin(angle), 0.0);
   }

   glVertex3f(r1 * cos(0), r1 * sin(0), width * 0.5);
   glVertex3f(r1 * cos(0), r1 * sin(0), -width * 0.5);

   glEnd();

   glShadeModel(GL_SMOOTH);

   /* draw inside radius cylinder */
   glBegin(GL_QUAD_STRIP);
   for (i = 0; i <= teeth; i++) {
      angle = i * 2.0 * M_PI / teeth;
      glNormal3f(-cos(angle), -sin(angle), 0.0);
      glVertex3f(r0 * cos(angle), r0 * sin(angle), -width * 0.5);
      glVertex3f(r0 * cos(angle), r0 * sin(angle), width * 0.5);
   }
   glEnd();

}

func sin(a float32) float32 {
	return float32(math.Sin(float64(a)))
}

func cos(a float32) float32 {
	return float32(math.Cos(float64(a)))
}
