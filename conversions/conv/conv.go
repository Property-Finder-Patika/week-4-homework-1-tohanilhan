package conversions

// Create an abstraction to make conversions among different temperature systems such as Celsius, Farhenheit and Kelvin.
// The conversion formula is:
// C = (F - 32) * 5 / 9
// K = C + 273.15
// F = C * 9 / 5 + 32
//

type Conversion interface {
	Calculate(arg float64) float64
	GetName() (string, string)
}

type Celsius struct {
	Name      string
	ConvertTo string
}

func (c Celsius) Calculate(arg float64) float64 {
	// return based on the conversion type
	if c.ConvertTo == "Farhenheit" {
		return (arg * 9 / 5) + 32

	}
	if c.ConvertTo == "Kelvin" {
		return arg + 273.15
	}
	return arg

}

func (c Celsius) GetName() (string, string) {
	return c.Name, c.ConvertTo
}

type Farhenheit struct {
	Name      string
	ConvertTo string
}

func (f Farhenheit) Calculate(arg float64) float64 {
	// return based on the conversion type
	if f.ConvertTo == "Celcius" {
		return (arg - 32) * 5 / 9
	}
	if f.ConvertTo == "Kelvin" {
		return (arg-32)*5/9 + 273.15
	}
	return arg
}

func (f Farhenheit) GetName() (string, string) {
	return f.Name, f.ConvertTo
}

type Kelvin struct {
	Name      string
	ConvertTo string
}

func (k Kelvin) Calculate(arg float64) float64 {
	// return based on the conversion type
	if k.ConvertTo == "Celcius" {
		return arg - 273.15
	}
	if k.ConvertTo == "Farhenheit" {
		return (arg-273.15)*9/5 + 32
	}
	return arg
}

func (k Kelvin) GetName() (string, string) {
	return k.Name, k.ConvertTo
}
