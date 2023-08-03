package types

import (
	"database/sql/driver"
	"fmt"
	"time"
)

type Date struct {
	time.Time
}

func (d *Date) UnmarshalJSON(b []byte) (err error) {
	if b[0] == '"' && b[len(b)-1] == '"' {
		b = b[1 : len(b)-1]
	}

	// take care of null..
	if len(b) == 0 || string(b) == "null" {
		d.Time = time.Time{}
		return
	}

	d.Time, err = time.Parse("2006-01-02", string(b))
	return
}

func (d *Date) Scan(b interface{}) (err error) {
	switch x := b.(type) {
	case time.Time:
		d.Time = x
	default:
		err = fmt.Errorf("unsupported scan type %T", b)
	}
	return
}

func (d Date) Value() (driver.Value, error) {
	// check if the date was not set..
	if d.Time.IsZero() {
		return nil, nil
	}
	return d.Time.Format("2006-01-02"), nil
}
