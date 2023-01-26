# passhash-argon2

Generate and verify salt/password-hash pairs using the argon2 password hashing algorithm (golang.org/x/crypto/argon2).

## CLI usage

The project includes a CLI utility where `phargon2 check` can be used to verify a password using the output of `phargon2 new`.

## Package Usage

`import "github.com/brandon-detty/passhash-argon2/passhash"`

### Functions

#### func GenerateCondiment() ([]byte, error)

This can be used to generate a secure "pepper" value. The pepper is to be shared across an application and stored outside of the database.

#### func CreateHash(password []byte, pepper []byte) (salt []byte, hash []byte, err error)

Create a salt/hash pair to store in the database for later user authentication.

#### func CheckPassword(password []byte, salt []byte, pepper []byte, correctHash []byte) bool

Authenticate a user based on an entered password, the application's pepper, and the salt and password hash stored in the database.

### Constants

ArgonKeyLength = 32 (byte-length of the generated hash)

CondimentLength = 16 (salt/pepper byte-length)
