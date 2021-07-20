package particles

import (
	"math/rand"

	"github.com/skycoin/cx-game/camera"
	"github.com/skycoin/cx-game/components/types"
	"github.com/skycoin/cx-game/cxmath"
	"github.com/skycoin/cx-game/spriteloader"
)

// a particle that bounces but is solid

// a particle that does not bounce but is alpha blended and not solid

// a particle that floats and is alpha blended (drifts in air, no gravity)

// And two draw modes

// 1> solid particle

// 2> alpha blended particle (has transparency, glowing, etc)

// And three physics modes

// 1> bounces, gravity

// 2> Does not bounce, stops on impact, gravity

// 3> "drifts" at fixed velocity, no gravity

// when you fire gun do "shell casing" particle coming out of player
// when you hit block, do "debris" particle, etc
// test a spark or type of glowing/alpha blended particle
// create a struct called "Emitter" which is just something that creates particles, we can throw in game

type Particle struct {
	Position         cxmath.Vec2
	Velocity         cxmath.Vec2
	Size             cxmath.Vec2
	SpriteId         spriteloader.SpriteID
	TimeToLive       float32
	Duration         float32
	DrawHandlerID    types.ParticleDrawHandlerId
	PhysicsHandlerID types.ParticlePhysicsHandlerID
}

func NewParticle(
	position,
	velocity,
	size cxmath.Vec2,
	spriteId spriteloader.SpriteID,
	duration float32,
	drawHandlerId types.ParticleDrawHandlerId,
	physicsHandlerId types.ParticlePhysicsHandlerID,

) *Particle {
	particle := Particle{
		Position:         position,
		Velocity:         velocity,
		Size:             size,
		SpriteId:         spriteId,
		Duration:         duration,
		DrawHandlerID:    drawHandlerId,
		PhysicsHandlerID: physicsHandlerId,
	}
	particle.TimeToLive = duration
	return &particle
}

func (particle *Particle) Update(dt float32) {
	particle.TimeToLive -= dt

	particle.Position = particle.Position.Add(particle.Velocity.Mult(dt))
}

func (particle *Particle) Draw(cam *camera.Camera) {
	spriteloader.DrawSpriteQuad(
		particle.Position.X,
		particle.Position.Y,
		particle.Size.X,
		particle.Size.Y,
		particle.SpriteId,
	)
}

func Init() {
	spriteloader.LoadSingleSprite("./assets/particles/circle.png", "circle")
	spriteloader.LoadSingleSprite("./assets/particles/star.png", "star")
	spriteloader.LoadSingleSprite("./assets/particles/diamond.png", "diamond")
}

func BinByDrawHandlerId(particleList []*Particle) map[types.ParticleDrawHandlerId][]*Particle {
	bins := make(map[types.ParticleDrawHandlerId][]*Particle)
	for _, par := range particleList {
		bins[par.DrawHandlerID] = append(bins[par.DrawHandlerID], par)
	}
	return bins
}
func getRandomParticleSprite() spriteloader.SpriteID {
	n := rand.Intn(3)
	switch n {
	case 0:
		return spriteloader.GetSpriteIdByName("circle")
	case 1:
		return spriteloader.GetSpriteIdByName("star")
	case 2:
		return spriteloader.GetSpriteIdByName("diamond")
	}
	return 0
}

func getDefaultDrawHandler() types.ParticleDrawHandlerId {
	return 0
}

func getDefaultPhysicsHandler() types.ParticlePhysicsHandlerID {
	return 0
}
