package nio

import "io"

func NewReader(reader io.Reader) io.Reader {
	r, w := Pipe()
	go func() {
		io.Copy(w, reader)
		w.Close()
	}()
	return r
}

func NewReadCloser(reader io.ReadCloser) io.Reader {
	r, w := Pipe()
	go func() {
		io.Copy(w, reader)
		w.Close()
		reader.Close()
	}()
	return r
}

func NewWriter(writer io.Writer) io.WriteCloser {
	r, w := Pipe()
	go io.Copy(writer, r)
	return w
}

func NewWriteCloser(writer io.WriteCloser) io.WriteCloser {
	r, w := Pipe()
	go func() {
		io.Copy(writer, r)
		writer.Close()
	}()
	return w
}
