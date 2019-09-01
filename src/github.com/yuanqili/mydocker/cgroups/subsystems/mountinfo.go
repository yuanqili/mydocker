// The code is adapted from
//   https://github.com/fntlnz/mountinfo/blob/master/mountinfo.go

/**
https://www.kernel.org/doc/Documentation/filesystems/proc.txt

3.5	/proc/<pid>/mountinfo - Information about mounts
--------------------------------------------------------

This file contains lines of the form:

36 35 98:0 /mnt1 /mnt2 rw,noatime master:1 - ext3 /dev/root rw,errors=continue
(1)(2)(3)   (4)   (5)      (6)      (7)   (8) (9)   (10)         (11)

(1) mount ID:  unique identifier of the mount (may be reused after umount)
(2) parent ID:  ID of parent (or of self for the top of the mount tree)
(3) major:minor:  value of st_dev for files on filesystem
(4) root:  root of the mount within the filesystem
(5) mount point:  mount point relative to the process's root
(6) mount options:  per mount options
(7) optional fields:  zero or more fields of the form "tag[:value]"
(8) separator:  marks the end of the optional fields
(9) filesystem type:  name of filesystem of the form "type[.subtype]"
(10) mount source:  filesystem specific information or "none"
(11) super options:  per super block options

Parsers should ignore all unrecognised optional fields.  Currently the
possible optional fields are:

shared:X  mount is shared in peer group X
master:X  mount is slave to peer group X
propagate_from:X  mount is slave and receives propagation from peer group X (*)
unbindable  mount is unbindable

(*) X is the closest dominant peer group under the process's root.  If
X is the immediate master of the mount, or if there's no dominant peer
group under the same root, then only the "master:X" field is present
and not the "propagate_from:X" field.
*/

package subsystems

import (
	"bufio"
	"io"
	"os"
	"strings"
)

type Mountinfo struct {
	MountID        string
	ParentID       string
	MajorMinor     string
	Root           string
	MountPoint     string
	MountOptions   string
	OptionalFields string
	FilesystemType string
	MountSource    string
	SuperOptions   string
}

func GetMountinfo(fd string) ([]Mountinfo, error) {
	f, err := os.Open(fd)
	if err != nil {
		return nil, err
	}

	defer f.Close()
	return ParseMountinfo(f)
}

func ParseMountinfo(buffer io.Reader) ([]Mountinfo, error) {
	var info []Mountinfo
	scanner := bufio.NewScanner(buffer)
	for scanner.Scan() {
		txt := scanner.Text()
		info = append(info, *ParseMountinfoString(txt))
	}
	err := scanner.Err()
	return info, err
}

func ParseMountinfoString(txt string) *Mountinfo {
	pieces := strings.Split(txt, " ")
	count := len(pieces)
	if count < 1 {
		return nil
	}

	i := strings.Index(txt, " - ")
	preFields := strings.Fields(txt[:i])
	postFields := strings.Fields(txt[i+3:])

	return &Mountinfo{
		MountID:        getMountPart(preFields, 0),
		ParentID:       getMountPart(preFields, 1),
		MajorMinor:     getMountPart(preFields, 2),
		Root:           getMountPart(preFields, 3),
		MountPoint:     getMountPart(preFields, 4),
		MountOptions:   getMountPart(preFields, 5),
		OptionalFields: getMountPart(preFields, 6),
		FilesystemType: getMountPart(postFields, 0),
		MountSource:    getMountPart(postFields, 1),
		SuperOptions:   getMountPart(postFields, 2),
	}
}

func getMountPart(pieces []string, index int) string {
	if len(pieces) > index {
		return pieces[index]
	} else {
		return ""
	}
}
