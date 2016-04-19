package servertime

import (
	"encoding/json"
	"testing"
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
	It("Marshals/unmarshals a time value idempotently", func() {
		happyNewYear1970 := ServerTime{time.Unix(1, 0).UTC(), false}
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

		It("Unmarshals as the local system's time", func() {
			setMe := ServerTime{SetTime: true}
			marshaled, err := json.Marshal(setMe)
			Expect(err).To(BeNil())

			var fromJSON ServerTime
			err = json.Unmarshal(marshaled, &fromJSON)
			Expect(err).To(BeNil())

			expectedTime := ServerTime{mockTime.UTC(), false}
			Expect(fromJSON).To(Equal(expectedTime))
		})
	})
})

func TestServertime(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Servertime Suite")
}
