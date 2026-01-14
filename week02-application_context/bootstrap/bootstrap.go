package bootstrap

import (
	"reflect"
	"week02-application_context/context"
	"week02-application_context/dispatcher"
	"week02-application_context/http"
)

func Run() {
	engine := http.NewEchoEngine()

	// ApplicationContext 생성 (IoC)
	ctx := context.NewApplicationContext()

	// ApplicationContext에 Dispatcher 등록
	ctx.RegisterBean(
		reflect.TypeFor[*dispatcher.Dispatcher](),
		func(ctx *context.ApplicationContext) any {
			return dispatcher.NewDispatcher()
		},
	)

	// ApplicationContext에서 Dispatcher 타입으로 Bean 가져오기
	dispatcher := ctx.GetBean(
		reflect.TypeFor[*dispatcher.Dispatcher](),
	).(*dispatcher.Dispatcher)

	// Dispatcher Bean을 Engine에 등록
	engine.RegisterDispatcher(dispatcher)
	engine.Start(":8080")
}
