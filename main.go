package main

/*
#include "postgres.h"
#include "utils/builtins.h"
#include "fmgr.h"
#include <string.h>

#cgo darwin LDFLAGS: -ldl -Wl,-undefined,dynamic_lookup
#cgo linux  LDFLAGS: -ldl -Wl,--unresolved-symbols=ignore-in-object-files

static Datum get_arg(PG_FUNCTION_ARGS, uint i) {
	return PG_GETARG_DATUM(i);
}

static char* datum_to_cstring(Datum val) {
    return DatumGetCString(text_to_cstring((struct varlena *)val));
}

static int datum_to_int16(Datum val) {
	return DatumGetInt16(val);
}

static float8 datum_to_float(Datum val) {
	return DatumGetFloat8(val);
}

*/
import "C"
import (
	"fmt"
	"unsafe"
)

func getParam(fcinfo *C.FunctionCallInfoBaseData, itemNum C.uint) C.Datum {
	return C.get_arg((*C.struct_FunctionCallInfoBaseData)(unsafe.Pointer(fcinfo)),
		itemNum)
}

//export getArgText
func getArgText(fcinfo *C.FunctionCallInfoBaseData) *C.text {

	//log.Print(C.GoString(C.datum_to_cstring(C.get_arg((*C.struct_FunctionCallInfoBaseData)(unsafe.Pointer(fcinfo)), C.uint(0)))))

	a := C.datum_to_cstring(getParam(fcinfo, 0))
	b := C.datum_to_int16(getParam(fcinfo, 1))
	c := C.datum_to_float(getParam(fcinfo, 2))

	return C.cstring_to_text(C.CString(fmt.Sprintf("Hello %s, int: %d, float: %f",
		C.GoString(a), b+1, c+1)))
}

func main() {
}
