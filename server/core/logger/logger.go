package logger

import (
	"github.com/gookit/slog"
	"github.com/gookit/slog/handler"
	"github.com/gookit/slog/rotatefile"
)

var Blogger *slog.SugaredLogger

func init() {
	Blogger = slog.NewStdLogger()
	myTemplate := "[GIN-REACT-ADMIN] [{{datetime}}] [{{level}}] {{message}}\n"
	textFormatter := &slog.TextFormatter{TimeFormat: "2006-01-02 15:04:05", EnableColor: true}
	textFormatter.SetTemplate(myTemplate)

	textFormatterWithOutColor := &slog.TextFormatter{TimeFormat: "2006-01-02 15:04:05", EnableColor: false}
	textFormatterWithOutColor.SetTemplate(myTemplate)
	// Set up file rotation handlers
	h1 := handler.MustRotateFile("logs/error.log", rotatefile.EveryDay, handler.WithLogLevels(slog.DangerLevels))
	h2 := handler.MustRotateFile("logs/info.log", rotatefile.EveryDay, handler.WithLogLevels(slog.NormalLevels))
	h1.SetFormatter(textFormatterWithOutColor)
	h2.SetFormatter(textFormatterWithOutColor)
	Blogger.PushHandlers(h1, h2)

	Blogger.Formatter = textFormatter
}
