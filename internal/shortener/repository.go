package shortener

type RedircetRepository interface {
	Find(code string) (*Redirect, error)
	Store(redirect *Redirect) error
}
