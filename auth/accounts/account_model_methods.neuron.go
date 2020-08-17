// Code generated by neuron/generator. DO NOT EDIT.
// This file was generated at:
// Sun, 16 Aug 2020 23:26:13 +0200

package account

import (
	"time"

	"github.com/google/uuid"
	"github.com/neuronlabs/neuron/errors"
	"github.com/neuronlabs/neuron/mapping"
	"github.com/neuronlabs/neuron/query"
)

// Compile time check if Account implements mapping.Model interface.
var _ mapping.Model = &Account{}

// NeuronCollectionName implements mapping.Model interface method.
// Returns the name of the collection for the 'Account'.
func (a *Account) NeuronCollectionName() string {
	return "accounts"
}

// IsPrimaryKeyZero implements mapping.Model interface method.
func (a *Account) IsPrimaryKeyZero() bool {
	return a.ID == uuid.UUID([16]byte{})
}

// GetPrimaryKeyValue implements mapping.Model interface method.
func (a *Account) GetPrimaryKeyValue() interface{} {
	return a.ID
}

// GetPrimaryKeyStringValue implements mapping.Model interface method.
func (a *Account) GetPrimaryKeyStringValue() (string, error) {
	id, err := a.ID.MarshalText()
	if err != nil {
		return "", errors.Wrapf(mapping.ErrFieldValue, "invalid primary field value: %v to parse string. Err: %v", a.ID, err)
	}
	return string(id), nil
}

// GetPrimaryKeyAddress implements mapping.Model interface method.
func (a *Account) GetPrimaryKeyAddress() interface{} {
	return &a.ID
}

// GetPrimaryKeyHashableValue implements mapping.Model interface method.
func (a *Account) GetPrimaryKeyHashableValue() interface{} {
	return a.ID
}

// GetPrimaryKeyZeroValue implements mapping.Model interface method.
func (a *Account) GetPrimaryKeyZeroValue() interface{} {
	return uuid.UUID([16]byte{})
}

// SetPrimaryKey implements mapping.Model interface method.
func (a *Account) SetPrimaryKeyValue(value interface{}) error {
	if v, ok := value.(uuid.UUID); ok {
		a.ID = v
		return nil
	} else if v, ok := value.([16]byte); ok {
		a.ID = uuid.UUID(v)
	}
	return errors.Wrapf(mapping.ErrFieldValue, "provided invalid value: '%T' for the primary field for model: '%T'",
		value, a)
}

// SetPrimaryKeyStringValue implements mapping.Model interface method.
func (a *Account) SetPrimaryKeyStringValue(value string) error {
	if err := a.ID.UnmarshalText([]byte(value)); err != nil {
		return errors.Wrapf(mapping.ErrFieldValue, "invalid primary field value: %v to parse string. Err: %v", a.ID, err)
	}
	return nil
}

// SetFrom implements FromSetter interface.
func (a *Account) SetFrom(model mapping.Model) error {
	if model == nil {
		return errors.Wrap(query.ErrInvalidInput, "provided nil model to set from")
	}
	from, ok := model.(*Account)
	if !ok {
		return errors.WrapDetf(mapping.ErrModelNotMatch, "provided model doesn't match the input: %T", model)
	}
	*a = *from
	return nil
}

// Compile time check if Account implements mapping.Fielder interface.
var _ mapping.Fielder = &Account{}

// GetFieldsAddress gets the address of provided 'field'.
func (a *Account) GetFieldsAddress(field *mapping.StructField) (interface{}, error) {
	switch field.Index[0] {
	case 0: // ID
		return &a.ID, nil
	case 1: // CreatedAt
		return &a.CreatedAt, nil
	case 2: // UpdatedAt
		return &a.UpdatedAt, nil
	case 3: // DeletedAt
		return &a.DeletedAt, nil
	case 4: // Username
		return &a.Username, nil
	case 5: // PasswordHash
		return &a.PasswordHash, nil
	case 6: // PasswordSalt
		return &a.PasswordSalt, nil
	}
	return nil, errors.Wrapf(mapping.ErrInvalidModelField, "provided invalid field: '%s' for given model: Account'", field.Name())
}

// GetFieldZeroValue implements mapping.Fielder interface.s
func (a *Account) GetFieldZeroValue(field *mapping.StructField) (interface{}, error) {
	switch field.Index[0] {
	case 0: // ID
		return uuid.UUID([16]byte{}), nil
	case 1: // CreatedAt
		return time.Time{}, nil
	case 2: // UpdatedAt
		return time.Time{}, nil
	case 3: // DeletedAt
		return nil, nil
	case 4: // Username
		return "", nil
	case 5: // PasswordHash
		return nil, nil
	case 6: // PasswordSalt
		return nil, nil
	default:
		return nil, errors.Wrapf(mapping.ErrInvalidModelField, "provided invalid field name: '%s'", field.Name())
	}
}

// IsFieldZero implements mapping.Fielder interface.
func (a *Account) IsFieldZero(field *mapping.StructField) (bool, error) {
	switch field.Index[0] {
	case 0: // ID
		return a.ID == uuid.UUID([16]byte{}), nil
	case 1: // CreatedAt
		return a.CreatedAt == time.Time{}, nil
	case 2: // UpdatedAt
		return a.UpdatedAt == time.Time{}, nil
	case 3: // DeletedAt
		return a.DeletedAt == nil, nil
	case 4: // Username
		return a.Username == "", nil
	case 5: // PasswordHash
		return len(a.PasswordHash) == 0, nil
	case 6: // PasswordSalt
		return len(a.PasswordSalt) == 0, nil
	}
	return false, errors.Wrapf(mapping.ErrInvalidModelField, "provided invalid field name: '%s'", field.Name())
}

// SetFieldZeroValue implements mapping.Fielder interface.s
func (a *Account) SetFieldZeroValue(field *mapping.StructField) error {
	switch field.Index[0] {
	case 0: // ID
		a.ID = uuid.UUID([16]byte{})
	case 1: // CreatedAt
		a.CreatedAt = time.Time{}
	case 2: // UpdatedAt
		a.UpdatedAt = time.Time{}
	case 3: // DeletedAt
		a.DeletedAt = nil
	case 4: // Username
		a.Username = ""
	case 5: // PasswordHash
		a.PasswordHash = nil
	case 6: // PasswordSalt
		a.PasswordSalt = nil
	default:
		return errors.Wrapf(mapping.ErrInvalidModelField, "provided invalid field name: '%s'", field.Name())
	}
	return nil
}

// GetHashableFieldValue implements mapping.Fielder interface.
func (a *Account) GetHashableFieldValue(field *mapping.StructField) (interface{}, error) {
	switch field.Index[0] {
	case 0: // ID
		return a.ID, nil
	case 1: // CreatedAt
		return a.CreatedAt, nil
	case 2: // UpdatedAt
		return a.UpdatedAt, nil
	case 3: // DeletedAt
		if a.DeletedAt == nil {
			return nil, nil
		}
		return *a.DeletedAt, nil
	case 4: // Username
		return a.Username, nil
	case 5: // PasswordHash
		return string(a.PasswordHash), nil
	case 6: // PasswordSalt
		return string(a.PasswordSalt), nil
	}
	return nil, errors.Wrapf(mapping.ErrInvalidModelField, "provided invalid field: '%s' for given model: 'Account'", field.Name())
}

// GetFieldValue implements mapping.Fielder interface.
func (a *Account) GetFieldValue(field *mapping.StructField) (interface{}, error) {
	switch field.Index[0] {
	case 0: // ID
		return a.ID, nil
	case 1: // CreatedAt
		return a.CreatedAt, nil
	case 2: // UpdatedAt
		return a.UpdatedAt, nil
	case 3: // DeletedAt
		return a.DeletedAt, nil
	case 4: // Username
		return a.Username, nil
	case 5: // PasswordHash
		return a.PasswordHash, nil
	case 6: // PasswordSalt
		return a.PasswordSalt, nil
	}
	return nil, errors.Wrapf(mapping.ErrInvalidModelField, "provided invalid field: '%s' for given model: Account'", field.Name())
}

// SetFieldValue implements mapping.Fielder interface.
func (a *Account) SetFieldValue(field *mapping.StructField, value interface{}) (err error) {
	switch field.Index[0] {
	case 0: // ID
		if v, ok := value.(uuid.UUID); ok {
			a.ID = v
			return nil
		}
		if generic, ok := value.([]interface{}); ok {
			if len(generic) > 16 {
				return errors.Wrapf(mapping.ErrFieldValue, "provided too many values for the field: 'ID")
			}
			for i, item := range generic {
				if v, ok := item.(byte); ok {
					a.ID[i] = v
					continue
				}

			}
			return nil
		}
		// Checked wrapped types.
		if v, ok := value.([16]byte); ok {
			a.ID = uuid.UUID(v)
			return nil
		}
		return errors.Wrapf(mapping.ErrFieldValue, "provided invalid field type: '%T' for the field: %s", value, field.Name())
	case 1: // CreatedAt
		if v, ok := value.(time.Time); ok {
			a.CreatedAt = v
			return nil
		}

		return errors.Wrapf(mapping.ErrFieldValue, "provided invalid field type: '%T' for the field: %s", value, field.Name())
	case 2: // UpdatedAt
		if v, ok := value.(time.Time); ok {
			a.UpdatedAt = v
			return nil
		}

		return errors.Wrapf(mapping.ErrFieldValue, "provided invalid field type: '%T' for the field: %s", value, field.Name())
	case 3: // DeletedAt
		if value == nil {
			a.DeletedAt = nil
			return nil
		}
		if v, ok := value.(*time.Time); ok {
			a.DeletedAt = v
			return nil
		}
		// Check if it is non-pointer value.
		if v, ok := value.(time.Time); ok {
			a.DeletedAt = &v
			return nil
		}

		return errors.Wrapf(mapping.ErrFieldValue, "provided invalid field type: '%T' for the field: %s", value, field.Name())
	case 4: // Username
		if v, ok := value.(string); ok {
			a.Username = v
			return nil
		}

		// Check alternate types for the Username.
		if v, ok := value.([]byte); ok {
			a.Username = string(v)
			return nil
		}
		return errors.Wrapf(mapping.ErrFieldValue, "provided invalid field type: '%T' for the field: %s", value, field.Name())
	case 5: // PasswordHash
		if value == nil {
			a.PasswordHash = nil
			return nil
		}
		if v, ok := value.([]byte); ok {
			a.PasswordHash = v
			return nil
		}
		if generic, ok := value.([]interface{}); ok {
			for _, item := range generic {
				if v, ok := item.(byte); ok {
					a.PasswordHash = append(a.PasswordHash, v)
					continue
				}

			}
			return nil
		}
		// Check alternate types for the PasswordHash.
		if v, ok := value.(string); ok {
			a.PasswordHash = []byte(v)
			return nil
		}
		return errors.Wrapf(mapping.ErrFieldValue, "provided invalid field type: '%T' for the field: %s", value, field.Name())
	case 6: // PasswordSalt
		if value == nil {
			a.PasswordSalt = nil
			return nil
		}
		if v, ok := value.([]byte); ok {
			a.PasswordSalt = v
			return nil
		}
		if generic, ok := value.([]interface{}); ok {
			for _, item := range generic {
				if v, ok := item.(byte); ok {
					a.PasswordSalt = append(a.PasswordSalt, v)
					continue
				}

			}
			return nil
		}
		// Check alternate types for the PasswordSalt.
		if v, ok := value.(string); ok {
			a.PasswordSalt = []byte(v)
			return nil
		}
		return errors.Wrapf(mapping.ErrFieldValue, "provided invalid field type: '%T' for the field: %s", value, field.Name())
	default:
		return errors.Wrapf(mapping.ErrInvalidModelField, "provided invalid field: '%s' for the model: 'Account'", field.Name())
	}
}

// SetPrimaryKeyStringValue implements mapping.Model interface method.
func (a *Account) ParseFieldsStringValue(field *mapping.StructField, value string) (interface{}, error) {
	switch field.Index[0] {
	case 0: // ID
		temp := a.ID
		if err := a.ID.UnmarshalText([]byte(value)); err != nil {
			return "", errors.Wrapf(mapping.ErrFieldValue, "invalid field 'ID' value: '%v' to parse string. Err: %v", a.ID, err)
		}
		bt, err := a.ID.MarshalText()
		if err != nil {
			return "", errors.Wrapf(mapping.ErrFieldValue, "invalid field 'ID' value: '%v' to parse string. Err: %v", a.ID, err)
		}
		a.ID = temp
		return string(bt), nil
	case 1: // CreatedAt
		temp := a.CreatedAt
		if err := a.CreatedAt.UnmarshalText([]byte(value)); err != nil {
			return "", errors.Wrapf(mapping.ErrFieldValue, "invalid field 'CreatedAt' value: '%v' to parse string. Err: %v", a.CreatedAt, err)
		}
		bt, err := a.CreatedAt.MarshalText()
		if err != nil {
			return "", errors.Wrapf(mapping.ErrFieldValue, "invalid field 'CreatedAt' value: '%v' to parse string. Err: %v", a.CreatedAt, err)
		}
		a.CreatedAt = temp
		return string(bt), nil
	case 2: // UpdatedAt
		temp := a.UpdatedAt
		if err := a.UpdatedAt.UnmarshalText([]byte(value)); err != nil {
			return "", errors.Wrapf(mapping.ErrFieldValue, "invalid field 'UpdatedAt' value: '%v' to parse string. Err: %v", a.UpdatedAt, err)
		}
		bt, err := a.UpdatedAt.MarshalText()
		if err != nil {
			return "", errors.Wrapf(mapping.ErrFieldValue, "invalid field 'UpdatedAt' value: '%v' to parse string. Err: %v", a.UpdatedAt, err)
		}
		a.UpdatedAt = temp
		return string(bt), nil
	case 3: // DeletedAt
		var base time.Time
		temp := &base
		if err := temp.UnmarshalText([]byte(value)); err != nil {
			return "", errors.Wrapf(mapping.ErrFieldValue, "invalid field 'DeletedAt' value: '%v' to parse string. Err: %v", a.DeletedAt, err)
		}
		bt, err := temp.MarshalText()
		if err != nil {
			return "", errors.Wrapf(mapping.ErrFieldValue, "invalid field 'DeletedAt' value: '%v' to parse string. Err: %v", a.DeletedAt, err)
		}

		return string(bt), nil
	case 4: // Username
		return value, nil
	case 5: // PasswordHash
		return "", errors.Wrap(mapping.ErrFieldNotParser, "field 'PasswordHash' doesn't have string setter.")
	case 6: // PasswordSalt
		return "", errors.Wrap(mapping.ErrFieldNotParser, "field 'PasswordSalt' doesn't have string setter.")
	}
	return nil, errors.Wrapf(mapping.ErrInvalidModelField, "provided invalid field: '%s' for given model: Account'", field.Name())
}