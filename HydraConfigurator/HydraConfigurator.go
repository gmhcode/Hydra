package HydraConfigurator

import (
	"errors"
	"reflect"
)

const (
	CUSTOM uint8 = iota
)

var wrongTypeError error = errors.New("Type must be a pointer to a struct")

// GetConfiguration gets the configuration from the config file
func GetConfiguration(confType uint8, obj interface{}, filename string) (err error) {

	mysRValue := reflect.ValueOf(obj)
	// mysRValue = &{ false 0 0}

	if mysRValue.Kind() != reflect.Ptr || mysRValue.IsNil() {
		return wrongTypeError
	}

	//getting rid of the pointer here
	mysRValue = mysRValue.Elem()
	// mysRValue = { false 0 0}

	// *object => object
	//reflection value of *object.Elem() => object() (Settable!!)
	if mysRValue.Kind() != reflect.Struct {
		return wrongTypeError
	}

	switch confType {
	case CUSTOM:
		err = MarshalCustomConfig(mysRValue, filename)
	}
	return err
}
