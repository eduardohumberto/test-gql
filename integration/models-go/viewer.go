package models

import "github.com/eduardohumberto/test-gql/integration/remote_api"

type Viewer struct {
	User *remote_api.User
}
