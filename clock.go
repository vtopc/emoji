package emoji

import "time"

// Clock returns clock emoji according to hours.
// Minutes are ignored.
func Clock(t time.Time) string {
	switch t.Hour() % 12 {
	case 0:
		return "🕛"
	case 1:
		return "🕐"
	case 2:
		return "🕑"
	case 3:
		return "🕒"
	case 4:
		return "🕓"
	case 5:
		return "🕔"
	case 6:
		return "🕕"
	case 7:
		return "🕖"
	case 8:
		return "🕗"
	case 9:
		return "🕘"
	case 10:
		return "🕙"
	case 11:
		return "🕚"
	default:
		return "🕤"
	}
}
