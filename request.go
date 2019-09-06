package goauth42

import (
	"net/http"
	"io"
)

func (f *goauth42)Do(urls string, body io.Reader, types string) (resp *http.Response, err error){
	r , err :=  http.NewRequest(types, urls, body);
	if err != nil{
		return
	}
	if body != nil {
		r.Header.Set("Content-Type", "application/json")
	}
	return f.Client.Do(r);
}

func (f *goauth42)Post(url string, body io.Reader) (resp *http.Response, err error){
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

func (f *goauth42)Patch(url string, body io.Reader) (resp *http.Response, err error){
		return f.Do(url, body,"patch");
}

func (f *goauth42)Put(url string, body io.Reader) (resp *http.Response, err error){
		return f.Do(url, body, "put");
}

func (f *goauth42)Delete(url string) (resp *http.Response, err error){
		return f.Do(url,nil,"delete")
}
