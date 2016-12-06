// lint_ext_authority_key_identifer_no_key_identifier.go
/***********************************************************************
RFC 5280: 4.2.1.1
The keyIdentifier field of the authorityKeyIdentifier extension MUST
   be included in all certificates generated by conforming CAs to
   facilitate certification path construction.  There is one exception;
   where a CA distributes its public key in the form of a "self-signed"
   certificate, the authority key identifier MAY be omitted.  The
   signature on a self-signed certificate is generated with the private
   key associated with the certificate's subject public key.  (This
   proves that the issuer possesses both the public and private keys.)
   In this case, the subject and authority key identifiers would be
   identical, but only the subject key identifier is needed for
   certification path building.
***********************************************************************/

package lints

import (

	"github.com/teamnsrg/zlint/util"
	"github.com/zmap/zgrab/ztools/x509"
)

type authorityKeyIdNoKeyIdField struct {
	// Internal data here
}

func (l *authorityKeyIdNoKeyIdField) Initialize() error {
	return nil
}

func (l *authorityKeyIdNoKeyIdField) CheckApplies(c *x509.Certificate) bool {
	return true
}

func (l *authorityKeyIdNoKeyIdField) RunTest(c *x509.Certificate) (ResultStruct, error) {
	if c.AuthorityKeyId == nil && !util.IsSelfSigned(c) { //will be nil by defualt if not found in x509.parseCert
		return ResultStruct{Result: Error}, nil
	} else {
		return ResultStruct{Result: Pass}, nil
	}
}

func init() {
	RegisterLint(&Lint{
		Name:          "ext_authority_key_identifier_no_key_identifier",
		Description:   "CAs must include keyIdentifer field of aki in all non-self-issued certs",
		Providence:    "RFC 5280: 4.2.1.1",
		EffectiveDate: util.RFC2459Date,
		Test:          &authorityKeyIdNoKeyIdField{}})
}
