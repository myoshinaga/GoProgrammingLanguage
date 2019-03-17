package tempconv

// CtoF は摂氏の温度を華氏へ変換
func CtoF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

// FtoC は華氏の温度を摂氏へ変換
func FtoC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

// KtoC は絶対温度を摂氏へ変換
func KtoC(k Kelvin) Celsius { return Celsius((k - 273.15)) }

// CtoK は摂氏を絶対温度へ変換
func CtoK(c Celsius) Kelvin { return Kelvin((c + 273.15)) }
