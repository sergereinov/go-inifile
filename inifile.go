package inifile

import (
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"gopkg.in/ini.v1"
)

const (
	// Default ini file path in case of os.Executable() error
	_DEFAULT_INI_PATH = "./config.ini"

	_DEFAULT_EXT   = ".ini"
	_DEFAULT_DELIM = ","
)

type iniFile struct {
	*ini.File
	path string
}

// Simplified constructor
func New(optPath ...string) *iniFile {
	file, _ := NewIniFile(optPath...)
	return file
}

// Main constructor
func NewIniFile(optPath ...string) (*iniFile, error) {
	path := buildIniPath(optPath...)
	file, err := ini.Load(path)
	if err != nil {
		return &iniFile{File: ini.Empty(), path: path}, err
	}
	return &iniFile{File: file, path: path}, nil
}

func buildIniPath(optPath ...string) string {
	// Return absolute path if given
	if len(optPath) > 0 && len(optPath[0]) > 0 {
		path := optPath[0]
		if !filepath.IsAbs(path) {
			execPath, _ := os.Executable()
			path = filepath.Clean(filepath.Join(filepath.Dir(execPath), path))
		}
		return path
	}

	// Build ini path from executable location
	path, err := os.Executable()
	if err != nil || len(path) == 0 {
		return _DEFAULT_INI_PATH
	}

	// Remove ext
	for i := len(path) - 1; i >= 0 && !os.IsPathSeparator(path[i]); i-- {
		if path[i] == '.' {
			path = path[:i]
			break
		}
	}

	return path + _DEFAULT_EXT
}

func (f *iniFile) Path() string {
	return f.path
}

func (f *iniFile) Save() {
	f.SaveTo(f.path)
}

func (f *iniFile) String(section, key, defaultVal string) string {
	return f.Section(section).Key(key).MustString(defaultVal)
}

func (f *iniFile) Int(section, key string, defaultVal int) int {
	return f.Section(section).Key(key).MustInt(defaultVal)
}

func (f *iniFile) Bool(section, key string, defaultVal bool) bool {
	return f.Section(section).Key(key).MustBool(defaultVal)
}

func (f *iniFile) Duration(section, key string, defaultVal time.Duration) time.Duration {
	return f.Section(section).Key(key).MustDuration(defaultVal)
}

func (f *iniFile) Strings(section, key string, defaultVal []string, optDelim ...string) []string {
	var delim string
	if len(optDelim) > 0 {
		delim = optDelim[0]
	} else {
		delim = _DEFAULT_DELIM
	}

	s := f.Section(section)
	if s.HasKey(key) {
		return s.Key(key).Strings(delim)
	}

	s.Key(key).SetValue(strings.Join(defaultVal, delim+" "))
	return defaultVal
}

func (f *iniFile) Ints(section, key string, defaultVal []int, optDelim ...string) []int {
	var delim string
	if len(optDelim) > 0 {
		delim = optDelim[0]
	} else {
		delim = _DEFAULT_DELIM
	}

	s := f.Section(section)
	if s.HasKey(key) {
		return s.Key(key).Ints(delim)
	}

	var sb strings.Builder
	for i, v := range defaultVal {
		s := strconv.FormatInt(int64(v), 10)
		sb.WriteString(s)
		if i < len(defaultVal)-1 {
			sb.WriteString(delim)
			sb.WriteString(" ")
		}
	}
	s.Key(key).SetValue(sb.String())
	return defaultVal
}
