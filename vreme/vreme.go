// Paket vreme nudi ustvarjanje vremenskih postaj z podatkovnim tipom Postaja
// Ponuja možnosti upravljanja s postajami in ustvarjanje poročila
package vreme

import (
	"fmt"
	"slices"
)

type Postaja struct {
	ime     string
	kraj    string
	meritve []float64
}

var postaje = make(map[string]Postaja)

// Doda postajo v slovar z oznako, imenom in krajem
func DodajPostajo(oznaka string, ime, kraj string) {
	postaje[oznaka] = Postaja{
		ime:  ime,
		kraj: kraj,
	}
}

// Doda meritev med min in max temperaturo
func DodajMeritev(oznaka string, temperatura, min, max float64) {
	if temperatura < min || temperatura > max {
		fmt.Printf("Temperatura ni v pravem obmocju")
		fmt.Println()
		return
	}
	e, ok := postaje[oznaka]
	if !ok {
		fmt.Printf("Postaja z to oznako ne obstaja.")
		fmt.Println()
		return
	}
	e.meritve = append(e.meritve, temperatura)
	postaje[oznaka] = e
}

func statistika(oznaka string, stMeritev int) (najnizja, najvisja, povprecna float64, ok bool) {
	e, ok := postaje[oznaka]
	if !ok {
		fmt.Println("Postaja z to oznako ne obstaja.")
		return 0, 0, 0, false
	}
	if len(e.meritve) < stMeritev {
		fmt.Println("Premalo meritev.")
		return 0, 0, 0, false
	}
	najnizja = slices.Min(e.meritve)
	najvisja = slices.Max(e.meritve)
	var sum float64 = 0
	for _, val := range e.meritve {
		sum += val
	}
	povprecna = sum / float64(len(e.meritve))
	ok = true
	return najnizja, najvisja, povprecna, ok
}

// Izpis meritev za vsako postajo
func IzpisMeritev() {
	fmt.Println("MERITVE:")
	for oznaka, postaja := range postaje {
		fmt.Printf("%s - %s (%s): %v", oznaka, postaja.ime, postaja.kraj, postaja.meritve)
		fmt.Println()
	}
}

// Izpis statističnega poročila za vsako postajo
func IzpisPorocila(stMeritev int) {
	fmt.Println("POROČILO:")
	for oznaka, postaja := range postaje {
		min, max, avg, ok := statistika(oznaka, stMeritev)
		if !ok {
			continue
		}
		var obcutek string
		switch {
		case avg >= float64(25):
			obcutek = "Vroče"
		case avg > float64(15) && avg < float64(25):
			obcutek = "Toplo"
		case avg >= float64(5) && avg <= float64(15):
			obcutek = "Hladno"
		case avg < float64(5):
			obcutek = "Mrzlo"
		}
		fmt.Printf("%s: min %g °C, max %g °C, povprecje %.1f °C -> %s", postaja.ime, min, max, avg, obcutek)
		fmt.Println()
	}
}
