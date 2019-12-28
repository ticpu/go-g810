libg810-led-bridge.so: libg810-bridge.cpp
	clang++ -o libg810-led-bridge.so libg810-bridge.cpp -lg810-led \
		-std=c++17 -fPIC -shared

clean:
	rm -f libg810-led-bridge.so
