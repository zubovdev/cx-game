package cursor

import (
	"errors"
	"fmt"
	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/skycoin/cx-game/engine/spriteloader"
	"image"
)

var (
	cursors = make(map[string]*glfw.Cursor)

	ErrUnknownCursor = errors.New("unknown cursor")
)

// Get returns the cursor by name.
// If the cursor does not exist, an error will be returned.
func Get(name string) (*glfw.Cursor, error) {
	// Check whether cursor exist.
	cursor, ok := cursors[name]
	if !ok {
		return nil, ErrUnknownCursor
	}
	return cursor, nil
}

// Load loads the cursors defined in the config.
func Load(configPath, spritePath string) error {
	// Reading the sprite sheet config.
	sheetConfig, err := spriteloader.ReadSpriteSheetConfig(configPath)
	if err != nil {
		return fmt.Errorf("failed to load sprite sheet config for the cursors: %v", err)
	}

	// Loading image.
	_, spriteImg, _ := spriteloader.LoadPng(spritePath)
	for cursorName, cfg := range sheetConfig.SpriteConfigs {
		// Creating an image from the loaded spite image.
		img := spriteImg.SubImage(image.Rect(cfg.Top, cfg.Left, cfg.Top+cfg.Height, cfg.Left+cfg.Width))

		// Registering the cursor.
		cursors[cursorName] = glfw.CreateCursor(img, 0, 0)
	}
	return nil
}
