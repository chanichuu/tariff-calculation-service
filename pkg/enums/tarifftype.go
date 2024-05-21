package enums

type TariffType uint8

const (
	Electricity TariffType = iota
	Water
	Gas
	Biogas
	Oil
)

func (tariffType TariffType) String() string {
	switch tariffType {
	case Electricity:
		return "Electricity"
	case Water:
		return "Water"
	case Gas:
		return "Gas"
	case Biogas:
		return "Biogas"
	case Oil:
		return "Oil"
	}
	return "unknown"
}
