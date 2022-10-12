package exception

type NotFoundError struct {
	Error string
}

// gak usah pake pointer karena tidak ada fungsi yang lain yang mengubah nilai dari struct ini
func NewNotFoundError(err string) NotFoundError {
	return NotFoundError{
		Error: err,
	}
}
