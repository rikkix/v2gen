package v2gen_test

//
//import (
//	"io/ioutil"
//	"testing"
//
//	"iochen.com/v2gen/v2"
//)
//
//type config map[string]string
//
//func (c *config) Get() map[string]string {
//	return map[string]string(*c)
//}
//
//func TestRender(t *testing.T) {
//	c := &config{
//		"AA": "TEXT1",
//		"BB": "TEXT2",
//	}
//	reader, err := v2gen.Render(template, c)
//	if err != nil {
//		t.Fatal(err)
//	}
//
//	bytes, err := ioutil.ReadAll(reader)
//	if err != nil {
//		t.Fatal(err)
//	}
//	t.Log(string(bytes))
//}
//
//var template = `
//{{.AA}}
//{{.BB}}
//`
