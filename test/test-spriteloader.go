package main

import (
	"fmt"
	"log"
	"math/rand"
	"runtime"
	"time"

	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/skycoin/cx-game/render"
	"github.com/skycoin/cx-game/spriteloader"
)

func init() {
	// This is needed to arrange that main() runs on main thread.
	// See documentation for functions that are only allowed to be called
	// from the main thread.
	runtime.LockOSThread()
}

func keyCallBack(w *glfw.Window, k glfw.Key, s int, a glfw.Action, mk glfw.ModifierKey) {
	if a == glfw.Press && k == glfw.KeyEscape {
		w.SetShouldClose(true)
	}
	if a == glfw.Press {
		if k == glfw.KeyW {
			ypos += 1
		} else if k == glfw.KeyS {
			ypos -= 1
		}
		if k == glfw.KeyA {
			xpos -= 1
		} else if k == glfw.KeyD {
			xpos += 1
		}

		fmt.Println(xpos, "   ", ypos)
	}
}

var xpos, ypos float32

func main() {
	log.Print("running test")
	log.Print("You should see an orange square rock.")
	win := render.NewWindow(800, 600, true)
	window := win.Window
	window.SetKeyCallback(keyCallBack)
	defer glfw.Terminate()
	spriteloader.InitSpriteloader(&win)
	spriteSheetId := spriteloader.
		LoadSpriteSheet("./assets/starfield/stars/planets.png")
	spriteloader.
		LoadSprite(spriteSheetId, "star", 2, 1)
	spriteId := spriteloader.
		GetSpriteIdByName("star")
	rand.Seed(time.Now().UnixNano())
	xpos, ypos = 0, 0
	for !window.ShouldClose() {
		gl.ClearColor(1, 1, 1, 1)
		gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

		spriteloader.DrawSpriteQuad(xpos, ypos, 1, 1, spriteId)
		glfw.PollEvents()
		window.SwapBuffers()
	}
}
