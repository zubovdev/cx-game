package game

import (
	"log"
	"runtime"

	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/skycoin/cx-game/camera"
	"github.com/skycoin/cx-game/components"
	"github.com/skycoin/cx-game/enemies"
	"github.com/skycoin/cx-game/input"
	"github.com/skycoin/cx-game/item"
	"github.com/skycoin/cx-game/models"
	"github.com/skycoin/cx-game/particles"
	"github.com/skycoin/cx-game/render"
	"github.com/skycoin/cx-game/sound"
	"github.com/skycoin/cx-game/spriteloader"
	"github.com/skycoin/cx-game/starfield"
	"github.com/skycoin/cx-game/ui"
	"github.com/skycoin/cx-game/world"
)

func init() {
	// This is needed to arrange that main() runs on main thread.
	// See documentation for functions that are only allowed to be called from the main thread.
	runtime.LockOSThread()
}

const (
	WINDOW_WIDTH  = 800
	WINDOW_HEIGHT = 600
)

var (
	Cam    *camera.Camera
	win    render.Window
	window *glfw.Window
	player *models.Player
	fps    *models.Fps

	CurrentPlanet      *world.Planet
	DrawCollisionBoxes = false
	FPS                int

	catIsScratching bool

	tilePaletteSelector ui.TilePaletteSelector

	inventoryId item.InventoryID

	//unused
	isTileSelectorVisible = false
	worldItem             *item.WorldItem
)

func Init() {
	win = render.NewWindow(WINDOW_WIDTH, WINDOW_HEIGHT, true)
	// defer glfw.Terminate()

	window = win.Window

	window.SetMouseButtonCallback(mouseButtonCallback)
	window.SetScrollCallback(scrollCallback)
	window.SetSizeCallback(windowSizeCallback)

	input.Init(&win)
	sound.Init()
	spriteloader.InitSpriteloader(&win)
	spriteloader.DEBUG = false
	world.RegisterTileTypes()
	item.InitWorldItem()
	ui.Init()
	enemies.InitBasicEnemies()
	particles.InitParticles()
	item.RegisterItemTypes()

	models.Init()
	player = models.NewPlayer()

	fps = models.NewFps(false)
	Cam = camera.NewCamera(&win)
	//CurrentPlanet = world.NewDevPlanet()
	CurrentPlanet = world.GeneratePlanet()
	Cam.PlanetWidth = float32(CurrentPlanet.Width)

	components.Init(CurrentPlanet, Cam, player)

	starfield.InitStarField(&win, player, Cam)

	inventoryId = item.NewDevInventory()

	worldTiles := CurrentPlanet.GetAllTilesUnique()
	log.Printf("Found [%v] unique tiles in the world", len(worldTiles))
	tilePaletteSelector = ui.
		NewDevTilePaleteSelector()

	//init cam and cat positions
	spawnX := int(20)
	Cam.SetCameraPosition(float32(spawnX), 5)
	// Cam.SetCameraZoomPosition(0)
	player.Pos.X = float32(spawnX)
	player.Pos.Y = float32(CurrentPlanet.GetHeight(spawnX) + 10)
	enemies.SpawnBasicEnemy(player.Pos.X+6, player.Pos.Y)
	enemies.SpawnLeapingEnemy(player.Pos.X-6, player.Pos.Y)

	sound.LoadSound("player_jump", "jump.wav")

}
