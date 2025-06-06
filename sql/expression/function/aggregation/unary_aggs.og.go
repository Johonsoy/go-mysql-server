// Code generated by optgen; DO NOT EDIT.

package aggregation

import (
	"fmt"

	"github.com/dolthub/go-mysql-server/sql"
	"github.com/dolthub/go-mysql-server/sql/expression"
	"github.com/dolthub/go-mysql-server/sql/transform"
	"github.com/dolthub/go-mysql-server/sql/types"
)

type AnyValue struct {
	unaryAggBase
}

var _ sql.FunctionExpression = (*AnyValue)(nil)
var _ sql.Aggregation = (*AnyValue)(nil)
var _ sql.WindowAdaptableExpression = (*AnyValue)(nil)

func NewAnyValue(e sql.Expression) *AnyValue {
	return &AnyValue{
		unaryAggBase{
			UnaryExpression: expression.UnaryExpression{Child: e},
			functionName:    "AnyValue",
			description:     "returns any single value in the grouped rows",
		},
	}
}

func (a *AnyValue) Type() sql.Type {
	return a.Child.Type()
}

func (a *AnyValue) IsNullable() bool {
	return true
}

func (a *AnyValue) String() string {
	if a.window != nil {
		pr := sql.NewTreePrinter()
		_ = pr.WriteNode("ANYVALUE")
		children := []string{a.window.String(), a.Child.String()}
		pr.WriteChildren(children...)
		return pr.String()
	}
	return fmt.Sprintf("ANYVALUE(%s)", a.Child)
}

func (a *AnyValue) DebugString() string {
	if a.window != nil {
		pr := sql.NewTreePrinter()
		_ = pr.WriteNode("ANYVALUE")
		children := []string{sql.DebugString(a.window), sql.DebugString(a.Child)}
		pr.WriteChildren(children...)
		return pr.String()
	}
	return fmt.Sprintf("ANYVALUE(%s)", sql.DebugString(a.Child))
}

func (a *AnyValue) WithWindow(window *sql.WindowDefinition) sql.WindowAdaptableExpression {
	res := a.unaryAggBase.WithWindow(window)
	return &AnyValue{unaryAggBase: *res.(*unaryAggBase)}
}

func (a *AnyValue) WithChildren(children ...sql.Expression) (sql.Expression, error) {
	res, err := a.unaryAggBase.WithChildren(children...)
	return &AnyValue{unaryAggBase: *res.(*unaryAggBase)}, err
}

func (a *AnyValue) WithId(id sql.ColumnId) sql.IdExpression {
	res := a.unaryAggBase.WithId(id)
	return &AnyValue{unaryAggBase: *res.(*unaryAggBase)}
}

func (a *AnyValue) NewBuffer() (sql.AggregationBuffer, error) {
	child, err := transform.Clone(a.Child)
	if err != nil {
		return nil, err
	}
	return NewAnyValueBuffer(child), nil
}

func (a *AnyValue) NewWindowFunction() (sql.WindowFunction, error) {
	child, err := transform.Clone(a.Child)
	if err != nil {
		return nil, err
	}
	return NewAnyValueAgg(child).WithWindow(a.Window())
}

type Avg struct {
	unaryAggBase
}

var _ sql.FunctionExpression = (*Avg)(nil)
var _ sql.Aggregation = (*Avg)(nil)
var _ sql.WindowAdaptableExpression = (*Avg)(nil)

func NewAvg(e sql.Expression) *Avg {
	return &Avg{
		unaryAggBase{
			UnaryExpression: expression.UnaryExpression{Child: e},
			functionName:    "Avg",
			description:     "returns the average value of expr in all rows.",
		},
	}
}

func (a *Avg) Type() sql.Type {
	return a.Child.Type()
}

func (a *Avg) IsNullable() bool {
	return true
}

func (a *Avg) String() string {
	if a.window != nil {
		pr := sql.NewTreePrinter()
		_ = pr.WriteNode("AVG")
		children := []string{a.window.String(), a.Child.String()}
		pr.WriteChildren(children...)
		return pr.String()
	}
	return fmt.Sprintf("AVG(%s)", a.Child)
}

func (a *Avg) DebugString() string {
	if a.window != nil {
		pr := sql.NewTreePrinter()
		_ = pr.WriteNode("AVG")
		children := []string{sql.DebugString(a.window), sql.DebugString(a.Child)}
		pr.WriteChildren(children...)
		return pr.String()
	}
	return fmt.Sprintf("AVG(%s)", sql.DebugString(a.Child))
}

func (a *Avg) WithWindow(window *sql.WindowDefinition) sql.WindowAdaptableExpression {
	res := a.unaryAggBase.WithWindow(window)
	return &Avg{unaryAggBase: *res.(*unaryAggBase)}
}

func (a *Avg) WithChildren(children ...sql.Expression) (sql.Expression, error) {
	res, err := a.unaryAggBase.WithChildren(children...)
	return &Avg{unaryAggBase: *res.(*unaryAggBase)}, err
}

func (a *Avg) WithId(id sql.ColumnId) sql.IdExpression {
	res := a.unaryAggBase.WithId(id)
	return &Avg{unaryAggBase: *res.(*unaryAggBase)}
}

func (a *Avg) NewBuffer() (sql.AggregationBuffer, error) {
	child, err := transform.Clone(a.Child)
	if err != nil {
		return nil, err
	}
	return NewAvgBuffer(child), nil
}

func (a *Avg) NewWindowFunction() (sql.WindowFunction, error) {
	child, err := transform.Clone(a.Child)
	if err != nil {
		return nil, err
	}
	return NewAvgAgg(child).WithWindow(a.Window())
}

type BitAnd struct {
	unaryAggBase
}

var _ sql.FunctionExpression = (*BitAnd)(nil)
var _ sql.Aggregation = (*BitAnd)(nil)
var _ sql.WindowAdaptableExpression = (*BitAnd)(nil)

func NewBitAnd(e sql.Expression) *BitAnd {
	return &BitAnd{
		unaryAggBase{
			UnaryExpression: expression.UnaryExpression{Child: e},
			functionName:    "BitAnd",
			description:     "returns the bitwise AND of all bits in expr.",
		},
	}
}

func (a *BitAnd) Type() sql.Type {
	return types.Uint64
}

func (a *BitAnd) IsNullable() bool {
	return false
}

func (a *BitAnd) String() string {
	if a.window != nil {
		pr := sql.NewTreePrinter()
		_ = pr.WriteNode("BITAND")
		children := []string{a.window.String(), a.Child.String()}
		pr.WriteChildren(children...)
		return pr.String()
	}
	return fmt.Sprintf("BITAND(%s)", a.Child)
}

func (a *BitAnd) DebugString() string {
	if a.window != nil {
		pr := sql.NewTreePrinter()
		_ = pr.WriteNode("BITAND")
		children := []string{sql.DebugString(a.window), sql.DebugString(a.Child)}
		pr.WriteChildren(children...)
		return pr.String()
	}
	return fmt.Sprintf("BITAND(%s)", sql.DebugString(a.Child))
}

func (a *BitAnd) WithWindow(window *sql.WindowDefinition) sql.WindowAdaptableExpression {
	res := a.unaryAggBase.WithWindow(window)
	return &BitAnd{unaryAggBase: *res.(*unaryAggBase)}
}

func (a *BitAnd) WithChildren(children ...sql.Expression) (sql.Expression, error) {
	res, err := a.unaryAggBase.WithChildren(children...)
	return &BitAnd{unaryAggBase: *res.(*unaryAggBase)}, err
}

func (a *BitAnd) WithId(id sql.ColumnId) sql.IdExpression {
	res := a.unaryAggBase.WithId(id)
	return &BitAnd{unaryAggBase: *res.(*unaryAggBase)}
}

func (a *BitAnd) NewBuffer() (sql.AggregationBuffer, error) {
	child, err := transform.Clone(a.Child)
	if err != nil {
		return nil, err
	}
	return NewBitAndBuffer(child), nil
}

func (a *BitAnd) NewWindowFunction() (sql.WindowFunction, error) {
	child, err := transform.Clone(a.Child)
	if err != nil {
		return nil, err
	}
	return NewBitAndAgg(child).WithWindow(a.Window())
}

type BitOr struct {
	unaryAggBase
}

var _ sql.FunctionExpression = (*BitOr)(nil)
var _ sql.Aggregation = (*BitOr)(nil)
var _ sql.WindowAdaptableExpression = (*BitOr)(nil)

func NewBitOr(e sql.Expression) *BitOr {
	return &BitOr{
		unaryAggBase{
			UnaryExpression: expression.UnaryExpression{Child: e},
			functionName:    "BitOr",
			description:     "returns the bitwise OR of all bits in expr.",
		},
	}
}

func (a *BitOr) Type() sql.Type {
	return types.Uint64
}

func (a *BitOr) IsNullable() bool {
	return false
}

func (a *BitOr) String() string {
	if a.window != nil {
		pr := sql.NewTreePrinter()
		_ = pr.WriteNode("BITOR")
		children := []string{a.window.String(), a.Child.String()}
		pr.WriteChildren(children...)
		return pr.String()
	}
	return fmt.Sprintf("BITOR(%s)", a.Child)
}

func (a *BitOr) DebugString() string {
	if a.window != nil {
		pr := sql.NewTreePrinter()
		_ = pr.WriteNode("BITOR")
		children := []string{sql.DebugString(a.window), sql.DebugString(a.Child)}
		pr.WriteChildren(children...)
		return pr.String()
	}
	return fmt.Sprintf("BITOR(%s)", sql.DebugString(a.Child))
}

func (a *BitOr) WithWindow(window *sql.WindowDefinition) sql.WindowAdaptableExpression {
	res := a.unaryAggBase.WithWindow(window)
	return &BitOr{unaryAggBase: *res.(*unaryAggBase)}
}

func (a *BitOr) WithChildren(children ...sql.Expression) (sql.Expression, error) {
	res, err := a.unaryAggBase.WithChildren(children...)
	return &BitOr{unaryAggBase: *res.(*unaryAggBase)}, err
}

func (a *BitOr) WithId(id sql.ColumnId) sql.IdExpression {
	res := a.unaryAggBase.WithId(id)
	return &BitOr{unaryAggBase: *res.(*unaryAggBase)}
}

func (a *BitOr) NewBuffer() (sql.AggregationBuffer, error) {
	child, err := transform.Clone(a.Child)
	if err != nil {
		return nil, err
	}
	return NewBitOrBuffer(child), nil
}

func (a *BitOr) NewWindowFunction() (sql.WindowFunction, error) {
	child, err := transform.Clone(a.Child)
	if err != nil {
		return nil, err
	}
	return NewBitOrAgg(child).WithWindow(a.Window())
}

type BitXor struct {
	unaryAggBase
}

var _ sql.FunctionExpression = (*BitXor)(nil)
var _ sql.Aggregation = (*BitXor)(nil)
var _ sql.WindowAdaptableExpression = (*BitXor)(nil)

func NewBitXor(e sql.Expression) *BitXor {
	return &BitXor{
		unaryAggBase{
			UnaryExpression: expression.UnaryExpression{Child: e},
			functionName:    "BitXor",
			description:     "returns the bitwise XOR of all bits in expr.",
		},
	}
}

func (a *BitXor) Type() sql.Type {
	return types.Uint64
}

func (a *BitXor) IsNullable() bool {
	return false
}

func (a *BitXor) String() string {
	if a.window != nil {
		pr := sql.NewTreePrinter()
		_ = pr.WriteNode("BITXOR")
		children := []string{a.window.String(), a.Child.String()}
		pr.WriteChildren(children...)
		return pr.String()
	}
	return fmt.Sprintf("BITXOR(%s)", a.Child)
}

func (a *BitXor) DebugString() string {
	if a.window != nil {
		pr := sql.NewTreePrinter()
		_ = pr.WriteNode("BITXOR")
		children := []string{sql.DebugString(a.window), sql.DebugString(a.Child)}
		pr.WriteChildren(children...)
		return pr.String()
	}
	return fmt.Sprintf("BITXOR(%s)", sql.DebugString(a.Child))
}

func (a *BitXor) WithWindow(window *sql.WindowDefinition) sql.WindowAdaptableExpression {
	res := a.unaryAggBase.WithWindow(window)
	return &BitXor{unaryAggBase: *res.(*unaryAggBase)}
}

func (a *BitXor) WithChildren(children ...sql.Expression) (sql.Expression, error) {
	res, err := a.unaryAggBase.WithChildren(children...)
	return &BitXor{unaryAggBase: *res.(*unaryAggBase)}, err
}

func (a *BitXor) WithId(id sql.ColumnId) sql.IdExpression {
	res := a.unaryAggBase.WithId(id)
	return &BitXor{unaryAggBase: *res.(*unaryAggBase)}
}

func (a *BitXor) NewBuffer() (sql.AggregationBuffer, error) {
	child, err := transform.Clone(a.Child)
	if err != nil {
		return nil, err
	}
	return NewBitXorBuffer(child), nil
}

func (a *BitXor) NewWindowFunction() (sql.WindowFunction, error) {
	child, err := transform.Clone(a.Child)
	if err != nil {
		return nil, err
	}
	return NewBitXorAgg(child).WithWindow(a.Window())
}

type Count struct {
	unaryAggBase
}

var _ sql.FunctionExpression = (*Count)(nil)
var _ sql.Aggregation = (*Count)(nil)
var _ sql.WindowAdaptableExpression = (*Count)(nil)

func NewCount(e sql.Expression) *Count {
	return &Count{
		unaryAggBase{
			UnaryExpression: expression.UnaryExpression{Child: e},
			functionName:    "Count",
			description:     "returns a count of the number of non-NULL values of expr in the rows retrieved by a SELECT statement.",
		},
	}
}

func (a *Count) Type() sql.Type {
	return types.Int64
}

func (a *Count) IsNullable() bool {
	return false
}

func (a *Count) String() string {
	if a.window != nil {
		pr := sql.NewTreePrinter()
		_ = pr.WriteNode("COUNT")
		children := []string{a.window.String(), a.Child.String()}
		pr.WriteChildren(children...)
		return pr.String()
	}
	return fmt.Sprintf("COUNT(%s)", a.Child)
}

func (a *Count) DebugString() string {
	if a.window != nil {
		pr := sql.NewTreePrinter()
		_ = pr.WriteNode("COUNT")
		children := []string{sql.DebugString(a.window), sql.DebugString(a.Child)}
		pr.WriteChildren(children...)
		return pr.String()
	}
	return fmt.Sprintf("COUNT(%s)", sql.DebugString(a.Child))
}

func (a *Count) WithWindow(window *sql.WindowDefinition) sql.WindowAdaptableExpression {
	res := a.unaryAggBase.WithWindow(window)
	return &Count{unaryAggBase: *res.(*unaryAggBase)}
}

func (a *Count) WithChildren(children ...sql.Expression) (sql.Expression, error) {
	res, err := a.unaryAggBase.WithChildren(children...)
	return &Count{unaryAggBase: *res.(*unaryAggBase)}, err
}

func (a *Count) WithId(id sql.ColumnId) sql.IdExpression {
	res := a.unaryAggBase.WithId(id)
	return &Count{unaryAggBase: *res.(*unaryAggBase)}
}

func (a *Count) NewBuffer() (sql.AggregationBuffer, error) {
	child, err := transform.Clone(a.Child)
	if err != nil {
		return nil, err
	}
	return NewCountBuffer(child), nil
}

func (a *Count) NewWindowFunction() (sql.WindowFunction, error) {
	child, err := transform.Clone(a.Child)
	if err != nil {
		return nil, err
	}
	return NewCountAgg(child).WithWindow(a.Window())
}

type First struct {
	unaryAggBase
}

var _ sql.FunctionExpression = (*First)(nil)
var _ sql.Aggregation = (*First)(nil)
var _ sql.WindowAdaptableExpression = (*First)(nil)

func NewFirst(e sql.Expression) *First {
	return &First{
		unaryAggBase{
			UnaryExpression: expression.UnaryExpression{Child: e},
			functionName:    "First",
			description:     "returns the first value in a sequence of elements of an aggregation.",
		},
	}
}

func (a *First) Type() sql.Type {
	return a.Child.Type()
}

func (a *First) IsNullable() bool {
	return false
}

func (a *First) String() string {
	if a.window != nil {
		pr := sql.NewTreePrinter()
		_ = pr.WriteNode("FIRST")
		children := []string{a.window.String(), a.Child.String()}
		pr.WriteChildren(children...)
		return pr.String()
	}
	return fmt.Sprintf("FIRST(%s)", a.Child)
}

func (a *First) DebugString() string {
	if a.window != nil {
		pr := sql.NewTreePrinter()
		_ = pr.WriteNode("FIRST")
		children := []string{sql.DebugString(a.window), sql.DebugString(a.Child)}
		pr.WriteChildren(children...)
		return pr.String()
	}
	return fmt.Sprintf("FIRST(%s)", sql.DebugString(a.Child))
}

func (a *First) WithWindow(window *sql.WindowDefinition) sql.WindowAdaptableExpression {
	res := a.unaryAggBase.WithWindow(window)
	return &First{unaryAggBase: *res.(*unaryAggBase)}
}

func (a *First) WithChildren(children ...sql.Expression) (sql.Expression, error) {
	res, err := a.unaryAggBase.WithChildren(children...)
	return &First{unaryAggBase: *res.(*unaryAggBase)}, err
}

func (a *First) WithId(id sql.ColumnId) sql.IdExpression {
	res := a.unaryAggBase.WithId(id)
	return &First{unaryAggBase: *res.(*unaryAggBase)}
}

func (a *First) NewBuffer() (sql.AggregationBuffer, error) {
	child, err := transform.Clone(a.Child)
	if err != nil {
		return nil, err
	}
	return NewFirstBuffer(child), nil
}

func (a *First) NewWindowFunction() (sql.WindowFunction, error) {
	child, err := transform.Clone(a.Child)
	if err != nil {
		return nil, err
	}
	return NewFirstAgg(child).WithWindow(a.Window())
}

type JsonArray struct {
	unaryAggBase
}

var _ sql.FunctionExpression = (*JsonArray)(nil)
var _ sql.Aggregation = (*JsonArray)(nil)
var _ sql.WindowAdaptableExpression = (*JsonArray)(nil)

func NewJsonArray(e sql.Expression) *JsonArray {
	return &JsonArray{
		unaryAggBase{
			UnaryExpression: expression.UnaryExpression{Child: e},
			functionName:    "JsonArray",
			description:     "returns result set as a single JSON array.",
		},
	}
}

func (a *JsonArray) Type() sql.Type {
	return types.JSON
}

func (a *JsonArray) IsNullable() bool {
	return false
}

func (a *JsonArray) String() string {
	if a.window != nil {
		pr := sql.NewTreePrinter()
		_ = pr.WriteNode("JSON_ARRAYAGG")
		children := []string{a.window.String(), a.Child.String()}
		pr.WriteChildren(children...)
		return pr.String()
	}
	return fmt.Sprintf("JSON_ARRAYAGG(%s)", a.Child)
}

func (a *JsonArray) DebugString() string {
	if a.window != nil {
		pr := sql.NewTreePrinter()
		_ = pr.WriteNode("JSON_ARRAYAGG")
		children := []string{sql.DebugString(a.window), sql.DebugString(a.Child)}
		pr.WriteChildren(children...)
		return pr.String()
	}
	return fmt.Sprintf("JSON_ARRAYAGG(%s)", sql.DebugString(a.Child))
}

func (a *JsonArray) WithWindow(window *sql.WindowDefinition) sql.WindowAdaptableExpression {
	res := a.unaryAggBase.WithWindow(window)
	return &JsonArray{unaryAggBase: *res.(*unaryAggBase)}
}

func (a *JsonArray) WithChildren(children ...sql.Expression) (sql.Expression, error) {
	res, err := a.unaryAggBase.WithChildren(children...)
	return &JsonArray{unaryAggBase: *res.(*unaryAggBase)}, err
}

func (a *JsonArray) WithId(id sql.ColumnId) sql.IdExpression {
	res := a.unaryAggBase.WithId(id)
	return &JsonArray{unaryAggBase: *res.(*unaryAggBase)}
}

func (a *JsonArray) NewBuffer() (sql.AggregationBuffer, error) {
	child, err := transform.Clone(a.Child)
	if err != nil {
		return nil, err
	}
	return NewJsonArrayBuffer(child), nil
}

func (a *JsonArray) NewWindowFunction() (sql.WindowFunction, error) {
	child, err := transform.Clone(a.Child)
	if err != nil {
		return nil, err
	}
	return NewJsonArrayAgg(child).WithWindow(a.Window())
}

type Last struct {
	unaryAggBase
}

var _ sql.FunctionExpression = (*Last)(nil)
var _ sql.Aggregation = (*Last)(nil)
var _ sql.WindowAdaptableExpression = (*Last)(nil)

func NewLast(e sql.Expression) *Last {
	return &Last{
		unaryAggBase{
			UnaryExpression: expression.UnaryExpression{Child: e},
			functionName:    "Last",
			description:     "returns the last value in a sequence of elements of an aggregation.",
		},
	}
}

func (a *Last) Type() sql.Type {
	return a.Child.Type()
}

func (a *Last) IsNullable() bool {
	return false
}

func (a *Last) String() string {
	if a.window != nil {
		pr := sql.NewTreePrinter()
		_ = pr.WriteNode("LAST")
		children := []string{a.window.String(), a.Child.String()}
		pr.WriteChildren(children...)
		return pr.String()
	}
	return fmt.Sprintf("LAST(%s)", a.Child)
}

func (a *Last) DebugString() string {
	if a.window != nil {
		pr := sql.NewTreePrinter()
		_ = pr.WriteNode("LAST")
		children := []string{sql.DebugString(a.window), sql.DebugString(a.Child)}
		pr.WriteChildren(children...)
		return pr.String()
	}
	return fmt.Sprintf("LAST(%s)", sql.DebugString(a.Child))
}

func (a *Last) WithWindow(window *sql.WindowDefinition) sql.WindowAdaptableExpression {
	res := a.unaryAggBase.WithWindow(window)
	return &Last{unaryAggBase: *res.(*unaryAggBase)}
}

func (a *Last) WithChildren(children ...sql.Expression) (sql.Expression, error) {
	res, err := a.unaryAggBase.WithChildren(children...)
	return &Last{unaryAggBase: *res.(*unaryAggBase)}, err
}

func (a *Last) WithId(id sql.ColumnId) sql.IdExpression {
	res := a.unaryAggBase.WithId(id)
	return &Last{unaryAggBase: *res.(*unaryAggBase)}
}

func (a *Last) NewBuffer() (sql.AggregationBuffer, error) {
	child, err := transform.Clone(a.Child)
	if err != nil {
		return nil, err
	}
	return NewLastBuffer(child), nil
}

func (a *Last) NewWindowFunction() (sql.WindowFunction, error) {
	child, err := transform.Clone(a.Child)
	if err != nil {
		return nil, err
	}
	return NewLastAgg(child).WithWindow(a.Window())
}

type Max struct {
	unaryAggBase
}

var _ sql.FunctionExpression = (*Max)(nil)
var _ sql.Aggregation = (*Max)(nil)
var _ sql.WindowAdaptableExpression = (*Max)(nil)

func NewMax(e sql.Expression) *Max {
	return &Max{
		unaryAggBase{
			UnaryExpression: expression.UnaryExpression{Child: e},
			functionName:    "Max",
			description:     "returns the maximum value of expr in all rows.",
		},
	}
}

func (a *Max) Type() sql.Type {
	return a.Child.Type()
}

func (a *Max) IsNullable() bool {
	return false
}

func (a *Max) String() string {
	if a.window != nil {
		pr := sql.NewTreePrinter()
		_ = pr.WriteNode("MAX")
		children := []string{a.window.String(), a.Child.String()}
		pr.WriteChildren(children...)
		return pr.String()
	}
	return fmt.Sprintf("MAX(%s)", a.Child)
}

func (a *Max) DebugString() string {
	if a.window != nil {
		pr := sql.NewTreePrinter()
		_ = pr.WriteNode("MAX")
		children := []string{sql.DebugString(a.window), sql.DebugString(a.Child)}
		pr.WriteChildren(children...)
		return pr.String()
	}
	return fmt.Sprintf("MAX(%s)", sql.DebugString(a.Child))
}

func (a *Max) WithWindow(window *sql.WindowDefinition) sql.WindowAdaptableExpression {
	res := a.unaryAggBase.WithWindow(window)
	return &Max{unaryAggBase: *res.(*unaryAggBase)}
}

func (a *Max) WithChildren(children ...sql.Expression) (sql.Expression, error) {
	res, err := a.unaryAggBase.WithChildren(children...)
	return &Max{unaryAggBase: *res.(*unaryAggBase)}, err
}

func (a *Max) WithId(id sql.ColumnId) sql.IdExpression {
	res := a.unaryAggBase.WithId(id)
	return &Max{unaryAggBase: *res.(*unaryAggBase)}
}

func (a *Max) NewBuffer() (sql.AggregationBuffer, error) {
	child, err := transform.Clone(a.Child)
	if err != nil {
		return nil, err
	}
	return NewMaxBuffer(child), nil
}

func (a *Max) NewWindowFunction() (sql.WindowFunction, error) {
	child, err := transform.Clone(a.Child)
	if err != nil {
		return nil, err
	}
	return NewMaxAgg(child).WithWindow(a.Window())
}

type Min struct {
	unaryAggBase
}

var _ sql.FunctionExpression = (*Min)(nil)
var _ sql.Aggregation = (*Min)(nil)
var _ sql.WindowAdaptableExpression = (*Min)(nil)

func NewMin(e sql.Expression) *Min {
	return &Min{
		unaryAggBase{
			UnaryExpression: expression.UnaryExpression{Child: e},
			functionName:    "Min",
			description:     "returns the minimum value of expr in all rows.",
		},
	}
}

func (a *Min) Type() sql.Type {
	return a.Child.Type()
}

func (a *Min) IsNullable() bool {
	return false
}

func (a *Min) String() string {
	if a.window != nil {
		pr := sql.NewTreePrinter()
		_ = pr.WriteNode("MIN")
		children := []string{a.window.String(), a.Child.String()}
		pr.WriteChildren(children...)
		return pr.String()
	}
	return fmt.Sprintf("MIN(%s)", a.Child)
}

func (a *Min) DebugString() string {
	if a.window != nil {
		pr := sql.NewTreePrinter()
		_ = pr.WriteNode("MIN")
		children := []string{sql.DebugString(a.window), sql.DebugString(a.Child)}
		pr.WriteChildren(children...)
		return pr.String()
	}
	return fmt.Sprintf("MIN(%s)", sql.DebugString(a.Child))
}

func (a *Min) WithWindow(window *sql.WindowDefinition) sql.WindowAdaptableExpression {
	res := a.unaryAggBase.WithWindow(window)
	return &Min{unaryAggBase: *res.(*unaryAggBase)}
}

func (a *Min) WithChildren(children ...sql.Expression) (sql.Expression, error) {
	res, err := a.unaryAggBase.WithChildren(children...)
	return &Min{unaryAggBase: *res.(*unaryAggBase)}, err
}

func (a *Min) WithId(id sql.ColumnId) sql.IdExpression {
	res := a.unaryAggBase.WithId(id)
	return &Min{unaryAggBase: *res.(*unaryAggBase)}
}

func (a *Min) NewBuffer() (sql.AggregationBuffer, error) {
	child, err := transform.Clone(a.Child)
	if err != nil {
		return nil, err
	}
	return NewMinBuffer(child), nil
}

func (a *Min) NewWindowFunction() (sql.WindowFunction, error) {
	child, err := transform.Clone(a.Child)
	if err != nil {
		return nil, err
	}
	return NewMinAgg(child).WithWindow(a.Window())
}

type Sum struct {
	unaryAggBase
}

var _ sql.FunctionExpression = (*Sum)(nil)
var _ sql.Aggregation = (*Sum)(nil)
var _ sql.WindowAdaptableExpression = (*Sum)(nil)

func NewSum(e sql.Expression) *Sum {
	return &Sum{
		unaryAggBase{
			UnaryExpression: expression.UnaryExpression{Child: e},
			functionName:    "Sum",
			description:     "returns the sum of expr in all rows",
		},
	}
}

func (a *Sum) Type() sql.Type {
	return a.Child.Type()
}

func (a *Sum) IsNullable() bool {
	return false
}

func (a *Sum) String() string {
	if a.window != nil {
		pr := sql.NewTreePrinter()
		_ = pr.WriteNode("SUM")
		children := []string{a.window.String(), a.Child.String()}
		pr.WriteChildren(children...)
		return pr.String()
	}
	return fmt.Sprintf("SUM(%s)", a.Child)
}

func (a *Sum) DebugString() string {
	if a.window != nil {
		pr := sql.NewTreePrinter()
		_ = pr.WriteNode("SUM")
		children := []string{sql.DebugString(a.window), sql.DebugString(a.Child)}
		pr.WriteChildren(children...)
		return pr.String()
	}
	return fmt.Sprintf("SUM(%s)", sql.DebugString(a.Child))
}

func (a *Sum) WithWindow(window *sql.WindowDefinition) sql.WindowAdaptableExpression {
	res := a.unaryAggBase.WithWindow(window)
	return &Sum{unaryAggBase: *res.(*unaryAggBase)}
}

func (a *Sum) WithChildren(children ...sql.Expression) (sql.Expression, error) {
	res, err := a.unaryAggBase.WithChildren(children...)
	return &Sum{unaryAggBase: *res.(*unaryAggBase)}, err
}

func (a *Sum) WithId(id sql.ColumnId) sql.IdExpression {
	res := a.unaryAggBase.WithId(id)
	return &Sum{unaryAggBase: *res.(*unaryAggBase)}
}

func (a *Sum) NewBuffer() (sql.AggregationBuffer, error) {
	child, err := transform.Clone(a.Child)
	if err != nil {
		return nil, err
	}
	return NewSumBuffer(child), nil
}

func (a *Sum) NewWindowFunction() (sql.WindowFunction, error) {
	child, err := transform.Clone(a.Child)
	if err != nil {
		return nil, err
	}
	return NewSumAgg(child).WithWindow(a.Window())
}

type StdDevPop struct {
	unaryAggBase
}

var _ sql.FunctionExpression = (*StdDevPop)(nil)
var _ sql.Aggregation = (*StdDevPop)(nil)
var _ sql.WindowAdaptableExpression = (*StdDevPop)(nil)

func NewStdDevPop(e sql.Expression) *StdDevPop {
	return &StdDevPop{
		unaryAggBase{
			UnaryExpression: expression.UnaryExpression{Child: e},
			functionName:    "StdDevPop",
			description:     "returns the population standard deviation of expr",
		},
	}
}

func (a *StdDevPop) Type() sql.Type {
	return a.Child.Type()
}

func (a *StdDevPop) IsNullable() bool {
	return false
}

func (a *StdDevPop) String() string {
	if a.window != nil {
		pr := sql.NewTreePrinter()
		_ = pr.WriteNode("STDDEVPOP")
		children := []string{a.window.String(), a.Child.String()}
		pr.WriteChildren(children...)
		return pr.String()
	}
	return fmt.Sprintf("STDDEVPOP(%s)", a.Child)
}

func (a *StdDevPop) DebugString() string {
	if a.window != nil {
		pr := sql.NewTreePrinter()
		_ = pr.WriteNode("STDDEVPOP")
		children := []string{sql.DebugString(a.window), sql.DebugString(a.Child)}
		pr.WriteChildren(children...)
		return pr.String()
	}
	return fmt.Sprintf("STDDEVPOP(%s)", sql.DebugString(a.Child))
}

func (a *StdDevPop) WithWindow(window *sql.WindowDefinition) sql.WindowAdaptableExpression {
	res := a.unaryAggBase.WithWindow(window)
	return &StdDevPop{unaryAggBase: *res.(*unaryAggBase)}
}

func (a *StdDevPop) WithChildren(children ...sql.Expression) (sql.Expression, error) {
	res, err := a.unaryAggBase.WithChildren(children...)
	return &StdDevPop{unaryAggBase: *res.(*unaryAggBase)}, err
}

func (a *StdDevPop) WithId(id sql.ColumnId) sql.IdExpression {
	res := a.unaryAggBase.WithId(id)
	return &StdDevPop{unaryAggBase: *res.(*unaryAggBase)}
}

func (a *StdDevPop) NewBuffer() (sql.AggregationBuffer, error) {
	child, err := transform.Clone(a.Child)
	if err != nil {
		return nil, err
	}
	return NewStdDevPopBuffer(child), nil
}

func (a *StdDevPop) NewWindowFunction() (sql.WindowFunction, error) {
	child, err := transform.Clone(a.Child)
	if err != nil {
		return nil, err
	}
	return NewStdDevPopAgg(child).WithWindow(a.Window())
}

type StdDevSamp struct {
	unaryAggBase
}

var _ sql.FunctionExpression = (*StdDevSamp)(nil)
var _ sql.Aggregation = (*StdDevSamp)(nil)
var _ sql.WindowAdaptableExpression = (*StdDevSamp)(nil)

func NewStdDevSamp(e sql.Expression) *StdDevSamp {
	return &StdDevSamp{
		unaryAggBase{
			UnaryExpression: expression.UnaryExpression{Child: e},
			functionName:    "StdDevSamp",
			description:     "returns the sample standard deviation of expr",
		},
	}
}

func (a *StdDevSamp) Type() sql.Type {
	return a.Child.Type()
}

func (a *StdDevSamp) IsNullable() bool {
	return false
}

func (a *StdDevSamp) String() string {
	if a.window != nil {
		pr := sql.NewTreePrinter()
		_ = pr.WriteNode("STDDEVSAMP")
		children := []string{a.window.String(), a.Child.String()}
		pr.WriteChildren(children...)
		return pr.String()
	}
	return fmt.Sprintf("STDDEVSAMP(%s)", a.Child)
}

func (a *StdDevSamp) DebugString() string {
	if a.window != nil {
		pr := sql.NewTreePrinter()
		_ = pr.WriteNode("STDDEVSAMP")
		children := []string{sql.DebugString(a.window), sql.DebugString(a.Child)}
		pr.WriteChildren(children...)
		return pr.String()
	}
	return fmt.Sprintf("STDDEVSAMP(%s)", sql.DebugString(a.Child))
}

func (a *StdDevSamp) WithWindow(window *sql.WindowDefinition) sql.WindowAdaptableExpression {
	res := a.unaryAggBase.WithWindow(window)
	return &StdDevSamp{unaryAggBase: *res.(*unaryAggBase)}
}

func (a *StdDevSamp) WithChildren(children ...sql.Expression) (sql.Expression, error) {
	res, err := a.unaryAggBase.WithChildren(children...)
	return &StdDevSamp{unaryAggBase: *res.(*unaryAggBase)}, err
}

func (a *StdDevSamp) WithId(id sql.ColumnId) sql.IdExpression {
	res := a.unaryAggBase.WithId(id)
	return &StdDevSamp{unaryAggBase: *res.(*unaryAggBase)}
}

func (a *StdDevSamp) NewBuffer() (sql.AggregationBuffer, error) {
	child, err := transform.Clone(a.Child)
	if err != nil {
		return nil, err
	}
	return NewStdDevSampBuffer(child), nil
}

func (a *StdDevSamp) NewWindowFunction() (sql.WindowFunction, error) {
	child, err := transform.Clone(a.Child)
	if err != nil {
		return nil, err
	}
	return NewStdDevSampAgg(child).WithWindow(a.Window())
}

type VarPop struct {
	unaryAggBase
}

var _ sql.FunctionExpression = (*VarPop)(nil)
var _ sql.Aggregation = (*VarPop)(nil)
var _ sql.WindowAdaptableExpression = (*VarPop)(nil)

func NewVarPop(e sql.Expression) *VarPop {
	return &VarPop{
		unaryAggBase{
			UnaryExpression: expression.UnaryExpression{Child: e},
			functionName:    "VarPop",
			description:     "returns the population variance of expr",
		},
	}
}

func (a *VarPop) Type() sql.Type {
	return a.Child.Type()
}

func (a *VarPop) IsNullable() bool {
	return false
}

func (a *VarPop) String() string {
	if a.window != nil {
		pr := sql.NewTreePrinter()
		_ = pr.WriteNode("VARPOP")
		children := []string{a.window.String(), a.Child.String()}
		pr.WriteChildren(children...)
		return pr.String()
	}
	return fmt.Sprintf("VARPOP(%s)", a.Child)
}

func (a *VarPop) DebugString() string {
	if a.window != nil {
		pr := sql.NewTreePrinter()
		_ = pr.WriteNode("VARPOP")
		children := []string{sql.DebugString(a.window), sql.DebugString(a.Child)}
		pr.WriteChildren(children...)
		return pr.String()
	}
	return fmt.Sprintf("VARPOP(%s)", sql.DebugString(a.Child))
}

func (a *VarPop) WithWindow(window *sql.WindowDefinition) sql.WindowAdaptableExpression {
	res := a.unaryAggBase.WithWindow(window)
	return &VarPop{unaryAggBase: *res.(*unaryAggBase)}
}

func (a *VarPop) WithChildren(children ...sql.Expression) (sql.Expression, error) {
	res, err := a.unaryAggBase.WithChildren(children...)
	return &VarPop{unaryAggBase: *res.(*unaryAggBase)}, err
}

func (a *VarPop) WithId(id sql.ColumnId) sql.IdExpression {
	res := a.unaryAggBase.WithId(id)
	return &VarPop{unaryAggBase: *res.(*unaryAggBase)}
}

func (a *VarPop) NewBuffer() (sql.AggregationBuffer, error) {
	child, err := transform.Clone(a.Child)
	if err != nil {
		return nil, err
	}
	return NewVarPopBuffer(child), nil
}

func (a *VarPop) NewWindowFunction() (sql.WindowFunction, error) {
	child, err := transform.Clone(a.Child)
	if err != nil {
		return nil, err
	}
	return NewVarPopAgg(child).WithWindow(a.Window())
}

type VarSamp struct {
	unaryAggBase
}

var _ sql.FunctionExpression = (*VarSamp)(nil)
var _ sql.Aggregation = (*VarSamp)(nil)
var _ sql.WindowAdaptableExpression = (*VarSamp)(nil)

func NewVarSamp(e sql.Expression) *VarSamp {
	return &VarSamp{
		unaryAggBase{
			UnaryExpression: expression.UnaryExpression{Child: e},
			functionName:    "VarSamp",
			description:     "returns the sample variance of expr",
		},
	}
}

func (a *VarSamp) Type() sql.Type {
	return a.Child.Type()
}

func (a *VarSamp) IsNullable() bool {
	return false
}

func (a *VarSamp) String() string {
	if a.window != nil {
		pr := sql.NewTreePrinter()
		_ = pr.WriteNode("VARSAMP")
		children := []string{a.window.String(), a.Child.String()}
		pr.WriteChildren(children...)
		return pr.String()
	}
	return fmt.Sprintf("VARSAMP(%s)", a.Child)
}

func (a *VarSamp) DebugString() string {
	if a.window != nil {
		pr := sql.NewTreePrinter()
		_ = pr.WriteNode("VARSAMP")
		children := []string{sql.DebugString(a.window), sql.DebugString(a.Child)}
		pr.WriteChildren(children...)
		return pr.String()
	}
	return fmt.Sprintf("VARSAMP(%s)", sql.DebugString(a.Child))
}

func (a *VarSamp) WithWindow(window *sql.WindowDefinition) sql.WindowAdaptableExpression {
	res := a.unaryAggBase.WithWindow(window)
	return &VarSamp{unaryAggBase: *res.(*unaryAggBase)}
}

func (a *VarSamp) WithChildren(children ...sql.Expression) (sql.Expression, error) {
	res, err := a.unaryAggBase.WithChildren(children...)
	return &VarSamp{unaryAggBase: *res.(*unaryAggBase)}, err
}

func (a *VarSamp) WithId(id sql.ColumnId) sql.IdExpression {
	res := a.unaryAggBase.WithId(id)
	return &VarSamp{unaryAggBase: *res.(*unaryAggBase)}
}

func (a *VarSamp) NewBuffer() (sql.AggregationBuffer, error) {
	child, err := transform.Clone(a.Child)
	if err != nil {
		return nil, err
	}
	return NewVarSampBuffer(child), nil
}

func (a *VarSamp) NewWindowFunction() (sql.WindowFunction, error) {
	child, err := transform.Clone(a.Child)
	if err != nil {
		return nil, err
	}
	return NewVarSampAgg(child).WithWindow(a.Window())
}
