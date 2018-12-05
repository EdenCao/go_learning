package tempconv

import "fmt"

// Celsius Custom Format
type Celsius float64
// Fahrenheit Custom Format
type Fahrenheit float64

const (
  // AbsoluteZero special tempo
  AbsoluteZero Celsius = -273.15
  // FreezingC special tempo
  FreezingC    Celsius = 0
  // BoilingC special tempo
	BoilingC     Celsius = 100
)

func (c Celsius) String() string    { return fmt.Sprintf("%g°C", c) }
func (f Fahrenheit) String() string { return fmt.Sprintf("%g°F", f) }
