package services

import (
	"context"
	"errors"
	"reflect"
)
func GetService(target interface{}) error{
return GetServiceForContext(context.Background(),target)
}
func GetServiceForContext(c context.Context ,target interface{})(err error){
	targetValue := reflect.ValueOf(target)
	if targetValue.Kind() == reflect.Pointer && targetValue.Elem().CanSet(){
		err = resolveServiceFromValue(c,targetValue)
	}else{
err = errors.New("Type cannot used as target")
	}
	return
}
