// ARM version of SHA256

TEXT	·block(SB), NOSPLIT, $0
	MOVD   message+0(FP), R0
	MOVD   0(R0), R1
	ADD    $1, R1
	MOVD   R1, 0(R0)
	RET
