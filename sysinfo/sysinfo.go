package sysinfo

import "runtime"

// SystemInfo encapsulates useful information about a system
type SystemInfo struct {
	OS      string `json:"os,omitempty"`
	Name    string `json:"name,omitempty"`
	Version string `json:"version,omitempty"`
	Arch    string `json:"arch,omitempty"`

	Model  string `json:"model,omitempty"`
	CPU    string `json:"cpu,omitempty"`
	Memory int64  `json:"memory,omitempty"`
}

// GetSystemInfo populates a new SystemInfo with both generic and
// platform specific information
func GetSystemInfo() SystemInfo {
	info := SystemInfo{
		Arch: runtime.GOARCH,
	}
	switch runtime.GOOS {
	case "darwin":
		info.OS = "osx"
	case "windows":
		info.OS = "win"
	default:
		info.OS = runtime.GOOS
	}
	return getPlatformSpecifics(info)
}

// List returns the system info as a list of strings
func (s SystemInfo) List() []string {
	l := []string{
		s.OS,
		s.Name,
		s.Version,
		s.Arch,
	}
	switch s.OS {
	case "osx", "win":
		l = append(l, runtime.GOOS)
	}
	return l
}
