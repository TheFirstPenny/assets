package currency

var (
	RussianRuble    = Currency{alphabeticCode: "RUB", numericCode: 643, minorUnitMultiplier: 100, entity: "Russian ruble"}
	BelarusianRuble = Currency{alphabeticCode: "BYN", numericCode: 933, minorUnitMultiplier: 100, entity: "Belarusian ruble"}
	UsDollar        = Currency{alphabeticCode: "USD", numericCode: 840, minorUnitMultiplier: 100, entity: "United States dollar"}
	Euro            = Currency{alphabeticCode: "EUR", numericCode: 978, minorUnitMultiplier: 100, entity: "Euro"}
)

type Currency struct {
	alphabeticCode      string
	numericCode         int
	minorUnitMultiplier int
	entity              string
}

func (c *Currency) AlphabeticCode() string {
	return c.alphabeticCode
}

func (c *Currency) NumericCode() int {
	return c.numericCode
}

func (c *Currency) MinorUnitMultiplier() int {
	return c.minorUnitMultiplier
}

func (c *Currency) Entity() string {
	return c.entity
}

func (c *Currency) IsEqual(cc *Currency) bool {
	return c.alphabeticCode == cc.alphabeticCode
}

type CurrencyPair struct {
	baseCurrency  *Currency
	qouteCurrency *Currency
}
