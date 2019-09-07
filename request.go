package goauth42

import (
	"net/http"
	"encoding/json"
	"bytes"
)

func (f *goauth42)Do(urls string, bodys interface{}, types string) (resp *http.Response, err error){
	var body []byte
	r := new(http.Request)
	if bodys != nil {
		body , err = json.Marshal(bodys);
		if (err != nil) {
			return
		}
	r , err =  http.NewRequest(types, urls, bytes.NewBuffer(body));
	}else {
		r , err =  http.NewRequest(types, urls, nil);
	}
	if err != nil{
		return
	}
	if bodys != nil {
		r.Header.Set("Content-Type", "application/json")
	}
	return f.Client.Do(r);
}

func (f *goauth42)Post(url string, body interface{}) (resp *http.Response, err error){
	return f.Do(url, body,"post");
}

func (f *goauth42)Get(url string,  b ...int) (resp *http.Response, err error){
	if url, err = Pages(url, b...); err != nil {
		return
	}
	return f.Client.Get(url)
}
func (f *goauth42)Head(url string) (resp *http.Response, err error){
	return f.Client.Head(url)
}

func (f *goauth42)Patch(url string, body interface{}) (resp *http.Response, err error){
		return f.Do(url, body,"patch");
}

func (f *goauth42)Put(url string, body interface{}) (resp *http.Response, err error){
		return f.Do(url, body, "put");
}

func (f *goauth42)Delete(url string) (resp *http.Response, err error){
		return f.Do(url,nil,"delete")
}
