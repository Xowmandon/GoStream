package ApexProfile

import (



)

type Userinfo struct {
	Userid          interface{}   `json:"userId"`
	Ispremium       bool          `json:"isPremium"`
	Isverified      bool          `json:"isVerified"`
	Isinfluencer    bool          `json:"isInfluencer"`
	Ispartner       bool          `json:"isPartner"`
	Countrycode     interface{}   `json:"countryCode"`
	Customavatarurl interface{}   `json:"customAvatarUrl"`
	Customherourl   interface{}   `json:"customHeroUrl"`
	Socialaccounts  []interface{} `json:"socialAccounts"`
	Pageviews       int           `json:"pageviews"`
	Issuspicious    interface{}   `json:"isSuspicious"`
} 