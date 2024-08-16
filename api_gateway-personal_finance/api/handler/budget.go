package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	pb "api-gateway/genproto"
)

// CreateBudget handles the creation of a new budget
// @Summary      Create Budget
// @Description  Create a new budget for a user
// @Tags         Budget
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        request body pb.CreateBudgetRequest true "Create budget request"
// @Success      200 {object} pb.MessageResponsee "Budget created successfully"
// @Failure      500 {string} string "Error while creating budget"
// @Router       /budget/create [post]
func (h *Handler) CreateBudget(ctx *gin.Context) {
	req := &pb.CreateBudgetRequest{}
	err := ctx.BindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	res, err := h.Budget.CreateBudget(ctx, req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, res)
}

// ListBudgets handles listing all budgets
// @Summary      List Budgets
// @Description  List all budgets for a user
// @Tags         Budget
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Success      200 {object} pb.ListBudgetsResponse "List of budgets"
// @Failure      500 {string} string "Error while listing budgets"
// @Router       /budget/list [get]
func (h *Handler) ListBudgets(ctx *gin.Context) {
	req := &pb.ListBudgetsRequest{}

	res, err := h.Budget.ListBudgets(ctx, req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, res)
}

// GetBudgetById handles retrieving a specific budget by ID
// @Summary      Get Budget by ID
// @Description  Retrieve a budget's details by its ID
// @Tags         Budget
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        budget_id query string true "Budget ID"
// @Success      200 {object} pb.BudgetResponse "Budget details"
// @Failure      400 {string} string "Missing required query parameter: budget_id"
// @Failure      500 {string} string "Error while fetching budget"
// @Router       /budget/get/{budget_id} [get]
func (h *Handler) GetBudgetById(ctx *gin.Context) {
	budgetId := ctx.Query("budget_id")
	if budgetId == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Missing required query parameter: budget_id"})
		return
	}

	req := &pb.GetBudgetByIdRequest{BudgetId: budgetId}

	res, err := h.Budget.GetBudgetById(ctx, req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, res)
}

// UpdateBudget handles updating an existing budget
// @Summary      Update Budget
// @Description  Update the details of an existing budget
// @Tags         Budget
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        request body pb.UpdateBudgetRequest true "Update budget request"
// @Success      200 {object} pb.MessageResponsee "Budget updated successfully"
// @Failure      400 {string} string "Invalid request or missing budget_id"
// @Failure      500 {string} string "Error while updating budget"
// @Router       /budget/update [put]
func (h *Handler) UpdateBudget(ctx *gin.Context) {
	req := &pb.UpdateBudgetRequest{}
	err := ctx.BindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	if req.BudgetId == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Missing budget_id"})
		return
	}

	res, err := h.Budget.UpdateBudget(ctx, req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, res)
}

// DeleteBudget handles deleting a budget by ID
// @Summary      Delete Budget
// @Description  Delete an existing budget by its ID
// @Tags         Budget
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        budget_id query string true "Budget ID"
// @Success      200 {object} pb.BudgetDeleteResponse "Budget deleted successfully"
// @Failure      400 {string} string "Missing required query parameter: budget_id"
// @Failure      500 {string} string "Error while deleting budget"
// @Router       /budget/delete/{budget_id} [delete]
func (h *Handler) DeleteBudget(ctx *gin.Context) {
	budgetId := ctx.Query("budget_id")
	if budgetId == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Missing required query parameter: budget_id"})
		return
	}

	req := &pb.DeleteBudgetRequest{BudgetId: budgetId}

	res, err := h.Budget.DeleteBudget(ctx, req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, res)
}
