// Поиск и работа с аккаунтами в Shikimori.
package shikimori

import (
	"github.com/devmoded/Terebi/internal/backend/defaults"
	types "github.com/devmoded/Terebi/internal/backend/types/shikimori"
	"github.com/devmoded/Terebi/internal/version"
)

type Search struct {
	// Title. Название аниме по которому будет выполнен поиск в базе Shikimori.
	Title string
	// Limit. Максимальное количество результатов в ответе.
	Limit int
	// Config. Конфигурация Shikimori.
	Config *Configuration
}

type Configuration struct {
	// Name. Название клиента.
	Name string
	// Token. Токен Shikimori необходимый для авторизации.
	Token string
}

func NewSearch() Search {
	return Search{
		Limit: defaults.ShikimoriLimit,
		Config: &Configuration{
			Name: version.Name,
		},
	}
}

func (s Search) WithTitle(title string) Search {
	s.Title = title
	return s
}

func (s Search) WithLimit(limit int) Search {
	s.Limit = limit
	return s
}

func (s Search) WithToken(token string) Search {
	s.Config.Token = token
	return s
}

func (s Search) Execute() types.Response
