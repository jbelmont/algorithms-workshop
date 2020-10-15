package table

import (
	"testing"
)

func TestTable(t *testing.T) {
	table := Table{}
	table.Name = "Order"
	table.CNames = []string{
		"ID", "Name", "Description",
	}

	rows := make([]Row, 2, 4)
	rows[0] = Row{}

	threeColumn := make([]Column, 3)
	threeColumn[0] = Column{
		ID:    0,
		Value: "Pizza",
	}
	threeColumn[1] = Column{
		ID:    1,
		Value: "Wings",
	}
	threeColumn[2] = Column{
		ID:    2,
		Value: "Coke",
	}
	rows[0].Columns = threeColumn

	rows[1] = Row{}
	threeColumn2 := make([]Column, 3)
	threeColumn2[0] = Column{
		ID:    3,
		Value: "Hamburger",
	}
	threeColumn2[1] = Column{
		ID:    4,
		Value: "Fries",
	}
	threeColumn2[2] = Column{
		ID:    5,
		Value: "Donut",
	}

	table.Rows = rows
	table.Print()
}
