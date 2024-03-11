package logus

import (
	"fmt"
	"log/slog"

	"github.com/darklab8/infra-game-servers/bot/utils/types"
)

func logGroupFiles() slog.Attr {
	return slog.Group("files",
		"file3", GetCallingFile(3),
		"file4", GetCallingFile(4),
	)
}

type slogGroup struct {
	params map[string]string
}

func (s slogGroup) Render() slog.Attr {
	anies := []any{}
	for key, value := range s.params {
		anies = append(anies, key)
		anies = append(anies, value)
	}

	return slog.Group("extras", anies...)
}

type slogParam func(r *slogGroup)

func newSlogGroup(opts ...slogParam) slog.Attr {
	client := &slogGroup{params: make(map[string]string)}
	for _, opt := range opts {
		opt(client)
	}

	return (*client).Render()
}

func TestParam(value int) slogParam {
	return func(c *slogGroup) {
		c.params["test_param"] = fmt.Sprintf("%d", value)
	}
}

func Expected(value any) slogParam {
	return func(c *slogGroup) {
		c.params["expected"] = fmt.Sprintf("%v", value)
	}
}
func Actual(value any) slogParam {
	return func(c *slogGroup) {
		c.params["actual"] = fmt.Sprintf("%v", value)
	}
}

func Regex(value types.RegexExpression) slogParam {
	return func(c *slogGroup) {
		c.params["regex"] = string(value)
	}
}

func OptError(err error) slogParam {
	return func(c *slogGroup) {
		c.params["error_msg"] = fmt.Sprintf("%v", err)
		c.params["error_type"] = fmt.Sprintf("%T", err)
	}
}

func FilePath(value types.FilePath) slogParam {
	return func(c *slogGroup) {
		c.params["file_path"] = string(value)
	}
}

func ChannelID(value types.DiscordChannelID) slogParam {
	return func(c *slogGroup) {
		c.params["channel_id"] = string(value)
	}
}
