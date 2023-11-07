package resolved_ast

import (
	"unsafe"

	"github.com/nutshelllabs/go-zetasql/types"
)

//go:linkname newValue github.com/nutshelllabs/go-zetasql/types.newValue
func newValue(unsafe.Pointer) types.Value

//go:linkname getRawValue github.com/nutshelllabs/go-zetasql/types.getRawValue
func getRawValue(types.Value) unsafe.Pointer

//go:linkname newType github.com/nutshelllabs/go-zetasql/types.newType
func newType(unsafe.Pointer) types.Type

//go:linkname getRawType github.com/nutshelllabs/go-zetasql/types.getRawType
func getRawType(types.Type) unsafe.Pointer

//go:linkname newTypeParameters github.com/nutshelllabs/go-zetasql/types.newTypeParameters
func newTypeParameters(unsafe.Pointer) *types.TypeParameters

//go:linkname getRawTypeParameters github.com/nutshelllabs/go-zetasql/types.getRawTypeParameters
func getRawTypeParameters(*types.TypeParameters) unsafe.Pointer

//go:linkname newAnnotationMap github.com/nutshelllabs/go-zetasql/types.newAnnotationMap
func newAnnotationMap(unsafe.Pointer) types.AnnotationMap

//go:linkname getRawAnnotationMap github.com/nutshelllabs/go-zetasql/types.getRawAnnotationMap
func getRawAnnotationMap(types.AnnotationMap) unsafe.Pointer

//go:linkname newAnnotatedType github.com/nutshelllabs/go-zetasql/types.newAnnotatedType
func newAnnotatedType(unsafe.Pointer) *types.AnnotatedType

//go:linkname getRawAnnotatedType github.com/nutshelllabs/go-zetasql/types.getRawAnnotatedType
func getRawAnnotatedType(*types.AnnotatedType) unsafe.Pointer

//go:linkname newConstant github.com/nutshelllabs/go-zetasql/types.newConstant
func newConstant(unsafe.Pointer) types.Constant

//go:linkname getRawConstant github.com/nutshelllabs/go-zetasql/types.getRawConstant
func getRawConstant(types.Constant) unsafe.Pointer

//go:linkname newFunction github.com/nutshelllabs/go-zetasql/types.newFunction
func newFunction(unsafe.Pointer) *types.Function

//go:linkname getRawFunction github.com/nutshelllabs/go-zetasql/types.getRawFunction
func getRawFunction(*types.Function) unsafe.Pointer

//go:linkname newFunctionSignature github.com/nutshelllabs/go-zetasql/types.newFunctionSignature
func newFunctionSignature(unsafe.Pointer) *types.FunctionSignature

//go:linkname getRawFunctionSignature github.com/nutshelllabs/go-zetasql/types.getRawFunctionSignature
func getRawFunctionSignature(*types.FunctionSignature) unsafe.Pointer

//go:linkname newModel github.com/nutshelllabs/go-zetasql/types.newModel
func newModel(unsafe.Pointer) types.Model

//go:linkname getRawModel github.com/nutshelllabs/go-zetasql/types.getRawModel
func getRawModel(types.Model) unsafe.Pointer

//go:linkname newConnection github.com/nutshelllabs/go-zetasql/types.newConnection
func newConnection(unsafe.Pointer) types.Connection

//go:linkname getRawConnection github.com/nutshelllabs/go-zetasql/types.getRawConnection
func getRawConnection(types.Connection) unsafe.Pointer

//go:linkname newTable github.com/nutshelllabs/go-zetasql/types.newTable
func newTable(unsafe.Pointer) types.Table

//go:linkname getRawTable github.com/nutshelllabs/go-zetasql/types.getRawTable
func getRawTable(types.Table) unsafe.Pointer

//go:linkname newTableValuedFunction github.com/nutshelllabs/go-zetasql/types.newSQLTableValuedFunction
func newTableValuedFunction(unsafe.Pointer) types.TableValuedFunction

//go:linkname getRawTableValuedFunction github.com/nutshelllabs/go-zetasql/types.getRawTableValuedFunction
func getRawTableValuedFunction(types.TableValuedFunction) unsafe.Pointer

//go:linkname newTVFSignature github.com/nutshelllabs/go-zetasql/types.newTVFSignature
func newTVFSignature(unsafe.Pointer) *types.TVFSignature

//go:linkname getRawTVFSignature github.com/nutshelllabs/go-zetasql/types.getRawTVFSignature
func getRawTVFSignature(*types.TVFSignature) unsafe.Pointer

//go:linkname newProcedure github.com/nutshelllabs/go-zetasql/types.newProcedure
func newProcedure(unsafe.Pointer) types.Procedure

//go:linkname getRawProcedure github.com/nutshelllabs/go-zetasql/types.getRawProcedure
func getRawProcedure(types.Procedure) unsafe.Pointer

//go:linkname newParseLocationRange github.com/nutshelllabs/go-zetasql/types.newParseLocationRange
func newParseLocationRange(unsafe.Pointer) *types.ParseLocationRange

//go:linkname getRawParseLocationRange github.com/nutshelllabs/go-zetasql/types.getRawParseLocationRange
func getRawParseLocationRange(*types.ParseLocationRange) unsafe.Pointer
