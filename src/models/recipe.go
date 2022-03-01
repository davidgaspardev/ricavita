package models

/// Recipe data model
type RecipeData struct {
	Id           string            `json:"id:omitempty" bson:"id,omitempty"`
	Name         string            `json:"name,omitempty" bson:"name,omitempty"`
	Cover        string            `json:"cover,omitempty" bson:"cover,omitempty"` // Image url
	Category     string            `json:"category,omitempty" bson:"category,omitempty"`
	Ingredients  []IngredientData  `json:"ingredients,omitempty" bson:"ingredients,omitempty"`
	Instructions []InstructionData `json:"instructions,omitempty" bson:"instructions,omitempty"`
}

/// Ingredient data model
type IngredientData struct {
	Name     string  `json:"name,omitempty" bson:"name,omitempty"`
	Quantity float32 `json:"quantity" bson:"quantity,omitempty"`
	Unit     string  `json:"unit,omitempty" bson:"unit,omitempty"`
}

/// Instruction data model
type InstructionData struct {
	Order       int    `json:"order,omitempty" bson:"order,omitempty"`
	Description string `json:"description,omitempty" bson:"order,omitempty"`
}
