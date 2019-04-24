package fanatick

const OfferStatusIDPending = 10
const OfferStatusIDAccepted = 50
const OfferStatusIDDeclined = 90
const OfferStatusIDCancelled = 99

type OfferStatus struct {
	OfferStatusID int    `json:"offer_status_id"`
	StatusKey     string `json:"key"`
	Description   string `json:"description"`
}

// Offer is a event offer.
type Offer struct {
	ID            string  `json:"offer_id"`
	Amount        float64 `json:"amount"`
	SeatID        int     `json:"seat_id"`
	EventID       int     `json:"event_id"`
	OfferStatusID int     `json:"offer_status_id"`

	CommonTSFields
}

// OfferGetter is the interface that wraps an offer get request.
type OfferGetter interface {
	Get(id string) (*Offer, error)
}

// OfferQueryer is the interface that wraps an offer query request.
type OfferQueryer interface {
	//Query(params map[OfferQueryParam]interface{}) ([]*Offer, error)
}

// OfferCreator is the interface that wraps an offer creation request.
type OfferCreator interface {
	Create(offer *Offer) error
}

// OfferUpdater is the interface that wraps an offer update request.
type OfferUpdater interface {
	Update(offer *Offer) error
}

// OfferDeleter is the interface that wrapts an offer delete request.
type OfferDeleter interface {
	Delete(id string) error
}

// OfferTxBeginner is the interface that wraps an offer transaction starter.
type OfferTxBeginner interface {
	BeginTx() OfferTx
}

// OfferStore defines the operations of an offer store.
type OfferStore interface {
	OfferGetter
	OfferQueryer
	OfferTxBeginner
}

// OfferTx defines the operations that may be performed on an offer update transaction.
type OfferTx interface {
	OfferCreator
	OfferUpdater
	OfferDeleter
}

type NegotiationOffer struct {
	OfferID  int     `json:"offer_id"`
	BySeller bool    `json:"by_seller"`
	ByBuyer  bool    `json:"by_buyer"`
	Amount   float64 `json:"amount"`

	CommonTSFields
}
