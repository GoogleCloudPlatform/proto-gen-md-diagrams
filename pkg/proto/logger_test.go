package proto

import "testing"

func TestLogger_Debug(t *testing.T) {
	type fields struct {
		debug bool
	}
	type args struct {
		in string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{name: "Write Debug", fields: fields{debug: true}, args: args{in: "test"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := Logger{
				debug: tt.fields.debug,
			}
			l.Debug(tt.args.in)
		})
	}
}

func TestLogger_Debugf(t *testing.T) {
	type fields struct {
		debug bool
	}
	type args struct {
		in   string
		args []any
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name:   "Write Debugf",
			fields: fields{debug: true},
			args: args{
				in:   "Test %s\n",
				args: []any{"test"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := Logger{
				debug: tt.fields.debug,
			}
			l.Debugf(tt.args.in, tt.args.args...)
		})
	}
}

func TestLogger_Error(t *testing.T) {
	type fields struct {
		debug bool
	}
	type args struct {
		in string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name:   "Test Error",
			fields: fields{debug: true},
			args: args{
				in: "Test Error",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := Logger{
				debug: tt.fields.debug,
			}
			l.Error(tt.args.in)
		})
	}
}

func TestLogger_Errorf(t *testing.T) {
	type fields struct {
		debug bool
	}
	type args struct {
		in   string
		args []any
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name:   "Test ErrorF",
			fields: fields{debug: true},
			args: args{
				in:   "Test Error: %s\n",
				args: []any{"error"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := Logger{
				debug: tt.fields.debug,
			}
			l.Errorf(tt.args.in, tt.args.args...)
		})
	}
}

func TestLogger_Info(t *testing.T) {
	type fields struct {
		debug bool
	}
	type args struct {
		in string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name:   "Test Info",
			fields: fields{debug: true},
			args: args{
				in: "test",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := Logger{
				debug: tt.fields.debug,
			}
			l.Info(tt.args.in)
		})
	}
}

func TestLogger_Infof(t *testing.T) {
	type fields struct {
		debug bool
	}
	type args struct {
		in   string
		args []any
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "Test InfoF",
			fields: fields{
				debug: true,
			},
			args: args{
				in:   "test info f: %s\n",
				args: []any{"test"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := Logger{
				debug: tt.fields.debug,
			}
			l.Infof(tt.args.in, tt.args.args...)
		})
	}
}
