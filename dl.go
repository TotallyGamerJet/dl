package dl

import (
	"errors"
	"runtime"
)

var (
	ErrNotFound = errors.New("not found")
)

const (
	// Each external function reference is bound the first time
	// the function is called.
	//
	// This is the default behavior.
	//
	BindLazy = _RTLD_LAZY

	// All external function references are bound immediately
	// during the call to Open.
	//
	// Lazy is normally preferred, for reasons of efficiency.
	// However, BindNow is useful to ensure that any undefined
	// symbols are discovered.
	//
	BindNow = _RTLD_NOW

	// One of the following may be ORed into the mode argument:

	// Exported symbols will be available to any images built with
	// -flat_namespace option to ld(1) or to calls to Lookup functions.
	//
	// This is the default behavior.
	//
	ScopeGlobal = _RTLD_GLOBAL

	// Exported symbols are generally hidden and only availble
	// to Lookup when directly using the Dylib returned by Open.
	//
	ScopeLocal = _RTLD_LOCAL

	// The image is not loaded. However, a valid Dylib is returned if
	// the image already exists in the process. This provides a way to
	// query if an image is already loaded.
	//
	// You eventually need a corresponding call to Close.
	//
	NoLoad = _RTLD_NOLOAD

	// The image will never be removed from the address space,
	// even after all clients have released it via Close.
	//
	NoDelete = _RTLD_NODELETE

	// Additionally, the following may be ORed into the mode argument:

	// Lookup calls will only search the image specified, and not
	// subsequent images.
	//
	LookupFirst = _RTLD_FIRST
)

const (
	_RTLD_LAZY     = 0x1
	_RTLD_NOW      = 0x2
	_RTLD_LOCAL    = 0x4
	_RTLD_GLOBAL   = 0x8
	_RTLD_NOLOAD   = 0x10
	_RTLD_NODELETE = 0x80
	_RTLD_FIRST    = 0x100
)

type Handle uintptr

type Dylib struct {
	Name   string
	Handle Handle
}

// MustOpen is like Open but panics if operation failes.
func MustOpen(path string, mode int) *Dylib {
	l, e := Open(path, mode)
	if e != nil {
		panic(e)
	}
	return l
}

// Open loads and links a dynamic library or bundle into the current process.
//
// Open examines the Mach-O file specified by path. If the file is compatible
// with the current process and has not already been loaded, it is loaded and
// linked. After being linked, if it contains any initializer functions, they
// are called, before Open returns.
//
// Open searches for a compatible Mach-O file in the directories specified by
// a set of environment variables and the process’s current working directory.
// When set, the environment variables must contain a colon-separated list of
// directory paths, which can be absolute or relative to the current working
// directory. The environment variables are LD_LIBRARY_PATH, DYLD_LIBRARY_PATH,
// and DYLD_FALLBACK_LIBRARY_PATH. The default value of the latter variable is
// $HOME/lib;/usr/local/lib;/usr/lib. The first two variables have no default
// value. Open searches the directories specified in the environment variables
// in the order they are listed.
//
// When path doesn’t contain a slash character (i.e. it is just a leaf name),
// Open searches the following the locations until it finds a compatible Mach-O
// file: $LD_LIBRARY_PATH, $DYLD_LIBRARY_PATH, current working directory,
// $DYLD_FALLBACK_LIBRARY_PATH.
//
// When path contains a slash (i.e. a full path or a partial path), Open
// searches the following the locations until it finds a compatible Mach-O
// file: $DYLD_LIBRARY_PATH (with leaf name from path), current working
// directory (for partial paths), $DYLD_FALLBACK_LIBRARY_PATH (with leaf
// name from path).
//
// Note: There are no configuration files that control dlopen searching.
//
// Note: If the main executable is a set[ug]id binary, then all environment
// variables are ignored, and only a full path can be used.
//
// Note: macOS uses "universal" files to combine multiarch libraries. This also
// means that there are no separate 32-bit and 64-bit search paths.
//
// See dlopen(3).
//
func Open(path string, mode int) (*Dylib, error) {
	p, err := cstring(path)
	if err != nil {
		return nil, err
	}

	runtime.LockOSThread()
	defer runtime.UnlockOSThread()

	handle := dlopen(p, mode)
	if handle <= 0 {
		return nil, lastError()
	}
	h := Handle(handle)
	d := &Dylib{path, h}
	runtime.SetFinalizer(d, (*Dylib).Close)
	return d, nil
}

// Lookup searches symbol with name.
//
// Returns the address of the code or data location specified by the symbol name.
//
// See dlsym(3).
//
func (d *Dylib) Lookup(name string) (uintptr, error) {
	p, err := cstring(name)
	if err != nil {
		return 0, err
	}

	runtime.LockOSThread()
	defer runtime.UnlockOSThread()

	ret := dlsym(uintptr(d.Handle), p)
	// We must check dlerrgl.go
	//gl.sor because symbol could be NULL.
	if err = lastError(); err != nil {
		return 0, err
	}
	return ret, nil
}

// Close closes a dynamic library or bundle.
//
// Close releases a reference to the dynamic library or bundle. If the reference
// count drops to 0, the bundle is removed from the address space. Just before
// removing a dynamic library or bundle in this way, any termination routines in
// it are called.
//
// There are a couple of cases in which a dynamic library will never be unloaded:
//
//   1) The main executable links against it,
//   2) An API that does not supoort unloading (e.g. NSAddImage()) was used
//      to load it or some other dynamic library that depends on it,
//   3) The dynamic library is in dyld’s shared cache.
//
// See dlclose(3).
//
func (d *Dylib) Close() (err error) {
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()

	ret := dlclose(uintptr(d.Handle))
	if ret != 0 {
		err = lastError()
	}

	// No need for a finalizer anymore.
	runtime.SetFinalizer(d, nil)
	return err
}

// lastError returns an error describing the last dyld error that occurred
// on this thread. At each call to lastError, the error indication is reset.
// Thus in the case of two calls to lastError, where the second call follows
// the first immediately, the second call will always return nil.
//
// See dlerror(3).
//
func lastError() error {
	ret := dlerror()
	if ret != 0 {
		s := gostring(ret)
		// TODO export error vars by known string suffixes.
		// E.g. strings.HasSuffix(s, ": symbol not found").
		err := errors.New(s)
		return err
	}
	return nil
}
