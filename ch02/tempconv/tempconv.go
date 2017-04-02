// tempcpmv package : convert Celsius and Fahrenheit
package tempconv

import "fmt"

type Celsius float64
type Fahrenheit float64
type Kelvin float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)

func (c Celsius) String() string    { return fmt.Sprintf("%g deg C", c) }
func (f Fahrenheit) String() string { return fmt.Sprintf("%g deg F", f) }
func (k Kelvin) String() string     { return fmt.Sprintf("%g deg K", k) }
