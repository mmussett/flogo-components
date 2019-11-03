package ems

type Settings struct {
	ServerURL   string `md:"serverURL,required"`
	Destination string `md:"destination"`
	Username    string `md:"username"`
	Password    string `md:"password"`
}

type Output struct {
	Data interface{} `md:"data"`
}

func (o *Output) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"data": o.Data,
	}
}

func (o *Output) FromMap(values map[string]interface{}) error {

	o.Data = values["data"]

	return nil
}
