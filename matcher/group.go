package matcher

//go:generate mockery -name=MatcherGroup

type MatcherGroup interface {
	AddMatcher(MessageMatcher) MatcherGroup
	Matchers() []MessageMatcher
	MatchAny(Message) bool
	FromConfig(MatcherGroupConfig) MatcherGroup
}

type matcherGroup struct {
	matchers []*MessageMatcher
}

func NewMatcherGroup() MatcherGroup {
	return &matcherGroup{}
}

func (g *matcherGroup) AddMatcher(matcher MessageMatcher) MatcherGroup {
	g.matchers = append(g.matchers, &matcher)
	return g
}

type MatcherGroupConfig []map[interface{}]interface{}

func (g *matcherGroup) FromConfig(cfg MatcherGroupConfig) MatcherGroup {
	for _, mc := range cfg {
		var matcher = NewMessageMatcher()
		for path, exp := range mc {
			rule := MatcherRule{
				KeyPath: path.(string),
				RegExp:  exp.(string),
			}
			matcher.AddRule(rule)
		}
		g.AddMatcher(matcher)
	}
	return g
}

func (g matcherGroup) Matchers() []MessageMatcher {
	var safe = []MessageMatcher{}
	for _, m := range g.matchers {
		safe = append(safe, *m)
	}
	return safe
}

func (g matcherGroup) MatchAny(msg Message) bool {
	for _, m := range g.Matchers() {
		if m.Match(msg) {
			return true
		}
	}
	return false
}
