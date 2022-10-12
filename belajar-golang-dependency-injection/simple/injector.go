//go:build wireinject
// +build wireinject

package simple

import (
	"github.com/google/wire"
	"io"
	"os"
)

func InitializedService(isError bool) (*SimpleService, error) {
	wire.Build(NewSimpleRepository, NewSimpleService)
	return nil, nil
}

func InitializeDatabase() *DatabaseRepository {
	wire.Build(NewDatabaseMongoDB, NewDatabasePostgreeSQL, NewDatabaseRepository)
	return nil
}

var fooSet = wire.NewSet(NewFooRepository, NewFooService)

var barSet = wire.NewSet(NewBarRepository, NewBarService)

func InitializeFooBarService() *FooBarService {
	wire.Build(fooSet, barSet, NewFooBarService)
	return nil
}

// // error
// func InitializedHelloService() *HelloService {
// 	wire.Build(NewHelloService, NewSayHelloImp)
// 	return nil
// }

// siapapun yang membutuhkan SayHello, maka akan dikirimkan SayHelloImpl
var helloSet = wire.NewSet(
	NewSayHelloImpl,
	wire.Bind(new(SayHello), new(*SayHelloImpl)),
)

func InitializeHelloService() *HelloService {
	wire.Build(helloSet, NewHelloService)
	return nil
}

// provider struct
// jarang digunakan
var fooBarSet = wire.NewSet(
	NewFoo,
	NewBar,
)

func InitializedFooBar() *FooBar {
	wire.Build(
		fooBarSet,
		// wire.Struct(new(FooBar), "Foo", "Bar"), // Foo dan Bar akan diinject ke foobar struct nya
		wire.Struct(new(FooBar), "*"), // all field
	)
	return nil
}

// value
var fooValue = &Foo{}
var barValue = &Bar{}

func InitializedFooBarUsingValue() *FooBar {
	wire.Build(
		wire.Value(fooValue),
		wire.Value(barValue),
		wire.Struct(new(FooBar), "*"),
	)
	return nil
}

// injector interface value
func InitializedReader() io.Reader {
	wire.Build(wire.InterfaceValue(new(io.Reader), os.Stdin))
	return nil
}

// mengambil field configuration pada struct application
func InitializedConfiguration() *Configuration {
	wire.Build(
		NewApplication,
		wire.FieldsOf(new(*Application), "Configuration"),
	)
	return nil
}

// func nya akan mereturn cleanup nya, liat di wire_gen
func InitializedConnection(name string) (*Connection, func()) {
	wire.Build(NewConnection, NewFile)
	return nil, nil
}
