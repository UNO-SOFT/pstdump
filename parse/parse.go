package parse

import (
	"encoding/json"
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
	for {
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
}
