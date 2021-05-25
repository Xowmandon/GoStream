package ApexProfile

import (
	"time"

)

const NumOfLegends = 16;

type ApexProfile struct {
	Data struct {

		PlatformInfo PlatformData `json:"platformInfo"`
		UserInfo Userinfo         `json:"userInfo"`

		Metadata struct {
			Currentseason     int         `json:"currentSeason"`
			Activelegend      string      `json:"activeLegend"`
			Activelegendname  string      `json:"activeLegendName"`
			Activelegendstats interface{} `json:"activeLegendStats"`
		} `json:"metadata"`

		Segments []struct {
			Type  string `json:"type"`

			LifetimeAttributes struct {} `json:"attributes,omitempty"`

			LifetimeMetadata struct {
				Name string `json:"name"`
			} `json:"metadata,omitempty"`

			Expirydate time.Time         `json:"expiryDate"`
			GlobalStats GlobalStats      `json:"stats,omitempty"`
			
			Legends [NumOfLegends]Legend

		} `json:"segments"`

		Availablesegments []struct {
			Type       string `json:"type"`
			Attributes struct {
			} `json:"attributes"`
			Metadata struct {
			} `json:"metadata"`
		} `json:"availableSegments"`

		Expirydate time.Time `json:"expiryDate"`

	} `json:"data"`
}