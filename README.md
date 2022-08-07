# simple-bank-ledger

*Very* simple bank ledger application in Go.

The goal is to create an interface that:

- [x] registers new users
- [x] logs in users
- [x] shows logged-in user's balance
- [x] shows logged-in user's transaction history
- [x] records a transaction for a logged-in user
- [x] logs out logged-in user
- [x] just uses local cache on a single client machine (doesn't handle concurrency, isn't persistent, does not check against a database)
- [x] could easily be extended to use a database, etc.

## Some cleanup goals

- [ ] Add some password validation rules
- [ ] Validate debit/credit/voids and negative amounts
- [x] Implement test coverage reporting
- [ ] Make UI strings internationalizeable
- [ ] Improve test coverage
- [ ] Use test tables and assertions for all unit tests where sensible

## Some stretch goals

- [ ] Use ledger "layers"
- [ ] Implement with file-based storage
- [ ] Add a GUI implementation
- [ ] Add an API implementation
