package models

import (
	"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
)

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

func CreateRecipeWithPhotosAndPost(recipe Recipe, user User, post Post, photoLinks []string, db *sqlx.DB) (postId int, err error) {
	tx, err := db.Begin()
	if err != nil {
		rollbackErr := tx.Rollback()
		if rollbackErr != nil {
			return -1, rollbackErr
		}
		return -1, err
	}
	err = tx.QueryRow("INSERT INTO recipes (title, instructions) VALUES ($1, $2) RETURNING id", recipe.Title, recipe.Instructions).Scan(&recipe.ID)
	if err != nil {
		rollbackErr := tx.Rollback()
		if rollbackErr != nil {
			return -1, rollbackErr
		}
		return -1, err
	}

	postId, err = createPostForRecipeTrx(recipe, user, post, photoLinks, tx)

	return
}

func CreatePostForRecipe(recipe Recipe, user User, post Post, photoLinks []string, db *sqlx.DB) (postId int, err error) {
	tx, err := db.Begin()
	postId, err = createPostForRecipeTrx(recipe, user, post, photoLinks, tx)
	return
}

func createPostForRecipeTrx(recipe Recipe, user User, post Post, photoLinks []string, tx *sql.Tx) (postId int, err error) {
	err = tx.QueryRow("INSERT INTO posts (user_id, recipe_id, description) VALUES ((SELECT id from users where username=$1 LIMIT 1), $2, $3) RETURNING id", user.UserName, recipe.ID, post.Description).Scan(&postId)
	if err != nil {
		rollbackErr := tx.Rollback()
		if rollbackErr != nil {
			return -1, rollbackErr
		}
		return -1, err
	}
	if err != nil {
		fmt.Printf("Failed to get postId from result\n")
		return -1, err
	}
	photoInsertQuery := "INSERT INTO recipe_photos (recipe_id, post_id, photo_link) VALUES"
	for i, link := range photoLinks {
		if i != 0 {
			photoInsertQuery += ","
		}
		photoInsertQuery += fmt.Sprintf(" (%d, %d, '%s')", recipe.ID, postId, link)
	}
	photoInsertQuery += ";"
	_, err = tx.Exec(photoInsertQuery)
	if err != nil {
		rollbackErr := tx.Rollback()
		if rollbackErr != nil {
			return -1, rollbackErr
		}
		return -1, err
	}
	err = tx.Commit()
	return postId, err
}
