package jwt

import "os"

var tokenEncodeString = os.Getenv("JWT_PSSWD")

// UserType use to indicate that this lib is being used for user token
const UserType string = "models.User"

// AdminType use to indicate that this lib is being used for admin token
const AdminType string = "models.Admin"

// DriverType use to indicate that this lib is being used for driver token
const DriverType string = "models.Driver"
