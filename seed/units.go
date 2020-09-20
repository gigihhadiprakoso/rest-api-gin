package seed

import (
	. "rest-api-gin/models"
)

var units = []Units{
	Units{
		Name: "Pieces",
		Symbol: "PCS",
		UnitTypeID: 1,
	},
	Units{
		Name: "Kilogram",
		Symbol: "Kg",
		UnitTypeID: 2,
	},
	Units{
		Name: "Gram",
		Symbol: "G",
		UnitTypeID: 2,
	},
	Units{
		Name: "Miligram",
		Symbol: "Mg",
		UnitTypeID: 2,
	},
	Units{
		Name: "Kilometer",
		Symbol: "Km",
		UnitTypeID: 3,
	},
	Units{
		Name: "Meter",
		Symbol: "M",
		UnitTypeID: 3,
	},
	Units{
		Name: "Centimeter",
		Symbol: "Cm",
		UnitTypeID: 3,
	},
	Units{
		Name: "Liter",
		Symbol: "L",
		UnitTypeID: 4,
	},
}