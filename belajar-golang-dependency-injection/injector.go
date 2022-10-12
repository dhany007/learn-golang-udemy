//go:build wireinject
// +build wireinject

package main

import (
	"dhany007/belajar-golang-restful-api/app"
	"dhany007/belajar-golang-restful-api/controller"
	"dhany007/belajar-golang-restful-api/middleware"
	"dhany007/belajar-golang-restful-api/repository"
	"dhany007/belajar-golang-restful-api/service"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/google/wire"
	"github.com/julienschmidt/httprouter"
)

var categorySet = wire.NewSet(
	repository.NewCategoryRespository,
	wire.Bind(new(repository.CategoryRepository), new(*repository.CategoryRespositoryImpl)),
	service.NewCategoryService,
	wire.Bind(new(service.CategoryService), new(*service.CategoryServiceImpl)),
	controller.NewCategoryController,
	wire.Bind(new(controller.CategoryController), new(*controller.CategoryControllerImpl)),
)

// ambil dari server semua. liat rest-full semua
func InitializedServer() *http.Server {
	wire.Build(
		app.NewDB,
		validator.New,
		categorySet,
		app.NewRouter,
		wire.Bind(new(http.Handler), new(*httprouter.Router)),
		middleware.NewAuthMiddleware,
		NewServer,
	)

	return nil
}
