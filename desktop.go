package desktop

import (
	"strings"

	"github.com/CalebQ42/desktop/ini"
)

type File struct {
	Base *ini.File
}

func (f File) Exec() string {
	return f.Base.Section("Desktop Entry").Value("Exec").String()
}

//Get's the value with the given key under the particular Locale.
//If the given locale is not present, tries to find the best fit.
//Must specify key and lang.
func (f File) ValueWithLocale(key string, lang, country, modifier string) ini.Value {
	if country != "" && modifier != "" {
		if f.Base.Section("Desktop Entry").HasKey(key + "[" + lang + "_" + strings.ToUpper(country) + "@" + modifier + "]") {
			return f.Base.Section("Desktop Entry").Value(key + "[" + lang + "_" + strings.ToUpper(country) + "@" + modifier + "]")
		}
	}
	if country != "" {
		if f.Base.Section("Desktop Entry").HasKey(key + "[" + lang + "_" + strings.ToUpper(country) + "]") {
			return f.Base.Section("Desktop Entry").Value(key + "[" + lang + "_" + strings.ToUpper(country) + "]")
		}
	}
	if modifier != "" {
		if f.Base.Section("Desktop Entry").HasKey(key + "[" + lang + "@" + modifier + "]") {
			return f.Base.Section("Desktop Entry").Value(key + "[" + lang + "@" + modifier + "]")
		}
	}
	if f.Base.Section("Desktop Entry").HasKey(key + "[" + lang + "]") {
		return f.Base.Section("Desktop Entry").Value(key + "[" + lang + "]")
	}
	return f.Base.Section("Desktop Entry").Value(key)
}
