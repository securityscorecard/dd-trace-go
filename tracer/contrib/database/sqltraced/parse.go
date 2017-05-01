package sqltraced

import (
	"github.com/DataDog/dd-trace-go/tracer/contrib/database/sqltraced/parsedsn"
)

// parseDSN returns all information passed through the DSN:
func parseDSN(driverType, dsn string) (meta map[string]string, err error) {
	switch driverType {
	case "*pq.Driver":
		meta, err = parsedsn.Postgres(dsn)
	case "*mysql.MySQLDriver":
		meta, err = parsedsn.MySQL(dsn)
	}
	meta = normalize(meta)
	return meta, err
}

func normalize(meta map[string]string) map[string]string {
	m := make(map[string]string)
	for k, v := range meta {
		if nk, ok := normalizeKey(k); ok {
			m[nk] = v
		}
	}
	return m
}

func normalizeKey(k string) (string, bool) {
	switch k {
	case "user":
		return "db.user", true
	case "application_name":
		return "db.application", true
	case "dbname":
		return "db.name", true
	case "host", "port":
		return "out." + k, true
	default:
		return "", false
	}
}
