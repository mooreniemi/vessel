package main

import (
	"fmt"
	gc "github.com/rthornton128/goncurses"
	"log"
	"strconv"
	ascii "vessel/ascii"
	vp "vessel/vesselparser"
)

func cleanUp(menu *gc.Menu) {
	menu.UnPost()
	for _, item := range menu.Items() {
		item.Free()
	}
	menu.Free()
}

func makeMenu(stdscr *gc.Window, chamber vp.Chamber, chambers []*vp.Chamber) (*gc.Menu, *gc.Window) {
	items := make([]*gc.MenuItem, len(chamber.Doors))
	for i, doorID := range chamber.Doors {
		items[i], _ = gc.NewItem(strconv.Itoa(chambers[doorID].ID),
			fmt.Sprintf("%s door", chambers[doorID].DoorDesc))
	}

	menu, _ := gc.NewMenu(items)

	_, x := stdscr.MaxYX()
	// NewWindow(lines, columns, y, x)
	menuwin, err := gc.NewWindow(10, 40, 0, x-40)
	if err != nil {
		log.Fatal(err)
	}

	menuwin.Keypad(true)
	menu.SetWindow(menuwin)

	dwin := menuwin.Derived(6, 38, 3, 1)
	menu.SubWindow(dwin)
	menu.Mark(" > ")

	// Print centered menu title
	_, x = menuwin.MaxYX()
	title := "Exits"
	menuwin.Box(0, 0)
	menuwin.ColorOn(1)
	menuwin.MovePrint(1, (x/2)-(len(title)/2), title)
	menuwin.ColorOff(1)
	menuwin.MoveAddChar(2, 0, gc.ACS_LTEE)
	menuwin.HLine(2, 1, gc.ACS_HLINE, x-3)
	menuwin.MoveAddChar(2, x-2, gc.ACS_RTEE)

	stdscr.Refresh()

	menu.Post()
	menuwin.Refresh()

	return menu, menuwin
}

func main() {
	vessel, err := vp.ParseVesselYaml()
	if err != nil {
		log.Fatal(err)
	}

	chambers := vessel.Chambers

	stdscr, err := gc.Init()
	if err != nil {
		log.Fatal(err)
	}
	defer gc.End()

	stdscr.Border(gc.ACS_VLINE, gc.ACS_VLINE, gc.ACS_HLINE, gc.ACS_HLINE,
		gc.ACS_ULCORNER, gc.ACS_URCORNER, gc.ACS_LLCORNER, gc.ACS_LRCORNER)

	gc.StartColor()
	gc.Raw(true)
	gc.Echo(false)
	gc.Cursor(0)
	stdscr.Keypad(true)
	gc.InitPair(1, gc.C_RED, gc.C_BLACK)

	menu, menuwin := makeMenu(stdscr, *chambers[0], chambers)

	y, _ := stdscr.MaxYX()

	viewwin, err := gc.NewWindow(y-10, 40, 1, 1)
	viewwin.MovePrint(0, 0, "viewwin")
	viewwin.MovePrint(1, 1, ascii.RandomArt())
	viewwin.Refresh()

	stdscr.MovePrint(y-1, 1, "'q' to quit")
	stdscr.Refresh()

	currentChamber := chambers[0]
	for {
		gc.Update()
		ch := menuwin.GetChar()

		switch ch {
		case 'q':
			cleanUp(menu)
			return
		case gc.KEY_DOWN:
			menu.Driver(gc.REQ_DOWN)
		case gc.KEY_UP:
			menu.Driver(gc.REQ_UP)
		case gc.KEY_RETURN, gc.KEY_ENTER, gc.Key('\r'):
			chamberID, _ := strconv.Atoi(menu.Current(nil).Name())
			currentChamber = chambers[chamberID]
			stdscr.Move(y-4, 1)
			stdscr.ClearToEOL()
			stdscr.MovePrint(y-4, 1, fmt.Sprintf("[CHAMBER %s] %s", menu.Current(nil).Name(), currentChamber.Desc))
			stdscr.Refresh()

			cleanUp(menu)
			menu, menuwin = makeMenu(stdscr, *currentChamber, chambers)
		}
	}
}
