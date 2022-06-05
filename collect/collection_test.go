package collect

import (
	"os"
	"strconv"
	"testing"
	"time"
)

var (
	mymap  map[string]time.Time
	toFind time.Time
)

func TestMain(m *testing.M) {
	os.Exit(func() int {
		mymap := make(map[string]time.Time)
		for i := 0; i < 1000; i++ {
			mymap[strconv.Itoa(i)] = time.Now().Add(time.Millisecond + time.Duration(i))
			if i == 500 {
				toFind = mymap[strconv.Itoa(i)]
			}
		}
		return m.Run()
	}())
}

func TestIndexOf(t *testing.T) {
	//index := IndexOf(mymap, toFind)
}
