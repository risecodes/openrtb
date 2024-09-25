package v1_0

import (
	"encoding/json"

	"github.com/risecodes/openrtb/native1"
	"github.com/risecodes/openrtb/native1/request"
)

// Request 1.0 4.1 Native Markup Request Object
//
// The Native Object defines the native1 advertising opportunity available for bid via this bid request.
// It will be included as a JSON-encoded string in the bid request’s imp.native1 field or as a direct JSON object, depending on the choice of the exchange.
// While OpenRTB 2.x officially supports only JSON-encoded strings, many exchanges have implemented a formal object.
// Check with your integration docs.
//
// Note: Prior to VERSION 1.1, the specification could be interpreted as requiring the native1 request to have a root node with a single field “native” that would contain the object above as its value.
// The Native Markup Request Object specified above is now the root object.
type Request struct {
	// Field:
	//   ver
	// Scope:
	//   optional
	// Type:
	//   string
	// Default:
	//   1.0
	// Description:
	//   Version of the Native Markup version in use.
	Ver string `json:"ver,omitempty"`

	// Field:
	//   layout
	// Scope:
	//   recommended in 1.0, deprecated/removed in 1.2
	// Type:
	//   integer
	// Description:
	//   The Layout ID of the native1 ad unit.
	//   See the Table of Layout IDs below.
	Layout native1.Layout `json:"layout,omitempty"`

	// Field:
	//   adunit
	// Scope:
	//   recommended in 1.0, deprecated/removed in 1.2
	// Type:
	//   integer
	// Description:
	//   The Ad unit ID of the native1 ad unit.
	//   See Table of Ad Unit IDs below for a list of supported core ad units.
	AdUnit native1.AdUnit `json:"adunit,omitempty"`

	// Field:
	//   plcmtcnt
	// Scope:
	//   optional
	// Type:
	//   integer
	// Default:
	//   1
	// Description:
	//   The number of identical placements in this Layout.
	//   Refer Section 8.1 Multiplacement Bid Requests for further detail.
	PlcmtCnt int64 `json:"plcmtcnt,omitempty"`

	// Field:
	//   seq
	// Scope:
	//   optional
	// Type:
	//   integer
	// Default:
	//   0
	// Description:
	//   0 for the first ad, 1 for the second ad, and so on.
	//   Note this would generally NOT be used in combination with plcmtcnt - either you are auctioning multiple identical placements (in which case plcmtcnt>1, seq=0) or you are holding separate auctions for distinct items in the feed (in which case plcmtcnt=1, seq=>=1)
	Seq int64 `json:"seq,omitempty"`

	// Field:
	//   assets
	// Scope:
	//   required
	// Type:
	//   array of objects
	// Description:
	//   An array of Asset Objects.
	//   Any bid response must comply with the array of elements expressed in the bid request.
	Assets []request.Asset `json:"assets"`

	// Field:
	//   ext
	// Scope:
	//   optional
	// Type:
	//   object
	// Description:
	// This object is a placeholder that may contain custom JSON agreed to by the parties to support flexibility beyond the standard defined in this specification
	Ext json.RawMessage `json:"ext,omitempty"`
}
