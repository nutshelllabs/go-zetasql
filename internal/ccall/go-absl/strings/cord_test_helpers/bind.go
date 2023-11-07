package cord_test_helpers

/*
#cgo CXXFLAGS: -std=c++11
#cgo CXXFLAGS: -I../../../
#cgo CXXFLAGS: -I../../../absl/strings
*/
import "C"

import (
	_ "github.com/nutshelllabs/go-zetasql/internal/ccall/go-absl/base/config"
	_ "github.com/nutshelllabs/go-zetasql/internal/ccall/go-absl/strings/cord"
	_ "github.com/nutshelllabs/go-zetasql/internal/ccall/go-absl/strings/cord_internal"
	_ "github.com/nutshelllabs/go-zetasql/internal/ccall/go-absl/strings/strings"
)
