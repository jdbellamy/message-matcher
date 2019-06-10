package matcher_test

import (
	"testing"

	"gitlab.com/jdbellamy/message-matcher/matcher"
	"gitlab.com/jdbellamy/message-matcher/matcher/mocks"
	
	"github.com/stretchr/testify/assert"
  )

  func TestMessageGroupMatchAll(t *testing.T) {
	mmg := new(mocks.MessageMatcher)
  
	m := map[interface{}]interface{}{
		"type": "WARN",
  		"details.depth2.depth3.foo3": "bar3",
	}

	cfg := matcher.MatcherGroupConfig{m}

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
  
	mmg.AssertExpectations(t)  
  }