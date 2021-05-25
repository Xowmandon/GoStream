package ApexProfile

import(


)

type GlobalStats struct {
	Level struct {
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
	} `json:"level"`
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
	Rankscore struct {
		Rank            interface{} `json:"rank"`
		Percentile      interface{} `json:"percentile"`
		Displayname     string      `json:"displayName"`
		Displaycategory string      `json:"displayCategory"`
		Category        interface{} `json:"category"`
		Metadata        struct {
			Iconurl  string `json:"iconUrl"`
			Rankname string `json:"rankName"`
		} `json:"metadata"`
		Value        float64 `json:"value"`
		Displayvalue string  `json:"displayValue"`
		Displaytype  string  `json:"displayType"`
	} `json:"rankScore"`

	GlobalSeasonData StatsBySeason
} 