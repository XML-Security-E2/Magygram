package handler

type AppHandler interface {
	UserHandler
	CollectionsHandler
	HighlightsHandler
}
