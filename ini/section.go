package ini

import "io"

type Section struct {
	m     map[string]Value
	order []string
}

func NewSection() *Section {
	return &Section{
		m: make(map[string]Value),
	}
}

//The keys with values present in the section.
func (s Section) Keys() []string {
	return s.order
}

//If the key is not present, adds the key with the given value.
//If the key is present, add the given value as an additional line (MultivalueArray)
func (s *Section) AddValue(key string, value string) {
	_, ok := s.m[key]
	if ok {
		s.m[key] = append(s.m[key], value)
	} else {
		s.m[key] = []string{value}
		s.order = append(s.order, key)
	}
}

//Sets the value of key to value.
//This will overwrite if the key already has a value.
//If the value is a multivalue array, ALL values (except for the new value) will be deleted.
func (s *Section) SetValue(key string, value string) {
	_, ok := s.m[key]
	s.m[key] = []string{value}
	if !ok {
		s.order = append(s.order, key)
	}
}

//Does the section have the given key.
func (s Section) HasKey(key string) bool {
	_, ok := s.m[key]
	return ok
}

//The value associated with the key.
//If the key is not present, returns an empty string value.
func (s Section) Value(key string) Value {
	v, ok := s.m[key]
	if !ok {
		return []string{""}
	}
	return v
}

//Deletes the given key from the section.
func (s *Section) DeleteValue(key string) {
	if s.HasKey(key) {
		delete(s.m, key)
		for i := range s.order {
			if s.order[i] == key {
				s.order = append(s.order[:i], s.order[i+1:]...)
			}
		}
	}
}

func (s Section) writeTo(name string, w io.Writer) (n int64, err error) {
	var i int
	if name != "" {
		i, err = w.Write([]byte("[" + name + "]"))
		n += int64(i)
		if err != nil {
			return
		}
	}
	var newN int64
	for _, v := range s.order {
		newN, err = s.Value(v).writeTo(v, w)
		n += newN
		if err != nil {
			return
		}
		i, err = w.Write([]byte("\n"))
		n += int64(i)
		if err != nil {
			return
		}
	}
	return
}
