package telapi

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestSendSMS(t *testing.T) {

	var (
		err           error
		telapi_helper TelapiHelper
	)

	Convey("Tests when SendSMS method called ", t, func() {

		Convey("Should not have an error, bc correct credentials", func() {
			telapi_helper, err = CreateClient(testing_telapi_sid, testing_telapi_auth_token)

			So(err, ShouldBeNil)
		})

		Convey("Should blow up to no 'To' ", func() {
			_, err = telapi_helper.SendSMS("", "hey", "")

			So(err, ShouldNotBeNil)
		})

		Convey("Should blow up to no 'From' ", func() {
			_, err = telapi_helper.SendSMS("+", "", "")

			So(err, ShouldNotBeNil)
		})

		Convey("Should blow up to no 'Body' ", func() {
			_, err = telapi_helper.SendSMS("372", "hey", "")

			So(err, ShouldNotBeNil)
		})

		Convey("Should have no errors", func() {
			sms, err := telapi_helper.SendSMS(testing_number_to, testing_number_from, `TrapCall New Transcription
Cell Phone   NJ
(848) 210-6084
NEW JERSEY NJ 
(06/30/14 4:05 PM)
I make phone calls because Ronnie talk to you soon.
http://v.trapcall.com/445jccko

fklnfdlkdnldskdfd

fdkldsnskln
`)

			So(err, ShouldBeNil)
			So(sms, ShouldNotBeNil)

			So(sms.Sid, ShouldNotEqual, "")
			So(sms.Body, ShouldNotEqual, "")
			So(sms.To, ShouldNotEqual, "")
			So(sms.From, ShouldNotEqual, "")

		})

	})
}
