func errorHandler(errp *error) {
	if e := recover(); e != nil {
		if se, ok := e.(scanError); ok { // catch local error
			*errp = se.err
		} else if eof, ok := e.(error); ok && eof == io.EOF { // out of input
			*errp = eof
		} else {
			panic(e)
		}
	}
}

func handler() {
	if destring {
		switch qv := d.valueQuoted().(type) {
		case nil:
			d.literalStore(nullLiteral, subv, false)
		case string:
			d.literalStore([]byte(qv), subv, true)
			// ... other code

		}
	}
}

func otherHandler() {
	if destring {
		q, err := d.valueQuoted()
		if err != nil {
			return err
		}
		switch qv := q.(type) {
		case nil:
			if err := d.literalStore(nullLiteral, subv, false); err != nil {
				return err
			}
		case string:
			if err := d.literalStore([]byte(qv), subv, true); err != nil {
				return err
			}

		}
	}
}