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
