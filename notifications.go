package awesomeminer

import "time"

// NotificationResult describes the result of an API request to Awesome Miner's
// notifications endpoint. e.g. GET http://mypc:17790/api/notifications
type NotificationResult struct {
	Notifications                   []Notification `json:"notificationList"`
	NotificationCount               int            `json:"notificationCount"`
	UnacknowledgedNotificationCount int            `json:"unacknowledgedNotificationCount"`
}

// Notification describes a single notification record as obtained from
// Awesome Miner's notifications endpoint, as part of a NotificationResult.
type Notification struct {
	ID             int       `json:"id"`
	MinerName      string    `json:"minerName"`
	Title          string    `json:"title,omitempty"`
	TimeUTC        time.Time `json:"timeUtc"`
	Time           string    `json:"time"`
	Source         string    `json:"source"`
	Message        string    `json:"message"`
	IsAcknowledged bool      `json:"isAcknowledged"`
	IsWarning      bool      `json:"isWarning"`
}

// GetNotifications fetches and returns all notifications from Awesome Miner's
// notifications endpoint.
func (c *Client) GetNotifications() (*NotificationResult, error) {
	result := &NotificationResult{}
	_, err := c.doGetRequest("notifications", result)

	if err != nil {
		return nil, err
	}

	return result, nil
}
