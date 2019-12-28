#include <stdbool.h>
#include <stdint.h>
#include <stdlib.h>

#pragma once
#ifdef __cplusplus
extern "C" {
#endif

typedef struct {
	uint16_t vendorID;
	uint16_t productID;
	char manufacturer[64];
	char product[64];
	char serialNumber[64];
	uint8_t model;
} GoDeviceInfo;

typedef struct {
	uint8_t red;
	uint8_t green;
	uint8_t blue;
} GoKeyColor;

typedef struct {
	uint16_t key;
	GoKeyColor color;
} GoKeyValue;

void* LIB_NewLedKeyboard();
void LIB_LedKeyboardDestroy(void* ledKeyboard);
int LIB_LedKeyboardOpen(void* ledKeyboard);
int LIB_LedKeyboardOpenEx(void* ledKeyboard, uint16_t vendorID, uint16_t productID, char const* serial);
int LIB_LedKeyboardClose(void* ledKeyboard);
void LIB_LedKeyboardGetDeviceInfo(void* ledKeyboard, GoDeviceInfo* pdi);
int LIB_LedKeyboardCommit(void* ledKeyboard);
int LIB_LedKeyboardSetKey(void* ledKeyboard, GoKeyValue keyValue);
int LIB_LedKeyboardSetKeys(void* ledKeyboard, GoKeyValue keyValues[], uint32_t nElements);

#ifdef __cplusplus
}
#endif
