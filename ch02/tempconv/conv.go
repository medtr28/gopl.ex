package tempconv

// CToF : convert from Celsius to Fahrenheit
func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

// FToC : convert from Fahrenheit to Celsius
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

// CToK : convert from Celsius to Kelvin
func CToK(c Celsius) Kelvin { return Kelvin(c - 273.15) }

// KToC : convert from Kelvin to Celsius
func KToC(k Kelvin) Celsius { return Celsius(k + 273.15) }

// FToK : convert from Celsius to Kelvin
func FToK(f Fahrenheit) Kelvin { return Kelvin(CToK(FToC(f))) }

// KToF : convert from Kelvin to Celsius
func KToF(k Kelvin) Fahrenheit { return Fahrenheit(CToF(KToC(k))) }
