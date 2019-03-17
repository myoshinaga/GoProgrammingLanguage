package conv

import "fmt"

type Celsius float64
type Fahrenheit float64
type Feet float64
type Meters float64
type Pondus float64
type Kilogram float64

func CtoF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }
func FtoC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

func FtoM(f Feet) Meters { return Meters(f / 3.28084) }
func MtoF(m Meters) Feet { return Feet(m * 3.28084) }

func PtoK(p Pondus) Kilogram { return Kilogram(p / 2.20462) }
func KtoP(k Kilogram) Pondus { return Pondus(k * 2.20462) }

func (c Celsius) String() string    { return fmt.Sprintf("%g℃", c) }
func (f Fahrenheit) String() string { return fmt.Sprintf("%g℉", f) }

func (f Feet) String() string   { return fmt.Sprintf("%gft", f) }
func (m Meters) String() string { return fmt.Sprintf("%gm", m) }

func (p Pondus) String() string   { return fmt.Sprintf("%glb", p) }
func (k Kilogram) String() string { return fmt.Sprintf("%gkg", k) }
