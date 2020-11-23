package main

// Arguments to format are:
//	[1]: type name
const gqlMarshalMethod = `func (i %[1]s) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(i.String()))
}
`

const gqlUnMarshalMethod = `func (i *%[1]s) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	if !i.IsA%[1]s() {
		return fmt.Errorf("%%s is not a valid %[1]s", str)
	}

	val, err := %[1]sString(str)
	if err != nil {
		return err
	}

	*i = val
	return nil
}
`

func (g *Generator) addGQLMethods(typeName string) {
	g.Printf("\n")
	g.Printf(gqlMarshalMethod, typeName)
	g.Printf("\n\n")
	g.Printf(gqlUnMarshalMethod, typeName)
}
