package aconfighcl

import (
	"github.com/hashicorp/hcl"
	"io/ioutil"
)

// Decoder of HCL files for aconfig.
type Decoder struct{}

// New HCL decoder for aconfig.
func New() *Decoder { return &Decoder{} }

// DecodeFile implements aconfig.FileDecoder.
func (d *Decoder) DecodeFile(filename string) (map[string]interface{}, error) {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	f, err := hcl.ParseBytes(b)
	if err != nil {
		return nil, err
	}

	var raw map[string]interface{}
	if err := hcl.DecodeObject(&raw, f); err != nil {
		return nil, err
	}

	res := map[string]interface{}{}

	for key, value := range raw {
		//flatten("", key, value, res)
		res[key] = value
	}
	return res, nil
}

// copied and adapted from aconfig/utils.go
//
//func flatten(prefix, key string, curr interface{}, res map[string]interface{}) {
//	switch curr := curr.(type) {
//	case []map[string]interface{}:
//		for _, v := range curr {
//			flatten(prefix+key, "", v, res)
//		}
//	case []map[interface{}]interface{}:
//		for k, v := range curr {
//			flatten(prefix+key+".", fmt.Sprint(k), v, res)
//		}
//
//	case map[string]interface{}:
//		for k, v := range curr {
//			flatten(prefix+key+".", k, v, res)
//		}
//
//	case map[interface{}]interface{}:
//		for k, v := range curr {
//			if k, ok := k.(string); ok {
//				flatten(prefix+key+".", k, v, res)
//			}
//		}
//	case []interface{}:
//		b := &strings.Builder{}
//		for i, v := range curr {
//			if i > 0 {
//				b.WriteByte(',')
//			}
//			b.WriteString(fmt.Sprint(v))
//		}
//		res[prefix+key] = b.String()
//	case bool:
//		res[prefix+key] = fmt.Sprint(curr)
//	case string:
//		res[prefix+key] = curr
//	case float64:
//		res[prefix+key] = fmt.Sprint(curr)
//	case int, int8, int16, int32:
//		res[prefix+key] = fmt.Sprint(curr)
//	default:
//		panic(fmt.Sprintf("%s::%s got %T %v", prefix, key, curr, curr))
//	}
//}
