package usbtmc

import (
	"testing"
	"unsafe"
)

func Test_genIoctl(t *testing.T) {
	type args struct {
		dir        uint8
		ioctl_type uint8
		ioctl_nr   uint8
		size       uint16
	}
	tests := []struct {
		name    string
		args    args
		wantOut uint32
		wantErr bool
	}{
		{
			name: "BLKRASET",
			args: args{
				dir:        0,
				ioctl_type: 0x12,
				ioctl_nr:   98,
			},
			wantOut: 0x00001262,
			wantErr: false,
		},
		{
			name: "FS_IOC_GETFLAGS",
			args: args{
				dir:        _IOC_READ,
				ioctl_type: 'f',
				ioctl_nr:   1,
				size:       uint16(unsafe.Sizeof(int(0))),
			},
			wantOut: 0x80086601,
			wantErr: false,
		},
		{
			name: "too big",
			args: args{
				dir:        _IOC_READ,
				ioctl_type: 'f',
				ioctl_nr:   1,
				size:       2^15,
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotOut, err := genIoctl(tt.args.dir, tt.args.ioctl_type, tt.args.ioctl_nr, tt.args.size)
			if (err != nil) != tt.wantErr {
				t.Errorf("genIoctl() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotOut != tt.wantOut {
				t.Errorf("genIoctl() = 0x%x, want 0x%x", gotOut, tt.wantOut)
			}
		})
	}
}
