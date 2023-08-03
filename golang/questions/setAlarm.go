package questions

// hard case
func SetAlarm(employed, vacation bool) bool {

	var results bool

	if employed && vacation {
		results = false
	} else if employed && !vacation {
		results = true
	} else if !employed && vacation {
		results = false
	} else {
		results = false
	}

	return results
}

// simple case
func SetAlarm2(employed, vacation bool) bool {
	return employed && !vacation
}

// 	employed | vacation
// true     | true     => false
// true     | false    => true
// false    | true     => false
// false    | false    => false
