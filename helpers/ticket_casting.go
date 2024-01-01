package helpers

func StatusToText(status int16) string {
	switch status {
	case 1:
		return "To do"
	case 2:
		return "In Progress"
	case 0:
		return "Closed"
	}

	return "Closed"
}

func PriorityToText(status int16) string {
	switch status {
	case 1:
		return "High"
	case 2:
		return "Medium"
	case 3:
		return "Low"
	}

	return "IDK"
}
