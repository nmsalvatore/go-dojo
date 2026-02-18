package inventory

import (
	"io"
)

type Notifier interface {
	Notify(string) error
}

type StoreInMemory struct {
	messages []string
}

func (s *StoreInMemory) Notify(msg string) error {
	s.messages = append(s.messages, msg)
	return nil
}

func (s *StoreInMemory) Messages() []string {
	return s.messages
}

func (s *StoreInMemory) Reset() {
	s.messages = []string{}
}

type WriteTo struct {
	writer io.Writer
}

func (w *WriteTo) Notify(msg string) error {
	_, err := w.writer.Write([]byte(msg))
	if err != nil {
		return err
	}
	return nil
}
