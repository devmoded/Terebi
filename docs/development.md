# Реализация

Для реализации используется язык Go со следующими библиотеками:

- `gotk4`, `gotk4-adwaita`: https://github.com/diamondburned/gotk4
  Для создания интерфейса. Подробно интерфейс описан в разделе [Интерфейс](ui.md).

- `go-mpv`: https://github.com/gen2brain/go-mpv
  Для создания плеера. Подробно плеер описан в разделе [Плеер](player.md).

- `go-ytdlp`: https://github.com/lrstanley/go-ytdlp
  Для загрузки видео из m3u8 с поддержкой продолжения с того же места при её прерывании.

- `goshikimori`: https://github.com/ghostemeow/goshikimori
  Для получения информации об аниме из Shikimori и профиля.

- `kodik-api`: https://github.com/hksm-app/kodik-api
  Для получения переводов, сезонов и эпизодов из Kodik, с ссылками на плеер. Для резолвинга необходим `BatogiX/kodik` (указан ниже).

Дополнительные проекты, необходимые для Terebi:

- `kodik`: https://github.com/BatogiX/kodik
  Для резолвинга ссылок на плеер Kodik. Будет использоваться через CLI.

## Примеры и ориентиры

Для реализации можно пользоваться примерами:

- `Cine`: https://github.com/diegopvlk/Cine
  Python плеер построенный на libmpv с GTK4 + Libadwaita. Может быть полезен для реализации аналогичного на Go.

План развития проекта описан в разделе [Этапы разработки](roadmap.md).
