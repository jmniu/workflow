package qall

import (
	"github.com/jmniu/workflow/qlang/lib/bufio"
	"github.com/jmniu/workflow/qlang/lib/bytes"
	"github.com/jmniu/workflow/qlang/lib/crypto/md5"
	"github.com/jmniu/workflow/qlang/lib/encoding/hex"
	"github.com/jmniu/workflow/qlang/lib/encoding/json"
	"github.com/jmniu/workflow/qlang/lib/eqlang"
	"github.com/jmniu/workflow/qlang/lib/errors"
	"github.com/jmniu/workflow/qlang/lib/io"
	"github.com/jmniu/workflow/qlang/lib/io/ioutil"
	"github.com/jmniu/workflow/qlang/lib/math"
	"github.com/jmniu/workflow/qlang/lib/meta"
	"github.com/jmniu/workflow/qlang/lib/net/http"
	"github.com/jmniu/workflow/qlang/lib/os"
	"github.com/jmniu/workflow/qlang/lib/path"
	"github.com/jmniu/workflow/qlang/lib/reflect"
	"github.com/jmniu/workflow/qlang/lib/runtime"
	"github.com/jmniu/workflow/qlang/lib/strconv"
	"github.com/jmniu/workflow/qlang/lib/strings"
	"github.com/jmniu/workflow/qlang/lib/sync"
	"github.com/jmniu/workflow/qlang/lib/terminal"
	"github.com/jmniu/workflow/qlang/lib/tpl/extractor"
	"github.com/jmniu/workflow/qlang/lib/version"
	qlang "github.com/jmniu/workflow/qlang/spec"

	// qlang builtin modules
	_ "github.com/jmniu/workflow/qlang/lib/builtin"
	_ "github.com/jmniu/workflow/qlang/lib/chan"
)

// -----------------------------------------------------------------------------

// Copyright prints qlang copyright information.
//
func Copyright() {
	version.Copyright()
}

// InitSafe inits qlang and imports modules.
//
func InitSafe(safeMode bool) {

	qlang.SafeMode = safeMode

	qlang.Import("", math.Exports) // import math as builtin package
	qlang.Import("", meta.Exports) // import meta package
	qlang.Import("bufio", bufio.Exports)
	qlang.Import("bytes", bytes.Exports)
	qlang.Import("md5", md5.Exports)
	qlang.Import("io", io.Exports)
	qlang.Import("ioutil", ioutil.Exports)
	qlang.Import("hex", hex.Exports)
	qlang.Import("json", json.Exports)
	qlang.Import("errors", errors.Exports)
	qlang.Import("eqlang", eqlang.Exports)
	qlang.Import("math", math.Exports)
	qlang.Import("os", os.Exports)
	qlang.Import("", os.InlineExports)
	qlang.Import("path", path.Exports)
	qlang.Import("http", http.Exports)
	qlang.Import("reflect", reflect.Exports)
	qlang.Import("runtime", runtime.Exports)
	qlang.Import("strconv", strconv.Exports)
	qlang.Import("strings", strings.Exports)
	qlang.Import("sync", sync.Exports)
	qlang.Import("terminal", terminal.Exports)
	qlang.Import("extractor", extractor.Exports)
}

// -----------------------------------------------------------------------------
