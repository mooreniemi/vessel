package main

import (
	"fmt"
	ascii "github.com/mooreniemi/vessel/ascii"
	vp "github.com/mooreniemi/vessel/parsing"
	v "github.com/mooreniemi/vessel/vessel"
	gc "github.com/rthornton128/goncurses"
	"log"
	"strconv"
	"strings"
	"time"
)

const longForm = "Jan 2, 2006 at 3:04pm (MST)"

func updateTimer(w *gc.Window) {
	t, _ := time.Parse(longForm, "Aug 29, 2016 at 7:54pm (EST)")
	for _ = range time.Tick(time.Second) {
		until := t.Sub(time.Now())
		w.MovePrint(0, 1, fmt.Sprintf("%v until vessel decay.", until.String()))
		w.Refresh()
	}
}

func cleanUp(menu *gc.Menu) {
	menu.UnPost()
	for _, item := range menu.Items() {
		item.Free()
	}
	menu.Free()
}

func visorView(y int) *gc.Window {
	viewwin, err := gc.NewWindow(y-10, 40, 1, 1)
	if err != nil {
		log.Fatal(err)
	}
	viewwin.MovePrint(1, 1, ascii.RandomArt())
	viewwin.Refresh()

	return viewwin
}

func vesselMap(x int, current v.Chamber, chambers []*v.Chamber) *gc.Window {
	vmap := vp.ParseVesselMap()

	// NewWindow(lines, columns, y, x)
	mapwin, err := gc.NewWindow(len(vmap)+2, 11, 0, x-11)
	if err != nil {
		log.Fatal(err)
	}
	_, x = mapwin.MaxYX()
	mapwin.Box(0, 0)

	vy, err := strconv.Atoi(current.Y)
	vx, err := strconv.Atoi(current.X)

	windowPadding := 1
	for i, row := range vmap {
		rowFormatted := strings.Trim(fmt.Sprint(row), "[]")
		rowFormatted = strings.Replace(rowFormatted, "0", " ", -1)
		// TODO use gc.ACS_CKBOARD instead of #?
		rowFormatted = strings.Replace(rowFormatted, "1", "#", -1)
		if vy == i {
			mapwin.Move(0, 0)
			mapwin.ClearToEOL()
			mapwin.MovePrint(windowPadding+i, 1, rowFormatted)

			mapwin.ColorOn(1)
			marker := "x"
			if vx >= 1 {
				mapwin.MovePrint(windowPadding+i, 1+vx*2, marker)
			} else {
				mapwin.MovePrint(windowPadding+i, 1+vx, marker)
			}
			mapwin.ColorOff(1)
		} else {
			mapwin.MovePrint(windowPadding+i, 1, rowFormatted)
		}
	}

	mapwin.Refresh()

	return mapwin
}

func chamberMenu(stdscr *gc.Window, chamber v.Chamber, chambers []*v.Chamber) (*gc.Menu, *gc.Window) {
	items := make([]*gc.MenuItem, len(chamber.Doors))
	for i, doorID := range chamber.Doors {
		// reversed so that places you just came from
		// appear below places you could navigate to
		items[len(items)-i-1], _ = gc.NewItem(strconv.Itoa(chambers[doorID].ID),
			fmt.Sprintf("%s door", chambers[doorID].DoorDesc))
	}

	menu, _ := gc.NewMenu(items)

	_, x := stdscr.MaxYX()
	// NewWindow(lines, columns, y, x)
	menuwin, err := gc.NewWindow(8, 40, 0, x-50)
	if err != nil {
		log.Fatal(err)
	}

	menuwin.Keypad(true)
	menu.SetWindow(menuwin)

	dwin := menuwin.Derived(0, 38, 3, 1)
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

	menu, menuwin := chamberMenu(stdscr, *chambers[0], chambers)

	y, x := stdscr.MaxYX()
	visorView(y)

	stdscr.MovePrint(y-1, 1, "'q' to quit")
	stdscr.MovePrint(0, 1, "Loading timer...")
	stdscr.Refresh()

	go updateTimer(stdscr)

	currentChamber := chambers[0]

	for {
		gc.Update()
		vesselMap(x, *currentChamber, chambers)
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
			stdscr.MovePrint(y-4, 1,
				fmt.Sprintf("[CHAMBER %s] %s", menu.Current(nil).Name(), currentChamber.Desc))
			stdscr.Refresh()

			cleanUp(menu)
			menu, menuwin = chamberMenu(stdscr, *currentChamber, chambers)

			visorView(y)
		}
	}
}
