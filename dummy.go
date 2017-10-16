/*
 * Dummy go file to make go get github.com/snapcore/squashfuse work
 */

package dummy

// #cgo CFLAGS: -wrapper ./dummy-gcc
// #cgo LDFLAGS: -wrapper ./dummy-ld
import "C"
