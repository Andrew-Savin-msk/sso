package auth

import ssov1 "github.com/Andrew-Savin-msk/protos/gen/go/sso"

type serverApi struct {
	ssov1.UnimplementedAuthServer
}
