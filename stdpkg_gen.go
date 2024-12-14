// This file is part of GoRE.
//
// Copyright (C) 2019-2021 GoRE Authors
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program. If not, see <http://www.gnu.org/licenses/>.

// Code generated by go generate; DO NOT EDIT.
// This file was generated at
// 2024-12-14 07:17:22.640735161 +0000 UTC
// Generate with version hash: b6e84dd3b37307351e3e9228bd09ab6a4cd2d14db20961537e0afdcf793e9527

package gore

var stdPkgs = map[string]struct{}{
	"archive":                              {},
	"archive/tar":                          {},
	"archive/zip":                          {},
	"arena":                                {},
	"bufio":                                {},
	"builtin":                              {},
	"bytes":                                {},
	"cmp":                                  {},
	"compress":                             {},
	"compress/bzip2":                       {},
	"compress/flate":                       {},
	"compress/gzip":                        {},
	"compress/lzw":                         {},
	"compress/zlib":                        {},
	"constraints":                          {},
	"container":                            {},
	"container/heap":                       {},
	"container/list":                       {},
	"container/ring":                       {},
	"context":                              {},
	"crypto":                               {},
	"crypto/aes":                           {},
	"crypto/boring":                        {},
	"crypto/cipher":                        {},
	"crypto/des":                           {},
	"crypto/dsa":                           {},
	"crypto/ecdh":                          {},
	"crypto/ecdsa":                         {},
	"crypto/ed25519":                       {},
	"crypto/ed25519/internal":              {},
	"crypto/ed25519/internal/edwards25519": {},
	"crypto/ed25519/internal/edwards25519/field": {},
	"crypto/elliptic":                            {},
	"crypto/elliptic/internal":                   {},
	"crypto/elliptic/internal/fiat":              {},
	"crypto/elliptic/internal/nistec":            {},
	"crypto/fips140":                             {},
	"crypto/hkdf":                                {},
	"crypto/hmac":                                {},
	"crypto/internal":                            {},
	"crypto/internal/alias":                      {},
	"crypto/internal/bigmod":                     {},
	"crypto/internal/boring":                     {},
	"crypto/internal/boring/bbig":                {},
	"crypto/internal/boring/bcache":              {},
	"crypto/internal/boring/fipstls":             {},
	"crypto/internal/boring/sig":                 {},
	"crypto/internal/boring/syso":                {},
	"crypto/internal/cipherhw":                   {},
	"crypto/internal/cryptotest":                 {},
	"crypto/internal/edwards25519":               {},
	"crypto/internal/edwards25519/field":         {},
	"crypto/internal/entropy":                    {},
	"crypto/internal/fips140":                    {},
	"crypto/internal/fips140/aes":                {},
	"crypto/internal/fips140/aes/_asm/ctr":       {},
	"crypto/internal/fips140/aes/_asm/standard":  {},
	"crypto/internal/fips140/aes/gcm":            {},
	"crypto/internal/fips140/aes/gcm/_asm/gcm":   {},
	"crypto/internal/fips140/alias":              {},
	"crypto/internal/fips140/bigmod":             {},
	"crypto/internal/fips140/check":              {},
	"crypto/internal/fips140/check/checktest":    {},
	"crypto/internal/fips140/drbg":               {},
	"crypto/internal/fips140/ecdh":               {},
	"crypto/internal/fips140/ecdsa":              {},
	"crypto/internal/fips140/ed25519":            {},
	"crypto/internal/fips140/edwards25519":       {},
	"crypto/internal/fips140/edwards25519/field": {},
	"crypto/internal/fips140/hkdf":               {},
	"crypto/internal/fips140/hmac":               {},
	"crypto/internal/fips140/mlkem":              {},
	"crypto/internal/fips140/nistec":             {},
	"crypto/internal/fips140/nistec/fiat":        {},
	"crypto/internal/fips140/pbkdf2":             {},
	"crypto/internal/fips140/rsa":                {},
	"crypto/internal/fips140/sha256":             {},
	"crypto/internal/fips140/sha3":               {},
	"crypto/internal/fips140/sha512":             {},
	"crypto/internal/fips140/ssh":                {},
	"crypto/internal/fips140/subtle":             {},
	"crypto/internal/fips140/tls12":              {},
	"crypto/internal/fips140/tls13":              {},
	"crypto/internal/fips140deps":                {},
	"crypto/internal/fips140deps/byteorder":      {},
	"crypto/internal/fips140deps/cpu":            {},
	"crypto/internal/fips140deps/godebug":        {},
	"crypto/internal/fips140only":                {},
	"crypto/internal/fips140test":                {},
	"crypto/internal/hpke":                       {},
	"crypto/internal/impl":                       {},
	"crypto/internal/mlkem768":                   {},
	"crypto/internal/nistec":                     {},
	"crypto/internal/nistec/fiat":                {},
	"crypto/internal/randutil":                   {},
	"crypto/internal/subtle":                     {},
	"crypto/internal/sysrand":                    {},
	"crypto/internal/sysrand/internal":           {},
	"crypto/internal/sysrand/internal/seccomp":   {},
	"crypto/md5":                                 {},
	"crypto/mlkem":                               {},
	"crypto/pbkdf2":                              {},
	"crypto/rand":                                {},
	"crypto/rc4":                                 {},
	"crypto/rsa":                                 {},
	"crypto/sha1":                                {},
	"crypto/sha256":                              {},
	"crypto/sha3":                                {},
	"crypto/sha512":                              {},
	"crypto/subtle":                              {},
	"crypto/tls":                                 {},
	"crypto/tls/fipsonly":                        {},
	"crypto/tls/internal":                        {},
	"crypto/tls/internal/fips140tls":             {},
	"crypto/x509":                                {},
	"crypto/x509/internal":                       {},
	"crypto/x509/internal/macOS":                 {},
	"crypto/x509/internal/macos":                 {},
	"crypto/x509/pkix":                           {},
	"database":                                   {},
	"database/sql":                               {},
	"database/sql/driver":                        {},
	"database/sql/internal":                      {},
	"debug":                                      {},
	"debug/buildinfo":                            {},
	"debug/dwarf":                                {},
	"debug/elf":                                  {},
	"debug/goobj":                                {},
	"debug/gosym":                                {},
	"debug/macho":                                {},
	"debug/pe":                                   {},
	"debug/plan9obj":                             {},
	"embed":                                      {},
	"embed/internal":                             {},
	"embed/internal/embedtest":                   {},
	"encoding":                                   {},
	"encoding/ascii85":                           {},
	"encoding/asn1":                              {},
	"encoding/base32":                            {},
	"encoding/base64":                            {},
	"encoding/binary":                            {},
	"encoding/csv":                               {},
	"encoding/gob":                               {},
	"encoding/hex":                               {},
	"encoding/json":                              {},
	"encoding/pem":                               {},
	"encoding/xml":                               {},
	"errors":                                     {},
	"expvar":                                     {},
	"flag":                                       {},
	"fmt":                                        {},
	"go":                                         {},
	"go/ast":                                     {},
	"go/ast/internal":                            {},
	"go/ast/internal/tests":                      {},
	"go/build":                                   {},
	"go/build/constraint":                        {},
	"go/constant":                                {},
	"go/doc":                                     {},
	"go/doc/comment":                             {},
	"go/format":                                  {},
	"go/importer":                                {},
	"go/internal":                                {},
	"go/internal/gccgoimporter":                  {},
	"go/internal/gcimporter":                     {},
	"go/internal/srcimporter":                    {},
	"go/internal/typeparams":                     {},
	"go/parser":                                  {},
	"go/printer":                                 {},
	"go/scanner":                                 {},
	"go/token":                                   {},
	"go/types":                                   {},
	"go/types/fixedbugs":                         {},
	"go/version":                                 {},
	"hash":                                       {},
	"hash/adler32":                               {},
	"hash/crc32":                                 {},
	"hash/crc64":                                 {},
	"hash/fnv":                                   {},
	"hash/maphash":                               {},
	"html":                                       {},
	"html/template":                              {},
	"image":                                      {},
	"image/color":                                {},
	"image/color/palette":                        {},
	"image/draw":                                 {},
	"image/gif":                                  {},
	"image/internal":                             {},
	"image/internal/imageutil":                   {},
	"image/jpeg":                                 {},
	"image/png":                                  {},
	"index":                                      {},
	"index/suffixarray":                          {},
	"internal":                                   {},
	"internal/abi":                               {},
	"internal/asan":                              {},
	"internal/bisect":                            {},
	"internal/buildcfg":                          {},
	"internal/bytealg":                           {},
	"internal/byteorder":                         {},
	"internal/cfg":                               {},
	"internal/chacha8rand":                       {},
	"internal/concurrent":                        {},
	"internal/copyright":                         {},
	"internal/coverage":                          {},
	"internal/coverage/calloc":                   {},
	"internal/coverage/cfile":                    {},
	"internal/coverage/cformat":                  {},
	"internal/coverage/cmerge":                   {},
	"internal/coverage/decodecounter":            {},
	"internal/coverage/decodemeta":               {},
	"internal/coverage/encodecounter":            {},
	"internal/coverage/encodemeta":               {},
	"internal/coverage/pods":                     {},
	"internal/coverage/rtcov":                    {},
	"internal/coverage/slicereader":              {},
	"internal/coverage/slicewriter":              {},
	"internal/coverage/stringtab":                {},
	"internal/coverage/test":                     {},
	"internal/coverage/uleb128":                  {},
	"internal/cpu":                               {},
	"internal/dag":                               {},
	"internal/diff":                              {},
	"internal/execabs":                           {},
	"internal/exportdata":                        {},
	"internal/filepathlite":                      {},
	"internal/fmtsort":                           {},
	"internal/format":                            {},
	"internal/fuzz":                              {},
	"internal/goarch":                            {},
	"internal/godebug":                           {},
	"internal/godebugs":                          {},
	"internal/goexperiment":                      {},
	"internal/golang.org":                        {},
	"internal/golang.org/x":                      {},
	"internal/golang.org/x/net":                  {},
	"internal/golang.org/x/net/http2":            {},
	"internal/golang.org/x/net/http2/hpack":      {},
	"internal/goos":                              {},
	"internal/goroot":                            {},
	"internal/gover":                             {},
	"internal/goversion":                         {},
	"internal/intern":                            {},
	"internal/itoa":                              {},
	"internal/lazyregexp":                        {},
	"internal/lazytemplate":                      {},
	"internal/msan":                              {},
	"internal/nettrace":                          {},
	"internal/oserror":                           {},
	"internal/pkgbits":                           {},
	"internal/platform":                          {},
	"internal/poll":                              {},
	"internal/pprof":                             {},
	"internal/pprof/profile":                     {},
	"internal/profile":                           {},
	"internal/profilerecord":                     {},
	"internal/race":                              {},
	"internal/reflectlite":                       {},
	"internal/runtime":                           {},
	"internal/runtime/atomic":                    {},
	"internal/runtime/exithook":                  {},
	"internal/runtime/maps":                      {},
	"internal/runtime/math":                      {},
	"internal/runtime/sys":                       {},
	"internal/runtime/syscall":                   {},
	"internal/safefilepath":                      {},
	"internal/saferio":                           {},
	"internal/singleflight":                      {},
	"internal/stringslite":                       {},
	"internal/sync":                              {},
	"internal/synctest":                          {},
	"internal/syscall":                           {},
	"internal/syscall/execenv":                   {},
	"internal/syscall/unix":                      {},
	"internal/syscall/windows":                   {},
	"internal/syscall/windows/registry":          {},
	"internal/syscall/windows/sysdll":            {},
	"internal/sysinfo":                           {},
	"internal/syslist":                           {},
	"internal/testenv":                           {},
	"internal/testlog":                           {},
	"internal/testpty":                           {},
	"internal/trace":                             {},
	"internal/trace/event":                       {},
	"internal/trace/event/go122":                 {},
	"internal/trace/internal":                    {},
	"internal/trace/internal/oldtrace":           {},
	"internal/trace/internal/testgen":            {},
	"internal/trace/internal/testgen/go122":      {},
	"internal/trace/raw":                         {},
	"internal/trace/testtrace":                   {},
	"internal/trace/traceviewer":                 {},
	"internal/trace/traceviewer/format":          {},
	"internal/trace/traceviewer/static":          {},
	"internal/trace/v2":                          {},
	"internal/trace/v2/event":                    {},
	"internal/trace/v2/event/go122":              {},
	"internal/trace/v2/internal":                 {},
	"internal/trace/v2/internal/testgen":         {},
	"internal/trace/v2/internal/testgen/go122":   {},
	"internal/trace/v2/raw":                      {},
	"internal/trace/v2/testtrace":                {},
	"internal/trace/v2/version":                  {},
	"internal/trace/version":                     {},
	"internal/txtar":                             {},
	"internal/types":                             {},
	"internal/types/errors":                      {},
	"internal/unsafeheader":                      {},
	"internal/weak":                              {},
	"internal/x":                                 {},
	"internal/x/crypto":                          {},
	"internal/x/crypto/chacha20poly1305":         {},
	"internal/x/crypto/cryptobyte":               {},
	"internal/x/crypto/cryptobyte/asn1":          {},
	"internal/x/crypto/curve25519":               {},
	"internal/x/crypto/hkdf":                     {},
	"internal/x/crypto/internal":                 {},
	"internal/x/crypto/internal/chacha20":        {},
	"internal/x/crypto/poly1305":                 {},
	"internal/x/net":                             {},
	"internal/x/net/dns":                         {},
	"internal/x/net/dns/dnsmessage":              {},
	"internal/x/net/http":                        {},
	"internal/x/net/http/httpguts":               {},
	"internal/x/net/http/httpproxy":              {},
	"internal/x/net/http2":                       {},
	"internal/x/net/http2/hpack":                 {},
	"internal/x/net/idna":                        {},
	"internal/x/net/internal":                    {},
	"internal/x/net/internal/nettest":            {},
	"internal/x/net/lif":                         {},
	"internal/x/net/nettest":                     {},
	"internal/x/net/route":                       {},
	"internal/x/text":                            {},
	"internal/x/text/secure":                     {},
	"internal/x/text/secure/bidirule":            {},
	"internal/x/text/transform":                  {},
	"internal/x/text/unicode":                    {},
	"internal/x/text/unicode/bidi":               {},
	"internal/x/text/unicode/norm":               {},
	"internal/xcoff":                             {},
	"internal/zstd":                              {},
	"io":                                         {},
	"io/fs":                                      {},
	"io/ioutil":                                  {},
	"iter":                                       {},
	"lib9":                                       {},
	"lib9/fmt":                                   {},
	"lib9/utf":                                   {},
	"libbio":                                     {},
	"liblink":                                    {},
	"libmach":                                    {},
	"log":                                        {},
	"log/internal":                               {},
	"log/slog":                                   {},
	"log/slog/internal":                          {},
	"log/slog/internal/benchmarks":               {},
	"log/slog/internal/buffer":                   {},
	"log/slog/internal/slogtest":                 {},
	"log/syslog":                                 {},
	"maps":                                       {},
	"math":                                       {},
	"math/big":                                   {},
	"math/bits":                                  {},
	"math/cmplx":                                 {},
	"math/rand":                                  {},
	"math/rand/v2":                               {},
	"mime":                                       {},
	"mime/multipart":                             {},
	"mime/quotedprintable":                       {},
	"net":                                        {},
	"net/http":                                   {},
	"net/http/cgi":                               {},
	"net/http/cookiejar":                         {},
	"net/http/fcgi":                              {},
	"net/http/httptest":                          {},
	"net/http/httptrace":                         {},
	"net/http/httputil":                          {},
	"net/http/internal":                          {},
	"net/http/internal/ascii":                    {},
	"net/http/internal/testcert":                 {},
	"net/http/pprof":                             {},
	"net/internal":                               {},
	"net/internal/cgotest":                       {},
	"net/internal/socktest":                      {},
	"net/mail":                                   {},
	"net/netip":                                  {},
	"net/rpc":                                    {},
	"net/rpc/jsonrpc":                            {},
	"net/smtp":                                   {},
	"net/textproto":                              {},
	"net/url":                                    {},
	"os":                                         {},
	"os/exec":                                    {},
	"os/exec/internal":                           {},
	"os/exec/internal/fdtest":                    {},
	"os/signal":                                  {},
	"os/signal/internal":                         {},
	"os/signal/internal/pty":                     {},
	"os/user":                                    {},
	"path":                                       {},
	"path/filepath":                              {},
	"pkg":                                        {},
	"pkg/archive":                                {},
	"pkg/archive/tar":                            {},
	"pkg/archive/zip":                            {},
	"pkg/bufio":                                  {},
	"pkg/builtin":                                {},
	"pkg/bytes":                                  {},
	"pkg/compress":                               {},
	"pkg/compress/bzip2":                         {},
	"pkg/compress/flate":                         {},
	"pkg/compress/gzip":                          {},
	"pkg/compress/lzw":                           {},
	"pkg/compress/zlib":                          {},
	"pkg/container":                              {},
	"pkg/container/heap":                         {},
	"pkg/container/list":                         {},
	"pkg/container/ring":                         {},
	"pkg/crypto":                                 {},
	"pkg/crypto/aes":                             {},
	"pkg/crypto/cipher":                          {},
	"pkg/crypto/des":                             {},
	"pkg/crypto/dsa":                             {},
	"pkg/crypto/ecdsa":                           {},
	"pkg/crypto/elliptic":                        {},
	"pkg/crypto/hmac":                            {},
	"pkg/crypto/md5":                             {},
	"pkg/crypto/rand":                            {},
	"pkg/crypto/rc4":                             {},
	"pkg/crypto/rsa":                             {},
	"pkg/crypto/sha1":                            {},
	"pkg/crypto/sha256":                          {},
	"pkg/crypto/sha512":                          {},
	"pkg/crypto/subtle":                          {},
	"pkg/crypto/tls":                             {},
	"pkg/crypto/x509":                            {},
	"pkg/crypto/x509/pkix":                       {},
	"pkg/database":                               {},
	"pkg/database/sql":                           {},
	"pkg/database/sql/driver":                    {},
	"pkg/debug":                                  {},
	"pkg/debug/dwarf":                            {},
	"pkg/debug/elf":                              {},
	"pkg/debug/goobj":                            {},
	"pkg/debug/gosym":                            {},
	"pkg/debug/macho":                            {},
	"pkg/debug/pe":                               {},
	"pkg/debug/plan9obj":                         {},
	"pkg/encoding":                               {},
	"pkg/encoding/ascii85":                       {},
	"pkg/encoding/asn1":                          {},
	"pkg/encoding/base32":                        {},
	"pkg/encoding/base64":                        {},
	"pkg/encoding/binary":                        {},
	"pkg/encoding/csv":                           {},
	"pkg/encoding/gob":                           {},
	"pkg/encoding/hex":                           {},
	"pkg/encoding/json":                          {},
	"pkg/encoding/pem":                           {},
	"pkg/encoding/xml":                           {},
	"pkg/errors":                                 {},
	"pkg/expvar":                                 {},
	"pkg/flag":                                   {},
	"pkg/fmt":                                    {},
	"pkg/go":                                     {},
	"pkg/go/ast":                                 {},
	"pkg/go/build":                               {},
	"pkg/go/doc":                                 {},
	"pkg/go/format":                              {},
	"pkg/go/parser":                              {},
	"pkg/go/printer":                             {},
	"pkg/go/scanner":                             {},
	"pkg/go/token":                               {},
	"pkg/hash":                                   {},
	"pkg/hash/adler32":                           {},
	"pkg/hash/crc32":                             {},
	"pkg/hash/crc64":                             {},
	"pkg/hash/fnv":                               {},
	"pkg/html":                                   {},
	"pkg/html/template":                          {},
	"pkg/image":                                  {},
	"pkg/image/color":                            {},
	"pkg/image/color/palette":                    {},
	"pkg/image/draw":                             {},
	"pkg/image/gif":                              {},
	"pkg/image/jpeg":                             {},
	"pkg/image/png":                              {},
	"pkg/index":                                  {},
	"pkg/index/suffixarray":                      {},
	"pkg/io":                                     {},
	"pkg/io/ioutil":                              {},
	"pkg/log":                                    {},
	"pkg/log/syslog":                             {},
	"pkg/math":                                   {},
	"pkg/math/big":                               {},
	"pkg/math/cmplx":                             {},
	"pkg/math/rand":                              {},
	"pkg/mime":                                   {},
	"pkg/mime/multipart":                         {},
	"pkg/net":                                    {},
	"pkg/net/http":                               {},
	"pkg/net/http/cgi":                           {},
	"pkg/net/http/cookiejar":                     {},
	"pkg/net/http/fcgi":                          {},
	"pkg/net/http/httptest":                      {},
	"pkg/net/http/httputil":                      {},
	"pkg/net/http/pprof":                         {},
	"pkg/net/mail":                               {},
	"pkg/net/rpc":                                {},
	"pkg/net/rpc/jsonrpc":                        {},
	"pkg/net/smtp":                               {},
	"pkg/net/textproto":                          {},
	"pkg/net/url":                                {},
	"pkg/os":                                     {},
	"pkg/os/exec":                                {},
	"pkg/os/signal":                              {},
	"pkg/os/user":                                {},
	"pkg/path":                                   {},
	"pkg/path/filepath":                          {},
	"pkg/reflect":                                {},
	"pkg/regexp":                                 {},
	"pkg/regexp/syntax":                          {},
	"pkg/runtime":                                {},
	"pkg/runtime/cgo":                            {},
	"pkg/runtime/debug":                          {},
	"pkg/runtime/pprof":                          {},
	"pkg/runtime/race":                           {},
	"pkg/sort":                                   {},
	"pkg/strconv":                                {},
	"pkg/strings":                                {},
	"pkg/sync":                                   {},
	"pkg/sync/atomic":                            {},
	"pkg/syscall":                                {},
	"pkg/testing":                                {},
	"pkg/testing/iotest":                         {},
	"pkg/testing/quick":                          {},
	"pkg/text":                                   {},
	"pkg/text/scanner":                           {},
	"pkg/text/tabwriter":                         {},
	"pkg/text/template":                          {},
	"pkg/text/template/parse":                    {},
	"pkg/time":                                   {},
	"pkg/unicode":                                {},
	"pkg/unicode/utf16":                          {},
	"pkg/unicode/utf8":                           {},
	"pkg/unsafe":                                 {},
	"plugin":                                     {},
	"reflect":                                    {},
	"reflect/internal":                           {},
	"reflect/internal/example1":                  {},
	"reflect/internal/example2":                  {},
	"regexp":                                     {},
	"regexp/syntax":                              {},
	"runtime":                                    {},
	"runtime/asan":                               {},
	"runtime/cgo":                                {},
	"runtime/coverage":                           {},
	"runtime/debug":                              {},
	"runtime/internal":                           {},
	"runtime/internal/atomic":                    {},
	"runtime/internal/math":                      {},
	"runtime/internal/startlinetest":             {},
	"runtime/internal/sys":                       {},
	"runtime/internal/syscall":                   {},
	"runtime/internal/wasitest":                  {},
	"runtime/metrics":                            {},
	"runtime/msan":                               {},
	"runtime/pprof":                              {},
	"runtime/pprof/internal":                     {},
	"runtime/pprof/internal/profile":             {},
	"runtime/pprof/internal/protopprof":          {},
	"runtime/race":                               {},
	"runtime/race/internal":                      {},
	"runtime/race/internal/amd64v1":              {},
	"runtime/race/internal/amd64v3":              {},
	"runtime/trace":                              {},
	"slices":                                     {},
	"sort":                                       {},
	"strconv":                                    {},
	"strings":                                    {},
	"structs":                                    {},
	"sync":                                       {},
	"sync/atomic":                                {},
	"syscall":                                    {},
	"syscall/js":                                 {},
	"testing":                                    {},
	"testing/fstest":                             {},
	"testing/internal":                           {},
	"testing/internal/testdeps":                  {},
	"testing/iotest":                             {},
	"testing/quick":                              {},
	"testing/slogtest":                           {},
	"testing/synctest":                           {},
	"text":                                       {},
	"text/scanner":                               {},
	"text/tabwriter":                             {},
	"text/template":                              {},
	"text/template/parse":                        {},
	"time":                                       {},
	"time/tzdata":                                {},
	"unicode":                                    {},
	"unicode/utf16":                              {},
	"unicode/utf8":                               {},
	"unique":                                     {},
	"unsafe":                                     {},
	"vendor":                                     {},
	"vendor/golang.org":                          {},
	"vendor/golang.org/x":                        {},
	"vendor/golang.org/x/crypto":                 {},
	"vendor/golang.org/x/crypto/chacha20":        {},
	"vendor/golang.org/x/crypto/chacha20poly1305":                   {},
	"vendor/golang.org/x/crypto/cryptobyte":                         {},
	"vendor/golang.org/x/crypto/cryptobyte/asn1":                    {},
	"vendor/golang.org/x/crypto/curve25519":                         {},
	"vendor/golang.org/x/crypto/curve25519/internal":                {},
	"vendor/golang.org/x/crypto/curve25519/internal/field":          {},
	"vendor/golang.org/x/crypto/hkdf":                               {},
	"vendor/golang.org/x/crypto/internal":                           {},
	"vendor/golang.org/x/crypto/internal/alias":                     {},
	"vendor/golang.org/x/crypto/internal/chacha20":                  {},
	"vendor/golang.org/x/crypto/internal/poly1305":                  {},
	"vendor/golang.org/x/crypto/internal/subtle":                    {},
	"vendor/golang.org/x/crypto/poly1305":                           {},
	"vendor/golang.org/x/crypto/sha3":                               {},
	"vendor/golang.org/x/net":                                       {},
	"vendor/golang.org/x/net/dns":                                   {},
	"vendor/golang.org/x/net/dns/dnsmessage":                        {},
	"vendor/golang.org/x/net/http":                                  {},
	"vendor/golang.org/x/net/http/httpguts":                         {},
	"vendor/golang.org/x/net/http/httpproxy":                        {},
	"vendor/golang.org/x/net/http2":                                 {},
	"vendor/golang.org/x/net/http2/hpack":                           {},
	"vendor/golang.org/x/net/idna":                                  {},
	"vendor/golang.org/x/net/lex":                                   {},
	"vendor/golang.org/x/net/lex/httplex":                           {},
	"vendor/golang.org/x/net/lif":                                   {},
	"vendor/golang.org/x/net/nettest":                               {},
	"vendor/golang.org/x/net/route":                                 {},
	"vendor/golang.org/x/sys":                                       {},
	"vendor/golang.org/x/sys/cpu":                                   {},
	"vendor/golang.org/x/text":                                      {},
	"vendor/golang.org/x/text/secure":                               {},
	"vendor/golang.org/x/text/secure/bidirule":                      {},
	"vendor/golang.org/x/text/transform":                            {},
	"vendor/golang.org/x/text/unicode":                              {},
	"vendor/golang.org/x/text/unicode/bidi":                         {},
	"vendor/golang.org/x/text/unicode/norm":                         {},
	"vendor/golang_org":                                             {},
	"vendor/golang_org/x":                                           {},
	"vendor/golang_org/x/crypto":                                    {},
	"vendor/golang_org/x/crypto/chacha20poly1305":                   {},
	"vendor/golang_org/x/crypto/chacha20poly1305/internal":          {},
	"vendor/golang_org/x/crypto/chacha20poly1305/internal/chacha20": {},
	"vendor/golang_org/x/crypto/cryptobyte":                         {},
	"vendor/golang_org/x/crypto/cryptobyte/asn1":                    {},
	"vendor/golang_org/x/crypto/curve25519":                         {},
	"vendor/golang_org/x/crypto/internal":                           {},
	"vendor/golang_org/x/crypto/internal/chacha20":                  {},
	"vendor/golang_org/x/crypto/poly1305":                           {},
	"vendor/golang_org/x/net":                                       {},
	"vendor/golang_org/x/net/dns":                                   {},
	"vendor/golang_org/x/net/dns/dnsmessage":                        {},
	"vendor/golang_org/x/net/http":                                  {},
	"vendor/golang_org/x/net/http/httpguts":                         {},
	"vendor/golang_org/x/net/http/httpproxy":                        {},
	"vendor/golang_org/x/net/http2":                                 {},
	"vendor/golang_org/x/net/http2/hpack":                           {},
	"vendor/golang_org/x/net/idna":                                  {},
	"vendor/golang_org/x/net/internal":                              {},
	"vendor/golang_org/x/net/internal/nettest":                      {},
	"vendor/golang_org/x/net/lex":                                   {},
	"vendor/golang_org/x/net/lex/httplex":                           {},
	"vendor/golang_org/x/net/lif":                                   {},
	"vendor/golang_org/x/net/nettest":                               {},
	"vendor/golang_org/x/net/proxy":                                 {},
	"vendor/golang_org/x/net/route":                                 {},
	"vendor/golang_org/x/text":                                      {},
	"vendor/golang_org/x/text/secure":                               {},
	"vendor/golang_org/x/text/secure/bidirule":                      {},
	"vendor/golang_org/x/text/transform":                            {},
	"vendor/golang_org/x/text/unicode":                              {},
	"vendor/golang_org/x/text/unicode/bidi":                         {},
	"vendor/golang_org/x/text/unicode/norm":                         {},
	"vendor/golang_org/x/text/width":                                {},
	"weak":                                                          {},
}
