package models

// type Model interface {
// 	Valid() bool
// }

func (recipe *RecipeData) Valid() bool {
	if recipe.Name == "" || recipe.Cover == "" || recipe.Category == "" {
		return false
	}

	numberOfIngredients := len(recipe.Ingredients)
	numberOfInstrucions := len(recipe.Instructions)
	if numberOfIngredients == 0 || numberOfInstrucions == 0 {
		return false
	}

	for i := 0; i < numberOfIngredients; i++ {
		if !recipe.Ingredients[i].Valid() {
			return false
		}
	}
	for i := 0; i < numberOfInstrucions; i++ {
		if !recipe.Instructions[i].Valid() {
			return false
		}
	}

	return true
}

// func isModelsValid(models []Model) bool {
// 	for i := 0; i < len(models); i++ {
// 		if !models[i].Valid() {
// 			return false
// 		}
// 	}
// 	return false
// }

func (ingredient *IngredientData) Valid() bool {
	if ingredient.Name == "" || ingredient.Unit == "" {
		return false
	}
	if ingredient.Quantity == 0 {
		return false
	}
	return true
}

func (instruction *InstructionData) Valid() bool {
	if instruction.Order == 0 {
		return false
	}
	if instruction.Description == "" {
		return false
	}
	return true
}
