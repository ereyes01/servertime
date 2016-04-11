package servertime

import (
	"encoding/json"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var mockTime = time.Unix(1, 0)

type mockTimeNow struct{}

func (m *mockTimeNow) GetCurrentTime() time.Time {
	return mockTime
}

var _ = Describe("Servertime", func() {
	It("Marshals/unmarshals a non-zero time value idempotently", func() {
		happyNewYear1970 := ServerTime(time.Unix(1, 0).UTC())
		marshaled, err := json.Marshal(happyNewYear1970)
		Expect(err).To(BeNil())

		var fromJSON ServerTime
		err = json.Unmarshal(marshaled, &fromJSON)
		Expect(err).To(BeNil())
		Expect(fromJSON).To(Equal(happyNewYear1970))
	})

	Context("When marshaling a timestamp to be set at the server", func() {
		BeforeEach(func() {
			now = new(mockTimeNow)
		})

		AfterEach(func() {
			now = new(realTimeNow)
		})

		It("Unmarshals a zero time into the local system's time", func() {
			zero := ServerTime{}
			marshaled, err := json.Marshal(zero)
			Expect(err).To(BeNil())

			var fromJSON ServerTime
			err = json.Unmarshal(marshaled, &fromJSON)
			Expect(err).To(BeNil())

			expectedTime := ServerTime(mockTime.UTC())
			Expect(fromJSON).To(Equal(expectedTime))
		})
	})
})
