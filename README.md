# GoBills
GoBills is a simple billing system created in [Golang v1.19.2](https://go.dev/)

GoBills does the following:

* Stores bills in a specified directory, can be changed
* Uses a structure for items for type safety
* Returns the saved file path when the function is called
* Calculates before and after sales tax
* Has a simple codebase, easy to read and edit
* Makes storing bills easier, in general

GoBills does **not** make any calls to any sort of API to actually transfer money,
I'm afraid if you're looking for this, you'll have to look elsewhere.