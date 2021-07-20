package particles

import (
	"fmt"
	"math/rand"

	"github.com/skycoin/cx-game/camera"
	"github.com/skycoin/cx-game/cxmath"
	"github.com/skycoin/cx-game/input"
)

var MyEmitter *Emitter
var enableDraw = false

type Emitter struct {
	Position  cxmath.Vec2
	Particles []*Particle
}

func NewEmitter() *Emitter {
	emitter := Emitter{}
	return &emitter
}

func (emitter *Emitter) Update(dt float32) {
	emitter.Position = cxmath.Vec2{
		float32(input.GetScreenX() / 32),
		float32(input.GetScreenY() / 32),
	}
	fmt.Println(len(emitter.Particles))

	for i := 0; i < 240; i++ {
		size := (rand.Float32() + 0.5) / 3
		newParticle := NewParticle(
			emitter.Position,
			cxmath.Vec2{rand.Float32() - 0.5, rand.Float32() - 0.5},
			cxmath.Vec2{size, size},
			getRandomParticleSprite(),
			float32(rand.Intn(3)+2),
			getDefaultDrawHandler(),
			getDefaultPhysicsHandler(),
		)
		emitter.Particles = append(emitter.Particles, newParticle)
	}

	newParticleList := make([]*Particle, 0)
	for _, par := range emitter.Particles {
		par.Update(dt)
		if par.TimeToLive > 0 {
			newParticleList = append(newParticleList, par)
		}
	}
	emitter.Particles = newParticleList
	// fmt.Println(len(emitter.Particles))
}

func (emitter *Emitter) Draw(cam *camera.Camera) {
	if enableDraw == false {
		return
	}
	for _, par := range emitter.Particles {
		par.Draw(cam)
	}
}
