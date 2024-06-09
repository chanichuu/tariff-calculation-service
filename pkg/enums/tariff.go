package enums

import (
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

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

func (wd WeekDays) MarshalDynamoDBAttributeValue() (types.AttributeValue, error) {
	s := wd.String()
	return &types.AttributeValueMemberS{Value: s}, nil
}

func (wd WeekDays) UnmarshalDynamoDBAttributeValue(av types.AttributeValue) error {
	s := av.(*types.AttributeValueMemberS).Value
	switch s {
	case "Monday":
		wd = MO
	case "Tuesday":
		wd = TU
	case "Wednesday":
		wd = WE
	case "Thursday":
		wd = TH
	case "Friday":
		wd = FR
	case "Saturday":
		wd = SA
	case "Sunday":
		wd = SU
	}
	return nil
}
