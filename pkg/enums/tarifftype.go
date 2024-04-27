package enums

type TariffType uint8

const (
	Electricity TariffType = iota
	Water
	Gas
	Biogas
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
	}
	return "unknown"
}
