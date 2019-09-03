package handler

import (
	"net/http"
	"recipe-api/models"
	"strconv"

	"github.com/labstack/echo"
)

type Env struct {
	DB models.Datastore
}

func (env *Env) GetAllRecipes(c echo.Context) error {
	recipes, err := env.DB.GetAllRecipes()
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, recipes)
}

func (env *Env) CreateRecipe(c echo.Context) error {
	// init a new recipe
	var recipe models.Recipe

	// map incoming JSON body to the new recipe
	c.Bind(&recipe)
	id, err := env.DB.CreateRecipe(recipe)

	// if creation is successful return a response
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, id)
}

func (env *Env) DeleteRecipe(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	// delete a recipe using our model
	deletedId, err := env.DB.DeleteRecipe(id)

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, deletedId)
}
