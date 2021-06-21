package particles

import (
	"log"

	"github.com/go-gl/mathgl/mgl32"
	"github.com/go-gl/gl/v4.1-core/gl"

	"github.com/skycoin/cx-game/render"
	"github.com/skycoin/cx-game/utility"
	"github.com/skycoin/cx-game/spriteloader"
	"github.com/skycoin/cx-game/cxmath"
	"github.com/skycoin/cx-game/physics"
)

type Bullet struct {
	transform mgl32.Mat4
	velocity mgl32.Vec2
}

var (
	bulletShader *utility.Shader
	bullets []Bullet
)

func InitBullets() {
	bulletShader = utility.NewShader(
		"./assets/shader/simple.vert", "./assets/shader/color.frag" )
	bulletShader.Use()
	bulletShader.SetVec4F("colour",1,0,0,1)
	bulletShader.StopUsing()
	bullets = []Bullet {}
}

func CreateBullet( origin mgl32.Vec2, velocity mgl32.Vec2 ) {
	bullets = append(bullets, Bullet {
		transform: mgl32.Translate3D(origin.X(), origin.Y(), 0),
		velocity: velocity,
	})
}

func (bullet Bullet) WorldTransform() mgl32.Mat4 {
	return bullet.transform.Mul4(cxmath.Scale(1.0/4))
}

func (bullet Bullet) draw(ctx render.Context) {
	world := bullet.WorldTransform()
	bulletShader.SetMat4("world", &world)

	gl.DrawArrays(gl.TRIANGLES, 0, 6)
}

func configureGlForBullet() {
	gl.Disable(gl.DEPTH_TEST)
	gl.Enable(gl.TEXTURE_2D)
	gl.ActiveTexture(gl.TEXTURE0)
	/*
	gl.BindTexture(gl.TEXTURE_2D, bulletTex.gpuTex)
	// blurry is better than jagged for a bullet
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.LINEAR)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.LINEAR)
	*/

	gl.BindVertexArray(spriteloader.QuadVao);
}

func DrawBullets(ctx render.Context) {
	bulletShader.Use()
	bulletShader.SetMat4("projection", &ctx.Projection)
	configureGlForBullet()
	// TODO add texture
	// bulletShader.SetUint("tex", bulletShader.gpuTex)

	for _,bullet := range bullets {
		bullet.draw(ctx)
	}
	bulletShader.StopUsing()
}

func TickBullets(dt float32) {
	newBullets := []Bullet{}
	for _,bullet := range bullets {
		//bullet.ttl -= dt
		bullet.transform = bullet.transform.Mul4(
			mgl32.Translate3D(bullet.velocity.X()*dt, bullet.velocity.Y()*dt, 0) )

		collision,collided := physics.CheckCollision(bullet.WorldTransform())
		_ = collision
		if collided {
			log.Print("bullet hit something")
		} else {
			newBullets = append(newBullets,bullet)
		}
	}
	bullets = newBullets
}
