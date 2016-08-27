package main

import (
	"fmt"
	hd "github.com/MakeNowJust/heredoc"
	gc "github.com/rthornton128/goncurses"
	"log"
	"strconv"
)

var ascii = hd.Doc(`

       /\
      {.-}
     ;_.-'\
    {    _.}_
     \.-' /  ',
      \  |    /
       \ |  ,/
        \|_/

`)

type chamber struct {
	desc     string
	doors    []*chamber
	doorDesc string
	id       int
	items    []int
}

func cleanUp(menu *gc.Menu) {
	menu.UnPost()
	for _, item := range menu.Items() {
		item.Free()
	}
	menu.Free()
}

func makeMenu(stdscr *gc.Window, chamber chamber) (*gc.Menu, *gc.Window) {
	items := make([]*gc.MenuItem, len(chamber.doors))
	for i, ch := range chamber.doors {
		items[i], _ = gc.NewItem(strconv.Itoa(ch.id), fmt.Sprintf("%s door", ch.doorDesc))
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

	one := chamber{id: 1, desc: "Too dark to see what, but something is dripping.", doorDesc: "moldy"}
	two := chamber{id: 2, desc: "It's so dry in here you feel your tongue sticky in your mouth.", doorDesc: "creaking"}
	three := chamber{id: 3, desc: "There's a rocking chair and a crib.", doorDesc: "warm"}
	four := chamber{id: 4, desc: "You knock over open tins of food, spilling dust across the floor.", doorDesc: "'Mess Hall'"}
	entryway := chamber{id: 0, doors: []*chamber{&one, &two}, doorDesc: "cold", desc: "The chamber walls bubble in slow motion."}
	two.doors = []*chamber{&entryway, &three, &four}
	one.doors = []*chamber{&entryway}
	three.doors = []*chamber{&two}
	four.doors = []*chamber{&two}

	chambers := []chamber{entryway, one, two, three, four}

	menu, menuwin := makeMenu(stdscr, entryway)

	y, _ := stdscr.MaxYX()

	viewwin, err := gc.NewWindow(y-10, 40, 1, 1)
	viewwin.MovePrint(0, 0, "viewwin")
	viewwin.MovePrint(1, 1, ascii)
	viewwin.Refresh()

	stdscr.MovePrint(y-1, 1, "'q' to quit")
	stdscr.Refresh()

	currentChamber := entryway
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
			stdscr.Move(y-3, 1)
			stdscr.ClearToEOL()
			stdscr.MovePrint(y-3, 1, fmt.Sprintf("[CHAMBER %s] %s", menu.Current(nil).Name(), currentChamber.desc))
			stdscr.Refresh()

			cleanUp(menu)
			menu, menuwin = makeMenu(stdscr, currentChamber)
		}
	}
}
