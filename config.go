package v2gen

type Config interface {
	Config() map[string]string
}

//// Render renders configs with template
//// important config should reorder after
//func Render(tpl string, configs ...Config) (io.Reader, error) {
//	config := make(map[string]string)
//	for i := range configs {
//		for k, v := range configs[i].Config() {
//			config[k] = v
//		}
//	}
//
//	parse, err := template.New("default").Parse(tpl)
//	if err != nil {
//		return nil, err
//	}
//	buf := new(bytes.Buffer)
//	err = parse.Execute(buf, config)
//	if err != nil {
//		return nil, err
//	}
//	return buf, nil
//}
