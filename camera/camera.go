package camera

import (
	"sort"

	"github.com/go-gl/mathgl/mgl32"

	"github.com/skycoin/cx-game/cxmath"
	"github.com/skycoin/cx-game/input"
	"github.com/skycoin/cx-game/render"
	"github.com/skycoin/cx-game/utility"
)

var (
	//variables for smooth zooming over time
	zooming      bool    = false // flag for to know if zooming is occuring
	zoomDuration float32 = 0.6   // in seconds
	zoomProgress float32

	// variables for interpolation
	zoomTarget  float32 = 1 // zoom value to end on
	zoomCurrent float32     // zoom value to start from
	// current zoom progress (from 0 to 1)

	currentZoomIndex int = 1

	zoomLevels = []float32{0.75, 1, 1.75}
	// firstTick    bool    = true
	focus_area focusArea
)

type Camera struct {
	X           float32
	Y           float32
	Vel         mgl32.Vec2
	Zoom        float32
	movSpeed    float32
	window      *render.Window
	Frustum     Frustum
	focus_area  focusArea
	freeCam     bool
	PlanetWidth float32
}

type focusArea struct {
	center mgl32.Vec2
	left   float32
	right  float32
	top    float32
	bottom float32
}

//Initiates Camera Instances given the window
func NewCamera(window *render.Window) *Camera {
	sort.SliceStable()
	size := mgl32.Vec2{3, 5}
	xPos := float32(0)
	yPos := float32(0)
	cam := Camera{
		//take X,Y pos as a center to frustrum
		X:        xPos,
		Y:        yPos,
		Vel:      mgl32.Vec2{0, 0},
		Zoom:     1,
		movSpeed: 5,
		window:   window,
		focus_area: focusArea{
			center: mgl32.Vec2{xPos, yPos},
			left:   xPos - size.X()/2,
			right:  xPos + size.X()/2,
			top:    yPos + size.Y()/2,
			bottom: yPos - size.Y()/2,
		},
		freeCam: false,
	}
	return &cam
}

//Updates Camera Positions
func (camera *Camera) MoveCam(dTime float32) {
	camera.X += input.GetAxis(input.HORIZONTAL) * dTime * camera.movSpeed
	camera.Y += input.GetAxis(input.VERTICAL) * dTime * camera.movSpeed
	camera.UpdateFrustum()
}

func (camera *Camera) GetView() mgl32.Mat4 {
	return mgl32.Translate3D(-camera.X, -camera.Y, -camera.Zoom)
}

func (camera *Camera) GetProjectionMatrix() mgl32.Mat4 {
	left := -float32(camera.window.Width) / 2 / 32 / camera.Zoom
	right := float32(camera.window.Width) / 2 / 32 / camera.Zoom
	bottom := -float32(camera.window.Height) / 2 / 32 / camera.Zoom
	top := float32(camera.window.Height) / 2 / 32 / camera.Zoom
	projection := mgl32.Ortho(left, right, bottom, top, -1, 1000)

	return projection
}

func (camera *Camera) SetCameraCenter() {
	camera.X = float32(camera.window.Width) / 2
	camera.Y = float32(camera.window.Height) / 2
	camera.UpdateFrustum()
}

//sets camera for current position
func (camera *Camera) SetCameraPosition(x, y float32) {
	camera.updateFocusArea(x, y)
	camera.UpdateFrustum()
}

// update focus area to include (x,y)
func (camera *Camera) updateFocusArea(x, y float32) {
	modular := cxmath.NewModular(camera.PlanetWidth)
	var shiftX, shiftY float32
	if modular.IsLeft(x, camera.focus_area.left) {
		shiftX = modular.Disp(camera.focus_area.left, x)
	} else if modular.IsRight(x, camera.focus_area.right) {
		shiftX = modular.Disp(camera.focus_area.right, x)
	}
	if y < camera.focus_area.bottom {
		shiftY = y - camera.focus_area.bottom
	} else if y > camera.focus_area.top {
		shiftY = y - camera.focus_area.top
	}
	camera.focus_area.left += shiftX
	camera.focus_area.right += shiftX
	camera.focus_area.bottom += shiftY
	camera.focus_area.top += shiftY
	camera.focus_area.center = mgl32.Vec2{
		(camera.focus_area.left + camera.focus_area.right) / 2,
		(camera.focus_area.top + camera.focus_area.bottom) / 2,
	}

	camera.Vel[0] = camera.focus_area.center.X() - camera.X
	camera.Vel[1] = camera.focus_area.center.Y() - camera.Y

	camera.X = camera.focus_area.center.X()
	camera.Y = camera.focus_area.center.Y()
}

//sets camera for target position
func (camera *Camera) SetCameraPositionTarget(x, y float32) {
	camera.SetCameraPosition(x, y)
	camera.UpdateFrustum()
}

//zooms on target
func (camera *Camera) SetCameraZoomTarget(zoomOffset float32) {
	camera.SetCameraZoomPosition(zoomOffset)
}

//zooms on current position
func (camera *Camera) SetCameraZoomPosition(zoomOffset float32) {
	if !zooming {
		zooming = true
		zoomCurrent = zoomLevels[currentZoomIndex]

		currentZoomIndex = utility.ClampI(currentZoomIndex+int(zoomOffset), 0, len(zoomLevels)-1)
		nextZoomIndex := currentZoomIndex

		zoomTarget = zoomLevels[nextZoomIndex]
		if zoomTarget == zoomCurrent {
			zooming = false
		}
	}
}

func (camera *Camera) DrawLines(
	lines []float32, color []float32, ctx render.Context,
) {
	camCtx := ctx.PushView(camera.GetView())
	camera.window.DrawLines(lines, color, camCtx)
}

func (camera Camera) GetTransform() mgl32.Mat4 {
	return mgl32.Translate3D(camera.X, camera.Y, 0)
}

func (camera *Camera) updateProjection() {
	// projection := camera.GetProjectionMatrix()
	// gl.UseProgram(camera.window.Program)
	// gl.UniformMatrix4fv(gl.GetUniformLocation(camera.window.Program, gl.Str("projection\x00")), 1, false, &projection[0])
	projection := camera.GetProjectionMatrix()
	camera.window.SetProjectionMatrix(projection)
}

func (camera *Camera) Tick(dt float32) {
	// TODO optimize this later if necessary
	// always update the projection matrix in case window got resized
	camera.updateProjection()

	// if firstTick {
	// 	camera.updateProjection()
	// 	firstTick = false
	// }
	// if not zooming nothing to do here
	if !zooming {
		return
	}

	zoomProgress += dt / zoomDuration

	camera.Zoom = utility.Lerp(zoomCurrent, zoomTarget, zoomProgress)
	camera.updateProjection()

	if camera.Zoom == zoomTarget {
		zooming = false
		zoomProgress = 0
	}
}

func (camera *Camera) IsFreeCam() bool {
	return camera.freeCam
}

func (camera *Camera) ToggleFreeCam() {

	camera.freeCam = !camera.freeCam

	//reset velocity for now
	camera.Vel = mgl32.Vec2{}
}
