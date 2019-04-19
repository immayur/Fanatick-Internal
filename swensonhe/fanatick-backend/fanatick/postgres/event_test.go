package postgres_test

import (
	"fmt"
	"os"
	"time"

	"github.com/swensonhe/fanatick-backend/fanatick"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/swensonhe/fanatick-backend/fanatick/postgres"
)

var _ = Describe("EventDB", func() {
	var db *EventDB

	BeforeEach(func() {
		db = &EventDB{
			DB: NewDB(os.Getenv("DB_URL")),
		}
	})

	AfterEach(func() {
		db.Close()
	})

	Describe("Create", func() {
		var tx *EventTx

		BeforeEach(func() {
			tx = db.BeginTx().(*EventTx)
		})

		AfterEach(func() {
			tx.Rollback()
		})

		It("creates an event", func() {
			event := &fanatick.Event{
				Name:    "Event Name",
				StartAt: time.Now(),
			}

			err := tx.Create(event)
			Expect(err).ToNot(HaveOccurred())
		})
	})

	Describe("Update", func() {
		var tx *EventTx

		BeforeEach(func() {
			tx = db.BeginTx().(*EventTx)
		})

		AfterEach(func() {
			tx.Rollback()
		})

		It("updates an event", func() {
			event := &fanatick.Event{
				Name:    "Event Name",
				StartAt: time.Now(),
			}

			tx.Create(event)

			event.Name = "Updated Event Name"

			err := tx.Update(event)
			Expect(err).ToNot(HaveOccurred())
			Expect(event.UpdatedAt).ToNot(BeNil())
		})
	})

	Describe("Delete", func() {
		var tx *EventTx

		BeforeEach(func() {
			tx = db.BeginTx().(*EventTx)
		})

		AfterEach(func() {
			tx.Rollback()
		})

		It("deletes an event", func() {
			event := &fanatick.Event{
				Name:    "Event Name",
				StartAt: time.Now(),
			}

			tx.Create(event)
			err := tx.Delete(event.ID)
			Expect(err).ToNot(HaveOccurred())

			_, err = db.Get(event.ID)
			Expect(err).To(Equal(fanatick.ErrorNotFound))
		})
	})

	Describe("Query", func() {
		var tx *EventTx
		var records = []*fanatick.Event{
			&fanatick.Event{
				Name:    "Event 1",
				StartAt: time.Now(),
			},
			&fanatick.Event{
				Name:    "Event 2",
				StartAt: time.Now(),
			},
		}

		BeforeEach(func() {
			tx = db.BeginTx().(*EventTx)
			for _, record := range records {
				tx.Create(record)
				fmt.Println(record.ID)
			}
			tx.Commit()
		})

		AfterEach(func() {
			tx = db.BeginTx().(*EventTx)
			for _, record := range records {
				tx.Delete(record.ID)
			}
			tx.Commit()
		})

		It("returns a list of events", func() {
			events, err := db.Query(nil)
			Expect(err).ToNot(HaveOccurred())
			Expect(events).ToNot(BeNil())
			Expect(len(events)).To(Equal(2))
		})

		Context("when a before ID is provided", func() {
			It("only returns events before the given ID", func() {
				page1, _ := db.Query(map[fanatick.EventQueryParam]interface{}{
					fanatick.EventQueryParamLimit: 1,
				})

				page2, err := db.Query(map[fanatick.EventQueryParam]interface{}{
					fanatick.EventQueryParamBefore: page1[0].ID,
				})

				Expect(err).ToNot(HaveOccurred())
				Expect(page2).ToNot(BeNil())
				Expect(page2).To(HaveLen(1))
				Expect(page2[0].ID).ToNot(Equal(page1[0].ID))
			})
		})

		Context("when a limit is provided", func() {
			It("limits the number of events returned", func() {
				events, err := db.Query(map[fanatick.EventQueryParam]interface{}{
					fanatick.EventQueryParamLimit: 1,
				})
				Expect(err).ToNot(HaveOccurred())
				Expect(events).ToNot(BeNil())
				Expect(events).To(HaveLen(1))
			})
		})
	})

	Describe("Get", func() {
		var tx *EventTx
		var record = &fanatick.Event{
			Name:    "Event Name",
			StartAt: time.Now(),
		}

		BeforeEach(func() {
			tx = db.BeginTx().(*EventTx)
			tx.Create(record)
			tx.Commit()
		})

		AfterEach(func() {
			tx = db.BeginTx().(*EventTx)
			tx.Delete(record.ID)
			tx.Commit()
		})

		It("returns an event", func() {
			event, err := db.Get(record.ID)
			Expect(err).ToNot(HaveOccurred())
			Expect(event).ToNot(BeNil())
			Expect(event.ID).To(Equal(record.ID))
		})

		Context("when the event does not exist", func() {
			It("returns nil", func() {
				_, err := db.Get("")
				Expect(err).To(HaveOccurred())
				Expect(err).To(Equal(fanatick.ErrorNotFound))
			})
		})
	})
})
