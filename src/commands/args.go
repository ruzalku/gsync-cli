package commands

import "errors"

func BoolToString(flag bool) string {
	if flag {
		return "1"
	}
	return "0"
}

func ArgsToCommand(c ...any) (string, error) {
	command := ""

	for i := 0; i < len(c); i++ {
		switch f := c[i].(type) {
		case string:
			command += f
		case bool:
			command += BoolToString(f)
		default:
			return "", errors.New("Undefined type")
		}
	}
	
	return command, nil
}