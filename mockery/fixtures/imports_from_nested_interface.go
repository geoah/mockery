package test

import (
	"github.com/geoah/mockery/mockery/fixtures/http"
)

type HasConflictingNestedImports interface {
	RequesterNS
	Z() http.MyStruct
}
