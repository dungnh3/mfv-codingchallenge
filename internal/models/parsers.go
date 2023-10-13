package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
)

func (a *AccountIDs) Value() (driver.Value, error) {
	return json.Marshal(a)
}

func (a *AccountIDs) Scan(value interface{}) (err error) {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("failed to unmarshal AccountIDs value: ", value))
	}
	return json.Unmarshal(bytes, &a)
}
