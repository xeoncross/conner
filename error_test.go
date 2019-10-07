package conner

import (
	"encoding/json"
	"errors"
	"fmt"
	"testing"
)

func TestValues(t *testing.T) {

	//
	// Each function can add structured context values to the error
	// without any lose of functionality
	//

	f3 := func() error {
		return errors.New("F3 Error")
	}

	f2 := func() error {
		err := f3()
		if err != nil {
			return Error(fmt.Errorf("F2: %w", err), map[string]interface{}{"a": "a"})
		}
		return nil
	}

	f1 := func() error {
		err := f2()
		if err != nil {
			return Error(fmt.Errorf("F1: %w", err), map[string]interface{}{"b": "b"})
		}
		return nil
	}

	err := f1()

	b, err := json.Marshal(Values(err))
	if err != nil {
		t.Fatal(err)
	}

	result := `{"a":"a","b":"b"}`

	if string(b) != result {
		t.Errorf("%s != %s", string(b), result)
	}

}
