package teamcity10

type ContextKey string

var (
	TeamCityCreds   ContextKey = "tc-creds"
	TeamCityToken   ContextKey = "tc-tok"
	TeamCityBaseURL ContextKey = "tc-url"
)
