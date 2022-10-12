package simple

type SayHello interface {
	Hello(name string) string
}

type HelloService struct {
	SayHello
}

type SayHelloImpl struct {
}

func (s *SayHelloImpl) Hello(name string) string {
	return "Hello " + name
}

// ingat yang direturn kan itu selalu hasil akhir yaitu sayhelloimpl
func NewSayHelloImpl() *SayHelloImpl {
	return &SayHelloImpl{}
}

// ingat yang direturn kan itu selalu hasil akhir yaitu sayhelloservice
func NewHelloService(sayHello SayHello) *HelloService {
	return &HelloService{
		SayHello: sayHello,
	}
}
