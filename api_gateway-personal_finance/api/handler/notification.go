package handler

import (
	pb "api-gateway/genproto"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetNotification handles retrieving a notification by user_id
// @Summary      Get Notification
// @Description  Retrieve a notification by user_id
// @Tags         Notification
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        user_id query string true "User ID"
// @Success      200 {object} pb.GetNotificationByidResponse "Notification retrieved successfully"
// @Failure      404 {string} string "Notification not found"
// @Failure      500 {string} string "Error while retrieving notification"
// @Router       /notification/get/{user_id} [get]
func (h *Handler) GetNotification(ctx *gin.Context) {
	user := ctx.Query("user_id")
	// fmt.Println("-----------------------------------------------")
	// fmt.Println(user)
	if user == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Missing required query parameter: user_id"})
		return
	}

	res, err := h.Notification.GetNotification(ctx, &pb.GetNotificationByidRequest{UserId: user})
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, res)
}

// // DeleteNotification handles deleting a notification by user_id
// // @Summary      Delete Notification
// // @Description  Delete a notification by user_id
// // @Tags         Notification
// // @Accept       json
// // @Produce      json
// // @Security     BearerAuth
// // @Param        id path string true "User ID"
// // @Success      200 {object} pb.NotificationsResponse "Notification deleted successfully"
// // @Failure      404 {string} string "Notification not found"
// // @Failure      500 {string} string "Error while deleting notification"
// // @Router       /notification/{id} [delete]
// func (h *Handler) DeleteNotification(ctx *gin.Context) {
// 	var req pb.GetNotificationByidRequest
// 	if err := ctx.ShouldBindUri(&req); err != nil {
// 		ctx.JSON(400, gin.H{"error": "Invalictxd ID"})
// 		return
// 	}

// 	res, err := h.Notification.DeleteNotification(ctx, &req)
// 	if err != nil {
// 		ctx.JSON(500, gin.H{"error": err.Error()})
// 		return
// 	}

// 	ctx.JSON(200, res)
// }

// ListNotification handles listing all notifications
// @Summary      List Notifications
// @Description  Retrieve a list of all notifications
// @Tags         Notification
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Success      200 {object} pb.ListNotificationResponse "List of notifications retrieved successfully"
// @Failure      500 {object} pb.Responsee "Error while retrieving notifications" // Use Responsee for general errors
// @Router       /notification/list [get]
func (h *Handler) ListNotification(ctx *gin.Context) {
	var req pb.Void

	res, err := h.Notification.ListNotification(ctx, &req)
	if err != nil {
		// Return a structured error response with message and success
		ctx.JSON(500, &pb.Responsee{
			Message: "Error while retrieving notifications",
		})
		return
	}

	ctx.JSON(200, res)
}
