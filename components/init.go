package components

import (
	"github.com/skycoin/cx-game/camera"
	"github.com/skycoin/cx-game/components/agents/agent_ai"
	"github.com/skycoin/cx-game/components/agents/agent_draw"
	"github.com/skycoin/cx-game/components/agents/agent_health"
	"github.com/skycoin/cx-game/components/particles"
	"github.com/skycoin/cx-game/components/particles/particle_draw"
	"github.com/skycoin/cx-game/components/particles/particle_physics"
	"github.com/skycoin/cx-game/models"
	"github.com/skycoin/cx-game/world"
)

var (
	currentWorldState *world.WorldState
	currentPlanet     *world.Planet
	currentCamera     *camera.Camera
	currentPlayer     *models.Player
)

func Init(planet *world.Planet, cam *camera.Camera, player *models.Player) {
	currentWorldState = planet.WorldState
	currentPlanet = planet
	currentCamera = cam
	currentPlayer = player

	agent_health.Init()
	agent_draw.Init()
	agent_ai.Init()

	particle_physics.Init()
	particle_draw.Init()

	particles.Init()

	particles.MyEmitter = particles.NewEmitter()
}

func ChangeCamera(newCamera *camera.Camera) {
	currentCamera = newCamera
}
func ChangePlanet(newPlanet *world.Planet) {
	currentPlanet = newPlanet
	currentWorldState = newPlanet.WorldState
}

func ChangePlayer(newPlayer *models.Player) {
	currentPlayer = newPlayer
}
