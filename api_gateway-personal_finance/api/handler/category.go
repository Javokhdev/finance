package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	pb "api-gateway/genproto"
)

// CreateCategory handles the creation of a new category
// @Summary      Create Category
// @Description  Create a new category for a user
// @Tags         Category
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        request body pb.CreateCategoryRequest true "Create category request"
// @Success      200 {object} pb.MessageResponse "Category created successfully"
// @Failure      500 {string} string "Error while creating category"
// @Router       /category/create [post]
func (h *Handler) CreateCategory(ctx *gin.Context) {
	req := &pb.CreateCategoryRequest{}
	err := ctx.BindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	res, err := h.Category.CreateCategory(ctx, req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, res)
}

// ListCategories handles listing all categories
// @Summary      List Categories
// @Description  List all categories for a user
// @Tags         Category
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Success      200 {object} pb.ListResponse "List of categories"
// @Failure      500 {string} string "Error while listing categories"
// @Router       /category/list [get]
func (h *Handler) ListCategories(ctx *gin.Context) {
	req := &pb.ListCategoriesRequest{}

	res, err := h.Category.ListCategories(ctx, req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, res)
}

// GetCategoryById handles retrieving a specific category by ID
// @Summary      Get Category by ID
// @Description  Retrieve a category's details by its ID
// @Tags         Category
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        category_id query string true "Category ID"
// @Success      200 {object} pb.CategoryResponse "Category details"
// @Failure      400 {string} string "Missing required query parameter: category_id"
// @Failure      500 {string} string "Error while fetching category"
// @Router       /category/get/{category_id} [get]
func (h *Handler) GetCategoryById(ctx *gin.Context) {
	categoryId := ctx.Query("category_id")
	if categoryId == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Missing required query parameter: category_id"})
		return
	}

	req := &pb.GetCategoryByIdRequest{CategoryId: categoryId}

	res, err := h.Category.GetCategoryById(ctx, req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, res)
}

// UpdateCategory handles updating an existing category
// @Summary      Update Category
// @Description  Update the details of an existing category
// @Tags         Category
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        request body pb.UpdateCategoryRequest true "Update category request"
// @Success      200 {object} pb.MessageResponse "Category updated successfully"
// @Failure      400 {string} string "Invalid request or missing category_id"
// @Failure      500 {string} string "Error while updating category"
// @Router       /category/update [put]
func (h *Handler) UpdateCategory(ctx *gin.Context) {
	req := &pb.UpdateCategoryRequest{}
	err := ctx.BindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	if req.CategoryId == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Missing category_id"})
		return
	}

	res, err := h.Category.UpdateCategory(ctx, req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, res)
}

// DeleteCategory handles deleting a category by ID
// @Summary      Delete Category
// @Description  Delete an existing category by its ID
// @Tags         Category
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        category_id query string true "Category ID"
// @Success      200 {object} pb.CategoryDeleteResponse "Category deleted successfully"
// @Failure      400 {string} string "Missing required query parameter: category_id"
// @Failure      500 {string} string "Error while deleting category"
// @Router       /category/delete/{category_id} [delete]
func (h *Handler) DeleteCategory(ctx *gin.Context) {
	categoryId := ctx.Query("category_id")
	if categoryId == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Missing required query parameter: category_id"})
		return
	}

	req := &pb.DeleteCategoryRequest{CategoryId: categoryId}

	res, err := h.Category.DeleteCategory(ctx, req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, res)
}
