package designed

import (
	"context"
	"encoding/json"
	"github.com/fatih/color"
	"io"
	stdLog "log"
	"log/slog"
	"os"
)

type PrettyHandlerOptions struct {
	SlogOpts *slog.HandlerOptions
}

type PrettyHandler struct {
	opts PrettyHandlerOptions
	slog.Handler
	l     *stdLog.Logger
	attrs []slog.Attr
}

func NewPrettySlog() *slog.Logger {
	opts := PrettyHandlerOptions{
		SlogOpts: &slog.HandlerOptions{
			Level: slog.LevelDebug,
		},
	}

	handler := opts.NewPrettyHandler(os.Stdout)

	return slog.New(handler)
}

func (opts PrettyHandlerOptions) NewPrettyHandler(out io.Writer) *PrettyHandler {
	h := &PrettyHandler{
		Handler: slog.NewJSONHandler(out, opts.SlogOpts),
		l:       stdLog.New(out, "", 0),
	}

	return h
}

func (h *PrettyHandler) Handle(_ context.Context, r slog.Record) error {
	level := r.Level.String() + ":"

	switch r.Level {
	case slog.LevelDebug:
		level = color.MagentaString(level)
	case slog.LevelInfo:
		level = color.BlueString(level)
	case slog.LevelWarn:
		level = color.YellowString(level)
	case slog.LevelError:
		level = color.RedString(level)
	}

	fields := make(map[string]interface{})
	r.Attrs(func(a slog.Attr) bool {
		h.processAttr(a, fields)
		return true
	})

	for _, a := range h.attrs {
		h.processAttr(a, fields)
	}

	var b []byte
	var err error
	if len(fields) > 0 {
		b, err = json.MarshalIndent(fields, "", "  ")
		if err != nil {
			return err
		}
	}

	timeStr := r.Time.Format("[15:05:05.000]")
	msg := color.CyanString(r.Message)

	h.l.Println(
		timeStr,
		level,
		msg,
		color.WhiteString(string(b)),
	)

	return nil
}

func (h *PrettyHandler) processAttr(attr slog.Attr, fields map[string]interface{}) {
	if attr.Value.Kind() == slog.KindGroup {
		groupFields := make(map[string]interface{})
		for _, g := range attr.Value.Group() {
			h.processAttr(g, groupFields)
		}
		fields[attr.Key] = groupFields
	} else {
		fields[attr.Key] = attr.Value.Any()
	}
}

func (h *PrettyHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return &PrettyHandler{
		Handler: h.Handler,
		l:       h.l,
		attrs:   attrs,
	}
}
