package common

import (
	"database/sql/driver"
	"encoding/json"
	"reflect"

	"github.com/google/uuid"
)

type NullUUID struct {
	UUID  uuid.UUID
	Valid bool
}

func (nd *NullUUID) Scan(value interface{}) (err error) {
	var s uuid.UUID
	if err := s.Scan(value); err != nil {
		return err
	}

	// if nil then make Valid false
	if reflect.TypeOf(value) == nil {
		*nd = NullUUID{Valid: false}
	} else {
		if err != nil {
			return err
		}
		*nd = NullUUID{s, true}
	}

	return nil
}

func (nd NullUUID) Value() (driver.Value, error) {
	if !nd.Valid {
		return nil, nil
	}

	return nd.Value, nil
}

func (nd NullUUID) MarshalJSON() ([]byte, error) {
	if !nd.Valid {
		return []byte("null"), nil
	}

	return json.Marshal(nd.UUID.String())
}

func (nd *NullUUID) UnmarshalJSON(b []byte) error {
	var str string

	err := json.Unmarshal(b, &str)
	if err != nil {
		nd.Valid = false
		return err
	}

	id, err := uuid.Parse(str)
	if err != nil {
		nd.Valid = false
		return err
	}

	nd.UUID = id
	nd.Valid = true
	return err
}
