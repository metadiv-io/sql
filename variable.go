package sql

var db_encrypt_key = ""

/*
SetEncryptKey sets the encryption key for the database.
*/
func SetEncryptKey(key string) {
	db_encrypt_key = key
}
