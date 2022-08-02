package ssh

import "github.com/project-flogo/core/data/coerce"

type Input struct {
	User     string `md:"user"`
	Password string `md:"password"`
	Host     string `md:"host"`
	Command  string `md:"command"`
}

type Output struct {
	Result string `md:"result"`
}

func (o *Output) FromMap(values map[string]interface{}) error {
	var err error
	o.Result, err = coerce.ToString(values["result"])
	if err != nil {
		return err
	}

	return nil
}

func (i *Input) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"user":     i.User,
		"password": i.Password,
		"host":     i.Host,
		"command":  i.Command,
	}
}

func (i *Input) FromMap(values map[string]interface{}) error {
	var err error

	i.User, err = coerce.ToString(values["user"])
	if err != nil {
		return err
	}
	i.Password, err = coerce.ToString(values["password"])
	if err != nil {
		return err
	}
	i.Host, err = coerce.ToString(values["host"])
	if err != nil {
		return err
	}
	i.Command, err = coerce.ToString(values["command"])
	if err != nil {
		return err
	}

	return nil
}

func (o *Output) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"result": o.Result,
	}
}
