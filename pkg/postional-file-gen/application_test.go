package postionalfilegen

import (
	"fmt"
	"strings"
	"testing"
)

/**

Discusssion on how to create a datamapper
see we will be having datamapper which will alread have some predefined
data like In header most of the data except batch_number and no_of records will
be availaible.

In details most of the data except linenumber and other needs to be created at runtime from predefined data.

what will be the data source this also needs to be taken care of we need to
think about his particular thingh what will be the datasource and how to deal with that
stuff.

First we need to get all active data from some table.
then we need to creat a predefined database set which will be static and
check for fields and add it to our table.



**/

func TestString(t *testing.T) {
	// Create a DataMapperObj instance with sample data
	mapper := DataMapperObj{
		Header: map[string]Record{
			"Name": {"John", 15, 10},
			"Age":  {"30", 25, 2},
			"City": {"New York", 45, 20},
		},
		HeaderLength: 60,
		Details: []map[string]Record{
			{
				"Name": {"Alice", 15, 10},
				"Age":  {"25", 25, 2},
				"City": {"Los Angeles", 45, 20},
			},
		},
		DetailLength: 60,
		FileName:     "example.txt",
		BatchNumber:  "12345",
	}
	val, error := mapper.CreateZipFile()

	fmt.Println(val, error)
	// Expected result
	expected := strings.Join([]string{
		"               John            30                    New York                    ",
		"               Alice           25                    Los Angeles                 ",
	}, "\n") + "\n"

	// Get the string representation using the String method
	result := mapper.String()

	// Compare the result with the expected value
	if result != expected {
		t.Errorf("String representation mismatch. Expected: %s, Got: %s", expected, result)
	}
}

// func TestString(t *testing.T) {
// 	// Create a DataMapperObj instance with sample data
// 	mapper := DataMapperObj{
// 		Header: map[string]Record{
// 			"Name": {"John", 10, 10},
// 			"Age":  {"30", 15, 2},
// 			"City": {"New York", 45, 20},
// 		},
// 		HeaderLength: 90,
// 		Details: []map[string]Record{
// 			{
// 				"Name": {"Alice", 10, 10},
// 				"Age":  {"25", 15, 2},
// 				"City": {"Los Angeles", 45, 20},
// 			},
// 		},
// 		DetailLength: 80,
// 		FileName:     "example.txt",
// 		BatchNumber:  "12345",
// 	}

// 	// Expected result
// 	expected := "    John 30          New York          \n    Alice25          Los Angeles       \n"

// 	// Get the string representation using the String method
// 	result := mapper.String()

// 	// Compare the result with the expected value
// 	if result != expected {
// 		t.Errorf("String representation mismatch. Expected: %s, Got: %s", expected, result)
// 	}
// }
