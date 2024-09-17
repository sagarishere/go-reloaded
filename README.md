# Go Reloaded

In go reloaded, we manipulate a string from a file, to make it correct according to the specifications provided to us.

## Detailed subject

<https://github.com/01-edu/public/tree/master/subjects/go-reloaded/>

## Audit questions

<https://github.com/01-edu/public/tree/master/subjects/go-reloaded/audit>

P.S. The audit questions are also included in the test cases as well as in test files in the source2copy folder.

## Learnings (in golang)

- read content of a file
- string manipulation methods in golang
- for loops
- switch statements
- if statements
- finding next token (in our case we do not do runes so our tokens are mostly words) and manipulating it according to prev token
- pointers
- write content to a file
- testing
- big O notation (our program scales well as we used 2 for loops only and no nested loops)

## Additional Scripts

Python Script to copy content of one file to another. This was done to run tests and audits faster.
To copy content of file ./source2copy/1.txt to ./sample.txt, run the following command:

```bash
python3 cpf2f.py 1
```

## Testing

```bash
go test -v
```

## Scope of development of project

The project was mostly done to pass the tests and audits. The code is not optimized and can be improved.

## Scope of improvmenet

- No benchmarking done
- No coverage report generated
- File contents read without buffer
- No extensive error handling
- Does not cover edge cases

## Date

14 Sep 2022

## Author

- [Sagar Yadav](https://linkedin.com/in/sagaryadav)
