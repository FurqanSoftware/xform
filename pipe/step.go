package pipe

import "gopkg.in/yaml.v3"

type Step interface {
	Do(Values) error
}

var Steps = map[string]func(*yaml.Node) (Step, error){
	"dump":       NewDump,
	"email":      NewEmail,
	"honeypot":   NewHoneypot,
	"log":        NewLog,
	"mattermost": NewMattermost,
	"parse":      NewParse,
	"unspam":     NewUnspam,
	"redirect":   NewRedirect,
}
