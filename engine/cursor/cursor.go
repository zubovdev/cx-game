package cursor

import (
	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/skycoin/cx-game/engine/spriteloader"
	"image"
	"log"
	"strings"
)

var cursors = make(map[string]*glfw.Cursor)

func Get(name string) *glfw.Cursor {
	cursor, ok := cursors[name]
	if !ok {
		log.Fatalf("Unknown cursor %s", name)
	}
	return cursor
}

func Load(filePath, spritePath string) {
	// Reading the sprite sheet config.
	sheetConfig, err := spriteloader.ReadSpriteSheetConfig(filePath)
	if err != nil {
		log.Fatalln("failed to load sprite sheet config for the cursors:", err)
	}

	// Loading image.
	_, spriteImg, _ := spriteloader.LoadPng(spritePath)
	for spriteName, cfg := range sheetConfig.SpriteConfigs {
		// Splitting the spriteName by the "/".
		// Example: actual spriteName is "cursors/default" while cursor name must be "default".
		elements := strings.Split(spriteName, "/")
		cursorName := elements[len(elements)-1]

		// Creating an image from the loaded spite image.
		img := spriteImg.SubImage(image.Rect(cfg.Top, cfg.Left, cfg.Top+cfg.Height, cfg.Left+cfg.Width))

		// Registering the cursor.
		cursors[cursorName] = glfw.CreateCursor(img, 0, 0)
	}
}
