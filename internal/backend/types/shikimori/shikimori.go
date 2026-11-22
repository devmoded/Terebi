// Типы данных для Shikimori.
package shikimori

type Response struct {
	// Animes. Аниме полученные из Shikimori.
	Animes []Anime
}

type Anime struct {
	Id            int          `json:"id"`
	Name          string       `json:"name"`
	Russian       string       `json:"russian"`
	Rating        string       `json:"rating"`
	Episodes      int          `json:"episodes"`
	EpisodesAired int          `json:"episodes_aired"`
	NextEpisodeAt string       `json:"next_episode_at"`
	Score         float32      `json:"score"`
	Poster        Poster       `json:"poster"`
	Genres        []Genre      `json:"genres"`
	Studios       []Studio     `json:"studios"`
	Screenshots   []Screenshot `json:"screenshots"`
	Videos        []Video      `json:"video"`
}

type Studio struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	ImageUrl string `json:"image_url"`
}

type Poster struct {
	Id          string `json:"id"`
	OriginalUrl string `json:"original_url"`
	MainUrl     string `json:"main_url"`
}

type Screenshot struct {
	Id          string `json:"id"`
	OriginalUrl string `json:"original_url"`
	X166Url     string `json:"x166url"`
	X332Url     string `json:"x332url"`
}

type Video struct {
	Id   string `json:"id"`
	Url  string `json:"url"`
	Name string `json:"name"`
	Kind string `json:"kind"`
}

type Genre struct {
	Id      string `json:"id"`
	Name    string `json:"name"`
	Russian string `json:"russian"`
	Kind    string `json:"kind"`
}
