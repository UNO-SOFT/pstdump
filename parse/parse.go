package parse

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/pkg/errors"
)

type Email struct {
	Folder, Body, BodyHTML, BodyPrefix string
	ClientSubmitTime, DeliveryTime     string
	ConversationID                     []byte
	ConversationTopic, InReplyTo       string
	ArticleNumber                      int64
	ID, Class                          string
	Size                               int64
	NextSendAcct, PrimarySendAccount   string
	ReceivedRepresenting, ReceivedBy   EmailAddress
	ReturnPath, RTFBody                string
	Sender, SetnRepresenting           EmailAddress
	Subject, BCC, CC, To               string
	CompName                           string
	Headers                            string
	Recipients                         []EmailAddress
	Attachments                        []Attachment
}

func (e *Email) WriteTo(w io.Writer) (int64, error) {
	var n int64
	i, err := io.WriteString(w, e.Headers)
	if err != nil {
		return int64(i), err
	}
	n += int64(i)
	for _, a := range e.Attachments {
		i, err = fmt.Fprintf(w, "%s\r\n", a.Data)
		n += int64(i)
		if err != nil {
			return n, err
		}
	}
	return n, err
}

type Attachment struct {
	ContentDisposition string
	Size               int64
	ID                 string
	Created            string
	FileName           string
	FileSize           int64
	ContentType        string
	Data               []byte
}

type EmailAddress struct {
	Address, Name string
}

func Parse(r io.Reader, f func(*Email) error) error {
	dec := json.NewDecoder(r)
	for dec.More() {
		var eml Email
		if err := dec.Decode(&eml); err != nil {
			if err == io.EOF {
				return nil
			}
			return errors.Wrap(err, "decode")
		}
		if err := f(&eml); err != nil {
			return err
		}
	}
	return nil
}
