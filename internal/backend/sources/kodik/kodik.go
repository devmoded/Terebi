// Поиск в каталоге Kodik.
package kodik

import (
	"github.com/devmoded/Terebi/internal/backend/defaults"
	types "github.com/devmoded/Terebi/internal/backend/types/catalog"
)

type Search struct {
	// ID. ID в базе Shikimori по которому будет выполнен поиск в каталоге.
	ID int
	// Title. Название аниме. Можно переопределить через OverrideTitle().
	// По умолчанию берётся из каталога.
	Title string
	// Limit. Максимальное количество результатов в ответе.
	Limit int
}

func NewSearch() Search {
	return Search{
		Limit: defaults.KodikLimit,
	}
}

func (s Search) WithLimit(limit int) Search {
	s.Limit = limit
	return s
}

func (s Search) WithShikimoriID(id int) Search {
	s.ID = id
	return s
}

func (s Search) OverrideTitle(title string) Search {
	s.Title = title
	return s
}

func (s Search) Execute() types.Response
