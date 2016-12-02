// go.mkdef -w sys_unix.go sys_linux.go
// MACHINE GENERATED BY go.mkdef (github.com/tredoe/gotool/go.mkdef); DO NOT EDIT
//
// Created by cgo -godefs - DO NOT EDIT
// cgo -godefs _z-sys_linux_amd64.go

package sys

const (
	TCSADRAIN = 0x1
	TCSAFLUSH = 0x2
	TCSANOW   = 0x0
)

const TIOCGWINSZ = 0x5413

const (
	VDISCARD = 0xd
	VEOF     = 0x4
	VEOL     = 0xb
	VEOL2    = 0x10
	VERASE   = 0x2
	VINTR    = 0x0
	VKILL    = 0x3
	VLNEXT   = 0xf
	VMIN     = 0x6
	VQUIT    = 0x1
	VREPRINT = 0xc
	VSTART   = 0x8
	VSTOP    = 0x9
	VSUSP    = 0xa
	VTIME    = 0x5
	VWERASE  = 0xe
)

const (
	BRKINT  = 0x2
	ICRNL   = 0x100
	IGNBRK  = 0x1
	IGNCR   = 0x80
	IGNPAR  = 0x4
	IMAXBEL = 0x2000
	INLCR   = 0x40
	INPCK   = 0x10
	ISTRIP  = 0x20
	IXANY   = 0x800
	IXOFF   = 0x1000
	IXON    = 0x400
	PARMRK  = 0x8
)

const (
	BS0    = 0x0
	BS1    = 0x2000
	CR0    = 0x0
	CR1    = 0x200
	CR2    = 0x400
	CR3    = 0x600
	FF0    = 0x0
	FF1    = 0x8000
	NL0    = 0x0
	NL1    = 0x100
	OCRNL  = 0x8
	ONLCR  = 0x4
	ONLRET = 0x20
	ONOCR  = 0x10
	OPOST  = 0x1
	TAB0   = 0x0
	TAB1   = 0x800
	TAB2   = 0x1000
	XTABS  = 0x1800
)

const (
	B0      = 0x0
	B110    = 0x3
	B115200 = 0x1002
	B1200   = 0x9
	B134    = 0x4
	B150    = 0x5
	B1800   = 0xa
	B19200  = 0xe
	B200    = 0x6
	B230400 = 0x1003
	B2400   = 0xb
	B300    = 0x7
	B38400  = 0xf
	B4800   = 0xc
	B50     = 0x1
	B57600  = 0x1001
	B600    = 0x8
	B75     = 0x2
	B9600   = 0xd
	CLOCAL  = 0x800
	CREAD   = 0x80
	CRTSCTS = 0x80000000
	CS5     = 0x0
	CS6     = 0x10
	CS7     = 0x20
	CS8     = 0x30
	CSIZE   = 0x30
	CSTOPB  = 0x40
	EXTA    = 0xe
	EXTB    = 0xf
	HUPCL   = 0x400
	PARENB  = 0x100
	PARODD  = 0x200
)

const (
	ECHO    = 0x8
	ECHOCTL = 0x200
	ECHOE   = 0x10
	ECHOK   = 0x20
	ECHOKE  = 0x800
	ECHONL  = 0x40
	ECHOPRT = 0x400
	EXTPROC = 0x10000
	FLUSHO  = 0x1000
	ICANON  = 0x2
	IEXTEN  = 0x8000
	ISIG    = 0x1
	NOFLSH  = 0x80
	PENDIN  = 0x4000
	TOSTOP  = 0x100
)

const (
	TCGETS  = 0x5401
	TCSETS  = 0x5402
	TCSETSF = 0x5404
	TCSETSW = 0x5403
)

type Termios struct {
	Iflag uint32
	Oflag uint32
	Cflag uint32
	Lflag uint32
	Line  uint8
	Cc    [19]uint8
}

type Winsize struct {
	Row    uint16
	Col    uint16
	Xpixel uint16
	Ypixel uint16
}
