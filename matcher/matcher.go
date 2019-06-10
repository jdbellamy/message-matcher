package matcher

//go:generate mockery -name=MessageMatcher

import (
	"github.com/tidwall/gjson"
)

type MessageMatcher interface {
	AddRule(MatcherRule) MessageMatcher
	Match(Message) bool
	Rules() []MatcherRule
}

type matcher struct {
	rules []*MatcherRule
}

func NewMessageMatcher() MessageMatcher {
	return &matcher{}
}

func (m *matcher) AddRule(rule MatcherRule) MessageMatcher {
	m.rules = append(m.rules, &rule)
	return m
}

func (m matcher) Rules() []MatcherRule {
	var safe = []MatcherRule{}
	for _, r := range m.rules {
		safe = append(safe, *r)
	}
	return safe
}

func (m *matcher) Match(msg Message) bool {
	for _, rule := range m.rules {
		var result = gjson.Get(msg.Value(), rule.KeyPath)
		// TODO - apply regexp on values
		if result.String() != rule.RegExp {
			return false
		}
		rule.Matched = true
	}
	return true
}
