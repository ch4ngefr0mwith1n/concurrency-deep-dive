package main

import (
	"fmt"
	"strings"
	"time"
)

// "gorutine" = "go coroutine"
func f(n int) {
	for i := 0; i < 10; i++ {
		fmt.Printf("%sf(%d): %d\n", strings.Repeat(" ", n), n, i)
		time.Sleep(1 * time.Microsecond)
	}
}

func main() {
	// preko "go" keyworda pravimo asinhroni poziv funkcije
	// odnosno, ova funkcija će se izvršavati nezavisno od ostatka koda
	go f(0)
	f(1)
	// kada se završi izvršavanje main funkcije, svaka gorutina će biti prekinuta
	// radi potreba testiranja koristimo "time.Sleep" funkciju
	// zbog toga će nastati čekanje, radi kog će se trenutna gorutina izvršiti do kraja
	time.Sleep(1 * time.Second)
}
