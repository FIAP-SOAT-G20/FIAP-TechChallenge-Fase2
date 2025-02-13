package handler

import (
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/infrastructure/routes"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(
		AsRoute(NewProductHandler),
	),
)

func AsRoute(f any) any {
	return fx.Annotate(
		f,
		fx.As(new(routes.IRouter)),
		fx.ResultTags(`group:"routes"`),
	)
}
