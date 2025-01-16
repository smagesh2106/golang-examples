package tempconv

func CToF(c Celsius) Farenheit { return Farenheit(c*9/5 + 32) }

func FToC(f Farenheit) Celsius { return Celsius((f - 32) * 5 / 9) }
