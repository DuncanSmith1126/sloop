package main

import (
	"encoding/json"
	"github.com/golang/glog"
	"github.com/salesforce/sloop/pkg/sloop/queries"
	"github.com/salesforce/sloop/pkg/sloop/store/typed"
	"net/url"
	"time"
)

var QueryName = "Example"
var Query = ExampleQuery
var IsTimeline = true

func ExampleQuery(params url.Values, tables typed.Tables, startTime time.Time, endTime time.Time, requestId string) ([]byte, error) {
	result := queries.TimelineRoot{
		ViewOpt: queries.ViewOptions{Sort:"MostEvents"},
		Rows: []queries.TimelineRow{
			{
				Text:       "I'm an example entry",
				Duration:   600,
				Kind:       "Sample",
				Namespace:  "default",
				Overlays:   []queries.Overlay{},
				ChangedAt:  []int64{},
				NoChangeAt: []int64{},
				StartDate:  startTime.Unix(),
			},
		},
	}
	glog.Infof("Hello my name is plugin")
	return json.Marshal(result)
}



