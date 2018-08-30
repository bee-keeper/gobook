// ex2.2 takes numerical input and shows conversions
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// weight conversion
type Grams float64
type Pounds float64

func (g Grams) String() string  { return fmt.Sprintf("%.2f grams", g) }
func (p Pounds) String() string { return fmt.Sprintf("%.2f lb", p) }

// PToG converts lb to grams
func PToG(p Pounds) Grams { return Grams(p * 453.59237) }

// GToP converts grams to pounds
func GToP(g Grams) Pounds { return Pounds(g / 453.59237) }

// distance conversions
type Metres float64
type Feet float64

func (f Feet) String() string   { return fmt.Sprintf("%.1f feet", f) }
func (m Metres) String() string { return fmt.Sprintf("%.1f metres", m) }

// FToM converts feet to metres
func FToM(f Feet) Metres { return Metres(f / 3.2808) }

// MToF converts metres to feet
func MToF(m Metres) Feet { return Feet(m * 3.2808) }

// temp conversions
type Celsius float64
type Fahrenheit float64

func (c Celsius) String() string    { return fmt.Sprintf("%.1f°C", c) }
func (f Fahrenheit) String() string { return fmt.Sprintf("%g°F", f) }

// CToF converts a Celsius temperature to Fahrenheit.
func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

// FToC converts a Fahrenheit temperature to Celsius.
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

func showConversions(s string) {
	i, err := strconv.ParseFloat(s, 64)
	if err != nil {
		fmt.Println("Expected numerical argument, got %s", s)
	} else {
		// show conversions
		cel := Celsius(i)
		far := Fahrenheit(i)
		fmt.Printf("%s is equal to %s or %s is equal to %s\n", cel, CToF(cel), far, FToC(far))

		m := Metres(i)
		f := Feet(i)
		fmt.Printf("%s is equal to %s or %s is equal to %s\n", f, FToM(f), m, MToF(m))

		g := Grams(i)
		p := Pounds(i)
		fmt.Printf("%s is equal to %s or %s is equal to %s\n", g, GToP(g), p, PToG(p))

		fmt.Println()
	}
}

func main() {
	if len(os.Args) > 1 {
		for _, arg := range os.Args[1:] {
			showConversions(arg)
		}
	} else {
		scan := bufio.NewScanner(os.Stdin)
		for scan.Scan() {
			showConversions(scan.Text())
		}
	}
}
