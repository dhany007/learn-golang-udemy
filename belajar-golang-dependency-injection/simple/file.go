package simple

import "fmt"

type File struct {
	Name string
}

// provider return nya ada fungsi, yaitu closure
func NewFile(name string) (*File, func()) {
	file := &File{
		Name: name,
	}
	// diterurnkan file dan fungsi
	return file, func() {
		file.Close()
	}
}

func (f *File) Close() {
	fmt.Print("Close File", f.Name)
}
