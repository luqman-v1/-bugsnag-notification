package gate

import (
	"bugsnag-notification/entity"
	"bugsnag-notification/repo/telegram"
	"bugsnag-notification/util"
	"encoding/json"

	"github.com/gin-gonic/gin"
)

func Run() {
	r := gin.Default()
	r.POST("/webhook/error", func(c *gin.Context) {
		request, err := c.GetRawData()
		if err != nil {
			return
		}
		var payload entity.Payload
		if err := json.Unmarshal(request, &payload); err != nil {
			err := util.ErrMessage(err.Error())
			if err != nil {
				return
			}
		}
		err = telegram.TelegramRepo{}.PushToChannel(payload)
		if err != nil {
			return
		}
		c.JSON(200, gin.H{
			"message": "success publish to channel",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
