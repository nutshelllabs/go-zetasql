package ast

import (
	"unsafe"

	"github.com/nutshelllabs/go-zetasql/types"
)

//go:linkname newParseLocationPoint github.com/nutshelllabs/go-zetasql/types.newParseLocationPoint
func newParseLocationPoint(unsafe.Pointer) *types.ParseLocationPoint

//go:linkname getRawParseLocationPoint github.com/nutshelllabs/go-zetasql/types.getRawParseLocationPoint
func getRawParseLocationPoint(*types.ParseLocationPoint) unsafe.Pointer

//go:linkname newParseLocationRange github.com/nutshelllabs/go-zetasql/types.newParseLocationRange
func newParseLocationRange(unsafe.Pointer) *types.ParseLocationRange

//go:linkname getRawParseLocationRange github.com/nutshelllabs/go-zetasql/types.getRawParseLocationRange
func getRawParseLocationRange(*types.ParseLocationRange) unsafe.Pointer
