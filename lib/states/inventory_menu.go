package states

import (
	"image"

	"github.com/ebitenui/ebitenui"
	e_image "github.com/ebitenui/ebitenui/image"
	"github.com/ebitenui/ebitenui/widget"
	"github.com/golang/freetype/truetype"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	gc "github.com/kijimaD/sokotwo/lib/components"
	"github.com/kijimaD/sokotwo/lib/effects"
	"github.com/kijimaD/sokotwo/lib/engine/loader"
	er "github.com/kijimaD/sokotwo/lib/engine/resources"
	"github.com/kijimaD/sokotwo/lib/engine/states"
	w "github.com/kijimaD/sokotwo/lib/engine/world"
	"github.com/kijimaD/sokotwo/lib/eui"
	"github.com/kijimaD/sokotwo/lib/resources"
	"github.com/kijimaD/sokotwo/lib/styles"
	ecs "github.com/x-hgg-x/goecs/v2"
	"golang.org/x/image/font"
)

type InventoryMenuState struct {
	selection     int
	inventoryMenu []ecs.Entity
	menuLen       int
	ui            *ebitenui.UI
	clickedItem   ecs.Entity
}

// State interface ================

func (st *InventoryMenuState) OnPause(world w.World) {}

func (st *InventoryMenuState) OnResume(world w.World) {}

func (st *InventoryMenuState) OnStart(world w.World) {
	prefabs := world.Resources.Prefabs.(*resources.Prefabs)
	st.inventoryMenu = append(st.inventoryMenu, loader.AddEntities(world, prefabs.Menu.InventoryMenu)...)
	st.ui = st.initUI(world)
}

func (st *InventoryMenuState) OnStop(world w.World) {
	world.Manager.DeleteEntities(st.inventoryMenu...)
}

func (st *InventoryMenuState) Update(world w.World) states.Transition {
	effects.RunEffectQueue(world)

	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		return states.Transition{Type: states.TransSwitch, NewStates: []states.State{&CampMenuState{}}}
	}

	if inpututil.IsKeyJustPressed(ebiten.KeySlash) {
		return states.Transition{Type: states.TransPush, NewStates: []states.State{&DebugMenuState{}}}
	}

	st.ui.Update()

	return updateMenu(st, world)
}

func (st *InventoryMenuState) Draw(world w.World, screen *ebiten.Image) {
	st.ui.Draw(screen)
}

// Menu Interface ================

func (st *InventoryMenuState) getSelection() int {
	return st.selection
}

func (st *InventoryMenuState) setSelection(selection int) {
	st.selection = selection
}

func (st *InventoryMenuState) confirmSelection(world w.World) states.Transition {
	return states.Transition{Type: states.TransNone}
}

func (st *InventoryMenuState) getMenuIDs() []string {
	return []string{""}
}

func (st *InventoryMenuState) getCursorMenuIDs() []string {
	return []string{""}
}

// ================

type entryStruct struct {
	entity      ecs.Entity
	name        string
	description string
}

func (st *InventoryMenuState) initUI(world w.World) *ebitenui.UI {
	ui := ebitenui.UI{}
	buttonImage, _ := loadButtonImage()
	face, _ := loadFont((*world.Resources.Fonts)["kappa"])

	gameComponents := world.Components.Game.(*gc.Components)
	var members []ecs.Entity
	world.Manager.Join(
		gameComponents.Member,
		gameComponents.InParty,
	).Visit(ecs.Visit(func(entity ecs.Entity) {
		members = append(members, entity)
	}))

	rootContainer := widget.NewContainer(
		widget.ContainerOpts.Layout(widget.NewGridLayout(
			widget.GridLayoutOpts.Columns(3),
			widget.GridLayoutOpts.Spacing(2, 0),
			widget.GridLayoutOpts.Stretch([]bool{true, false, true}, []bool{false, true, false}),
		)),
	)

	title := widget.NewText(
		widget.TextOpts.Text("インベントリ", face, styles.TextColor),
		widget.TextOpts.WidgetOpts(
			widget.WidgetOpts.LayoutData(widget.RowLayoutData{
				Position: widget.RowLayoutPositionCenter,
			}),
		),
	)
	rootContainer.AddChild(title)
	rootContainer.AddChild(eui.EmptyContainer())
	rootContainer.AddChild(eui.EmptyContainer())

	content := widget.NewContainer(widget.ContainerOpts.Layout(widget.NewRowLayout(
		widget.RowLayoutOpts.Direction(widget.DirectionVertical),
		widget.RowLayoutOpts.Spacing(20),
	)))

	itemDescContainer := widget.NewContainer(
		widget.ContainerOpts.Layout(widget.NewRowLayout(
			widget.RowLayoutOpts.Padding(widget.Insets{
				Top:    20,
				Bottom: 20,
			}),
		)),
		widget.ContainerOpts.WidgetOpts(
			widget.WidgetOpts.MinSize(0, 40),
		),
	)

	itemDesc := widget.NewText(
		widget.TextOpts.Text(" ", face, styles.TextColor),
		widget.TextOpts.WidgetOpts(widget.WidgetOpts.LayoutData(widget.RowLayoutData{
			Stretch: true,
		})),
	)
	itemDescContainer.AddChild(itemDesc)

	var items []ecs.Entity
	world.Manager.Join(
		gameComponents.Item,
		gameComponents.Name,
		gameComponents.Description,
		gameComponents.InBackpack,
		gameComponents.Consumable,
	).Visit(ecs.Visit(func(entity ecs.Entity) {
		items = append(items, entity)
	}))

	for _, entity := range items {
		entity := entity
		name := gameComponents.Name.Get(entity).(*gc.Name)

		windowContainer := widget.NewContainer(
			widget.ContainerOpts.BackgroundImage(e_image.NewNineSliceColor(styles.WindowBodyColor)),
			widget.ContainerOpts.Layout(
				widget.NewGridLayout(
					widget.GridLayoutOpts.Columns(1),
					widget.GridLayoutOpts.Stretch([]bool{true}, []bool{false, false, false}),
					widget.GridLayoutOpts.Padding(widget.Insets{
						Top:    20,
						Bottom: 20,
						Left:   10,
						Right:  10,
					}),
					widget.GridLayoutOpts.Spacing(0, 15),
				),
			),
		)

		titleContainer := widget.NewContainer(
			widget.ContainerOpts.BackgroundImage(e_image.NewNineSliceColor(styles.WindowHeaderColor)),
			widget.ContainerOpts.Layout(widget.NewAnchorLayout()),
		)
		titleContainer.AddChild(widget.NewText(
			widget.TextOpts.Text("アクション", face, styles.TextColor),
			widget.TextOpts.WidgetOpts(widget.WidgetOpts.LayoutData(widget.AnchorLayoutData{
				HorizontalPosition: widget.AnchorLayoutPositionCenter,
				VerticalPosition:   widget.AnchorLayoutPositionCenter,
			})),
		))

		window := widget.NewWindow(
			widget.WindowOpts.Contents(windowContainer),
			widget.WindowOpts.TitleBar(titleContainer, 25),
			widget.WindowOpts.Modal(),
			widget.WindowOpts.CloseMode(widget.CLICK_OUT),
			widget.WindowOpts.Draggable(),
			widget.WindowOpts.Resizeable(),
			widget.WindowOpts.MinSize(200, 200),
			widget.WindowOpts.MaxSize(300, 400),
		)

		button := widget.NewButton(
			widget.ButtonOpts.Image(buttonImage),
			widget.ButtonOpts.Text(name.Name, face, &widget.ButtonTextColor{
				Idle: styles.TextColor,
			}),
			widget.ButtonOpts.TextPadding(widget.Insets{
				Left:   30,
				Right:  30,
				Top:    5,
				Bottom: 5,
			}),
			widget.ButtonOpts.ClickedHandler(func(args *widget.ButtonClickedEventArgs) {
				x, y := ebiten.CursorPosition()
				r := image.Rect(0, 0, x, y)
				r = r.Add(image.Point{x + 20, y + 20})
				window.SetLocation(r)
				ui.AddWindow(window)

				st.clickedItem = entity
			}),
		)

		button.GetWidget().CursorEnterEvent.AddHandler(func(args interface{}) {
			if st.clickedItem != entity {
				st.clickedItem = entity
			}

			var description string
			world.Manager.Join(gameComponents.Description).Visit(ecs.Visit(func(entity ecs.Entity) {
				if entity == st.clickedItem && entity.HasComponent(gameComponents.Description) {
					c := gameComponents.Description.Get(entity).(*gc.Description)
					description = c.Description
				}
			}))
			itemDesc.Label = description
		})
		content.AddChild(button)

		windowContainer.AddChild(widget.NewButton(
			widget.ButtonOpts.Image(buttonImage),
			widget.ButtonOpts.Text("使う", face, &widget.ButtonTextColor{
				Idle: styles.TextColor,
			}),
			widget.ButtonOpts.ClickedHandler(func(args *widget.ButtonClickedEventArgs) {
				effects.ItemTrigger(nil, entity, effects.Single{members[0]}, world)
				content.RemoveChild(button)
				window.Close()
			}),
		))

		windowContainer.AddChild(widget.NewButton(
			widget.ButtonOpts.Image(buttonImage),
			widget.ButtonOpts.Text("捨てる", face, &widget.ButtonTextColor{
				Idle: styles.TextColor,
			}),
			widget.ButtonOpts.ClickedHandler(func(args *widget.ButtonClickedEventArgs) {
				world.Manager.DeleteEntity(entity)
				content.RemoveChild(button)
				window.Close()
			}),
		))

		windowContainer.AddChild(widget.NewButton(
			widget.ButtonOpts.Image(buttonImage),
			widget.ButtonOpts.Text("閉じる", face, &widget.ButtonTextColor{
				Idle: styles.TextColor,
			}),
			widget.ButtonOpts.ClickedHandler(func(args *widget.ButtonClickedEventArgs) {
				window.Close()
			}),
		))
	}

	sc, v := eui.NewScrollContainer(content)
	rootContainer.AddChild(sc)
	rootContainer.AddChild(v)

	itemSpec := widget.NewText(
		widget.TextOpts.Text("性能", face, styles.TextColor),
		widget.TextOpts.WidgetOpts(
			widget.WidgetOpts.LayoutData(widget.RowLayoutData{
				Position: widget.RowLayoutPositionCenter,
			}),
		),
	)
	rootContainer.AddChild(itemSpec)

	rootContainer.AddChild(itemDescContainer)

	ui = ebitenui.UI{
		Container: rootContainer,
	}

	ui.Container = rootContainer

	return &ui
}

func loadButtonImage() (*widget.ButtonImage, error) {
	idle := e_image.NewNineSliceColor(styles.ButtonIdleColor)
	hover := e_image.NewNineSliceColor(styles.ButtonHoverColor)
	pressed := e_image.NewNineSliceColor(styles.ButtonPressedColor)

	return &widget.ButtonImage{
		Idle:    idle,
		Hover:   hover,
		Pressed: pressed,
	}, nil
}

func loadFont(font er.Font) (font.Face, error) {
	return truetype.NewFace(font.Font, &truetype.Options{
		Size: 24,
		DPI:  72,
	}), nil
}
