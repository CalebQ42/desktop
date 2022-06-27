package ini

import "io"

type File struct {
	pre   *Section
	m     map[string]*Section
	order []string
}

func (f File) Sections() []string {
	return f.order
}

//Returns a Section with the key=value pairs that appear before the first section.
func (f File) PreSection() *Section {
	return f.pre
}

//Does the File have the given section.
func (f File) HasSection(name string) bool {
	_, ok := f.m[name]
	return ok
}

//The section with the given name.
//If the section is not already present, a new section is created.
func (f *File) Section(name string) *Section {
	v, ok := f.m[name]
	if !ok {
		f.m[name] = NewSection()
	}
	return v
}

//Deletes the given key from the section.
func (f *File) DeleteSection(name string) {
	if f.HasSection(name) {
		delete(f.m, name)
		for i := range f.order {
			if f.order[i] == name {
				f.order = append(f.order[:i], f.order[i+1:]...)
			}
		}
	}
}

func (f File) WriteTo(w io.Writer) (n int64, err error) {
	n, err = f.pre.writeTo("", w)
	if err != nil {
		return
	}
	var i int
	var newN int64
	for _, v := range f.order {
		i, err = w.Write([]byte("\n"))
		n += int64(i)
		if err != nil {
			return
		}
		newN, err = f.Section(v).writeTo(v, w)
		n += newN
		if err != nil {
			return
		}
	}
	return
}
