package mysql

import (
	"fmt"

	"go-migrate/pkg/interfaces"

	sk "go-migrate/go_solve_kit"
)

type Blueprint struct {
	metadata []*meta
}

func NewBlueprint() interfaces.Blueprint {
	return &Blueprint{}
}

func (bp *Blueprint) Id(name string, length int) {
	if length == 0 {
		length = 11 // Default length for INT
	}
	bp.metadata = append(bp.metadata, &meta{
		Name:          name,
		Type:          "INT",
		Length:        length,
		AutoIncrement: true,
		Primary:       true,
	})
}

func (bp *Blueprint) String(name string, length int) interfaces.Blueprint {
	if length == 0 {
		length = 255 // Default length for VARCHAR
	}
	bp.metadata = append(bp.metadata, &meta{
		Name:   name,
		Type:   "VARCHAR",
		Length: length,
	})
	return bp
}

func (bp *Blueprint) Text(name string) interfaces.Blueprint {
	bp.metadata = append(bp.metadata, &meta{
		Name: name,
		Type: "TEXT",
	})
	return bp
}

func (bp *Blueprint) Integer(name string, length int) interfaces.Blueprint {
	if length == 0 {
		length = 11 // Default length for INT
	}
	bp.metadata = append(bp.metadata, &meta{
		Name:   name,
		Type:   "INT",
		Length: length,
	})
	return bp
}

func (bp *Blueprint) Float(name string, length int, precision int) interfaces.Blueprint {
	if precision == 0 {
		precision = 2 // Default precision for FLOAT
	}
	bp.metadata = append(bp.metadata, &meta{
		Name:      name,
		Type:      "FLOAT",
		Length:    length,
		Precision: precision,
	})
	return bp
}

func (bp *Blueprint) Double(name string, length int, precision int) interfaces.Blueprint {
	if precision == 0 {
		precision = 2 // Default precision for DOUBLE
	}
	bp.metadata = append(bp.metadata, &meta{
		Name:      name,
		Type:      "DOUBLE",
		Length:    length,
		Precision: precision,
	})
	return bp
}

func (bp *Blueprint) Decimal(name string, length int, precision int) interfaces.Blueprint {
	if precision == 0 {
		precision = 2 // Default precision for DECIMAL
	}
	bp.metadata = append(bp.metadata, &meta{
		Name:      name,
		Type:      "DECIMAL",
		Length:    length,
		Precision: precision,
	})
	return bp
}

func (bp *Blueprint) Date(name string) interfaces.Blueprint {
	bp.metadata = append(bp.metadata, &meta{
		Name: name,
		Type: "DATE",
	})
	return bp
}

func (bp *Blueprint) Boolean(name string) interfaces.Blueprint {
	bp.metadata = append(bp.metadata, &meta{
		Name: name,
		Type: "TINYINT",
	})
	return bp
}

func (bp *Blueprint) DateTime(name string) interfaces.Blueprint {
	bp.metadata = append(bp.metadata, &meta{
		Name: name,
		Type: "DATETIME",
	})
	return bp
}

func (bp *Blueprint) Timestamps() {
	bp.metadata = append(bp.metadata, &meta{
		Name:    "created_at",
		Type:    "DATETIME",
		Default: "CURRENT_TIMESTAMP",
	})

	bp.metadata = append(bp.metadata, &meta{
		Name:     "updated_at",
		Type:     "DATETIME",
		Nullable: true,
		Default:  "NULL",
	})
}

func (bp *Blueprint) Nullable() interfaces.Blueprint {
	bp.metadata[len(bp.metadata)-1].Nullable = true
	return bp
}

func (bp *Blueprint) Unique(column ...string) interfaces.Blueprint {
	if len(column) == 0 {
		bp.metadata[len(bp.metadata)-1].Unique = true
	} else {
		for _, c := range column {
			bp.metadata = append(bp.metadata, &meta{
				Name:   c,
				Unique: true,
			})
		}
	}
	return bp
}

func (bp *Blueprint) Index(column ...string) interfaces.Blueprint {
	if len(column) == 0 {
		bp.metadata[len(bp.metadata)-1].Index = true
	} else {
		for _, c := range column {
			bp.metadata = append(bp.metadata, &meta{
				Name:  c,
				Index: true,
			})
		}
	}
	return bp
}

func (bp *Blueprint) Default(value interface{}) interfaces.Blueprint {
	bp.metadata[len(bp.metadata)-1].Default = fmt.Sprintf("'%v'", value)
	return bp
}

func (bp *Blueprint) UnsignedBigInteger(name string) interfaces.Blueprint {
	bp.metadata = append(bp.metadata, &meta{
		Name:   name,
		Type:   "UNSIGNED BIGINT",
		Length: 20,
	})
	return bp
}
func (bp *Blueprint) UnsignedInteger(name string, length int) interfaces.Blueprint {
	bp.metadata = append(bp.metadata, &meta{
		Name:   name,
		Type:   "UNSIGNED INTEGER",
		Length: length,
	})
	return bp
}
func (bp *Blueprint) Foreign(name string) interfaces.ForeignBlueprint {
	fb := newForeignBlueprint().(*foreignBlueprint)
	bp.metadata = append(bp.metadata, &meta{
		Name:    name,
		Foreign: fb.meta,
	})
	return fb
}

func (bp *Blueprint) Primary(name ...string) interfaces.Blueprint {
	bp.metadata = append(bp.metadata, &meta{
		Name:    sk.FromStringArray(name).Join("`, `").ValueOf(),
		Primary: true,
	})
	return bp
}

func (bp *Blueprint) DropColumn(column string) {
	bp.metadata = append(bp.metadata, &meta{
		Name: column,
		Type: "DROP",
	})
}

func (bp *Blueprint) DropUnique(name string) {
	bp.metadata = append(bp.metadata, &meta{
		Name:   name,
		Type:   "DROP",
		Unique: true,
	})
}
func (bp *Blueprint) DropIndex(name string) {
	bp.metadata = append(bp.metadata, &meta{
		Name:  name,
		Type:  "DROP",
		Index: true,
	})
}
func (bp *Blueprint) DropForeign(name string) {
	bp.metadata = append(bp.metadata, &meta{
		Name:    name,
		Type:    "DROP",
		Foreign: newForeignBlueprint().(*foreignBlueprint).meta,
	})
}
func (bp *Blueprint) DropPrimary() {
	bp.metadata = append(bp.metadata, &meta{
		Type:    "DROP",
		Primary: true,
	})
}

func (bp *Blueprint) GetSqls(table string, operation operation) []string {
	return operation.generateSql(table, bp.metadata)
}
