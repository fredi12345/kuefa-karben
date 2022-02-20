package extensions

import (
	"bytes"
	"fmt"
	"io"
	"reflect"
	"strings"

	"github.com/labstack/echo/v4"
)

type BindFunc func(interface{}, echo.Context) error

func (fn BindFunc) Bind(i interface{}, ctx echo.Context) error {
	return fn(i, ctx)
}

func BinderChain(binders ...echo.Binder) echo.Binder {
	return BindFunc(func(i interface{}, c echo.Context) error {
		for _, b := range binders {
			err := b.Bind(i, c)
			if err != nil {
				return err
			}
		}

		return nil
	})
}

func MultipartFileBinder() echo.Binder {
	return BindFunc(func(i interface{}, c echo.Context) error {
		iValue := reflect.Indirect(reflect.ValueOf(i))
		if iValue.Kind() != reflect.Struct {
			return fmt.Errorf("input not is struct pointer, indirect type is %s", iValue.Type().String())
		}

		ctype := c.Request().Header.Get(echo.HeaderContentType)
		if !strings.HasPrefix(ctype, echo.MIMEApplicationForm) && !strings.HasPrefix(ctype, echo.MIMEMultipartForm) {
			return nil
		}

		iType := iValue.Type()
		for i := 0; i < iType.NumField(); i++ {
			fType := iType.Field(i)
			if fType.Type != typeBufferP {
				continue
			}

			fValue := iValue.Field(i)
			if !fValue.CanSet() {
				continue
			}

			tagName := fType.Tag.Get("form")
			fileHeader, err := c.FormFile(tagName)
			if err != nil {
				return fmt.Errorf("could not parse form file '%s': %w", tagName, err)
			}

			formFile, err := fileHeader.Open()
			if err != nil {
				return fmt.Errorf("could not open form file header: %w", err)
			}

			var b bytes.Buffer
			_, err = io.Copy(&b, formFile)
			if err != nil {
				return fmt.Errorf("could not load file into buffer: %w", err)
			}

			fValue.Set(reflect.ValueOf(&b))

			err = formFile.Close()
			if err != nil {
				return fmt.Errorf("could not close file: %w", err)
			}
		}
		return nil
	})
}

var (
	typeBufferP = reflect.TypeOf((*bytes.Buffer)(nil))
)
