package excelize

import (
	"encoding/xml"

	"github.com/richardlehane/mscfb"
)

var (
	blockKey                    = []byte{0x14, 0x6e, 0x0b, 0xe7, 0xab, 0xac, 0xd0, 0xd6}
	oleIdentifier               = []byte{0xd0, 0xcf, 0x11, 0xe0, 0xa1, 0xb1, 0x1a, 0xe1}
	headerCLSID                 = make([]byte, 16)
	difSect                     = -4
	endOfChain                  = -2
	fatSect                     = -3
	iterCount                   = 50000
	packageEncryptionChunkSize  = 4096
	packageOffset               = 8
	sheetProtectionSpinCount    = 1e5
	workbookProtectionSpinCount = 1e5
)

type Encryption struct {
	XMLName       xml.Name      `xml:"encryption"`
	KeyData       KeyData       `xml:"keyData"`
	DataIntegrity DataIntegrity `xml:"dataIntegrity"`
	KeyEncryptors KeyEncryptors `xml:"keyEncryptors"`
}

type KeyData struct {
	SaltSize        int    `xml:"saltSize,attr"`
	BlockSize       int    `xml:"blockSize,attr"`
	KeyBits         int    `xml:"keyBits,attr"`
	HashSize        int    `xml:"hashSize,attr"`
	CipherAlgorithm string `xml:"cipherAlgorithm,attr"`
	CipherChaining  string `xml:"cipherChaining,attr"`
	HashAlgorithm   string `xml:"hashAlgorithm,attr"`
	SaltValue       string `xml:"saltValue,attr"`
}

type DataIntegrity struct {
	EncryptedHmacKey   string `xml:"encryptedHmacKey,attr"`
	EncryptedHmacValue string `xml:"encryptedHmacValue,attr"`
}

type KeyEncryptors struct {
	KeyEncryptor []KeyEncryptor `xml:"keyEncryptor"`
}

type KeyEncryptor struct {
	XMLName      xml.Name     `xml:"keyEncryptor"`
	URI          string       `xml:"uri,attr"`
	EncryptedKey EncryptedKey `xml:"encryptedKey"`
}

type EncryptedKey struct {
	XMLName                    xml.Name `xml:"http://schemas.microsoft.com/office/2006/keyEncryptor/password encryptedKey"`
	SpinCount                  int      `xml:"spinCount,attr"`
	EncryptedVerifierHashInput string   `xml:"encryptedVerifierHashInput,attr"`
	EncryptedVerifierHashValue string   `xml:"encryptedVerifierHashValue,attr"`
	EncryptedKeyValue          string   `xml:"encryptedKeyValue,attr"`
	KeyData
}

type StandardEncryptionHeader struct {
	Flags        uint32
	SizeExtra    uint32
	AlgID        uint32
	AlgIDHash    uint32
	KeySize      uint32
	ProviderType uint32
	Reserved1    uint32
	Reserved2    uint32
	CspName      string
}

type StandardEncryptionVerifier struct {
	SaltSize              uint32
	Salt                  []byte
	EncryptedVerifier     []byte
	VerifierHashSize      uint32
	EncryptedVerifierHash []byte
}

type encryption struct {
	BlockSize, SaltSize                                                                  int
	EncryptedKeyValue, EncryptedVerifierHashInput, EncryptedVerifierHashValue, SaltValue []byte
	KeyBits                                                                              uint32
}

func Decrypt(raw []byte, opts *Options) (packageBuf []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func Encrypt(raw []byte, opts *Options) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func extractPart(doc *mscfb.Reader) ([]byte, []byte, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func encryptionMechanism(buffer []byte) (mechanism string, err error) {
	_ = "STUB: not implemented"
	return "", nil
}

func standardDecrypt(encryptionInfoBuf, encryptedPackageBuf []byte, opts *Options) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func standardEncryptionVerifier(algorithm string, blob []byte) StandardEncryptionVerifier {
	_ = "STUB: not implemented"
	return *new(StandardEncryptionVerifier)
}

func standardConvertPasswdToKey(header StandardEncryptionHeader, verifier StandardEncryptionVerifier, opts *Options) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func standardXORBytes(a, b []byte) []byte { _ = "STUB: not implemented"; return nil }

func (e *encryption) encrypt(input []byte) []byte { _ = "STUB: not implemented"; return nil }

func (e *encryption) standardKeyEncryption(password string) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func agileDecrypt(encryptionInfoBuf, encryptedPackageBuf []byte, opts *Options) (packageBuf []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func convertPasswdToKey(passwd string, blockKey []byte, encryption Encryption) (key []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func hashing(hashAlgorithm string, buffer ...[]byte) (key []byte) {
	_ = "STUB: not implemented"
	return nil
}

func createUInt32LEBuffer(value int, bufferSize int) []byte { _ = "STUB: not implemented"; return nil }

func parseEncryptionInfo(encryptionInfo []byte) (encryption Encryption, err error) {
	_ = "STUB: not implemented"
	return *new(Encryption), nil
}

func decrypt(key, iv, input []byte) (packageKey []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func decryptPackage(packageKey, input []byte, encryption Encryption) (outputChunks []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func createIV(blockKey interface{}, encryption Encryption) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func randomBytes(n int) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func genISOPasswdHash(passwd, hashAlgorithm, salt string, spinCount int) (hashValue, saltValue string, err error) {
	_ = "STUB: not implemented"
	return "", "", nil
}

type cfb struct {
	stream   []byte
	position int
	paths    []string
	sectors  []sector
}

type sector struct {
	clsID, content                             []byte
	name                                       string
	C, L, R, color, size, start, state, typeID int
}

func (c *cfb) writeBytes(value []byte) { _ = "STUB: not implemented"; return }

func (c *cfb) writeUint16(value int) { _ = "STUB: not implemented"; return }

func (c *cfb) writeUint32(value int) { _ = "STUB: not implemented"; return }

func (c *cfb) writeUint64(value int) { _ = "STUB: not implemented"; return }

func (c *cfb) writeStrings(value string) { _ = "STUB: not implemented"; return }

func (c *cfb) put(name string, content []byte) { _ = "STUB: not implemented"; return }

func (c *cfb) compare(left, right string) int { _ = "STUB: not implemented"; return 0 }

func (c *cfb) prepare() { _ = "STUB: not implemented"; return }

func (c *cfb) locate() []int { _ = "STUB: not implemented"; return nil }

func (c *cfb) writeMSAT(location []int) { _ = "STUB: not implemented"; return }

func (c *cfb) writeDirectoryEntry(location []int) { _ = "STUB: not implemented"; return }

func (c *cfb) writeSectorChains(location []int) sector {
	_ = "STUB: not implemented"
	return *new(sector)
}

func (c *cfb) write() []byte { _ = "STUB: not implemented"; return nil }
