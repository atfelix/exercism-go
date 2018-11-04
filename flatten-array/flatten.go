package flatten

func Flatten(list interface{}) []interface{} {
	if list == nil {
		return []interface{}{}
	} else if _, ok := list.([]interface{}); !ok {
		return []interface{}{list}
	}

	collection := make([]interface{}, 0)

	for _, element := range list.([]interface{}) {
		collection = append(collection, Flatten(element)...)
	}
	return collection
}