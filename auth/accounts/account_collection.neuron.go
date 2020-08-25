// Code generated by neurogonesis. DO NOT EDIT.
// This file was generated at:
// Tue, 25 Aug 2020 13:58:20 +0200

package accounts

import (
	"context"

	"github.com/neuronlabs/neuron/core"
	"github.com/neuronlabs/neuron/database"
	"github.com/neuronlabs/neuron/errors"
	"github.com/neuronlabs/neuron/mapping"
	"github.com/neuronlabs/neuron/query"
	"github.com/neuronlabs/neuron/query/filter"
)

// NRN_Accounts is the collection used to query Account model.
var NRN_Accounts = &_Accounts{}

type _Accounts struct {
	Model *mapping.ModelStruct
}

// InitializeCollection implements core.Initializer interface.
func (a *_Accounts) Initialize(_ctrl *core.Controller) error {
	mStruct, err := _ctrl.ModelStruct(&Account{})
	if err != nil {
		return err
	}
	NRN_Accounts = &_Accounts{Model: mStruct}
	return nil
}

// Query creates the query for the Account.
func (a *_Accounts) Query(db database.DB, models ...*Account) *accountsQueryBuilder {
	var queryModels []mapping.Model
	if len(models) > 0 {
		queryModels = make([]mapping.Model, len(models))
		for i, model := range models {
			queryModels[i] = model
		}
	}
	builder := db.Query(a.Model, queryModels...)
	return &accountsQueryBuilder{builder: builder}
}

// QueryCtx creates the query for the Account with provided 'ctx' context.
func (a *_Accounts) QueryCtx(ctx context.Context, db database.DB, models ...*Account) *accountsQueryBuilder {
	var queryModels []mapping.Model
	if len(models) > 0 {
		queryModels = make([]mapping.Model, len(models))
		for i, model := range models {
			queryModels[i] = model
		}
	}
	builder := db.QueryCtx(ctx, a.Model, queryModels...)
	return &accountsQueryBuilder{builder: builder}
}

// Insert inserts Account into database.
func (a *_Accounts) Insert(ctx context.Context, db database.DB, models ...*Account) error {
	if len(models) == 0 {
		return errors.Wrap(query.ErrNoModels, "nothing to insert")
	}
	queryModels := make([]mapping.Model, len(models))
	for i, model := range models {
		queryModels[i] = model
	}
	return db.Insert(ctx, a.Model, queryModels...)
}

// Update updates Account models into database.
func (a *_Accounts) Update(ctx context.Context, db database.DB, models ...*Account) (int64, error) {
	if len(models) == 0 {
		return 0, errors.Wrap(query.ErrNoModels, "nothing to insert")
	}
	queryModels := make([]mapping.Model, len(models))
	for i, model := range models {
		queryModels[i] = model
	}
	return db.Update(ctx, a.Model, queryModels...)
}

// Delete deletes Account models in database.
func (a *_Accounts) Delete(ctx context.Context, db database.DB, models ...*Account) (int64, error) {
	if len(models) == 0 {
		return 0, errors.Wrap(query.ErrNoModels, "nothing to insert")
	}
	queryModels := make([]mapping.Model, len(models))
	for i, model := range models {
		queryModels[i] = model
	}
	return db.Delete(ctx, a.Model, queryModels...)
}

// Refresh creates the query for the Account with provided 'ctx' context.
func (a *_Accounts) Refresh(ctx context.Context, db database.DB, models ...*Account) error {
	var queryModels []mapping.Model
	if len(models) == 0 {
		return errors.Wrap(query.ErrNoModels, "nothing to refresh")
	}
	queryModels = make([]mapping.Model, len(models))
	for i, model := range models {
		queryModels[i] = model
	}
	return db.Refresh(ctx, a.Model, queryModels...)
}

// accountsQueryBuilder is the query builder used to create and execute
// queries for the Accountmodel.
type accountsQueryBuilder struct {
	builder database.Builder
	err     error
}

// Scope returns given query scope.
func (a *accountsQueryBuilder) Scope() *query.Scope {
	return a.builder.Scope()
}

// Err returns errors that occurred during query building process.
func (a *accountsQueryBuilder) Err() error {
	if a.err != nil {
		return a.err
	}
	return a.builder.Err()
}

// Ctx returns the context of given query builder.
func (a *accountsQueryBuilder) Ctx() context.Context {
	return a.builder.Ctx()
}

// Count returns the number of model instances for provided query.
func (a *accountsQueryBuilder) Count() (int64, error) {
	if a.err != nil {
		return 0, a.err
	}
	return a.builder.Count()
}

// Insert new 'Account' instance(s) into the store.
func (a *accountsQueryBuilder) Insert() error {
	if a.err != nil {
		return a.err
	}
	return a.builder.Insert()
}

// Update updates given 'Account' instances.
func (a *accountsQueryBuilder) Update() (int64, error) {
	if a.err != nil {
		return 0, a.err
	}
	return a.builder.Update()
}

// Find returns all Account models that matches to given query.
func (a *accountsQueryBuilder) Find() ([]*Account, error) {
	if a.err != nil {
		return nil, a.err
	}
	queryModels, err := a.builder.Find()
	if err != nil {
		return nil, err
	}
	models := make([]*Account, len(queryModels))
	for i := range queryModels {
		models[i] = queryModels[i].(*Account)
	}
	return models, nil
}

// Refresh refreshes input 'Account' model fields. It might be combine with the included relations.
func (a *accountsQueryBuilder) Refresh() error {
	if a.err != nil {
		return a.err
	}
	return a.builder.Refresh()
}

// Get returns single Account model that matches given query.
// If the model is not found the function returns error of query.ErrQueryNoResult.
func (a *accountsQueryBuilder) Get() (*Account, error) {
	if a.err != nil {
		return nil, a.err
	}
	model, err := a.builder.Get()
	if err != nil {
		return nil, err
	}
	return model.(*Account), nil
}

// Delete deletes Account instances that matches given query.
func (a *accountsQueryBuilder) Delete() (int64, error) {
	if a.err != nil {
		return 0, a.err
	}
	return a.builder.Delete()
}

// Filter adds the 'filter' to the given query.
func (a *accountsQueryBuilder) Filter(filter filter.Filter) *accountsQueryBuilder {
	if a.err != nil {
		return a
	}
	a.builder.Filter(filter)
	return a
}

// Where creates query with given 'filter' and 'values'.
func (a *accountsQueryBuilder) Where(filter string, values ...interface{}) *accountsQueryBuilder {
	if a.err != nil {
		return a
	}
	a.builder.Where(filter, values...)
	return a
}

// Limit sets the maximum number of objects returned by the Find process,
// Returns error if the given scope has already different type of pagination.
func (a *accountsQueryBuilder) Limit(limit int64) *accountsQueryBuilder {
	if a.err != nil {
		return a
	}
	a.builder.Limit(limit)
	return a
}

// Offset sets the query result's offset. It says to skip as many object's from the repository
// before beginning to return the result. 'Offset' 0 is the same as omitting the 'Offset' clause.
// Returns error if the given scope has already different type of pagination.
func (a *accountsQueryBuilder) Offset(offset int64) *accountsQueryBuilder {
	if a.err != nil {
		return a
	}
	a.builder.Offset(offset)
	return a
}

// Select adds the fields to the scope's fieldset.
// Allowed fields to select:
//  - ID / id
//  - ID / id
//  - CreatedAt / created_at
//  - UpdatedAt / updated_at
//  - DeletedAt / deleted_at
//  - Username / username
//  - PasswordHash / password_hash
//  - PasswordSalt / password_salt
func (a *accountsQueryBuilder) Select(fields ...string) *accountsQueryBuilder {
	if a.err != nil {
		return a
	}
	var fieldSet []*mapping.StructField
	for _, field := range fields {
		structField, ok := NRN_Accounts.Model.FieldByName(field)
		if !ok {
			a.err = errors.Wrapf(mapping.ErrInvalidModelField, "field: '%s' is not valid for model: '_Accounts'", field)
			return a
		}
		fieldSet = append(fieldSet, structField)
	}
	a.builder.Select(fieldSet...)
	return a
}

// OrderBy adds the sort fields into query scope. By default field is ordered ascending. In order to sort descending
// add '-' before the field name i.e. '-id'. The order of the fields relates to the significance of the sorting order.
// Allowed fields to sort:
//  - ID
//  - ID
//  - CreatedAt
//  - UpdatedAt
//  - DeletedAt
//  - Username
//  - PasswordHash
//  - PasswordSalt
func (a *accountsQueryBuilder) OrderBy(fields ...string) *accountsQueryBuilder {
	if a.err != nil {
		return a
	}
	sortFields := make([]query.Sort, len(fields))
	for i, field := range fields {
		if len(field) == 0 {
			a.err = errors.Wrap(mapping.ErrInvalidModelField, "cannot set sorting order for an empty field for model: '_Accounts'")
			return a
		}
		var order query.SortOrder
		if field[0] == '-' {
			order = query.DescendingOrder
			field = field[1:]
		}
		structField, ok := NRN_Accounts.Model.FieldByName(field)
		if !ok {
			a.err = errors.Wrapf(mapping.ErrInvalidModelField, "field: '%s' is not valid for model: '_Accounts'", field)
			return a
		}
		sortFields[i] = query.SortField{StructField: structField, SortOrder: order}
	}
	a.builder.OrderBy(sortFields...)
	return a
}
