/*
This is a copy of the minimum amount of the original file to be able to test the inspector linter.
*/
package types

// UID is a type that holds unique ID values, including UUIDs.  Because we
// don't ONLY use UUIDs, this is an alias to string.  Being a type captures
// intent and helps make sure that UIDs and names do not get conflated.
type UID string
