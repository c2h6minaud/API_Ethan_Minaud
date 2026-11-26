package handlers

import (
	"github.com/gofiber/fiber/v3"
	"API/model"
	"API/store"
	"strconv"
	"strings"
)

type Handler struct {
	st *store.Store
}

func NewHandler(store *store.Store) *Handler {
	return &Handler{st: store}
}

func respondError(c fiber.Ctx, status int, message string) error {
	return c.Status(status).JSON(fiber.Map{
		"error": message,
	})
}

func (h *Handler) ListArtists(c fiber.Ctx) error {
	return c.JSON(h.st.Artists())
}

func (h *Handler) GetArtist(c fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return respondError(c, fiber.StatusBadRequest, "Invalid artist ID")
	}
	artist, err := h.st.GetByID(id)
	if err != nil {
		return respondError(c, fiber.StatusNotFound, "Artist not found")
	}
	return c.JSON(artist)
}

func (h *Handler) CreateArtist(c fiber.Ctx) error {
	var artist model.Artist
	err := c.Bind().Body(&artist)
	if err != nil {
		return respondError(c, fiber.StatusBadRequest, "Invalid request body")
	}
	artist.Name = strings.TrimSpace(artist.Name)
	artist.Genre = strings.TrimSpace(artist.Genre)
	if artist.Name == "" || artist.Genre == "" {
		return respondError(c, fiber.StatusBadRequest, "Name and genre are required")
	}
	createdArtist := h.st.Create(artist)
	return c.Status(fiber.StatusCreated).JSON(createdArtist)
}

func (h *Handler) UpdateArtist(c fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.Atoi(idStr)	
    if err != nil {
		return respondError(c, fiber.StatusBadRequest, "Invalid artist ID")
	}
	var artist model.Artist
	err = c.Bind().Body(&artist)
	if err != nil {
		return respondError(c, fiber.StatusBadRequest, "Invalid request body")
	}
	artist.Name = strings.TrimSpace(artist.Name)
	artist.Genre = strings.TrimSpace(artist.Genre)
	if artist.Name == "" || artist.Genre == "" {
		return respondError(c, fiber.StatusBadRequest, "Name and genre are required")
	}
	updatedArtist, err := h.st.Update(id, artist)
	if err != nil {
		return respondError(c, fiber.StatusNotFound, "Artist not found")
	}
	return c.JSON(updatedArtist)
}

func (h *Handler) DeleteArtist(c fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.Atoi(idStr)	
    if err != nil {
		return respondError(c, fiber.StatusBadRequest, "Invalid artist ID")
	}
	err = h.st.Delete(id)
	if err != nil {
		return respondError(c, fiber.StatusNotFound, "Artist not found")
	}
	return c.JSON(fiber.Map{
		"message": "Artist deleted successfully",
	})
}