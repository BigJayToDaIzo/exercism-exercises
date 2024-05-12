package lasagna

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, prepPerLayer int) int {
	prepTime := 0
	if prepPerLayer == 0 {
		prepPerLayer = 2
	}
	for i := 0; i < len(layers); i++ {
		prepTime += prepPerLayer
	}
	return prepTime
}

// TODO: define the 'Quantities()' function
func Quantities(layers []string) (noodles int, sauce float64) {
	noodles = 0
	sauce = 0.0
	for _, v := range layers {
		switch v {
		case "noodles":
			noodles += 50
		case "sauce":
			sauce += 0.2
		}
	}
	// naked returns are fragile, especially in longer functions
	return noodles, sauce
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(fi, mi []string) {
	mi[len(mi)-1] = fi[len(fi)-1]
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(qBase []float64, qNeeded int) []float64 {
	if len(qBase) == 0 {
		return []float64{}
	}
	var qScaled []float64
	for i := range qBase {
		qScaled = append(qScaled, qBase[i]/2*float64(qNeeded))
	}
	return qScaled
}
