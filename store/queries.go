package store

import (
	_ "embed"
)

//go:embed queries/init_users_table.sql
var INIT_USERS_TABLE string

//go:embed queries/select_user_by_email.sql
var SELECT_USER_BY_EMAIL string

//go:embed queries/put_user.sql
var PUT_USER string
