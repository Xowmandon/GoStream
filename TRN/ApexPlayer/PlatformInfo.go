package ApexProfile

import (


)

type PlatformData struct {

	Platformslug           string      `json:"platformSlug"`
	Platformuserid         string      `json:"platformUserId"`
	Platformuserhandle     string      `json:"platformUserHandle"`
	Platformuseridentifier string      `json:"platformUserIdentifier"`
	Avatarurl              string      `json:"avatarUrl"`
	Additionalparameters   interface{} `json:"additionalParameters"`

}