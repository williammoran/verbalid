# verbalid
Convert number IDs to/from alphanumeric representation that can be unambiguously communicated by humans.

The entire point is to represent IDs in a manner that is as easy as possible for humans to communicated correctly.

Think someone giving an ID over the phone.

* The alphabetic part of the ID is case-insensative
* Any characters that could be confusing are omitted. Thus there is no 0 or O, etc.

# Use

Do the typical stuff to include the library in your go project, then:

`IntToVerbalID(int)` returns a string that meets the above requirements.

`VerbalIDToInt(string)` returns the int that was originally use to generate the Verbal ID.
