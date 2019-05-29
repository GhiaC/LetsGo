package messages

import "github.com/aoacademy/letsgo/models"

var BadRequest = models.Response{Description: "Bad Request!"}

var ServiceUnavailable = models.Response{Description: "Service Unavailable. Try Again!"}
