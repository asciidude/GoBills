# GoBills
GoBills is a simple billing system created in [Golang v1.19.2](https://go.dev/)

GoBills does the following:

* Stores bills in a specified directory, can be changed
* Uses a structure for items for type safety
* Returns the saved file path, tax, after tax total, before tax total, and time 
* Calculates before and after sales tax
* Has a simple codebase, easy to read and edit
* Makes storing bills easier, in general

GoBills does **not** actually transfer money, I'm afraid if you're looking for that,
you'll have to look elsewhere. GoBills is simply an easy way to store bills.