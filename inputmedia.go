package tgbotapi

type inputMedia interface {
	medias() []fileInfo
	MarshalJSON() ([]byte, error)
}
