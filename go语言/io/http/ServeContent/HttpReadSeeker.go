package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

type HttpReadSeeker struct {
	offset        int64
	url           string
	contentLength int64
	client        *http.Client
}

func NewHttpReadSeeker(client *http.Client, url string) *HttpReadSeeker {
	return &HttpReadSeeker{
		offset:        0,
		url:           url,
		contentLength: -1,
		client:        client,
	}
}

func (h *HttpReadSeeker) Read(p []byte) (n int, err error) {
	rangeHeader := fmt.Sprintf("bytes=%d-%d", h.offset, h.offset+int64(len(p))-1)
	req, err := http.NewRequest(http.MethodGet, h.url, nil)
	if err != nil {
		return 0, err
	}
	req.Header.Set("Range", rangeHeader)
	resp, err := h.client.Do(req)
	if err != nil {
		return 0, err
	}
	if resp != nil {
		defer resp.Body.Close()
	}
	n, err = io.ReadFull(resp.Body, p)
	h.offset += int64(n)
	return n, err
}

func (h *HttpReadSeeker) Seek(offset int64, whence int) (int64, error) {
	switch whence {
	case io.SeekStart:
		h.offset = offset
	case io.SeekCurrent:
		h.offset += offset
	case io.SeekEnd:
		if h.contentLength == -1 {
			resp, err := http.Head(h.url)
			if err != nil {
				return 0, err
			}
			if resp != nil {
				defer func() {
					io.Copy(io.Discard, resp.Body)
					resp.Body.Close()
				}()
			}

			contentLengthString := resp.Header.Get("Content-Length")
			h.contentLength, err = strconv.ParseInt(contentLengthString, 10, 64)
			if err != nil {
				return 0, err
			}
		}
		h.offset = h.contentLength - offset
	default:
		return 0, errors.New("whence value error")
	}
	return h.offset, nil
}
