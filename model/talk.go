package model

// Talk defines a devopsdays event's talk
type Talk struct {
	Name         string
	Title        string
	Description  string `toml:"description,omitempty"`
	Speakers     []Speakers
	YouTube      string `toml:"youtube,omitempty"`
	Vimeo        string `toml:"vimeo,omitempty"`
	Speakerdeck  string `toml:"speakerdeck,omitempty"`
	Slideshare   string `toml:"slideshare,omitempty"`
	Googleslides string `toml:"googleslides,omitempty"`
	PDF          string `toml:"pdf,omitempty"`
	Slides       string `toml:"slides,omitempty"`
}

// Speakers is the array of speakers on a talk
type Speakers struct {
	Name string
}
