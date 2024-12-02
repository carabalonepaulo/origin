package service

import (
	"fmt"
	"reflect"
	"strings"
)

type Services map[string]Service

func Get[T Service](s Services) (value T, err error) {
	name := NameFor[T]()

	service, ok := s[name]
	if !ok {
		err = fmt.Errorf("service `%s` not registered", name)
		return
	}

	value, ok = service.(T)
	if !ok {
		err = fmt.Errorf("service `%s` not registered, wrong type", name)
	}

	return
}

func NameOf(instance Service) string {
	ty := reflect.TypeOf(instance)
	parts := strings.Split(ty.Elem().PkgPath(), "/")
	return parts[len(parts)-1]
}

func NameFor[T Service]() string {
	ty := reflect.TypeFor[T]()
	parts := strings.Split(ty.Elem().PkgPath(), "/")
	return parts[len(parts)-1]
}

type Service interface {
	Start(Services, func()) error
	Stop()
}

type UpdatableService interface {
	Service
	Update(float64)
}
