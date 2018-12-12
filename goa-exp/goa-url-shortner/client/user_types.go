// Code generated by goagen v1.3.1, DO NOT EDIT.
//
// API "URL shortner API": Application User Types
//
// Command:
// $ goagen
// --design=github.com/nikhil-thomas/technik-dojo/goa-exp/goa-url-shortner/design
// --out=$(GOPATH)/src/github.com/nikhil-thomas/technik-dojo/goa-exp/goa-url-shortner
// --version=v1.3.1

package client

import (
	"github.com/goadesign/goa"
	"unicode/utf8"
)

// createLinkPayload user type.
type createLinkPayload struct {
	Path *string `form:"path,omitempty" json:"path,omitempty" yaml:"path,omitempty" xml:"path,omitempty"`
	URL  *string `form:"url,omitempty" json:"url,omitempty" yaml:"url,omitempty" xml:"url,omitempty"`
}

// Validate validates the createLinkPayload type instance.
func (ut *createLinkPayload) Validate() (err error) {
	if ut.Path == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`request`, "path"))
	}
	if ut.URL == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`request`, "url"))
	}
	if ut.Path != nil {
		if utf8.RuneCountInString(*ut.Path) < 6 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`request.path`, *ut.Path, utf8.RuneCountInString(*ut.Path), 6, true))
		}
	}
	if ut.Path != nil {
		if utf8.RuneCountInString(*ut.Path) > 100 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`request.path`, *ut.Path, utf8.RuneCountInString(*ut.Path), 100, false))
		}
	}
	if ut.URL != nil {
		if utf8.RuneCountInString(*ut.URL) < 8 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`request.url`, *ut.URL, utf8.RuneCountInString(*ut.URL), 8, true))
		}
	}
	if ut.URL != nil {
		if utf8.RuneCountInString(*ut.URL) > 2000 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`request.url`, *ut.URL, utf8.RuneCountInString(*ut.URL), 2000, false))
		}
	}
	return
}

// Publicize creates CreateLinkPayload from createLinkPayload
func (ut *createLinkPayload) Publicize() *CreateLinkPayload {
	var pub CreateLinkPayload
	if ut.Path != nil {
		pub.Path = *ut.Path
	}
	if ut.URL != nil {
		pub.URL = *ut.URL
	}
	return &pub
}

// CreateLinkPayload user type.
type CreateLinkPayload struct {
	Path string `form:"path" json:"path" yaml:"path" xml:"path"`
	URL  string `form:"url" json:"url" yaml:"url" xml:"url"`
}

// Validate validates the CreateLinkPayload type instance.
func (ut *CreateLinkPayload) Validate() (err error) {
	if ut.Path == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`type`, "path"))
	}
	if ut.URL == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`type`, "url"))
	}
	if utf8.RuneCountInString(ut.Path) < 6 {
		err = goa.MergeErrors(err, goa.InvalidLengthError(`type.path`, ut.Path, utf8.RuneCountInString(ut.Path), 6, true))
	}
	if utf8.RuneCountInString(ut.Path) > 100 {
		err = goa.MergeErrors(err, goa.InvalidLengthError(`type.path`, ut.Path, utf8.RuneCountInString(ut.Path), 100, false))
	}
	if utf8.RuneCountInString(ut.URL) < 8 {
		err = goa.MergeErrors(err, goa.InvalidLengthError(`type.url`, ut.URL, utf8.RuneCountInString(ut.URL), 8, true))
	}
	if utf8.RuneCountInString(ut.URL) > 2000 {
		err = goa.MergeErrors(err, goa.InvalidLengthError(`type.url`, ut.URL, utf8.RuneCountInString(ut.URL), 2000, false))
	}
	return
}

// updateLinkPayload user type.
type updateLinkPayload struct {
	Path *string `form:"path,omitempty" json:"path,omitempty" yaml:"path,omitempty" xml:"path,omitempty"`
	URL  *string `form:"url,omitempty" json:"url,omitempty" yaml:"url,omitempty" xml:"url,omitempty"`
}

// Validate validates the updateLinkPayload type instance.
func (ut *updateLinkPayload) Validate() (err error) {
	if ut.Path == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`request`, "path"))
	}
	if ut.URL == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`request`, "url"))
	}
	if ut.Path != nil {
		if utf8.RuneCountInString(*ut.Path) < 6 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`request.path`, *ut.Path, utf8.RuneCountInString(*ut.Path), 6, true))
		}
	}
	if ut.Path != nil {
		if utf8.RuneCountInString(*ut.Path) > 100 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`request.path`, *ut.Path, utf8.RuneCountInString(*ut.Path), 100, false))
		}
	}
	if ut.URL != nil {
		if utf8.RuneCountInString(*ut.URL) < 8 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`request.url`, *ut.URL, utf8.RuneCountInString(*ut.URL), 8, true))
		}
	}
	if ut.URL != nil {
		if utf8.RuneCountInString(*ut.URL) > 2000 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`request.url`, *ut.URL, utf8.RuneCountInString(*ut.URL), 2000, false))
		}
	}
	return
}

// Publicize creates UpdateLinkPayload from updateLinkPayload
func (ut *updateLinkPayload) Publicize() *UpdateLinkPayload {
	var pub UpdateLinkPayload
	if ut.Path != nil {
		pub.Path = *ut.Path
	}
	if ut.URL != nil {
		pub.URL = *ut.URL
	}
	return &pub
}

// UpdateLinkPayload user type.
type UpdateLinkPayload struct {
	Path string `form:"path" json:"path" yaml:"path" xml:"path"`
	URL  string `form:"url" json:"url" yaml:"url" xml:"url"`
}

// Validate validates the UpdateLinkPayload type instance.
func (ut *UpdateLinkPayload) Validate() (err error) {
	if ut.Path == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`type`, "path"))
	}
	if ut.URL == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`type`, "url"))
	}
	if utf8.RuneCountInString(ut.Path) < 6 {
		err = goa.MergeErrors(err, goa.InvalidLengthError(`type.path`, ut.Path, utf8.RuneCountInString(ut.Path), 6, true))
	}
	if utf8.RuneCountInString(ut.Path) > 100 {
		err = goa.MergeErrors(err, goa.InvalidLengthError(`type.path`, ut.Path, utf8.RuneCountInString(ut.Path), 100, false))
	}
	if utf8.RuneCountInString(ut.URL) < 8 {
		err = goa.MergeErrors(err, goa.InvalidLengthError(`type.url`, ut.URL, utf8.RuneCountInString(ut.URL), 8, true))
	}
	if utf8.RuneCountInString(ut.URL) > 2000 {
		err = goa.MergeErrors(err, goa.InvalidLengthError(`type.url`, ut.URL, utf8.RuneCountInString(ut.URL), 2000, false))
	}
	return
}
