#include "libg810-bridge.h"
#include <g810-led/Keyboard.h>
#include <string.h>

void* LIB_NewLedKeyboard() {
	LedKeyboard* ledKeyboard = new LedKeyboard();
	return ledKeyboard;
}

LedKeyboard *AsLedKeyboard(void* ledKeyboard) { return reinterpret_cast<LedKeyboard*>(ledKeyboard); }

void LIB_LedKeyboardDestroy(void* ledKeyboard) {
	AsLedKeyboard(ledKeyboard)->~LedKeyboard();
}

int LIB_LedKeyboardOpen(void* ledKeyboard) {
	return AsLedKeyboard(ledKeyboard)->open();
}

int LIB_LedKeyboardOpenEx(void* ledKeyboard, uint16_t vendorID, uint16_t productID, char const* serial) {
	return AsLedKeyboard(ledKeyboard)->open(vendorID, productID, serial);
}

int LIB_LedKeyboardClose(void* ledKeyboard) {
	return AsLedKeyboard(ledKeyboard)->close();
}

int LIB_LedKeyboardCommit(void* ledKeyboard) {
	return AsLedKeyboard(ledKeyboard)->commit();
}

void LIB_LedKeyboardGetDeviceInfo(void* ledKeyboard, GoDeviceInfo* pdi) {
	LedKeyboard::DeviceInfo di;
		
	di = AsLedKeyboard(ledKeyboard)->getCurrentDevice();

	pdi->vendorID = di.vendorID;
	pdi->productID = di.productID;
	pdi->model = (uint8_t)di.model;
	strncpy(pdi->manufacturer, di.manufacturer.c_str(), sizeof(pdi->manufacturer));
	strncpy(pdi->product, di.product.c_str(), sizeof(pdi->product));
	strncpy(pdi->serialNumber, di.serialNumber.c_str(), sizeof(pdi->serialNumber));
}

int LIB_LedKeyboardSetKey(void* ledKeyboard, GoKeyValue goKeyValue) {
	return AsLedKeyboard(ledKeyboard)->setKey({
		(LedKeyboard::Key)goKeyValue.key,
		{
			goKeyValue.color.red,
			goKeyValue.color.green,
			goKeyValue.color.blue,
		},
	});
}

int LIB_LedKeyboardSetKeys(void* ledKeyboard, GoKeyValue goKeyValues[], uint32_t nElements) {
	LedKeyboard::KeyValueArray keyValues = {};

	for (uint32_t n = 0; n < nElements; n++) {
		keyValues.push_back({
			(LedKeyboard::Key)goKeyValues[n].key,
			{
				goKeyValues[n].color.red,
				goKeyValues[n].color.green,
				goKeyValues[n].color.blue,
			}
		});
	}

	return AsLedKeyboard(ledKeyboard)->setKeys(keyValues);
}

int LIB_LedKeyboardSetAllKeys(void* ledKeyboard, GoKeyColor color) {
	return AsLedKeyboard(ledKeyboard)->setAllKeys({
		color.red,
		color.green,
		color.blue,
	});
}

int LIB_LedKeyboardSetGroupKeys(void* ledKeyboard, uint8_t keyGroup, GoKeyColor color) {
	return AsLedKeyboard(ledKeyboard)->setGroupKeys((LedKeyboard::KeyGroup)keyGroup, {
		color.red,
		color.green,
		color.blue,
	});
}

