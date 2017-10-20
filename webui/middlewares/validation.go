package middlewares



import (
	"bytes"
	"net/http"

	validator "gopkg.in/bluesuncorp/validator.v9"

	"encoding/json"

	"github.com/labstack/echo"
)

// ValidatorMiddleware  ...
type ValidatorMiddleware struct {
	Target   interface{}
	Validate *validator.Validate
}

// Handler ...
func (m *ValidatorMiddleware) Handler(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {

		// request body
		body := c.Request().Body()

		//new buffer
		buf := new(bytes.Buffer)

		//store body into buffer
		_, err := buf.ReadFrom(body)
		if err != nil {
			return err
		}

		//Unmarshal
		if err = json.Unmarshal(buf.Bytes(), &m.Target); err != nil {
			return err
		}

		//validate
		errs := m.Validate.Struct(m.Target)
		if errs != nil {
			return c.JSON(http.StatusNotAcceptable, errs)
		}

		// send back to the request body
		c.Request().SetBody(buf)

		return next(c)
	}
}