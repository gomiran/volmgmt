package usnfilter

import (
	"github.com/gomiran/volmgmt/fileattr"
	"github.com/gomiran/volmgmt/usn"
)

// IsDir is a filter that returns true returns true when a record represents
// a directory.
func IsDir(record usn.Record) bool {
	return record.FileAttributes.Match(fileattr.Directory)
}
