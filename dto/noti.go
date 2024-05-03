package dto

import "poc_json/models"

type SNotificationResponse struct {
	ID       uint                   `json:"id"`
	Title    uint                   `json:"title"`
	Subtitle string                 `json:"subtitle"`
	Status   models.ST_NOTIFICATION `json:"status"`
}
