package query

import (
	_ "embed"
)

//go:embed insert.sql
var Insert string

//go:embed list.sql
var List string
