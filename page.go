package  goauth42

import (
		"strconv"
		"strings"
		"errors"
)

func Page(url string, i int) (s string){
	b := []string{url, "?page[number]=",  strconv.Itoa(i)}
	s = strings.Join(b, "");
	return
}

func Pages(c string, b ...int) (url string ,err error){
	url  = c
	err = nil
	size := len(b)
	if size > 1 {
		err = errors.New("vous devez precisez juste une page")
		return
	}
	if size == 1 {
		e := b[0]
		if (e < 0) {
			err = errors.New("une page ne peut etre negative")
			return
		}
		url = Page(url, e)
	}
	return
}
