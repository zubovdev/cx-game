package input

import (
	"log"

	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/skycoin/cx-game/utility"
)

//continuos keys, holding
func GetButton(button string) bool {
	key, ok := ButtonsToKeys[button]
	if !ok {
		log.Printf("KEY IS NOT MAPPED!")
		return false
	}
	pressed, ok := KeysPressed[key]
	if !ok {
		// log.Printf("ERROR!")
		return false
	}
	return pressed
}

//action keys, if pressed once
func GetButtonDown(button string) bool {
	key, ok := ButtonsToKeys[button]
	if !ok {
		log.Printf("KEY [%s] IS NOT MAPPED!", button)
		return false
	}
	pressed, ok := KeysPressedDown[key]
	if !ok {
		return false
	}
	KeysPressedDown[key] = false
	return pressed
}

func GetButtonUp(button string) bool {
	key, ok := ButtonsToKeys[button]
	if !ok {
		log.Printf("KEY IS NOT MAPPED")
		return false
	}
	pressed, ok := KeysPressedUp[key]
	if !ok {
		return false
	}
	KeysPressedUp[key] = false
	return pressed
}
func GetKey(key glfw.Key) bool {
	return KeysPressed[key]
}
func GetKeyDown(key glfw.Key) bool {
	return KeysPressedDown[key]
}
func GetKeyUp(key glfw.Key) bool {
	return KeysPressedUp[key]
}

func GetLastKey() glfw.Key {
	return lastKeyPressed
}

func GetAxis(axis Axis) float32 {
	if axis == HORIZONTAL {
		return utility.BoolToFloat(GetButton("right")) - utility.BoolToFloat(GetButton("left"))
	} else { // VERTICAL
		return utility.BoolToFloat(GetButton("up")) - utility.BoolToFloat(GetButton("down"))
	}

}

func GetMouseX() float32 {
	return float32(mouseCoords.X)
}

func GetMouseY() float32 {
	return float32(mouseCoords.Y)
}

func GetScreenX() float32 {
	screenX := float32(((GetMouseX()-float32(widthOffset))/scale - float32(window_.Width)/2)) / Zoom // adjust mouse position with zoom
	// fmt.Println("       | ", widthOffset, "   ", heightOffset, "  ", scale)
	return screenX
}

func GetScreenY() float32 {
	screenY := float32(((GetMouseY()-float32(heightOffset))/scale-float32(window_.Height)/2)*-1) / Zoom
	return screenY

}
