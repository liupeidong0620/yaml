package yaml_test

import (
	"fmt"
	"github.com/Netpas/yaml"
	"reflect"
	"testing"
)

var data0 = `
anchors:
  list:
    - &CENTER { "x": 1, "y": 2 }
    - &LEFT   { "x": 0, "y": 2 }
    - &BIG    { "r": 10 }
    - &SMALL  { "r": 1 }

# All the following maps are equal:

plain:
  # Explicit keys
  "x": 1
  "y": 2
  "r": 10
  label: center/big

mergeOne:
  # Merge one map
  << : *CENTER
  "r": 10
  label: center/big

mergeMultiple:
  # Merge multiple maps
  << : [ *CENTER, *BIG ]
  label: center/big

override:
  # Override
  << : [ *BIG, *LEFT, *SMALL ]
  "x": 1
  label: center/big # override lablel data

shortTag:
  # Explicit short merge tag
  !!merge "<<" : [ *CENTER, *BIG ]
  label: center/big

longTag:
  # Explicit merge long tag
  !<tag:yaml.org,2002:merge> "<<" : [ *CENTER, *BIG ]
  label: center/big

inlineMap:
  # Inlined map 
  << : {"x": 1, "y": 2, "r": 10}
  label: center/big

inlineSequenceMap:
  # Inlined map in sequence
  << : [ *CENTER, {"r": 10} ]
  label: center/big
`
var data = `
# liu test
database:
  # Database path
  path: 'path/to/db'

  # test_1
  sl:
  - one
  # Two
  - two
# test_2
netpas: letvpn

# test_3
`

var data1 = `
# liu test
database:
  # Database path
  path: 'path/to/db'
# netpas data
netpas:
  # app info
  app: letvpn
# qq data
teng:
  # app info1
  # app info
  app:
   # in data
    in:
    # qq data
      qq: true

# test_34
`

var data2 = `
# hhhhll dad
- tabitha:
    name: Tabitha Bitumen
    job: Developer
    # data liu
    liu:
      # netpas data1
      # netpas data
      netpas:
        # app data
        app: letvpn
    # skills data
    skills:
    - lisp
    # for data
    - fortran
    - erlang
# test_3
# test_5
- tabitha:
    # test_4
    netpas: letvpn

# test _1 
# test _2

`

var data3 = `
# teng data
teng:
  xun:
    app: uc
# test_1
`

var data4 = `
# zz
# start
`

func TestComment(t *testing.T) {
	b := []byte(data2)
	yaml.DefaultCommentsEnable = true
	yaml.DefaultMapType = reflect.TypeOf(yaml.MapSlice{})
	var doc interface{}
	err := yaml.Unmarshal(b, &doc)
	if err != nil {
		t.Errorf("%v", err)
	}

	ob, err := yaml.Marshal(doc)
	if err != nil {
		t.Errorf("%v", err)
	}
	fmt.Println("---------------------------------------------------")
	fmt.Println(doc)
	fmt.Println("+++++++++++++++++++++++++++++++++++++++++++++++++++")
	fmt.Println(string(ob))
}
