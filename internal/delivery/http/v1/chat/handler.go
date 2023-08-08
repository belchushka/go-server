package chat

import (
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"strings"
	"time"
	"github.com/gin-gonic/gin"
	"gopkg.in/loremipsum.v1"
	"lab18.matbea.xyz/savva/test-backend/internal/delivery/http/v1/errors"
)

type chatHandler struct{}

func NewChatHandler() *chatHandler {
	return &chatHandler{}
}

type SendMessageDto struct {
	Message string `binding:"required" json:"message"`
}

func sendChunkedLorem(ctx *gin.Context) {
	lorem := loremipsum.New().Words(rand.Intn(20) + 5)

	reader := strings.NewReader(lorem)

	buf := make([]byte, 1)

	ctx.Header("Transfer-Encoding", "chunked")

  controller := http.NewResponseController(ctx.Writer)

	for {
		controller.Flush()

		_, err := reader.Read(buf)

		if err == io.EOF {
			jsonBytes, _ := json.Marshal(map[string]interface{}{
				"status": "done",
				"value":  nil,
			})

			ctx.Writer.Write(jsonBytes)
			ctx.Status(200)
			break
		}

		jsonBytes, err := json.Marshal(map[string]interface{}{
			"status": "content",
			"value":  string(buf),
		})

		if err != nil {
			continue
		}

		ctx.Writer.Write(jsonBytes)
	}
}

func (c *chatHandler) SendMessage(ctx *gin.Context) {
	dto := &SendMessageDto{}

	if err := ctx.Bind(dto); err != nil {
		fmt.Println(err)
		errors.AbortWithValidationErrors(ctx, err)
		return
	}

	time.Sleep(time.Second * 2)

	sendChunkedLorem(ctx)
}

func (c *chatHandler) RegisterRoutes(group *gin.RouterGroup) {

	g := group.Group("/chat")

	g.POST("/send-message", c.SendMessage)
}
