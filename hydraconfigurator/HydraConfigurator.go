package hydraconfigurator
//entry point for package
import (
	"errors"
	"reflect"
)

//will allow for other file formats
const (
	CUSTOM uint8 = iota
	JSON
	XML
)

// if type is nto a struct then reflect cannot be set
var wrongTypeError error = errors.New("Type must be a pointer to a struct")

func GetConfiguration(confType uint8, obj interface{}, filename string) (err error){
	//check if this is a type pointer
	mysRValue := reflect.ValueOf(obj)
	if mysRValue.Kind() != reflect.Ptr || mysRValue.IsNil(){
		return wrongTypeError
	}
	//get an dconfirm the struct value
	mysRValue = mysRValue.Elem()
	if mysRValue.Kind() != reflect.Struct{
		return wrongTypeError
	}

	switch confType{
	case CUSTOM:
		err = MarshalCustomConfig(mysRValue, filename)
	case JSON:
		err = decodeJSONConfig(obj, filename)
	case XML:
		err = decodeXMLConfig(obj, filename)
	}
	return err
}