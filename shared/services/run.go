package services

import (
	"log"
	"slices"

	"github.com/carabalonepaulo/origin/shared/service"
)

func castAndAdd[T any](slice []T, value any) []T {
	v, ok := value.(T)
	if ok {
		return append(slice, v)
	}
	return slice
}

func Run(ctors ...func() service.Service) error {
	updatable := make([]service.UpdatableService, 0)
	instances := make([]service.Service, 0)

	names := make(map[service.Service]string)
	services := make(map[string]service.Service)
	for _, ctor := range ctors {
		instance := ctor()
		name := service.NameOf(instance)

		services[name] = instance
		names[instance] = name
		instances = append(instances, instance)
		updatable = castAndAdd(updatable, instance)
	}

	running := true
	shutdown := func() { running = false }
	alreadyStarted := make([]service.Service, 0)

	var startErr error
	for _, instance := range instances {
		err := instance.Start(services, shutdown)
		if err != nil {
			log.Printf("service `%s` failed to start with error: %s", names[instance], err)
			for _, activeInstance := range alreadyStarted {
				activeInstance.Stop()
			}
			startErr = err
			break
		} else {
			log.Printf("service `%s` started!", names[instance])
			alreadyStarted = append(alreadyStarted, instance)
		}
	}
	if startErr != nil {
		return startErr
	}

	for running {
		for _, service := range updatable {
			service.Update(0.0)
		}
	}

	slices.Reverse(instances)
	for _, instance := range instances {
		instance.Stop()
		log.Printf("service `%s` stopped!", names[instance])
	}

	return nil
}
