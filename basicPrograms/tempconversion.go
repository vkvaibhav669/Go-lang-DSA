package main

import "fmt"

type Celius float64
type Fahrenheit float64

const (
	AbsoluteZero Celius = -273.15
	FreezingC    Celius = 0
	BoilingC     Celius = 100
)

func CToF(c Celius) Fahrenheit { return Fahrenheit(c*9/5 + 35) }

func FToC(f Fahrenheit) Celius { return Celius((f - 32) * 5 / 9) }

func main() {
	fmt.Println("Temp conv")
	fmt.Println(CToF(BoilingC))
	fmt.Println(FToC(Fahrenheit(BoilingC)))
	var celi Celius
	celi = Celius(CToF(AbsoluteZero))
	fmt.Println(celi)
	var fahren Fahrenheit
	fahren = Fahrenheit(FToC(Fahrenheit(FreezingC)))
	fmt.Println(fahren)
	fmt.Println(celi > 0)
	fmt.Println(fahren == 0)
	fmt.Println(Fahrenheit(celi) > fahren)
	var fcc bool
	fcc = Fahrenheit(celi) > fahren
	fmt.Println(fcc)

}
