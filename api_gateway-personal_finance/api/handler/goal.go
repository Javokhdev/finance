package handler

import (
	"fmt"
	"net/http"

	pb "api-gateway/genproto"

	"github.com/gin-gonic/gin"
)

// CreateGoal handles creating a new goal
// @Summary      Create Goal
// @Description  Create a new financial goal for a user
// @Tags         Goal
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        request body pb.CreateGoalRequest true "Create goal request"
// @Success      200 {object} pb.Responsee "Goal created successfully"
// @Failure      500 {string} string "Error while creating goal"
// @Router       /goal/create [post]
func (h *Handler) CreateGoal(ctx *gin.Context) {
	req := &pb.CreateGoalRequest{}
	if err := ctx.BindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	res, err := h.Goal.CreateGoal(ctx, req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, res)
}

// ListGoals handles listing all goals
// @Summary      List Goals
// @Description  List all financial goals for a user
// @Tags         Goal
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Success      200 {object} pb.ListGoalsResponse "List of goals"
// @Failure      500 {string} string "Error while listing goals"
// @Router       /goal/list [get]
func (h *Handler) ListGoals(ctx *gin.Context) {
	req := &pb.ListGoalsRequest{}

	res, err := h.Goal.ListGoals(ctx, req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, res)
}

// GetGoalById handles retrieving a specific goal by ID
// @Summary      Get Goal by ID
// @Description  Retrieve a goal's details by its ID
// @Tags         Goal
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        goal_id query string true "Goal ID"
// @Success      200 {object} pb.GoalResponse "Goal details"
// @Failure      400 {string} string "Missing required query parameter: goal_id"
// @Failure      500 {string} string "Error while fetching goal"
// @Router       /goal/get/{goal_id} [get]
func (h *Handler) GetGoalById(ctx *gin.Context) {
	goalId := ctx.Query("goal_id")
	fmt.Print(goalId)
	if goalId == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Missing required query parameter: goal_id"})
		return
	}

	req := &pb.GetGoalByIdRequest{GoalId: goalId}

	res, err := h.Goal.GetGoalById(ctx, req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, res)
}

// UpdateGoal handles updating an existing goal
// @Summary      Update Goal
// @Description  Update the details of an existing goal
// @Tags         Goal
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        request body pb.UpdateGoalRequest true "Update goal request"
// @Success      200 {object} pb.Responsee "Goal updated successfully"
// @Failure      400 {string} string "Invalid request or missing goal_id"
// @Failure      500 {string} string "Error while updating goal"
// @Router       /goal/update [put]
func (h *Handler) UpdateGoal(ctx *gin.Context) {
	req := &pb.UpdateGoalRequest{}
	if err := ctx.BindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	if req.GoalId == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Missing goal_id"})
		return
	}

	res, err := h.Goal.UpdateGoal(ctx, req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, res)
}

// DeleteGoal handles deleting a goal by ID
// @Summary      Delete Goal
// @Description  Delete an existing goal by its ID
// @Tags         Goal
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        goal_id query string true "Goal ID"
// @Success      200 {object} pb.GoalDeleteResponse "Goal deleted successfully"
// @Failure      400 {string} string "Missing required query parameter: goal_id"
// @Failure      500 {string} string "Error while deleting goal"
// @Router       /goal/delete/{goal_id} [delete]
func (h *Handler) DeleteGoal(ctx *gin.Context) {
	goalId := ctx.Query("goal_id")
	if goalId == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Missing required query parameter: goal_id"})
		return
	}

	req := &pb.DeleteGoalRequest{GoalId: goalId}

	res, err := h.Goal.DeleteGoal(ctx, req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, res)
}
