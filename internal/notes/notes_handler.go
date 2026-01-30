package notes

import (
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Handler struct {
	repo *Repo
}

func NewHandler(repo *Repo) *Handler {
	return &Handler{repo: repo}
}

func (h *Handler) CreateNote(c *gin.Context) {

	var req CreateNoteRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid json",
		})
		return
	}

	now := time.Now().UTC()

	note := Note{
		ID:        primitive.NewObjectID(),
		Title:     req.Title,
		Content:   req.Content,
		Pinned:    req.Pinned,
		CreateAt:  now,
		UpdatedAt: now,
	}

	created, err := h.repo.Create(c.Request.Context(), note)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to create note Here!",
		})
		return
	}

	c.JSON(http.StatusCreated, created)
}

func (h *Handler) ListNotes(c *gin.Context) {
	notes, err := h.repo.List(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to fetch all notes",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"notes": notes,
	})
}

func (h *Handler) GetNoteByID(c *gin.Context) {
	idstr := c.Param("id")

	objID, err := primitive.ObjectIDFromHex(idstr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid ID",
		})
		return
	}

	note, err := h.repo.GetByID(c.Request.Context(), objID)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			c.JSON(http.StatusFound, gin.H{
				"error": "Note not found for that given ID",
			})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to fetch the note",
		})
		return

	}

	c.JSON(http.StatusOK, note)
}


func (h *Handler) UpdateByID(c *gin.Context){
	idstr := c.Param("id")

	objID, err := primitive.ObjectIDFromHex(idstr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":"Invalid ID",
		})
		return
	}

	var req UpdateNoteRequest

	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":"Invalid json format",
		
		})
		return
	}

	updated, err := h.repo.UpdateByID(c.Request.Context(), objID, req)

	c.JSON(http.StatusOK, updated)
}

func (h *Handler) DeleteByID(c *gin.Context) {
	idstr := c.Param("id")

	objID, err := primitive.ObjectIDFromHex(idstr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":"Invalid ID",
		})
		return
	}
	deleted, err := h.repo.DeleteByID(c.Request.Context(),objID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":"failed to delete the note",
		})
		return
	}

	if !deleted {
		c.JSON(http.StatusNotFound,gin.H{
			"error" :"Note not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"ok":true,
	})
}