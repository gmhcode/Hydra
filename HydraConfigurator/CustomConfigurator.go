package HydraConfigurator

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"
)

//ConfigFields - dictionary which holds configuration fields
type ConfigFields map[string]reflect.Value

// Add - Adds the ConfigFields Dictionary
func (f ConfigFields) Add(name, value, typ string) error {
	switch typ {
	case "STRING":
		f[name] = reflect.ValueOf(value)
		// value = hello world

	case "INTEGER":
		i, err := strconv.Atoi(value)
		// value = 4
		if err != nil {
			return err
		}
		f[name] = reflect.ValueOf(i)
	case "FLOAT":
		// value = 3.5
		fl, err := strconv.ParseFloat(value, 64)
		if err != nil {
			return err
		}
		f[name] = reflect.ValueOf(fl)
	case "BOOL":
		b, err := strconv.ParseBool(value)
		// value = true
		if err != nil {
			return err
		}
		f[name] = reflect.ValueOf(b)
	}
	return nil
}

// MarshalCustomConfig - reads from file and creates structs
func MarshalCustomConfig(value reflect.Value, filename string) error {

	// filename = configfile.conf
	// value =  { false 0 0}

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Panic occured", r)
		}
	}()

	if !value.CanSet() {
		return errors.New("Value passed not settable")
	}
	file, err := os.Open(filename)

	if err != nil {
		return err
	}
	defer file.Close()
	fields := make(ConfigFields) //make(map[string]reflect.Value)
	//fields = map[]

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		fmt.Println("\nProcessing Line", line)
		// prints
		// Processing Line testString|hello world;string
		// Processing Line testBool|true;bool
		// Processing Line testFloat|3.5;float
		// Processing Line TestInt|4;INTEGER

		if strings.Count(line, "|") != 1 || strings.Count(line, ";") != 1 {
			continue
		}
		args := strings.Split(line, "|")
		// args =
		// [testString hello world;string]
		// [testBool true;bool]
		// [testFloat 3.5;float]
		// [TestInt 4;INTEGER]

		valuetype := strings.Split(args[1], ";")
		// valuetype =
		// [hello world string]
		// [true bool]
		// [3.5 float]
		// [4 INTEGER]

		name, value, vtype := strings.TrimSpace(args[0]), strings.TrimSpace(valuetype[0]), strings.ToUpper(strings.TrimSpace(valuetype[1]))
		// fmt.Println("name ", name, "value ", value, "vtype ", vtype)
		// prints
		// name:  testString, value:  hello world, vtype:  STRING
		// name:  testBool, value:  true, vtype:  BOOL
		// name:  testFloat, value:  3.5, vtype:  FLOAT
		// name:  TestInt, value: 4, vtype:  INTEGER

		fields.Add(name, value, vtype)
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	valueType := value.Type()
	// valueType = main.ConfS

	// value.NumField() = 4
	for i := 0; i < value.NumField(); i++ {
		fieldType := valueType.Field(i)

		// fieldType=
		// {TS  string name:"testString" 0 [0] false}
		// {TB  bool name:"testBool" 16 [1] false}
		// {TF  float64 name:"testFloat" 24 [2] false}
		fieldValue := value.Field(i)
		// fieldValue =
		// false
		// 0
		// 0
		name := fieldType.Tag.Get("name")
		// name =
		// testString
		// testBool
		// testFloat

		if name == "" {
			name = fieldType.Name
			// TS
		}

		if value, ok := fields[name]; ok {
			// value =
			// hello world
			// true
			// 3.5
			// 4

			// fieldValue = emtpy
			fieldValue.Set(value)
			// fieldValue = hello world
		}
	}
	return nil
}
