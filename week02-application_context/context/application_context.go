package context

import "reflect"

type ApplicationContext struct {
	// 등록된 Bean 정의
	definitions map[reflect.Type]BeanDefinition
	// 이미 생성된 Bean 캐시
	singletons map[reflect.Type]any
	// 순환 의존성 감지용 플래그
	creating map[reflect.Type]bool
}

func NewApplicationContext() *ApplicationContext {
	return &ApplicationContext{
		definitions: map[reflect.Type]BeanDefinition{},
		singletons:  map[reflect.Type]any{},
		creating:    map[reflect.Type]bool{},
	}
}

/*
Bean 등록 메서드
Bean의 '정의'만 등록
객체를 여기서 만들지 않습니다.
*/
func (ctx *ApplicationContext) RegisterBean(beanType reflect.Type, factory func(ctx *ApplicationContext) any)

/*
Bean 획득 메서드
GetBean의 동작 규칙을 지켜서 구현해야 합니다.

1. 이미 생성된 Bean이면 -> 캐시를 반환합니다.
2. 지금 생성중이면 순환 의존성이므로, 즉시 실패합니다.
3. 정의가 없으면 실패합니다.
4. Factory를 실행하고 Bean을 반환받습니다.
5. Bean 결과를 Singleton 캐시에 저장합니다.
6. 생성된 Bean을 반환합니다.
*/
func (ctx *ApplicationContext) GetBean(beanType reflect.Type) any
