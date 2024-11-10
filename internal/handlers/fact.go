package handlers

import (
	"github.com/gofiber/fiber/v2"
	"goFiber/internal/database"
	"goFiber/internal/models"
)

func CreateFact(c *fiber.Ctx) error {
	fact := new(models.Fact)
	// Parse request body
	if err := c.BodyParser(fact); err != nil {
		return NewFactView(c)
	}

	// Create fact in database
	result := database.FactDB.Create(&fact)
	if result.Error != nil {
		return NewFactView(c)
	}

	return ListFacts(c)
}

func ListFacts(c *fiber.Ctx) error {
	var facts []models.Fact
	database.FactDB.Find(&facts)
	return c.JSON(facts)
}

func ShowFact(c *fiber.Ctx) error {
	fact := models.Fact{}
	id := c.Params("id")

	database.FactDB.Where("id = ?", id).First(&fact)

	return c.Render("show", fiber.Map{
		"Title": "Single Fact",
		"Fact":  fact,
	})
}

func NewFactView(c *fiber.Ctx) error {
	return c.Render("new-fact", fiber.Map{})
}

func ShowFactView(c *fiber.Ctx) error {
	id := c.Params("id")
	var fact models.Fact
	database.FactDB.Find(&fact, id)
	return c.Render("show-fact", fiber.Map{
		"Fact": fact,
	})
}

func EditFactView(c *fiber.Ctx) error {
	id := c.Params("id")
	var fact models.Fact
	database.FactDB.Find(&fact, id)
	return c.Render("edit-fact", fiber.Map{
		"Fact": fact,
	})
}

func UpdateFact(c *fiber.Ctx) error {
	id := c.Params("id")
	var fact models.Fact
	database.FactDB.Where("id = ?", id).First(&fact)
	if err := c.BodyParser(&fact); err != nil {
		return EditFactView(c)
	}
	database.FactDB.Save(&fact)
	return ShowFactView(c)
}

func DeleteFact(c *fiber.Ctx) error {
	id := c.Params("id")
	database.FactDB.Delete(&models.Fact{}, id)
	return ListFacts(c)
}
