package ApexProfile

import(


)


type StatsBySeason struct {
	Season7Wins struct {
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
	} `json:"season7Wins"`
	Season7Kills struct {
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
	} `json:"season7Kills"`
	Season8Wins struct {
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
	} `json:"season8Wins"`
	Season8Kills struct {
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
	} 
}