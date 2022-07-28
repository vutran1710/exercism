package lasagna

// TODO: define the 'PreparationTime()' function
// PreparationTime ...
func PreparationTime(layers []string, time int) int {
	prepareTime := 2
	if time > 0 {
		prepareTime = time
	}

	return prepareTime * len(layers)
}

// TODO: define the 'Quantities()' function
// Quantities ...
func Quantities(layers []string) (int, float64) {
	var noodles int
	var sauce float64

	for _, value := range layers {
		if value == "sauce" {
			sauce += 0.2

		}
		if value == "noodles" {
			noodles += 50

		}
	}

	return noodles, sauce

}

// TODO: define the 'AddSecretIngredient()' function
// AddSecretIngredient ...
func AddSecretIngredient(friendList []string, myList []string) []string {
	secretIngredient := friendList[len(friendList)-1]
	myList[len(myList)-1] = secretIngredient
	return myList
}

// TODO: define the 'ScaleRecipe()' function
// ScaleRecipe ...
func ScaleRecipe(quantities []float64, portions int) []float64 {
	times := float64(portions) / 2
	result := make([]float64, len(quantities), len(quantities))

	for key, value := range quantities {
		result[key] = value * times
	}
	return result

}
