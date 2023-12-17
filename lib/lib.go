package lib

import "go.uber.org/fx"

// Module exports dependency
var Module = fx.Options(
	fx.Provide(NewRequestHandler),
	fx.Provide(NewEnv),
	fx.Provide(GetLogger),
	fx.Provide(NewDatabase),
	fx.Provide(NewSuperBaseClient),
	fx.Provide(NewPasswordHash),
	fx.Provide(NewSnowflakeService),
	fx.Provide(NewSessionStore),
)
