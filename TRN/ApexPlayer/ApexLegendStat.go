package ApexProfile

import (


)

type Legend struct {

	Attributes struct {
		ID string `json:"id"`
		} `json:"attributes,omitempty"`
	Metadata struct {
		Name         string `json:"name"`
		Imageurl     string `json:"imageUrl"`
		Tallimageurl string `json:"tallImageUrl"`
		Bgimageurl   string `json:"bgImageUrl"`
		Isactive     bool   `json:"isActive"`
	} `json:"metadata,omitempty"`
	
	Stats struct {
		Kills struct {
			Rank            interface{} `json:"rank"`
			Percentile      float64     `json:"percentile"`
			Displayname     string      `json:"displayName"`
			Displaycategory string      `json:"displayCategory"`
			Category        interface{} `json:"category"`
			Metadata        struct {
			} `json:"metadata"`
			Value        float64 `json:"value"`
			Displayvalue string  `json:"displayValue"`
			Displaytype  string  `json:"displayType"`
		} `json:"kills"`

		StatsBySeasonLegend StatsBySeason

	} `json:"stats,omitempty"`

}
