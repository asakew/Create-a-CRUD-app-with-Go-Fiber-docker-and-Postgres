package handlers

import "github.com/gofiber/fiber/v2"

func CreateFact(c *fiber.Ctx) error {
	fact := new(models.Fact)
	// Parse request body
	if err := c.BodyParser(fact); err != nil {
		return NewFactView(c)
	}

	// Create fact in database
	result := database.DB.Db.Create(&fact)
	if result.Error != nil {
		return NewFactView(c)
	}

	return ListFacts(c)
}
