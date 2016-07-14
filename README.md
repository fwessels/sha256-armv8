# sha256-armv8

sources
+ https://github.com/jocover/sha256-armv8
+ http://lxr.free-electrons.com/source/arch/arm64/crypto/sha2-ce-core.S
+ https://github.com/openssl/openssl/blob/master/crypto/sha/asm/sha512-armv8.pl

```
linaro@linaro-alip:~/gopath/src/github.com/minio/blake2b-simd$ go test -bench .
PASS
BenchmarkComparisonMD5-8    	      50	  25276760 ns/op	  41.48 MB/s
BenchmarkComparisonSHA1-8   	      10	 105594200 ns/op	   9.93 MB/s
BenchmarkComparisonSHA256-8 	      10	 166012400 ns/op	   6.32 MB/s
BenchmarkComparisonSHA512-8 	      20	  64282600 ns/op	  16.31 MB/s
BenchmarkComparisonBlake2B-8	      30	  41451466 ns/op	  25.30 MB/s
```

```
linaro@linaro-alip:~/gopath/src/github.com/minio$ more /proc/cpuinfo 
processor	: 0
BogoMIPS	: 2.40
Features	: fp asimd evtstrm aes pmull sha1 sha2 crc32
CPU implementer	: 0x41
CPU architecture: 8
CPU variant	: 0x0
CPU part	: 0xd03
CPU revision	: 3

processor	: 1...2...3...4...5...6

processor	: 7
BogoMIPS	: 2.40
Features	: fp asimd evtstrm aes pmull sha1 sha2 crc32
CPU implementer	: 0x41
CPU architecture: 8
CPU variant	: 0x0
CPU part	: 0xd03
```

## ARM v8

```
linaro@linaro-alip:~/gopath/src/github.com/minio/blake2b-simd$ openssl speed sha256
Doing sha256 for 3s on 16 size blocks: 3114891 sha256's in 3.00s
Doing sha256 for 3s on 64 size blocks: 2779840 sha256's in 3.00s
Doing sha256 for 3s on 256 size blocks: 2085521 sha256's in 3.00s
Doing sha256 for 3s on 1024 size blocks: 1039140 sha256's in 3.00s
Doing sha256 for 3s on 8192 size blocks: 183896 sha256's in 3.00s
OpenSSL 1.0.2h  3 May 2016
built on: reproducible build, date unspecified
options:bn(64,64) rc4(ptr,char) des(idx,cisc,16,int) aes(partial) blowfish(ptr) 
compiler: gcc -I. -I.. -I../include  -fPIC -DOPENSSL_PIC -DOPENSSL_THREADS -D_REENTRANT -DDSO_DLFCN -DHAVE_DLFCN_H -DL_ENDIAN -g -O2 -fstack-protector-strong -Wformat -Werror=format-security -D_FORTIFY_SOURCE=2 -Wl,-z,relro -Wa,--noexecstack -Wall -DSHA1_ASM -DSHA256_ASM -DSHA512_ASM
The 'numbers' are in 1000s of bytes per second processed.
type             16 bytes     64 bytes    256 bytes   1024 bytes   8192 bytes
sha256           16612.75k    59303.25k   177964.46k   354693.12k   502158.68k
linaro@linaro-alip:~/gopath/src/github.com/minio/blake2b-simd$ 
```

## MacBook Pro

```
franks-mbp:~ frankw$ openssl speed sha256
To get the most accurate results, try to run this
program when this computer is idle.
Doing sha256 for 3s on 16 size blocks: 6128519 sha256's in 3.00s
Doing sha256 for 3s on 64 size blocks: 3608096 sha256's in 3.00s
Doing sha256 for 3s on 256 size blocks: 1598532 sha256's in 3.00s
Doing sha256 for 3s on 1024 size blocks: 503032 sha256's in 3.00s
Doing sha256 for 3s on 8192 size blocks: 67747 sha256's in 3.00s
OpenSSL 0.9.8zg 14 July 2015
built on: Aug 22 2015
options:bn(64,64) md2(int) rc4(ptr,char) des(idx,cisc,16,int) aes(partial) blowfish(idx) 
compiler: -arch x86_64 -fmessage-length=0 -pipe -Wno-trigraphs -fpascal-strings -fasm-blocks -O3 -D_REENTRANT -DDSO_DLFCN -DHAVE_DLFCN_H -DL_ENDIAN -DMD32_REG_T=int -DOPENSSL_NO_IDEA -DOPENSSL_PIC -DOPENSSL_THREADS -DZLIB -mmacosx-version-min=10.6
available timing options: TIMEB USE_TOD HZ=100 [sysconf value]
timing function used: getrusage
The 'numbers' are in 1000s of bytes per second processed.
type             16 bytes     64 bytes    256 bytes   1024 bytes   8192 bytes
sha256           32706.17k    77068.82k   136447.68k   171700.04k   185236.21k
```
