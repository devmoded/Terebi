// Типы данных для каталогов аниме.
package catalog

type Response struct {
	// Translations. Переводы из каталога.
	Translations []Translation
}

type Translation struct {
	// ShikimoriID. ID в базе Shikimori.
	ShikimoriID int
	// ID. ID перевода в каталоге.
	ID int
	// Name. Название перевода в каталоге.
	Name string
	// Seasons. Массив сезонов в каталоге.
	Seasons []Season
}

type Season struct {
	// Name. Название сезона в каталоге.
	Name string
	// Episodes. Массив эпизодов в каталоге.
	Episodes []Episode
}

type Episode struct {
	// Num. Номер эпизода в сезоне.
	Num int
	// Name. Название эпизода.
	Name string
	// Link. Ссылка на плеер для эпизода.
	Link string
}
