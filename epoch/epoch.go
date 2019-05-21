package epoch

import (
	"encoding/json"
	"fmt"
	"time"
)

// Epoch represents an epoch object with timestamp
// and a human readable representation
type Epoch struct {
	Timestamp int64  `json:"timestamp,omitempty"`
	Human     string `json:"human_date,omitempty"`
}

// List is an Epoch array
type List []*Epoch

// Print prints a single epoch timestamp
func (e *Epoch) Print(format string) {
	en := List{e}
	en.Print(format)
}

// Print prints multiple epoch timestamps
func (en *List) Print(format string) {
	// Print json timestamps
	if format == "json" {
		js, err := json.MarshalIndent(en, "", "	")
		if err != nil {
			panic(err.Error())
		}
		fmt.Printf("\n%s\n", string(js))
		return
	}

	// Print plain timestamps
	for _, e := range *en {
		fmt.Println(
			fmt.Sprintf(
				"\nTimestamp:\t%d\nHuman Date:\t%s",
				e.Timestamp,
				e.Human,
			),
		)
	}
}

// Convert converts a Unix timestamp to its human date
func Convert(timestamp int64) (converted *Epoch) {
	return &Epoch{
		Timestamp: timestamp,
		Human:     time.Unix(timestamp, 0).Format(time.UnixDate),
	}
}

// ConvertN converts multiple Unix timestamps
//  to their human dates
func ConvertN(timestamps []int64) (converted List) {
	converted = []*Epoch{}
	if len(timestamps) == 0 {
		timestamps = append(timestamps, time.Now().Unix())
	}
	for _, timestamp := range timestamps {
		converted = append(converted, Convert(timestamp))
	}
	return converted
}
