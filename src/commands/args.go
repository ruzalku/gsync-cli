package commands

import "errors"

func BoolToBytes(flag bool) []byte {
	if flag {
		return []byte("1")
	}
	return []byte("0")
}

func ArgsToCommand(c ...any) ([]byte, error) {
	command := make([]byte, 0, 1024)

	for i := 0; i < len(c); i++ {
		switch f := c[i].(type) {
		case string:
			command = append(command, []byte(f)...)
		case bool:
			command = append(command, BoolToBytes(f)...)
		default:
			return []byte{}, errors.New("Undefined type")
		}
		command = append(command, 59)
	}
	
	return command[:len(command) - 1], nil
}