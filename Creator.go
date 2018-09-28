package gotten

import "net/http"

type (
	Builder struct {
		baseUrl string
		cookies []*http.Cookie
		headers http.Header
		client  Client
	}

	Creator struct {
		baseUrl string
		cookies []*http.Cookie
		headers http.Header
		client  Client
	}
)

func NewBuilder() *Builder {
	return &Builder{
		cookies: make([]*http.Cookie, 0),
		headers: make(map[string][]string),
	}
}

func (builder *Builder) SetBaseUrl(url string) *Builder {
	builder.baseUrl = url
	return builder
}

func (builder *Builder) AddCookie(cookie *http.Cookie) *Builder {
	builder.cookies = append(builder.cookies, cookie)
	return builder
}

func (builder *Builder) AddCookies(cookies []*http.Cookie) *Builder {
	builder.cookies = append(builder.cookies, cookies...)
	return builder
}

func (builder *Builder) SetHeader(key, value string) *Builder {
	builder.headers.Set(key, value)
	return builder
}

func (builder *Builder) AddHeader(key, value string) *Builder {
	builder.headers.Add(key, value)
	return builder
}

func (builder *Builder) Build() *Creator {
	if builder.baseUrl == "" {
		panic(BASE_URL_CANNOT_BE_EMPTY)
	}

	if builder.client == nil {
		builder.client = &http.Client{}
	}

	return &Creator{
		baseUrl: builder.baseUrl,
		cookies: builder.cookies,
		headers: builder.headers,
		client:  builder.client,
	}
}

func (creator *Creator) Impl(interface{}) (err error) {
	return err
}
