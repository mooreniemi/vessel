package viewcomponents

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

// UpdateTimer countdown until decay
func UpdateTimer(w *gc.Window) {
	t, _ := time.Parse(longForm, "Aug 29, 2016 at 7:54pm (EST)")
	for _ = range time.Tick(time.Second) {
		until := t.Sub(time.Now())
		w.MovePrint(0, 1, fmt.Sprintf("%v until vessel decay.", until.String()))
		w.Refresh()
	}
}

// CleanUp frees menu items, helper
func CleanUp(menu *gc.Menu) {
	menu.UnPost()
	for _, item := range menu.Items() {
		item.Free()
	}
	menu.Free()
}

// VisorView is where we put ascii art
func VisorView(y int) *gc.Window {
	viewwin, err := gc.NewWindow(y-10, 40, 1, 1)
	if err != nil {
		log.Fatal(err)
	}
	viewwin.MovePrint(1, 1, ascii.RandomArt())
	viewwin.Refresh()

	return viewwin
}

// VesselMap parses and displays the vessel
func VesselMap(x int, current v.Chamber, chambers []*v.Chamber) *gc.Window {
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

// ChamberMenu is our navigational system
func ChamberMenu(stdscr *gc.Window, chamber v.Chamber, chambers []*v.Chamber) (*gc.Menu, *gc.Window) {
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
