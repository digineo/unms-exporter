package config

import "strings"

type tokenMap map[string]string

// Decode implements the envconfig.Decoder interface. Its purpose is to
// parse the UNMS_EXPORTER_TOKEN environment variable, which takes the
// following form:
//
//	UNMS_EXPORTER_TOKEN ::= <dns-token> ("," <host-token>)*
//	<dns-token> ::= <dns-name> "=" <token>
//
// The "dns-name" follows common convention (for details, see RFC 1123).
// The "token" is the API access token generated by UNMS and usually
// takes the shape of a UUID.
func (toks *tokenMap) Decode(v string) error {
	t := make(map[string]string)
	for _, pair := range strings.Split(v, ",") {
		kv := strings.SplitN(pair, "=", 2)
		t[kv[0]] = kv[1]
	}
	*toks = tokenMap(t)
	return nil
}
