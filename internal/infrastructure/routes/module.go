package routes

import "go.uber.org/fx"

var Module = fx.Options(
	fx.Provide(NewRouter),
	fx.Invoke(InitGinEngine),
	fx.Provide(fx.Annotate(RegisterRoutes, fx.ParamTags(`group:"routes"`))),
)