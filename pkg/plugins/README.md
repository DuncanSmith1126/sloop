### How to write your own custom Sloop Queries

Create a directory in pkg/plugins 

Create a Go file in that directory in the 'main' package

The symbol 'Query' must resolve to an exported function with precisely this signature (Naming does not matter)

func YOUR_FUNCTION_NAME(params url.Values, tables typed.Tables, startTime time.Time, endTime time.Time, requestId string) (bytes []byte, e error) {

The symbol QueryName must resolve to a string - this string is what users in the UI will 
see and select to use your query. Plugin customers will have access to the sloop backend 
as well (documentation on how to use the Table interface is forthcoming).

Make should generate .so files in each of the directories under pkg/plugins - these
must exist at runtime in some child of the directory pointed to by the 'plugins' command line flag (defaults 
to sloop/pkg/plugins)

