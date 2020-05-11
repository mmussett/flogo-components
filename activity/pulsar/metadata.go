package pulsar

import (
	"github.com/project-flogo/core/data/coerce"
)

type Settings struct {
}

type Input struct {
	Url string `md:"url,required"`
	Topic string `md:"topic,required"`
	Payload string `md:"payload,required"`
}

type Output struct {
}

func (r *Input) FromMap(values map[string]interface{}) error {
	url, _ := coerce.ToString(values["url"])
	topic, _ := coerce.ToString(values["topic"])
	payload, _ := coerce.ToString(values["payload"])
	r.Url = url
	r.Topic = topic
	r.Payload = payload
	return nil
}

func (r *Input) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"url": r.Url,
		"topic": r.Topic,
		"payload" : r.Payload,
	}
}

func (o *Output) FromMap(values map[string]interface{}) error {
	return nil
}

func (o *Output) ToMap() map[string]interface{} {
	return map[string]interface{}{
	}
}
