package storage

// A Query defines a set of conditions, pagination, and other options used to
// fetch data from a datastore.
type Query interface {
	Conditions() []Condition
	Range() (offset, limit int)
	Count() bool
}

// A QueryBuilder provides a simple means to additively construct new Queries.
// When and where possible, implementors should attempt to validate the query
// against a schema, i.e., if the query uses nonexistant properties or values
// which do not match the underlying type, an error should be returned from the
// Build() method.
type QueryBuilder interface {
	ConditionBuilder
	Range(offset, limit int) QueryBuilder
	Count(bool) QueryBuilder
	Build() (q Query, err error)
}

// A ConditionBuilder provides a simple means of building complex value
// assertions that limit a query result.
type ConditionBuilder interface {
	Condition(property string, value interface{}, operator Operator) ConditionBuilder
	ConditionGroup(conjunction Conjunction) ConditionBuilder
	Negate(bool) ConditionBuilder
}

// Condition records an assertion to which fetched data must adhere. A
// condition may contain child conditions and a conjunction to create nested
// queries. Implementors should always default to the 'And' conjunction.
type Condition interface {
	Conditions() []Condition
	Conjunction() string
	Assertion() Assertion
	IsGroup() bool
}

// An Assertion is the basic value of a condition. Unless a condition contains
// a group of conditions, it must provide an assertion.
type Assertion struct {
	Property string
	Value    interface{}
	Operator Operator
	Negate   bool
}

// A conjunction defines the method for evaluation groups of conditions.
type Conjunction int

const (
	And Conjunction = iota
	Or
)

// An operator defines the method of comparison for an Assertion.
type Operator int

const (
	Eq Operator = iota
	LtEq
	GtEq
	Lt
	Gt
	Contains
	Exists
)
