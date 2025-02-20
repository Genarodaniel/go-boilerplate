package healthcheck

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestHandleHealthcheck(t *testing.T) {
	gin.SetMode(gin.TestMode)
	Router(&gin.Default().RouterGroup)
	path := "/v1/healthcheck"

	t.Run("Should return health", func(t *testing.T) {
		handler := NewHealthcheckHandler()

		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)
		ctx.Request = httptest.NewRequest(http.MethodPost, path, nil)
		handler.HealthCheck(ctx)

		response, _ := io.ReadAll(w.Body)

		assert.Equal(t, http.StatusOK, w.Code)
		assert.Equal(t, "{\"message\":\"Health!\"}", string(response))

	})

}
