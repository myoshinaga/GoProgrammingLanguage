// Package tempconv は摂氏（Celsius）と華氏（Fahrenheit）の温度変換を行う
package tempconv

import "fmt"

type Celsius float64
type Fahrenheit float64
type Kelvin float64

const (
	AbsoluteZeroC Celsius = -273.5
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)

func (c Celsius) String() string    { return fmt.Sprintf("%g℃", c) }
func (f Fahrenheit) String() string { return fmt.Sprintf("%g℉", f) }
func (k Kelvin) String() string     { return fmt.Sprintf("%gK", k) }
