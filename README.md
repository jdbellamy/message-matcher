# message-matcher

A go library for defining and applying pattern matching rules against arbitrary JSON messages. Intended for situations where the message structure can be dissimilar between instances or the distinguishing elements of the message may be deeply nested.           

## todos
- [x] apply regex checks to matching logic
- [ ] messaging format codecs
- [x] rethink MatchAll (maybe MatchAny?)
- [ ] consider a more meaningful return type for Message.Value()

## example usage

Create a file `messages.json` in the current directory containing the following.

```json
[{
  "type": "INFO",
  "message": "test message one",
  "timestamp": 12345678,
  "details": {
    "foo1": "bar1",
    "depth2": {
      "foo2": "bar2"
    }
  }
},
{
  "type": "INFO",
  "message": "info message two",
  "timestamp": 12345679,
  "details": {
    "foo1": "bar1",
    "depth2": {
      "foo2": "bar2"
    }
  }
},
{
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
}]
```

Create a file called `.matcher.yml` in the current directory containing the following.

```yaml
example:
  file: messages.json

matches:
- type: INFO
  details.depth2.foo2: bar2
- type: WARN
  details.depth2.depth3.foo3: bar3
- type: DEBUG
- type: WARN
  message: test message three
```

Create a file called `main.go` in the current directory containing the following.

```go
package main

import (
	"io/ioutil"
	"log"

	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
	"github.com/tidwall/gjson"

	. "gitlab.com/jdbellamy/message-matcher/matcher"
)

func main() {
	matchers := loadConfig()
	messages := loadMessages()
	for i, msg := range messages {
		matched := matchers.MatchAny(msg)
		log.Printf("[%d: %t]", i, matched)
	}
}

// Loads the `messages.json` fixture data
func loadMessages() []Message {
	filename := viper.GetString("example.file")
	file, _ := ioutil.ReadFile(filename)
	data := gjson.Parse(string(file)).Array()
	var messages []Message
	for _, message := range data {
		msg, _ := NewMessage(message.String())
		messages = append(messages, msg)
	}
	return messages
}

// Loads the `.matcher.yml` config file
func loadConfig() MatcherGroup {
	viper.SetConfigName(".matcher")
	viper.AddConfigPath(".")
	_ = viper.ReadInConfig()
	cfg := MatcherGroupConfig{}
	defs := viper.Get("matchers")
	_ = mapstructure.Decode(defs, &cfg)
	return NewMatcherGroup().FromConfig(cfg)
}
```

Run the following commands inside the current directory.

`go get`

`go run main.go`

The output should resemble the following:
```shell
2019/06/08 04:07:34 [0: true]
2019/06/08 04:07:34 [1: true]
2019/06/08 04:07:34 [2: true]
```

## path syntax
This project uses [gjson](https://github.com/tidwall/gjson) for searching JSON elements.  The syntax for which can be found [here](https://github.com/tidwall/gjson#path-syntax).

## regular expression matching
PCRE style regular expressions can be used to match element values.

#### example regex usage:
```go
cfg := matcher.MatcherGroupConfig{{
    "type": "WARN",
    "details.depth2.depth3.foo3": "\\bbar[0-9]\\b",
}}
```
_NOTE: All forward slashes must be escaped._