package main

import (
	"fmt"

	"github.com/Hydra/HydraConfigurator"
)

// ConfS - configuration struct
type ConfS struct {
	TS      string  `name:"testString" xml:"testString" json:"testString"`
	TB      bool    `name:"testBool" xml:"testBool" json:"testBool"`
	TF      float64 `name:"testFloat" xml:"testFloat" json:"testFloat"`
	TestInt int
}

//Tests to make sure everything was decoded correctly fromt he various file types
func main() {
	configstruct := new(ConfS)
	HydraConfigurator.GetConfiguration(HydraConfigurator.XML, configstruct, "configfile.xml")
	// HydraConfigurator.GetConfiguration(HydraConfigurator.CUSTOM, configstruct, "configfile.conf")
	fmt.Println(*configstruct)
	if configstruct.TB {
		fmt.Println("bool is true")
	}

	fmt.Println("\n configstruct.TestFloat -- TF")
	fmt.Println(float64(4.8 * configstruct.TF))

	fmt.Println("\n configstruct.TestInt")
	fmt.Println(5 * configstruct.TestInt)

	fmt.Println("\n configstruct.TestString --TS ")
	fmt.Println(configstruct.TS)
}
