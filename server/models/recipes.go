package models

type Recipe struct {
	ID           int    `json:"id" db:"id"`
	Title        string `json:"title" db:"title"`
	Instructions string `json:"instructions" db:"instructions"`
}

type RecipePhoto struct {
	ID        int    `json:"id" db:"id"`
	RecipeID  int    `json:"recipeId" db:"recipe_id"`
	PostID    int    `json:"postId" db:"post_id"`
	PhotoLink string `json:"photoLink" db:"photo_link"`
}
