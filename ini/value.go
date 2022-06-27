package ini

import (
	"io"
	"os"
	"strconv"
	"strings"
)

//Value holds the value of a Key=Value pair.
//If multiple lines have the same key, only the first value is used, except for with MultivalueArray.
type Value []string

func NewValue(s string) Value {
	return []string{s}
}

//Returns the value as a string.
func (v Value) String() string {
	return v[0]
}

//Returns if value is a bool value (only true and false are allowed). Case insensitive.
func (v Value) IsBool() bool {
	return strings.ToLower(v[0]) == "true" || strings.ToLower(v[0]) == "false"
}

//Returns if value == "true". Other values returns false. Case insensitive.
func (v Value) Bool() bool {
	return strings.ToLower(v[0]) == "true"
}

//Returns if there are multiple value with this same key.
func (v Value) IsMultivalueArray() bool {
	return len(v) > 1
}

//Returns all values with the same key.
func (v Value) MultivalueArray() []string {
	return v
}

//Returns if value is an int.
func (v Value) IsInt() bool {
	_, err := strconv.Atoi(v[0])
	return err == nil
}

//Returns as an int. If ! IsInt, returns 0.
func (v Value) Int() int {
	i, err := strconv.Atoi(v[0])
	if err == nil {
		return i
	}
	return 0
}

//Returns if the value contains any commas, perhaps indicating that it's a comma delineated array.
func (v Value) IsCommaArray() bool {
	return strings.Contains(v[0], ",")
}

//Returns the value as a comma delineated array. Each output is trimmed.
func (v Value) CommaArray() (out []string) {
	out = strings.Split(v[0], ",")
	for i := range out {
		out[i] = strings.TrimSpace(out[i])
	}
	return
}

//Returns if the value contains any semicolons, perhaps indicating that it's a semicolon delineated array.
func (v Value) IsSemicolonArray() bool {
	return strings.Contains(v[0], ";")
}

//Returns the value as a semicolon delineated array. Each output is trimmed.
func (v Value) SemicolonArray() (out []string) {
	out = strings.Split(v[0], ";")
	for i := range out {
		out[i] = strings.TrimSpace(out[i])
	}
	return
}

//Attempts to open the file pointed to by the value.
func (v Value) File() (*os.File, error) {
	return os.Open(v[0])
}

func (v Value) writeTo(key string, w io.Writer) (n int64, err error) {
	i, err := w.Write([]byte(key + "=" + v[0]))
	n += int64(i)
	if err != nil || len(v) == 1 {
		return
	}
	for j := 1; j < len(v); j++ {
		i, err = w.Write([]byte("\n" + key + "=" + v[j]))
		n += int64(i)
		if err != nil {
			return
		}
	}
	return
}
