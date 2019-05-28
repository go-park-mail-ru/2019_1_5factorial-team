package ResponseWriter

import "net/http"

// https://www.reddit.com/r/golang/comments/7p35s4/how_do_i_get_the_response_status_for_my_middleware/dse5y4g?utm_source=share&utm_medium=web2x
type statusWriter struct {
	http.ResponseWriter
	status int
	length int
}

func NewStatusWriter(res http.ResponseWriter) *statusWriter {
	return &statusWriter{
		ResponseWriter: res,
	}
}

func (w *statusWriter) WriteHeader(status int) {
	w.status = status
	w.ResponseWriter.WriteHeader(status)
}

func (w *statusWriter) Write(b []byte) (int, error) {
	if w.status == 0 {
		w.status = 200
	}
	n, err := w.ResponseWriter.Write(b)
	w.length += n
	return n, err
}

func (w *statusWriter) GetStatusCode() int {
	return w.status
}
