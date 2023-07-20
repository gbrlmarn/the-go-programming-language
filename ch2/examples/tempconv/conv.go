package tempconv

// CToF converts a Celsius temperature to Fahrenheit
func CToF(c Celisius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

// FToC converts a Fahrenheit temperature to Celisius
func FToC(f Fahrenheit) Celisius { return Celisius((f - 32) * 5 / 9) }
