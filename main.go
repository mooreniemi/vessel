package main

import (
	"fmt"
	vp "github.com/mooreniemi/vessel/parsing"
	vc "github.com/mooreniemi/vessel/viewcomponents"
	gc "github.com/rthornton128/goncurses"
	"log"
	"strconv"
)

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

	menu, menuwin := vc.ChamberMenu(stdscr, *chambers[0], chambers)

	y, x := stdscr.MaxYX()
	vc.VisorView(y)

	stdscr.MovePrint(y-1, 1, "'q' to quit")
	stdscr.MovePrint(0, 1, "Loading timer...")
	stdscr.Refresh()

	go vc.UpdateTimer(stdscr)

	currentChamber := chambers[0]

	for {
		gc.Update()
		vc.VesselMap(x, *currentChamber, chambers)
		ch := menuwin.GetChar()

		switch ch {
		case 'q':
			vc.CleanUp(menu)
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

			vc.CleanUp(menu)
			menu, menuwin = vc.ChamberMenu(stdscr, *currentChamber, chambers)

			vc.VisorView(y)
		}
	}
}
