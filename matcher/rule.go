package matcher

type MatcherRule struct {
	KeyPath string
	RegExp  string
	Matched bool
}
