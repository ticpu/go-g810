package g810

// #cgo CXXFLAGS: -Dhidapi -std=gnu++14
// #cgo LDFLAGS: -L. -lg810-led-bridge -lg810-led
// #include "libg810-bridge.h"
import "C"

import (
	"errors"
	"unsafe"
)

type LedKeyboard struct {
	ptr unsafe.Pointer
}

type KeyColor struct {
	Red uint8
	Green uint8
	Blue uint8
}

type KeyValue struct {
	ID Key
	Color KeyColor
}

type DeviceInfo struct {
	VendorID uint16
	ProductID uint16
	Manufacturer string
	Product string
	SerialNumber string
	KeyboardModel string
}

var KeyboardModelName = map[uint8]string{
	KB_UNKNOWN: "Unknown",
	KB_G213: "G213",
	KB_G410: "G410",
	KB_G413: "G413",
	KB_G513: "G513",
	KB_G610: "G610",
	KB_G810: "G810",
	KB_G910: "G910",
	KB_GPRO: "GPro",
}

var KeyboardModel = map[uint16]map[uint16]uint8{
	0x46d: {
		0xc336: KB_G213,
		0xc330: KB_G410,
		0xc33a: KB_G413,
		0xc33c: KB_G513,
		0xc333: KB_G610,
		0xc338: KB_G610,
		0xc331: KB_G810,
		0xc337: KB_G810,
		0xc32b: KB_G910,
		0xc335: KB_G910,
		0xc339: KB_GPRO,
	},
}

func NewLedKeyboard() LedKeyboard {
	var lk LedKeyboard
	lk.ptr = C.LIB_NewLedKeyboard()
	return lk
}

func (lk LedKeyboard) Free() {
	C.LIB_LedKeyboardDestroy(lk.ptr)
}

func (lk LedKeyboard) Open() error {
	if C.LIB_LedKeyboardOpen(lk.ptr) == 1 {
		return nil
	} else {
		return errors.New("C.LIB_LedKeyboardOpen returned an error.")
	}
}

func (lk LedKeyboard) OpenEx(vendor_id uint16, product_id uint16, serial string) error {
	cSerial := C.CString(serial)
	defer C.free(unsafe.Pointer(cSerial))

	if C.LIB_LedKeyboardOpenEx(lk.ptr, C.ushort(vendor_id), C.ushort(product_id), cSerial) == 1 {
		return nil
	} else {
		return errors.New("C.LIB_LedKeyboardOpenEx returned an error.")
	}
}

func (lk LedKeyboard) Commit() error {
	if C.LIB_LedKeyboardCommit(lk.ptr) == 1 {
		return nil
	} else {
		return errors.New("C.LIB_LedKeyboardCommit returned an error.")
	}
}

func (lk LedKeyboard) Close() error {
	if C.LIB_LedKeyboardClose(lk.ptr) == 1 {
		return nil
	} else {
		return errors.New("C.LIB_LedKeyboardClose returned an error.")
	}
}

func (lk LedKeyboard) GetDeviceInfo() DeviceInfo {
	var deviceInfo DeviceInfo
	var cDeviceInfo *C.GoDeviceInfo

	cDeviceInfoPtr := C.malloc(C.sizeof_GoDeviceInfo)
	defer C.free(cDeviceInfoPtr)

	C.LIB_LedKeyboardGetDeviceInfo(lk.ptr, (*C.GoDeviceInfo)(cDeviceInfoPtr))
	cDeviceInfo = (*C.GoDeviceInfo)(cDeviceInfoPtr)

	deviceInfo = DeviceInfo{
		VendorID: uint16(cDeviceInfo.vendorID),
		ProductID: uint16(cDeviceInfo.productID),
		Manufacturer: C.GoStringN(&cDeviceInfo.manufacturer[0], 64),
		Product: C.GoStringN(&cDeviceInfo.product[0], 64),
		SerialNumber: C.GoStringN(&cDeviceInfo.serialNumber[0], 64),
		KeyboardModel: KeyboardModelName[KeyboardModel[uint16(cDeviceInfo.vendorID)][uint16(cDeviceInfo.productID)]],
	}

	return deviceInfo
}

func (lk LedKeyboard) SetKey(key KeyValue) error {
	cKeyValue := C.GoKeyValue{
		key: (C.ushort)(key.ID),
		color: C.GoKeyColor{
			(C.uchar)(key.Color.Red),
			(C.uchar)(key.Color.Green),
			(C.uchar)(key.Color.Blue),
		},
	}

	if C.LIB_LedKeyboardSetKey(lk.ptr, cKeyValue) == 1 {
		return nil
	} else {
		return errors.New("C.LIB_LedKeyboardSetKey returned an error.")
	}
}

func (lk LedKeyboard) SetKeys(keys []KeyValue) error {
	cKeyValuesCount := len(keys)

	if cKeyValuesCount < 1 {
		return errors.New("No keys in dataset.")
	}

	cKeyValuesPtr := C.malloc((C.ulong)(C.sizeof_GoKeyValue*cKeyValuesCount))
	defer C.free(cKeyValuesPtr)
	cKeyValues := (*[1<<16]C.GoKeyValue)(cKeyValuesPtr)

	for n, key := range keys {
		cKeyValues[n].key = (C.ushort)(key.ID)
		cKeyValues[n].color = C.GoKeyColor{
			(C.uchar)(key.Color.Red),
			(C.uchar)(key.Color.Green),
			(C.uchar)(key.Color.Blue),
		}
	}

	if C.LIB_LedKeyboardSetKeys(lk.ptr, (*C.GoKeyValue)(cKeyValuesPtr), C.uint(cKeyValuesCount)) == 1 {
		return nil
	} else {
		return errors.New("C.LIB_LedKeyboardSetKeys returned an error.")
	}
}

func (lk LedKeyboard) SetAllKeys(color KeyColor) error {
	cKeyColor := C.GoKeyColor{
		(C.uchar)(color.Red),
		(C.uchar)(color.Green),
		(C.uchar)(color.Blue),
	}

	if C.LIB_LedKeyboardSetAllKeys(lk.ptr, cKeyColor) == 1 {
		return nil
	} else {
		return errors.New("C.LIB_LedKeyboardSetAllKeys returned an error.")
	}
}

func (lk LedKeyboard) SetGroupKeys(keyGroup KeyGroup, color KeyColor) error {
	cKeyColor := C.GoKeyColor{
		(C.uchar)(color.Red),
		(C.uchar)(color.Green),
		(C.uchar)(color.Blue),
	}

	if C.LIB_LedKeyboardSetGroupKeys(lk.ptr, (C.uchar)(keyGroup), cKeyColor) == 1 {
		return nil
	} else {
		return errors.New("C.LIB_LedKeyboardSetGroupKeys returned an error.")
	}
}

