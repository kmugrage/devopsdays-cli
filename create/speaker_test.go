package create

import (
	"testing"

	"github.com/devopsdays/devopsdays-cli/model"
	. "github.com/smartystreets/goconvey/convey"
)

func TestNewSpeaker(t *testing.T) {

	Convey("Given a new object of type Speaker", t, func() {
		mySpeaker := model.Speaker{
			Name:      "matt-stratton",
			Title:     "Matt Stratton",
			Website:   "http://www.mattstratton.com",
			Twitter:   "mattstratton",
			Facebook:  "https://www.facebook.com/matt.stratton",
			Linkedin:  "https://www.linkedin.com/in/mattstratton/",
			Github:    "mattstratton",
			Gitlab:    "mattstratton",
			ImagePath: "matt-stratton.jpg",
		}
		Convey("When the object is passed to NewSpeaker function", func() {
			err := NewSpeaker(mySpeaker, "ponyville", "2017")
			Convey("There should be no errors", func() {
				So(err, ShouldEqual, nil)
			})
		})

	})

}
