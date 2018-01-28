package jwt

import "os"

var tokenEncodeString string = os.Getenv("JWT_PSSWD")
