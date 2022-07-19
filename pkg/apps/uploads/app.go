package uploads

import (
	"context"
	"errors"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	
	"github.com/jfyne/live"
)

const (
	validate = "validate"
	save     = "save"
)

// customError formats upload validation errors.
func customError(u *live.Upload, err error) string {
	msg := []string{}
	if u.Name != "" {
		msg = append(msg, u.Name)
	}
	
	switch {
	case errors.Is(err, live.ErrUploadTooLarge):
		msg = append(msg, "This is a custom too large message: "+err.Error())
	case errors.Is(err, live.ErrUploadTooManyFiles):
		msg = append(msg, "This is a custom too many files message: "+err.Error())
	default:
		msg = append(msg, err.Error())
	}
	return strings.Join(msg, " - ")
}

func onSave(ctx context.Context, s live.Socket, p live.Params) (interface{}, error) {
	m := newState(s)
	
	live.ConsumeUploads(s, "photos", func(upload *live.Upload) error {
		file, err := upload.File()
		if err != nil {
			return err
		}
		
		// When we are done close the file, and remove it from staging.
		defer func() {
			_ = file.Close()
			_ = os.Remove(file.Name())
		}()
		
		// Create a new file in our static directory to copy the staged file into.
		dst, err := os.Create(filepath.Join("uploads", "static", upload.Name))
		if err != nil {
			return err
		}
		defer dst.Close()
		
		// Do the copy
		if _, err := io.Copy(dst, file); err != nil {
			return err
		}
		
		// Record the name of the file so we can show the link to it.
		m.Uploads = append(m.Uploads, upload.Name)
		
		return nil
	})
	
	return m, nil
}

func onValidate(ctx context.Context, s live.Socket, p live.Params) (interface{}, error) {
	m := newState(s)
	live.ValidateUploads(s, p)
	
	return m, nil
}

func NewEngine(ctx context.Context, store live.HttpSessionStore) http.Handler {
	max := int64(2 * 1024 * 1024)
	h := live.NewHandler(withRenderConfig())
	h.HandleMount(onMount)
	h.HandleEvent(validate, onValidate)
	h.HandleEvent(save, onSave)
	
	return live.NewHttpHandler(store, h, live.WithMaxUploadSize(max))
}

func withRenderConfig() live.HandlerConfig {
	t, err := template.New("root.html").
		Funcs(template.FuncMap{"customError": customError}).
		ParseFiles("root.html", "pkg/apps/uploads/view.html")
	if err != nil {
		log.Fatal(err)
	}
	
	return live.WithTemplateRenderer(t)
}
