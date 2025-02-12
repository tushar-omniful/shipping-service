package mappings

func MappedStatus(status string) string {
	statusMap := map[string]string{
		"CREATED":    "created",
		"PICKED_UP":  "in_transit",
		"IN_TRANSIT": "in_transit",
		"DELIVERED":  "delivered",
		"CANCELLED":  "cancelled",
		"FAILED":     "failed",
	}

	if mappedStatus, ok := statusMap[status]; ok {
		return mappedStatus
	}
	return "unknown"
}
