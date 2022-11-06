package rook

import (
	"github.com/reiver/go-fck"

	"io"
	"net/url"
	"strings"
)

const (
	errBadSentence      = fck.Error("bad sentence")
	errObjectFound      = fck.Error("object found")
	errPunctuationFound = fck.Error("punctuation found")
	errVerbFound        = fck.Error("verb found")
)

func (receiver Sentence) String() string {

	var storage strings.Builder

	_, err := receiver.WriteTo(&storage)
	if nil != err {
		return "⧼error⧽"
	}

	return storage.String()
}

func (receiver Sentence) WriteTo(w io.Writer) (n int64, err error) {

	if "" == receiver.Punctuation {
		switch {
		case "GET" == receiver.Verb && "" != receiver.Object:
			return receiver.writeToForGet(w)
		case ""    == receiver.Verb && "" == receiver.Object:
			return receiver.writeToForList(w)
		case ""    == receiver.Verb && "" != receiver.Object:
			return receiver.writeToForUser(w)
		default:
			return 0, errBadSentence
		}
	}

	if "" == receiver.Verb {
		return 0, errBadSentence
	}

	return receiver.writeTo(w)
}


func (receiver Sentence) writeTo(writer io.Writer) (int64, error) {

	var storage strings.Builder
	{
		storage.WriteString(receiver.Punctuation)
		storage.WriteString(receiver.Verb)
		if ("" != receiver.Punctuation || "" != receiver.Verb) && "" != receiver.Object {
			storage.WriteRune(' ')
		}
		storage.WriteString(receiver.Object)
		switch receiver.Verb {
		case "W":
			storage.WriteString("\r\n")
		default:
			storage.WriteRune('\u0085')
		}
	}

	var n64 int64
	var err error
	{
		var n int

		n, err = io.WriteString(writer, storage.String())
		n64 = int64(n)
	}

	return n64, err
}

func (receiver Sentence) writeToForGet(writer io.Writer) (int64, error) {

	var requestURI string
	var hostname string
	{
		uri, err := url.Parse(receiver.Object)
		if nil != err {
			return 0, err
		}

		hostname = uri.Hostname()
		requestURI = uri.RequestURI()
	}

	var storage strings.Builder
	{
		storage.WriteString(receiver.Verb)
		storage.WriteRune(' ')
		storage.WriteString(requestURI)
		storage.WriteRune(' ')
		storage.WriteString("HTTP/1.1\r\n")

		storage.WriteString("Host: ")
		storage.WriteString(hostname)
		storage.WriteString("\r\n")

		storage.WriteString("Connection: close\r\n")

		storage.WriteString("\r\n")
	}

	var n64 int64
	var err error
	{
		var n int

		n, err = io.WriteString(writer, storage.String())
		n64 = int64(n)
	}

	return n64, err
}

func (receiver Sentence) writeToForUser(writer io.Writer) (int64, error) {

	if "" != receiver.Punctuation {
		return 0, errPunctuationFound
	}

	if "" != receiver.Verb {
		return 0, errVerbFound
	}


	var storage strings.Builder
	{
		storage.WriteString(receiver.Object)
		storage.WriteString("\r\n")
	}

	var n64 int64
	var err error
	{
		var n int

		n, err = io.WriteString(writer, storage.String())
		n64 = int64(n)
	}

	return n64, err
}

func (receiver Sentence) writeToForList(writer io.Writer) (int64, error) {

	if "" != receiver.Punctuation {
		return 0, errPunctuationFound
	}

	if "" != receiver.Verb {
		return 0, errVerbFound
	}

	if "" != receiver.Object {
		return 0, errObjectFound
	}


	var storage strings.Builder
	{
		storage.WriteString("\r\n")
	}

	var n64 int64
	var err error
	{
		var n int

		n, err = io.WriteString(writer, storage.String())
		n64 = int64(n)
	}

	return n64, err
}
