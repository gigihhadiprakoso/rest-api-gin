package seed

import (
	. "rest-api-gin/models"
)

var unit_types = []UnitTypes{
	UnitTypes{
		Name: "General",
	},
	UnitTypes{
		Name: "Weight",
	},
	UnitTypes{
		Name: "Length",
	},
	UnitTypes{
		Name: "Volume",
	},
	UnitTypes{
		Name: "Area",
	},
}