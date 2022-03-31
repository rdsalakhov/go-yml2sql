package main

import (
	"flag"
	"fmt"
	"github.com/rdsalakhov/go-yml2sql/yml2sql"
)

var (
	// table naming from dir name or file name
	_namingTyp = flag.String("t", yml2sql.NamingTypeDir, "table naming type")
	// target yaml file
	_target = flag.String("y", "", "yaml file")
	// is table name plural or not
	_plural = flag.Bool("plural", true, "pluralize table name")
	// allow insert null
	_null = flag.Bool("null", true, "allow null when data is missing")

	target = ""
)

// cli entry point
func main() {
	parseArgs()

	if target == "" {
		fmt.Printf("no target\nuse -y <yaml file path>\n")
		return
	}

	stmt := yml2sql.CreateStatementByFile(target)
	fmt.Println(stmt)
}

func parseArgs() {
	flag.Parse()

	target = *_target
	yml2sql.SetNamingTypeDir(*_namingTyp == yml2sql.NamingTypeDir)
	yml2sql.SetPlural(*_plural)
	yml2sql.SetNullable(*_null)
}
