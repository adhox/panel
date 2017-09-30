package panel

import (
	"math/rand"
	"time"
)

// util.go is meant for all miscellaneous contructs
// that are not exported and exist only to work
// behind the scenes

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func getStringKeys(m interface{}) (keys []string) {
	switch t := m.(type) {
	case map[string][]interface{}:
		for head := range t {
			if !stringInSlice(head, keys) {
				keys = append(keys, head)
			}
		}
	case map[string][]string:
		for head := range t {
			if !stringInSlice(head, keys) {
				keys = append(keys, head)
			}
		}
	case map[string][]bool:
		for head := range t {
			if !stringInSlice(head, keys) {
				keys = append(keys, head)
			}
		}

	case map[string][]int:
		for head := range t {
			if !stringInSlice(head, keys) {
				keys = append(keys, head)
			}
		}
	case map[string][]float64:
		for head := range t {
			if !stringInSlice(head, keys) {
				keys = append(keys, head)
			}
		}

	case [][]string:

		for h := range t[0] {
			head := t[0][h]
			if !stringInSlice(head, keys) {
				keys = append(keys, head)
			}
		}
	case []map[string]interface{}:
		for i := range t {
			for head := range t[i] {
				if !stringInSlice(head, keys) {
					keys = append(keys, head)
				}
			}
		}
	}
	return
}
<<<<<<< HEAD

func createRandSlice(n, lim int) (list []int) {
	for len(list) < n {
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		// x := rand.Intn(lim)
		x := r.Intn(lim)
		if !intInSlice(x, list) {
			list = append(list, x)
		}
	}
	return
}

func intInSlice(a int, list []int) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}
=======
>>>>>>> b7d41b383db941fc86e7c093029e2c5b2be6c9e8
