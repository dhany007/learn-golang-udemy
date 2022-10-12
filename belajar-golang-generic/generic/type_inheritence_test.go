package generic

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Employee interface {
	GetName() string
}

func GetName[T Employee](param T) string {
	return param.GetName()
}

type Manager interface {
	GetName() string
	GetManagerName() string
}

type MyManager struct {
	Name string
}

func (m *MyManager) GetName() string {
	return m.Name
}

func (m *MyManager) GetManagerName() string {
	return m.Name
}

type VicePresident interface {
	GetName() string
	GetVicePresidentName() string
}

type MyVicePresident struct {
	Name string
}

func (v *MyVicePresident) GetName() string {
	return v.Name
}

func (v *MyVicePresident) GetVicePresidentName() string {
	return v.Name
}

func TestInheritence(t *testing.T) {
	// Disini GetName bisa dipakai oleh manager dan vicepresident
	// Karena kontrak GetName pada employe sama seperti GetName pada Manager dan VicePresident
	assert.Equal(t, "Dhany", GetName[Manager](&MyManager{Name: "Dhany"}))
	assert.Equal(t, "Dhany", GetName[VicePresident](&MyVicePresident{Name: "Dhany"}))
}
