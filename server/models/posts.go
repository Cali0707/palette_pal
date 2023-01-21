package models

type Post struct {
	ID          int    `json:"id" db:"id"`
	UserID      int    `json:"userID" db:"user_id"`
	RecipeID    int    `json:"recipeID" db:"recipe_id"`
	Description string `json:"description" db:"description"`
}
