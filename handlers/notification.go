package handlers

import (
	"net/http"

	"github.com/eesoymilk/health-statistic-api/ent/notification"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

//	@Summary				Get Notifications
//	@Description.markdown	notifications.get
//	@Tags					Notification
//	@Produce				json
//	@Success				200	{object}	[]types.NotificationWithRecipient
//	@Router					/notifications [get]
func (h *NotificationHandler) GetNotifications(c *gin.Context) {
	notifications, err := h.DB.Notification.
		Query().
		WithRecipient().
		All(c.Request.Context())

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, notifications)
}

//	@Summary				Get Notification
//	@Description.markdown	notification.get
//	@Tags					Notification
//	@Produce				json
//	@Param					id	path		string	true	"The notification's ID"
//	@Success				200	{object}	types.NotificationWithRecipient
//	@Router					/notifications/{id} [get]
func (h *NotificationHandler) GetNotification(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	notifications, err := h.DB.Notification.
		Query().
		Where(notification.ID((id))).
		WithRecipient().
		Only(c.Request.Context())

	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{"error": err.Error()},
		)
		return
	}

	c.JSON(http.StatusOK, notifications)
}

//	@Summary				Delete Notification
//	@Description.markdown	notification.delete
//	@Tags					Notification
//	@Produce				json
//	@Param					id	path	string	true	"The notification's ID."
//	@Success				200
//	@Router					/notifications/{id} [delete]
func (h *NotificationHandler) DeleteNotification(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err := h.DB.Notification.DeleteOneID(id).Exec(c.Request.Context()); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
