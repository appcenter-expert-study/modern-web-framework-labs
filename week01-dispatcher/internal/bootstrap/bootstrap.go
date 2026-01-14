package bootstrap

import (
	"modern-web-framework-labs/week0-dispatcher/internal/dispatcher"
	"modern-web-framework-labs/week0-dispatcher/internal/http"
)

/*
bootstrap은 조립만 담당합니다.
요청 처리 로직은 여기에 포함하지 않습니다.
*/
func Run() {
	// HTTP Engine 생성 (한번 추상화된 EchoEngine 구조체 반환됨)
	engine := http.NewEchoEngine()
	// Dispatcher 생성
	dispatcher := dispatcher.NewDispatcher()
	// Engine에 Dispatcher 등록
	engine.RegisterDispatcher(dispatcher)
	// 서버 시작
	engine.Start(":8080")
}
