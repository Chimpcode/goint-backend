package api

import (
	"github.com/kataras/iris/sessions"
	"time"
)

var TimeToExpires =  5*time.Hour
var cookieNameForSessionID = "gointlogin"


var sess = sessions.New(sessions.Config{
	Cookie: cookieNameForSessionID,
	Expires: TimeToExpires,

})

