package ascii

import (
	hd "github.com/MakeNowJust/heredoc"
	"math/rand"
	"time"
)

// Shell shell- whelk 11/96 by jgs
var Shell = hd.Doc(`

       /\
      {.-}
     ;_.-'\
    {    _.}_
     \.-' /  ',
      \  |    /
       \ |  ,/
        \|_/

`)

// Octopus by Carl Pilcher
var Octopus = hd.Doc(`

        ___
       / o \
  __   \   /   _
    \__/ | \__/ \
   \___//|\\___/\
    ___/ | \___  
         |     \
        /

`)

// Bird by mrf
var Bird = hd.Doc(`

     .--.  
    /( @ >    ,-. 
   / ' .'--._/  /
   :   ,    , .'
   '. (___.'_/
    ((-((-''   

`)

// RandomArt returns random ascii string
func RandomArt() string {
	art := [3]string{Shell, Octopus, Bird}
	rand.Seed(time.Now().Unix())
	return art[rand.Int()%len(art)]
}

// ArtByID is modulo existing art
func ArtByID(id int) string {
	art := [3]string{Shell, Octopus, Bird}
	return art[id%len(art)]
}
