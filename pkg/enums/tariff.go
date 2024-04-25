package enums

type WeekDays uint8

const (
	MO WeekDays = iota
	TU
	WE
	TH
	FR
	SA
	SU
)

func (weekDay WeekDays) String() string {
	switch weekDay {
	case MO:
		return "Monday"
	case TU:
		return "Tuesday"
	case WE:
		return "Wednesday"
	case TH:
		return "Thursday"
	case FR:
		return "Friday"
	case SA:
		return "Saturday"
	case SU:
		return "Sunday"
	}
	return "unknown"
}
