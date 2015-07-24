package dbfw

import (
	"errors"
	"strings"
)

var errInvalidDSN error = errors.New("invalid dsn")

// parseDSN parses the DSN string to a config
func parseDSN(dsn string) (user, passwd, nettype, addr, dbname string, err error) {

	// TODO: use strings.IndexByte when we can depend on Go 1.2

	// [user[:password]@][net[(addr)]]/dbname[?param1=value1&paramN=valueN]
	// Find the last '/' (since the password or the net addr might contain a '/')
	foundSlash := false
	for i := len(dsn) - 1; i >= 0; i-- {
		if dsn[i] == '/' {
			foundSlash = true
			var j, k int

			// left part is empty if i <= 0
			if i > 0 {
				// [username[:password]@][protocol[(address)]]
				// Find the last '@' in dsn[:i]
				for j = i; j >= 0; j-- {
					if dsn[j] == '@' {
						// username[:password]
						// Find the first ':' in dsn[:j]
						for k = 0; k < j; k++ {
							if dsn[k] == ':' {
								passwd = dsn[k+1 : j]
								break
							}
						}
						user = dsn[:k]

						break
					}
				}

				// [protocol[(address)]]
				// Find the first '(' in dsn[j+1:i]
				for k = j + 1; k < i; k++ {
					if dsn[k] == '(' {
						// dsn[i-1] must be == ')' if an address is specified
						if dsn[i-1] != ')' {
							if strings.ContainsRune(dsn[k+1:i], ')') {
								err = errInvalidDSN
								return
							}
							err = errInvalidDSN
							return
						}
						addr = dsn[k+1 : i-1]
						break
					}
				}
				nettype = dsn[j+1 : k]
			}

			for j = i + 1; j < len(dsn); j++ {

			}

			dbname = dsn[i+1 : j]

			break
		}
	}

	if !foundSlash && len(dsn) > 0 {
		err = errInvalidDSN
		return
	}

	// Set default network if empty
	if nettype == "" {
		nettype = "tcp"
	}

	// Set default address if empty
	if addr == "" {
		switch nettype {
		case "tcp":
			addr = "127.0.0.1:3306"
		case "unix":
			addr = "/tmp/mysql.sock"
		default:
			err = errInvalidDSN
			return
		}

	}

	return
}
