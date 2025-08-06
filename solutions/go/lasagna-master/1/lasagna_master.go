package lasagna

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, preparationTime int) int {
	if preparationTime <= 0 {
		preparationTime = 2
	}
	return len(layers) * preparationTime
}

// TODO: define the 'Quantities()' function
func Quantities(layers []string) (int, float64) {
	var noodles int
	var sauce float64

	for i := 0; i < len(layers); i++ {
		if layers[i] == "noodles" {
			noodles += 1
		}
		if layers[i] == "sauce" {
			sauce += 1
		}
	}

	return noodles * 50, sauce * 0.2
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendsIngredients []string, myIngredients *[]string) {
	(*myIngredients)[len(*myIngredients)-1] = friendsIngredients[len(friendsIngredients)-1]
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(portion []float64, portionsToCook int) []float64 {
	var amountsNeeded []float64
	portionsToIncrease := float64(portionsToCook) - float64(2)

	for i := 0; i < len(portion); i++ {
		amountsNeeded = append(amountsNeeded, portion[i]*portionsToIncrease)
	}

	return amountsNeeded
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
