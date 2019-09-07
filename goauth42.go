package  goauth42

import (
		"golang.org/x/oauth2"
		"context"
		"net/http"
       )

type goauth42 struct {
	Ctx context.Context
	Conf *oauth2.Config
	Client *http.Client
	Tok    *oauth2.Token;
	Url string
}

var (
	Endpoint42 = oauth2.Endpoint{
		TokenURL: "https://api.intra.42.fr/oauth/token",
		AuthURL:  "https://api.intra.42.fr/oauth/authorize",
	}
)

func Start(uid, secret, urls string, scope []string)(f goauth42) {
	f.Ctx = context.Background()
	     f.Conf = &oauth2.Config {
			ClientID:     uid,
			ClientSecret: secret,
		      Scopes:       scope,
		      RedirectURL:  urls,
		      Endpoint: Endpoint42,
	     }
	f.Url = "https://api.intra.42.fr/";
	return
 }

 func (f *goauth42)Clients(s string, opts ...oauth2.AuthCodeOption) (er error){
	     f.Tok, er = f.Conf.Exchange(f.Ctx, s, opts...)
	     if er == nil {
		     f.Client = f.Conf.Client(f.Ctx, f.Tok)
	}
	return
}

