package matcher_test

import (
	"testing"

	"gitlab.com/jdbellamy/message-matcher/matcher"
	"gitlab.com/jdbellamy/message-matcher/matcher/mocks"

	"github.com/stretchr/testify/assert"
)

func TestMessageGroupMatchAll(t *testing.T) {  
	cfg := matcher.MatcherGroupConfig{{
		"type": "WARN",
		"details.depth2.depth3.foo3": "bar3",
	}}

	mg := matcher.NewMatcherGroup().FromConfig(cfg)

	msg, err := matcher.NewMessage(`{
		"type": "WARN",
		"message": "warning message",
		"timestamp": 12345680,
		"details": {
			"foo1": "bar1",
			"depth2": {
				"foo2": "bar2",
				"depth3": {
					"foo3": "bar3"
				}
			}
		}
	}`)

	assert.NoError(t, err)

	actual := mg.MatchAll(msg)

	assert.Equal(t, true, actual)  
}

func TestMessageGroupMatchAllIsFalse(t *testing.T) {  
	cfg := matcher.MatcherGroupConfig{{
		"type": "WARN",
		"details.foo1": "bar1",
	}}

	mg := matcher.NewMatcherGroup().FromConfig(cfg)

	msg, err := matcher.NewMessage(`{
		"type": "WARN",
		"details": {
			"foo1": "bar!1"
		}
	}`)

	assert.NoError(t, err)

	actual := mg.MatchAll(msg)

	assert.Equal(t, false, actual)
}

func TestMessageGroupMatchAllRegExp(t *testing.T) {  
	cfg := matcher.MatcherGroupConfig{{
		"type": "WARN",
  		"details.depth2.depth3.foo3": "\\bbar[0-9]\\b",
	}}

	mg := matcher.NewMatcherGroup().FromConfig(cfg)

	msg, err := matcher.NewMessage(`{
		"type": "WARN",
		"message": "warning message",
		"timestamp": 12345680,
		"details": {
			"foo1": "bar1",
			"depth2": {
				"foo2": "bar2",
				"depth3": {
			  		"foo3": "bar3"
				}
			}
		}
	}`)

	assert.NoError(t, err)

	actual := mg.MatchAll(msg)

	assert.Equal(t, true, actual)  
}
