package api

import (
	. "../../../webui"
)


type ApiContext struct {
	*Context
	AccessToken string
}
