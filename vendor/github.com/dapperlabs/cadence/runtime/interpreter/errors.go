package interpreter

import (
	"fmt"

	"github.com/dapperlabs/cadence/runtime/ast"
	"github.com/dapperlabs/cadence/runtime/common"
	"github.com/dapperlabs/cadence/runtime/sema"
)

// unsupportedOperation

type unsupportedOperation struct {
	kind      common.OperationKind
	operation ast.Operation
	ast.Range
}

func (e *unsupportedOperation) Error() string {
	return fmt.Sprintf(
		"cannot evaluate unsupported %s operation: %s",
		e.kind.Name(),
		e.operation.Symbol(),
	)
}

// NotDeclaredError

type NotDeclaredError struct {
	ExpectedKind common.DeclarationKind
	Name         string
}

func (e *NotDeclaredError) Error() string {
	return fmt.Sprintf(
		"cannot find %s in this scope: `%s`",
		e.ExpectedKind.Name(),
		e.Name,
	)
}

func (e *NotDeclaredError) SecondaryError() string {
	return "not found in this scope"
}

// NotInvokableError

type NotInvokableError struct {
	Value Value
}

func (e *NotInvokableError) Error() string {
	return fmt.Sprintf("cannot call value: %#+v", e.Value)
}

// ArgumentCountError

type ArgumentCountError struct {
	ParameterCount int
	ArgumentCount  int
}

func (e *ArgumentCountError) Error() string {
	return fmt.Sprintf(
		"incorrect number of arguments: expected %d, got %d",
		e.ParameterCount,
		e.ArgumentCount,
	)
}

// InvalidParameterTypeInInvocationError

type InvalidParameterTypeInInvocationError struct {
	InvalidParameterType sema.Type
}

func (e *InvalidParameterTypeInInvocationError) Error() string {
	return fmt.Sprintf("cannot invoke functions with parameter type: `%s`", e.InvalidParameterType)
}

// TransactionNotDeclaredError

type TransactionNotDeclaredError struct {
	Index int
}

func (e *TransactionNotDeclaredError) Error() string {
	return fmt.Sprintf(
		"cannot find transaction with index %d in this scope",
		e.Index,
	)
}

// ConditionError

type ConditionError struct {
	ConditionKind ast.ConditionKind
	Message       string
	LocationRange
}

func (e *ConditionError) Error() string {
	if e.Message == "" {
		return fmt.Sprintf("%s failed", e.ConditionKind.Name())
	}
	return fmt.Sprintf("%s failed: %s", e.ConditionKind.Name(), e.Message)
}

// RedeclarationError

type RedeclarationError struct {
	Name string
}

func (e *RedeclarationError) Error() string {
	return fmt.Sprintf("cannot redeclare: `%s` is already declared", e.Name)
}

// DereferenceError

type DereferenceError struct {
	LocationRange
}

func (e *DereferenceError) Error() string {
	return "dereference failed"
}

// OverflowError

type OverflowError struct{}

func (e OverflowError) Error() string {
	return "overflow"
}

// UnderflowError

type UnderflowError struct{}

func (e UnderflowError) Error() string {
	return "underflow"
}

// UnderflowError

type DivisionByZeroError struct{}

func (e DivisionByZeroError) Error() string {
	return "division by zero"
}

// DestroyedCompositeError

type DestroyedCompositeError struct {
	CompositeKind common.CompositeKind
	LocationRange
}

func (e *DestroyedCompositeError) Error() string {
	return fmt.Sprintf("%s is destroyed", e.CompositeKind)
}

// ForceAssignmentToNonNilResourceError

type ForceAssignmentToNonNilResourceError struct {
	LocationRange
}

func (e *ForceAssignmentToNonNilResourceError) Error() string {
	return "force assignment to non-nil resource-typed value"
}

// ForceNilError

type ForceNilError struct {
	LocationRange
}

func (e *ForceNilError) Error() string {
	return "unexpectedly found nil while forcing an Optional value"
}

// TypeMismatchError

type TypeMismatchError struct {
	ExpectedType sema.Type
	LocationRange
}

func (e *TypeMismatchError) Error() string {
	return fmt.Sprintf(
		"unexpectedly found non-`%s` while force-casting value",
		e.ExpectedType.QualifiedString(),
	)
}

// InvalidSavePathDomainError

type InvalidSavePathDomainError struct {
	ActualDomain   common.PathDomain
	ExpectedDomain common.PathDomain
	LocationRange
}

func (e *InvalidSavePathDomainError) Error() string {
	return fmt.Sprintf(
		"invalid path domain when saving value: expected `%s`, got `%s`",
		e.ExpectedDomain.Identifier(),
		e.ActualDomain.Identifier(),
	)
}

// OverwriteError

type OverwriteError struct {
	Address common.Address
	Path    PathValue
	LocationRange
}

func (e *OverwriteError) Error() string {
	return fmt.Sprintf(
		"failed to save object: path %s in account %s already stores an object",
		e.Path,
		e.Address,
	)
}