package main

// Pub is the authentication sender
func (e Engine) Pub(msg string) error {
	err := e.Send(msg)
	if err != nil {
		return err
	}

	return nil
}
